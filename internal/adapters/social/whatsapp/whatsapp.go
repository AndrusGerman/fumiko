package whatsapp

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AndrusGerman/fumiko/internal/core/ports"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	"go.uber.org/fx"
)

type whatsapp struct {
	client *whatsmeow.Client
}

func New(lc fx.Lifecycle) ports.Social {
	var social = new(whatsapp)
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
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	container, err := sqlstore.New("sqlite3", "file:whatsapp.db?_foreign_keys=on", dbLog)
	if err != nil {
		return err
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		return err
	}
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(w.eventHandler)
	w.client = client

	if w.client.Store.ID != nil {
		return w.client.Connect()
	}
	return nil
}

func (w *whatsapp) eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())
	}
}
