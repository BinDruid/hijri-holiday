package structs

import (
	"fmt"
	"strings"
)

type Holiday struct {
	Month string `json:"month"`
	Day   string `json:"day"`
}

func (h *Holiday) ConvertMonth() error {
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
		return fmt.Errorf("invalid Persian month name")
	}
	h.Month = month
	return nil
}

func (h *Holiday) ConvertDay() {
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
	var day strings.Builder
	for _, char := range h.Day {
		day.WriteString(persianToEnglishDigits[string(char)])
	}
	h.Day = day.String()
}
