package utils

import (
	"dropx/pkg/constants"
	"fmt"
	"github.com/gofrs/uuid"
	"math/rand"
	"time"
)

func StringToUUID(s string) uuid.UUID {
	return uuid.Must(uuid.FromString(s))
}

func GetId() uuid.UUID {
	id := uuid.Must(uuid.NewV4())
	return id
}

func GenerateOrderCode(orderType string) string {
	now := time.Now()

	timestamp := now.Format("20060102-150405")
	randomPart := randSeq(4)

	return fmt.Sprintf("%s-%s-%s", constants.OrderCodePrefix[orderType], timestamp, randomPart)
}

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func StringToTime(timeString string) time.Time {
	parseTime, _ := time.Parse("2006-01-02", timeString)
	return parseTime
}
