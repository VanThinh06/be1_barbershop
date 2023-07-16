package api

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/gin-gonic/gin"

	"google.golang.org/api/option"
)

func sdkFirebaseAdmin(ctx *gin.Context) {
	opt := option.WithCredentialsFile("./barbershop-75055-firebase-adminsdk-dvogj-080ebc30c9.json")
	config := &firebase.Config{ProjectID: "barbershop-75055"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Obtain a messaging.Client from the App.
	// ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := "d6P6zpFJTyKFwysgFbxzF7:APA91bHZFRwKl_R2PfVTeUzp8GHaiIZEeGImduWrSa41shhFeCtFJA-VUUQslukdYddRqO19sOC9DboSsKuJnnyd-QLyJ1cqGiYMPwKBh-xs2oJQUA-MAliTMGXddaaBjux4LAsFYAPr"

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Notification: &messaging.Notification{
			Title: "Breaking News",
			Body:  "New news story available.",
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)

}
