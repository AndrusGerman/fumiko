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
	client         *whatsmeow.Client
	deviceStore    *store.Device
	socailMessages chan ports.SocialMessage
	storage        ports.Storage
	socialHandlers []ports.SocialHandler
}

func (w *whatsapp) AddHandlers(handlers ...ports.SocialHandler) {
	w.socialHandlers = handlers
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
			qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
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
	var err error

	// process messages
	go func() {
		for message := range w.socailMessages {
			w.processMessage(message)
		}
	}()

	// create database
	container := sqlstore.NewWithDB(w.storage.GetDB(), w.storage.GetDialect(), waLog.Noop)

	// get device store
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

func (w *whatsapp) processMessage(socialMessage ports.SocialMessage) {
	fmt.Println("Received a whatsapp message!", socialMessage.GetText())
	for i := range w.socialHandlers {
		var isValid = w.socialHandlers[i].IsValid(socialMessage)
		if !isValid {
			continue
		}
		w.socialHandlers[i].Message(socialMessage)
		return
	}
}

func (w *whatsapp) eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		if w.isValidWhatsapp(v) {
			return
		}
		var sm = newSocialMessage(v, w.client)

		w.socailMessages <- sm
	}
}

func (w *whatsapp) isValidWhatsapp(event *events.Message) bool {
	if event.Message.GetConversation() == "" {
		return false
	}
	return string(event.Message.GetConversation()[0]) == ","
}

func New(lc fx.Lifecycle, storage ports.Storage) ports.Social {
	var social = &whatsapp{
		socailMessages: make(chan ports.SocialMessage),
		storage:        storage,
	}

	lc.Append(fx.StopHook(social.Close))
	lc.Append(fx.StartHook(social.Start))
	return social
}
