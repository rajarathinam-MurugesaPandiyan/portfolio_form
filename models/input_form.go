package models

import "time"

type FormInputs struct {
	OwnerMail   string    `json:"owner_mail"`
	Title       string    `json:"title"`
	Name        string    `json:"name"`
	Phone       string    `json:"mobile"`
	SenderEmail string    `json:"sender_email"`
	Description string    `json:"description"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
}
