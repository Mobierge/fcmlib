package fcmlib_test

import (
	"fmt"

	fcm "github.com/mobierge/fcmlib"
)

func ExampleClient_Send() {
	client := fcm.NewClient(fcm.Config{
		APIKey:     "your-fcm-server-key",
		MaxRetries: 4,
	})

	message := &fcm.Message{
		RegistrationIDs: []string{"registrationID1", "registrationID2"},
		Notification: &fcm.Notification{
			Title: "Example FCM message",
			Body:  "Hello world",
		},
		Data: map[string]string{
			"customKey": "custom value",
		},
	}

	response, err := client.Send(message)
	if err != nil {
		fmt.Printf("Error: %#v\n", err)
		return
	}

	fmt.Printf("Success: %#v\n", response)
}
