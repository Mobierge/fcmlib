# fcmlib


Golang Google Cloud Messaging(FCM) library.


## Installation
```
$ go get github.com/mobierge/fcmlib
```

Then

```go
import "github.com/mobierge/fcmlib"
```


## Example Usage

```go
package main

import (
	"fmt"

	fcm "github.com/mobierge/fcmlib"
)

func main() {
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

```


## Tests

`fcmlib` provides both unit and integration tests.

To run unit tests:

```
$ go test
```

To run integration tests you need to pass your android application's `FCM` **server key** and a **registration id** for this application.

```
$ go test -v -tags=integration -key=$FCM_KEY -regid=$FCM_REGID
```

By default, push messages in integration tests will only be sent to the google servers in `dry_run` mode. If you actually want to deliver messages to the device, set ```-dry=false```

```
$ go test -v -tags=integration -key=$FCM_KEY -regid=$FCM_REGID -dry=false
```



## License

MIT. See [LICENSE](./LICENSE).
