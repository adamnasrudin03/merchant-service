package utils

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(bytes)
}

func DateBetween(start, end time.Time) ([]string, []time.Time) {
	resString := []string{}
	resTime := []time.Time{}
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		resTime = append(resTime, d)
		resString = append(resString, d.Format("2006-01-02"))
	}

	return resString, resTime
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
