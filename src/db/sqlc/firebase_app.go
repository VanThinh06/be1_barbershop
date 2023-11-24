package db

import firebase "firebase.google.com/go/v4"

type FirebaseApp struct {
	FirebaseApp *firebase.App
}

func NewFirebase(firebase *firebase.App) FirebaseApp {
	return FirebaseApp{
		firebase,
	}
}
