package wbots

import (
	"context"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

type MyClient struct {
	WAClient *whatsmeow.Client
	eventHandlerID uint32
}

var Clients = make(map[string]*MyClient)
var Store *sqlstore.Container


func (mycli *MyClient) register() {
	mycli.eventHandlerID = mycli.WAClient.AddEventHandler(mycli.myEventHandler)
}

// event listener



func (mycli *MyClient) myEventHandler(evt interface{}) {
	// Handle event and access mycli.WAClient
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())
		msgExtra := whatsmeow.SendRequestExtra{
			ID: mycli.WAClient.GenerateMessageID(),
		}
		msgContent := &waProto.Message{
			Conversation: proto.String("Hello from Go!"),
		}

		fmt.Println("Sending a message to" , v.Info.Sender.ToNonAD(), "with content", "Hello from Go!")
		mycli.WAClient.SendMessage(context.Background(),v.Info.Sender.ToNonAD(), msgContent, msgExtra)
	}
}


func init(){
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
	container, err := sqlstore.New("sqlite3", "file:examplestore.db?_foreign_keys=on", dbLog)
	if err != nil {
		fmt.Println("Error creating database", err)
		panic(err)
	}

	Store = container
}

func InitSession(device *store.Device, jid string) error {

	if(Clients[jid] == nil) {
		if device == nil {
			// Initialize New WhatsApp Client Device in Datastore
			device = Store.NewDevice()
		}

	}

	// Create a new session


	myClient := &MyClient{}
	myClient.WAClient = whatsmeow.NewClient(device, nil)
	myClient.register()

	myClient.WAClient.EnableAutoReconnect = true
	myClient.WAClient.AutoTrustIdentity = true

	Clients[jid] = myClient

	// Start the session

	if myClient.WAClient.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := myClient.WAClient.GetQRChannel(context.Background())
		err := myClient.WAClient.Connect()
		if err != nil {
			fmt.Println("Error connecting to WhatsApp:", err)
			return err
		}

		for evt := range qrChan {
			if evt.Event == "code" {
				// Render the QR code here
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				// or just manually `echo 2@... | qrencode -t ansiutf8` in a terminal
				fmt.Println("QR code:", evt.Code)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		err := myClient.WAClient.Connect()
		if err != nil {
			fmt.Println("Error connecting to WhatsApp:", err)
			return err
		}

	}

	return nil
}

func startSessions () {
	
	fakeDb := []string{"1234567890"}

	fmt.Println("Starting sessions", fakeDb)
}

