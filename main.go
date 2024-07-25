package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/BinDruid/hijri-holiday/structs"
	"github.com/gocolly/colly/v2"
	"os"
	"strings"
)

func main() {
	jsonPath := flag.String("o", "holiday.json", "Path to the output JSON file")
	flag.Parse()
	c := colly.NewCollector()
	var dates []structs.Holiday
	c.OnXML("//li[contains(@class, 'eventHoliday')]//span[contains(@id, 'EventYearCalendar')]", func(e *colly.XMLElement) {
		persianDate := strings.TrimSpace(e.Text)
		parts := strings.Split(persianDate, " ")
		if len(parts) != 2 {
			fmt.Println("invalid Persian date format")
			return
		}

		holiday := structs.Holiday{Month: parts[1], Day: parts[0]}
		err := holiday.ConvertMonth()
		if err != nil {
			fmt.Println("Error converting date:", err)
			return
		}
		holiday.ConvertDay()
		dates = append(dates, holiday)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	err := c.Visit("https://www.time.ir/fa/eventyear-%d8%aa%d9%82%d9%88%db%8c%d9%85-%d8%b3%d8%a7%d9%84%db%8c%d8%a7%d9%86%d9%87")
	if err != nil {
		fmt.Println("Error visiting the URL:", err)
		return
	}
	err = saveToJSON(dates, jsonPath)
	if err != nil {
		fmt.Println("Error saving to JSON file:", err)
	}
	fmt.Println("Saved holidays to JSON file", *jsonPath)
}

func saveToJSON(dates []structs.Holiday, filename *string) error {
	file, err := os.Create(*filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(dates)
}
