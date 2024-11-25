package whatsapp

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AndrusGerman/fumiko/internal/core/ports"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	"go.uber.org/fx"
)

type whatsapp struct {
	client      *whatsmeow.Client
	deviceStore *store.Device
	events      chan *events.Message

	socialHandlers []ports.SocialHandler
}

func (w *whatsapp) AddHandlers(handlers ...ports.SocialHandler) {
	w.socialHandlers = handlers
}

func New(lc fx.Lifecycle) ports.Social {
	var social = &whatsapp{
		events: make(chan *events.Message),
	}

	lc.Append(fx.StopHook(social.Close))
	lc.Append(fx.StartHook(social.Start))
	return social
}

func (w *whatsapp) Register() error {
	if w.client.Store.ID != nil {
		log.Println("is alredy exist")
		return nil
	}

	// No ID stored, new login
	qrChan, _ := w.client.GetQRChannel(context.Background())
	err := w.client.Connect()
	if err != nil {
		return err
	}
	for evt := range qrChan {
		if evt.Event == "code" {
			// Render the QR code here
			qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			// or just manually `echo 2@... | qrencode -t ansiutf8` in a terminal
			//fmt.Println("QR code:", evt.Code)
		} else {
			fmt.Println("Login event:", evt.Event)
		}
	}

	time.Sleep(time.Second * 1)
	w.deviceStore.Save()
	return nil
}

func (w *whatsapp) Close() error {
	if w.client == nil {
		return nil
	}
	w.client.Disconnect()
	return nil
}

func (w *whatsapp) Start() error {
	// process messages
	go func() {
		for message := range w.events {
			w.processMessage(message)
		}
	}()

	container, err := sqlstore.New("sqlite3", "file:whatsapp.db?_foreign_keys=on", waLog.Noop)
	if err != nil {
		return err
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	w.deviceStore, err = container.GetFirstDevice()
	if err != nil {
		return err
	}
	client := whatsmeow.NewClient(w.deviceStore, waLog.Noop)
	client.AddEventHandler(w.eventHandler)
	w.client = client

	if w.client.Store.ID != nil {
		return w.client.Connect()
	}
	return nil
}

func (w *whatsapp) processMessage(event *events.Message) {
	fmt.Println("Received a message!", event.Message.GetConversation())

	var socialMessage = newSocialMessage(event, w.client)

	for _, handler := range w.socialHandlers {
		if handler.IsValid(socialMessage) {
			handler.Message(socialMessage)
		}
	}
}

func (w *whatsapp) eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		var message = v.Message.GetConversation()
		if message == "" {
			return
		}
		w.events <- v
	}
}
