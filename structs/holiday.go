package structs

import (
	"log"
	"strings"
	"time"
)

func ConvertNumber(n string) string {
	persianToEnglishDigits := map[string]string{
		"۰": "0",
		"۱": "1",
		"۲": "2",
		"۳": "3",
		"۴": "4",
		"۵": "5",
		"۶": "6",
		"۷": "7",
		"۸": "8",
		"۹": "9",
	}
	var converted strings.Builder
	for _, char := range n {
		converted.WriteString(persianToEnglishDigits[string(char)])
	}
	return converted.String()
}

type Holiday struct {
	Month string `json:"month"`
	Day   string `json:"day"`
}

func (h *Holiday) convertMonth() {
	persianMonths := map[string]string{
		"فروردین":  "1",
		"اردیبهشت": "2",
		"خرداد":    "3",
		"تیر":      "4",
		"مرداد":    "5",
		"شهریور":   "6",
		"مهر":      "7",
		"آبان":     "8",
		"آذر":      "9",
		"دی":       "10",
		"بهمن":     "11",
		"اسفند":    "12",
	}
	month, exists := persianMonths[h.Month]
	if !exists {
		log.Fatal("invalid Persian month name")
	}
	h.Month = month
}

func (h *Holiday) Convert() {
	h.convertMonth()
	h.Day = ConvertNumber(h.Day)
}

type ScrapResult struct {
	CrawlTime time.Time `json:"crawl_time"`
	Year      string    `json:"year"`
	Holidays  []Holiday `json:"holidays"`
}

func (s *ScrapResult) ConvertYear() {
	s.Year = ConvertNumber(s.Year)
}
