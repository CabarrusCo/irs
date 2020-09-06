package irs

import (
	"fmt"
	"testing"
	"time"
)

func TestScraper(t *testing.T) {
	year, _, _ := time.Now().Date()

	mileageRates, err := GrabStandardMileageRatesByYear(year)
	if err != nil {
		t.Errorf("Error encountered in Scrape Test %s", err)
	}

	fmt.Printf("%+v\n\n", mileageRates)
}
