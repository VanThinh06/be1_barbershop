package api

import (
	"context"
	"log"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)
  func sdkFirebaseAdmin()  {
	opt := option.WithCredentialsFile("./barbershop-75055-firebase-adminsdk-dvogj-080ebc30c9.json")
	config := &firebase.Config{ProjectID: "barbershop-75055"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	
	
  }
