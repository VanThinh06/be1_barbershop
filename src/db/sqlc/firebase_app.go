package db

import (
	firebase "firebase.google.com/go/v4"
	"github.com/cloudinary/cloudinary-go/v2"
)

type ServiceApp struct {
	FirebaseApp   *firebase.App
	CloudinaryApp *cloudinary.Cloudinary
}

func NewServiceApp(firebase *firebase.App, cloudinary *cloudinary.Cloudinary) ServiceApp {
	return ServiceApp{
		FirebaseApp:   firebase,
		CloudinaryApp: cloudinary,
	}
}
