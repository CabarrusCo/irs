package irs

import (
	"fmt"
	"strconv"

	"github.com/anaskhan96/soup"
)

type standardMileageRates struct {
	Year          int     `json:"year"`
	BusinessRate  float32 `json:"businessRate"`
	CharityRate   float32 `json:"charityRate"`
	MedicalMoving float32 `json:"medicalMoving"`
}

func validateAndSetMileageRatesByYear(year int, br string, cr string, mm string) (standardMileageRates, error) {
	var smr standardMileageRates

	brFloat, err := strconv.ParseFloat(br, 32)
	if err != nil {
		return smr, err
	}

	crFloat, err := strconv.ParseFloat(cr, 32)
	if err != nil {
		return smr, err
	}

	mmFloat, err := strconv.ParseFloat(mm, 32)
	if err != nil {
		return smr, err
	}

	smr.Year = year
	smr.BusinessRate = float32(brFloat)
	smr.CharityRate = float32(crFloat)
	smr.MedicalMoving = float32(mmFloat)

	return smr, nil
}

func GrabStandardMileageRatesByYear(year int) (standardMileageRates, error) {
	var smr standardMileageRates

	yearString := strconv.Itoa(year)
	if len(yearString) != 4 {
		return smr, fmt.Errorf("Year passed is not 4 in length")
	}

	resp, err := soup.Get("https://www.irs.gov/tax-professionals/standard-mileage-rates")
	if err != nil {
		return smr, fmt.Errorf("Error returned during initial connection of scrape attempt %v", err)
	}

	doc := soup.HTMLParse(resp)

	mileageRatesStrong := doc.Find("table").FindAll("strong")
	mileageRatesBold := doc.Find("table").FindAll("b")
	mileageRates := doc.Find("table").FindAll("td")
	indexCount := 1

	//We have to do this in two for loops because the IRS switched Table tags around 2016. They went from Strong to Bold. Part of the table is Strong, part is in bold

	for _, v := range mileageRatesStrong {
		if v.Text() == yearString {
			smr, err = validateAndSetMileageRatesByYear(year, mileageRates[indexCount].Text(), mileageRates[indexCount+1].Text(), mileageRates[indexCount+2].Text())
			if err != nil {
				return smr, err
			}
			return smr, nil
		}

		indexCount = indexCount + 5
	}

	for _, v := range mileageRatesBold {
		if v.Text() == yearString {
			smr, err = validateAndSetMileageRatesByYear(year, mileageRates[indexCount].Text(), mileageRates[indexCount+1].Text(), mileageRates[indexCount+2].Text())
			if err != nil {
				return smr, err
			}
			return smr, nil
		}

		indexCount = indexCount + 5
	}

	return smr, fmt.Errorf("Unable to find year in IRS table")
}
