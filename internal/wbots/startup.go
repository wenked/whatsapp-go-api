package internal

import (
	"fmt"
	"whatsapp-go-api/pkg/wbots"
)



func StartUp() {

	devices,err := wbots.Store.GetAllDevices()

	if err != nil {
		fmt.Println("Error getting devices from database", err)
		panic(err)
	}

	for _,device := range devices {

		jid := device.ID.User
		

		err := wbots.InitSession(device,jid)

		if err != nil {
			fmt.Println("Error connecting to WhatsApp:" ,jid, err)
			continue
		}

	}
}