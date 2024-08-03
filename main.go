package main

import (
	"encoding/json"
	"flag"
	"github.com/BinDruid/hijri-holiday/structs"
	"github.com/gocolly/colly/v2"
	"github.com/pterm/pterm"
	"os"
	"strings"
	"time"
)

const baseURL = "https://www.time.ir/fa/eventyear-%d8%aa%d9%82%d9%88%db%8c%d9%85-%d8%b3%d8%a7%d9%84%db%8c%d8%a7%d9%86%d9%87"
const todayPath = "//div[contains(@class, 'today-shamsi')]//span[contains(@id, 'ShamsiNumeral')]"
const holidayPath = "//li[contains(@class, 'eventHoliday')]//span[contains(@id, 'EventYearCalendar')]"

func main() {
	spinner, _ := pterm.DefaultSpinner.Start("Crawling time.ir")
	jsonPath := flag.String("o", "holidays.json", "Path to the output JSON file")
	flag.Parse()
	c := colly.NewCollector()
	var holidays []structs.Holiday
	var year string

	c.OnXML(todayPath, func(e *colly.XMLElement) {
		today := strings.TrimSpace(e.Text)
		parts := strings.Split(today, "/")
		year = parts[0]
	})

	c.OnXML(holidayPath, func(e *colly.XMLElement) {
		persianDate := strings.TrimSpace(e.Text)
		parts := strings.Split(persianDate, " ")
		if len(parts) != 2 {
			pterm.Fatal.WithFatal(true).Printf("invalid Persian date format")
		}

		holiday := structs.Holiday{Month: parts[1], Day: parts[0]}
		holiday.Convert()
		holidays = append(holidays, holiday)
	})

	c.OnError(func(_ *colly.Response, err error) {
		pterm.Fatal.WithFatal(true).Printf("error visiting the time.ir: %s", err)
	})

	c.OnScraped(func(r *colly.Response) {
		spinner.Success()
	})
	err := c.Visit(baseURL)
	if err != nil {
		pterm.Fatal.WithFatal(true).Printf("error visiting the time.ir: %s", err)
	}
	result := structs.ScrapResult{CrawlTime: time.Now(), Year: year, Holidays: holidays}
	result.ConvertYear()
	err = saveToJSON(result, jsonPath)
	if err != nil {
		pterm.Fatal.WithFatal(true).Printf("Error saving to JSON file: %s", err)
	}
	pterm.Info.Println("Saved holidays to JSON file: ", *jsonPath)
}

func saveToJSON(result structs.ScrapResult, filename *string) error {
	file, err := os.Create(*filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}
