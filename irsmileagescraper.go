package irsmileagescraper

import (
	"fmt"
	"strconv"

	"github.com/anaskhan96/soup"
)

type irsResponse struct {
	BusinessRate  string `json:"businessRate"`
	CharityRate   string `json:"charityRate"`
	MedicalMoving string `json:"medicalMoving"`
}

func validateIRSResponse(br string, cr string, mm string) error {
	_, err := strconv.ParseFloat(br, 32)
	if err != nil {
		return err
	}

	_, err = strconv.ParseFloat(cr, 32)
	if err != nil {
		return err
	}

	_, err = strconv.ParseFloat(mm, 32)
	if err != nil {
		return err
	}

	return nil
}

func GrabMileageByYear(year int) (irsResponse, error) {
	var ir irsResponse

	yearString := strconv.Itoa(year)
	if len(yearString) != 4 {
		return ir, fmt.Errorf("Year passed is not 4 in length")
	}

	resp, err := soup.Get("https://www.irs.gov/tax-professionals/standard-mileage-rates")
	if err != nil {
		return ir, fmt.Errorf("Error returned during initial connection of scrape attempt %v", err)
	}

	doc := soup.HTMLParse(resp)

	mileageRatesStrong := doc.Find("table").FindAll("strong")
	mileageRatesBold := doc.Find("table").FindAll("b")
	mileageRates := doc.Find("table").FindAll("td")
	indexCount := 1

	//We have to do this in two for loops because the IRS switched Table tags around 2016. They went from Strong to Bold. Part of the table is Strong, part is in bold

	for _, v := range mileageRatesStrong {
		if v.Text() == yearString {

			err := validateIRSResponse(mileageRates[indexCount].Text(), mileageRates[indexCount+1].Text(), mileageRates[indexCount+2].Text())
			if err != nil {
				return ir, err
			}

			ir.BusinessRate = mileageRates[indexCount].Text()
			ir.CharityRate = mileageRates[indexCount+1].Text()
			ir.MedicalMoving = mileageRates[indexCount+2].Text()
			return ir, nil
		}

		indexCount = indexCount + 5
	}

	for _, v := range mileageRatesBold {
		if v.Text() == yearString {

			err := validateIRSResponse(mileageRates[indexCount].Text(), mileageRates[indexCount+1].Text(), mileageRates[indexCount+2].Text())
			if err != nil {
				return ir, err
			}

			ir.BusinessRate = mileageRates[indexCount].Text()
			ir.CharityRate = mileageRates[indexCount+1].Text()
			ir.MedicalMoving = mileageRates[indexCount+2].Text()
			return ir, nil
		}

		indexCount = indexCount + 5
	}

	return ir, fmt.Errorf("Unable to find year in IRS table")
}
