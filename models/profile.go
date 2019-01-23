package models

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type Profile struct {
	UserID			int64 `bson:"user_id,omitempty"`
	CityID			int64 `bson:"city_id,omitempty"`
}

type Channel struct {
	ID primitive.ObjectID `bson:"_id"`
	Title string `bson:"title"`
	Slug string `bson:"slug"`
	Platform string `bson:"platform"`
	AppName	string `bson:"app_name"`
	Type string `bson:"type"`
	Archived bool `bson:"archived"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}