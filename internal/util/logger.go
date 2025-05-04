package util

import (
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:    true,
		QuoteEmptyFields: true,
	})
	return log
}

func GenerateOTP(length int) string {
    rand.Seed(time.Now().UnixNano())
    otp := ""
    for i := 0; i < length; i++ {
        otp += string('0' + rand.Intn(10))
    }
    return otp
}
