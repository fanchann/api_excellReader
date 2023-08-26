package utils

import (
	"fmt"
	"time"
)

func DateLogHistory() string {
	dateNow := time.Now()
	return fmt.Sprintf("%v%v%d0%d%d", dateNow.Hour(), dateNow.Minute(), dateNow.Day(), dateNow.Month(), dateNow.Year())
}
