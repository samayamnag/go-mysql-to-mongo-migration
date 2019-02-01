package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/samayamnag/icmyc-migration/models"
	"github.com/spf13/viper"
)

type IcmycUsers []models.IcmycUser

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	//fmt.Println(viper.GetString(`app.env`))
	//fmt.Println(viper.GetString("app.env"))
	//fmt.Println(viper.GetBool("app.debug"))
}

func fetchIcmycUsers(dbSession *gorm.DB, fromId, toId int) *IcmycUsers {
	dbSession.Where("id >= ? AND id <= ? AND swachh_manch_user_id = ?", fromId, toId, 0).Find(&IcmycUsers{})
	return &IcmycUsers{}
}

func findOrCreateSmUser(swachhaDBSession *gorm.DB, icmycUser models.IcmycUser) (user models.User, newUser bool) {
	user = models.User{}
	newUser = false
	swachhaDBSession.Where("mobile_number = ? AND deleted = ?", icmycUser.MobileNumber, 0).First(&user)
	if user.ID == 0 {
		user = models.User{
			MobileNumber:           icmycUser.MobileNumber,
			Otp:                    icmycUser.Otp,
			OtpSentAt:              icmycUser.OtpSentAt,
			MobileNumberVerified:   icmycUser.OtpVerified,
			MobileNumberVerifiedAt: icmycUser.OtpverifiedAt,
			MacAddress:             icmycUser.MacAddress,
			LastLoginAt:            icmycUser.LastLoginAt,
			LastLoginIp:            icmycUser.SignInIp,
			LastLoginUserAgent:     icmycUser.SignInUserAgent,
			CreatedAt:              icmycUser.CreatedAt,
			UpdatedAt:              icmycUser.UpdatedAt,
			MigratedAt:             time.Now(),
			IcmycUserId:            icmycUser.ID,
			Birthday:               icmycUser.Birthday,
		}
		swachhaDBSession.Create(&user)
		newUser = true
		return
	}
	return
}

func mapChannel(channelId uint32) string {
	var slug string
	if channelId == 1 {
		slug = "icmyc-portal"
	} else if channelId == 2 {
		slug = "icmyc-citizen-android"
	} else if channelId == 3 {
		slug = "icmyc-citizen-ios"
	} else {
		slug = "icmyc-portal"
	}
	return slug
}

func objectId2Hex(objectId interface{}) string {
	return objectId.(primitive.ObjectID).Hex()
}

func main() {
	smDBSession, smDBErro := gorm.Open("mysql", "nagesh:Welcome123!@tcp(localhost)/swachh_manch?charset=utf8mb4&parseTime=True&loc=Local")
	icmycDBSession, icmycDBErro := gorm.Open("mysql", "nagesh:Welcome123!@tcp(localhost)/ichangemycity_v4?charset=utf8mb4&parseTime=True&loc=Local")
	// Create Mongo client
	mongoClient, smMongoDBErr := mongo.NewClient("mongodb://localhost:27017")

	if smMongoDBErr != nil {
		log.Fatal(smMongoDBErr)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer smDBSession.Close()
	defer icmycDBSession.Close()
	defer cancel()

	if smDBErro != nil {
		fmt.Println("Swachh Manch Connection Failed to Open")
	} else {
		log.Println("Connected to Swachh Manch")
	}
	if icmycDBErro != nil {
		fmt.Println("ICMYC Connection Failed to Open")
	} else {
		log.Println("Connected to ICMYC")
	}

	// Connect to Mongo server
	smMongoDBConnErr := mongoClient.Connect(ctx)
	if smMongoDBConnErr != nil {
		log.Fatal(smMongoDBConnErr)
	}
	profileCollection := mongoClient.Database("swachh_manch").Collection("profiles")
	channelCollection := mongoClient.Database("swachh_manch").Collection("channels")
	user := models.User{}
	channel := models.Channel{}

	if !smDBSession.HasTable(&user) {
		smDBSession.CreateTable(&user)
	}

	if !icmycDBSession.HasTable(&models.IcmycUser{}) {
		icmycDBSession.CreateTable(&models.IcmycUser{})
	}

	// Extarct user id from command line args
	fromIdPtr := flag.Int("fromId", 0, "From user ID")
	toIdPtr := flag.Int("toId", 0, "To user ID")
	flag.Parse()

	icmycUsers := []models.IcmycUser{}
	icmycDBSession.Where("id >= ? AND id <= ? AND swachh_manch_user_id = ?", *fromIdPtr, *toIdPtr, 0).Find(&icmycUsers)
	fmt.Println("******************")
	for _, icmycUser := range icmycUsers {
		channelErr := channelCollection.FindOne(ctx, bson.M{"slug": mapChannel(icmycUser.SignUpChannelId)}).Decode(&channel)

		if channelErr != nil {
			panic(channelErr)
		}
		// Store into MYSQl
		smUser, newUser := findOrCreateSmUser(smDBSession, icmycUser)
		if newUser {
			// Store into Mongo Profiles
			profileRes, err := profileCollection.InsertOne(ctx, bson.M{
				"user_id":            smUser.ID,
				"full_name":          icmycUser.FullName,
				"sign_up_with":       "mobile_number",
				"sign_up_ip_address": icmycUser.SignUpIp,
				"channels": bson.A{
					bson.M{
						"id":            channel.ID,
						"slug":          channel.Slug,
						"sign_up":       true,
						"device_token":  icmycUser.DeviceToken,
						"last_login_at": icmycUser.LastLoginAt,
						"settings": bson.M{
							"email_notifications_preferred": true,
							"sms_notifications_preferred":   true,
							"push_notifications_preferred":  true,
						},
					},
				},
				"mobile_number":                   icmycUser.MobileNumber,
				"mobile_number_verified":          icmycUser.OtpVerified,
				"mobile_number_verified_at":       icmycUser.OtpverifiedAt,
				"mac_address":                     icmycUser.MacAddress,
				"last_login_at":                   icmycUser.LastLoginAt,
				"last_login_ip":                   icmycUser.SignInIp,
				"last_login_user_agent":           icmycUser.SignInUserAgent,
				"created_at":                      icmycUser.CreatedAt,
				"updated_at":                      icmycUser.UpdatedAt,
				"has_at_least_one_social_account": icmycUser.SocialUser,
				"icmyc_user_id":                   icmycUser.ID,
			})
			if err != nil {
				panic(err)
			}

			fmt.Printf("Inserted Profile Id: %s SM user id: %d and ICMYC user id: %d \n", objectId2Hex(profileRes.InsertedID), smUser.ID, icmycUser.ID)
		}
	}
	fmt.Println("******************")
}

// go run main.go -fromId=25080879 -toId=25082284
