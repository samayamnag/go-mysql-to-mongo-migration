package models

import (
	"time"
)

type User struct {
	ID uint
	Username string `sql:"type:VARCHAR(50);DEFAULT:NULL"`
	Email string `sql:"type:VARCHAR(255);DEFAULT:NULL"`
	Password string `sql:"type:VARCHAR(80);DEFAULT:NULL"`
	MobileNumber string `sql:"type:varchar(80);index;unique_index:idx_sm_users_mobile_number_deleted"`
	Otp string `sql:"type:VARCHAR(80);DEFAULT:NULL"`
	OtpSentAt *time.Time
	MobileNumberVerified bool `gorm:"DEFAULT:0"`
	MobileNumberVerifiedAt *time.Time
	EmailActivationToken string `gorm:"type:VARCHAR(255);DEFAULT:NULL"`
	EmailActivationTokenSentAt *time.Time
	EmailVerified bool `gorm:"DEFAULT:0"`
	EmailVerifiedAt *time.Time
	RememberToken string `gorm:"type:varchar(100);DEFAULT:NULL"`
	MacAddress string `sql:"type:varchar(255);DEFAULT:NULL"`
	Banned bool `sql:"DEFAULT:0;index"`
	BannedAt *time.Time
	LastLoginAt *time.Time
	LastLoginIp string `gorm:"varchar(40);DEFAULT:NULL"`
	LastLoginUserAgent string
	CreatedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
	DeletedAt *time.Time `gorm:"type:timestamp"`
	Deleted bool `sql:"DEFAULT:0;index;unique_index:idx_sm_users_mobile_number_deleted"`
	NonVerifiedMobileNumber string `sql:"type:varchar(20);DEFAULT:NULL"`
	NonVerifiedMobileNumberOtp string `sql:"type:varchar(80);DEFAULT:NULL"`
	NonVerifiedMobileNumberOtpSentAt *time.Time
	OtpSource string `sql:"type:varchar(50) NOT NULL;DEFAULT:'mvaayoo'"`
	LoginToken string `sql:"type:varchar(100);DEFAULT:NULL"`
	LoginTokenSentAt *time.Time
	MigratedAt time.Time
	IcmycUserId uint `sql:"index"`
	Birthday *time.Time `gorm:"type:date;DEFAULT:NULL"`
  }

func (u *User) TableName() string {
	return "sm_users"
}