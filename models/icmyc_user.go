package models

import (
	"time"
)

type IcmycUser struct {
	ID uint
	RoleId uint32 `sql:"DEFAULT:1;index"`
	SwachhManchUserId uint `sql:"index;DEFAULT:0"`
	FullName string `gorm:"size:191"`
	Email string `sql:"type:VARCHAR(255);DEFAULT:NULL"`
	Password string `sql:"type:VARCHAR(80);DEFAULT:NULL"`
	MobileNumber string `sql:"type:varchar(80);index"`
	Description string `sql:"type:text;DEFAULT NULL"`
	CityId uint32
	CityTitle string `gorm:"size:191;DEFAULT:NULL"`
	WardId uint32
	WardTitle string `gorm:"size:191;DEFAULT:NULL"`
	Location string `gorm:"size:191;DEFAULT:NULL"`
	Address string `gorm:"size:191;DEFAULT:NULL"`
	PostalCode string `gorm:"size:25;DEFAULT:NULL"`
	Birthday *time.Time `gorm:"type:date;DEFAULT:NULL"`
	Gender string `gorm:"size:1"`
	Latitude string `gorm:"size:30;DEFAULT:NULL"`
	Longitude string `gorm:"size:30;DEFAULT:NULL"`
	Otp string `sql:"type:VARCHAR(10);DEFAULT:NULL"`
	OtpSentAt *time.Time
	OtpVerified bool `gorm:"DEFAULT:0"`
	OtpverifiedAt *time.Time
	OtpVerifiedPageUri string `gorm:"size:1000;DEFAULT:NULL"`
	ActivationCode string `gorm:"size:1000;DEFAULT:NULL"`
	ActivationCodeSentAt *time.Time
	EmailVerified bool `gorm:"DEFAULT:0"`
	EmailVerifiedAt *time.Time
	WelcomePageVisitedAt *time.Time
	NearComplaintsPageVisitedAt *time.Time
	InviteFriendsPageVisitedAt *time.Time
	DashboardPageVisitedAt *time.Time
	PrefferedAnnoyingCategoryId uint32
	LastLoginAt *time.Time
	LastUserId uint `grom:"DEFAULT:NULL"`
	DashboardBannerStatus uint16
	NeighbourhoodsBannerStatus uint16
	PagesBannerStatus uint16
	SignUpIp string `gorm:"size:40;DEFAULT:NULL"`
	SignUpUserAgent string `gorm:"size:255;DEFAULT:NULL"`
	SignUpChannelId uint32
	SignInIp string `gorm:"size:40;DEFAULT:NULL"`
	SignInUserAgent string `gorm:"size:255;DEFAULT:NULL"`
	SignInChannelId uint32
	DeviceToken string `gorm:"size:255;DEFAULT:NULL"`
	MacAddress string `sql:"type:varchar(255);DEFAULT:NULL"`
	ReffererUserId uint `gorm:"DEFAULT:NULL"`
	SessionCount uint32 `gorm:"DEFAULT:0"`
	AgeForPb string `gorm:"size:50"`
	SignedUpWith string `gorm:"size:50"`
	Status bool `sql:"DEFAULT:1"`
	Activated bool `sql:"DEFAULT:0"`
	Avatar string `gorm:"size:255;DEFAULT:NULL"`
	AvatarId uint64 `gorm:"DEFAULT:NULL"`
	Slug string `gorm:"size:255;DEFAULT:NULL"`
	SocialUser bool `gorm:"DEFAULT:0"`
	FailedIoginCount uint8 `gorm:"DEFAULT:0"`
	RememberToken string `gorm:"type:varchar(100);DEFAULT:NULL"`
	LoggedIn bool `sql:"DEFAULT:0"`
	LoggedOutAt *time.Time
	PostedComplaintCount uint `sql:"DEFAULT:0"`
	VoteupCount uint `sql:"DEFAULT:0"`
	SharedComplaintCount uint `sql:"DEFAULT:0"`
	FollowingCount uint `sql:"DEFAULT:0"`
	FollowerCount uint `sql:"DEFAULT:0"`
	EmailNotifyEnabled bool `sql:"DEFAULT:1"`
	PushNotificationsEnabled bool `sql:"DEFAULT:1"`
	LastActivityAt *time.Time
	BounceStatus uint32
	CreatedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
	DeletedAt *time.Time `gorm:"type:timestamp"`	
  }

func (u *IcmycUser) TableName() string {
	return "icmyc_users"
}