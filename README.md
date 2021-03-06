# IRS Scraper

### About Cabarrus County
---
Cabarrus is an ever-growing county in the southcentral area of North Carolina. Cabarrus is part of the Charlotte/Concord/Gastonia NC-SC Metropolitan Statistical Area and has a population of about 210,000. Cabarrus is known for its rich stock car racing history and is home to Reed Gold Mine, the site of the first documented commercial gold find in the United States.

### About our team
---
The Business & Location Innovative Services (BLIS) team for Cabarrus County consists of five members:

+ Joseph Battinelli - Team Supervisor
+ Mark McIntyre - Software Developer
+ Landon Patterson - Software Developer
+ Brittany Yoder - Software Developer
+ Marci Jones - Software Developer
+ Jared Poe - GIS Analyst

Our team is responsible for software development and support for the [County](https://www.cabarruscounty.us/departments/information-technology). We work under the direction of the Chief Information Officer.

### About
---
At Cabarrus County we use the IRS mileage for travel rates and reimbursement. The problem we kept encountering is that the IRS has no API for this data and updates the rate once a year. In the past we've had to manually enter this data somewhere at a central location. This Go package scrapes the data located [here](https://www.irs.gov/tax-professionals/standard-mileage-rates). and returns the information as a struct. This package is minimal in nature but has helped us out in our organization because we are able to build a real time API on top of it.

### How it Works
---
This package uses Soup to scrape the IRS table and finds the nearest indexes for the nearest year the end user passes. For more information on the Soup package, please check out it's Github Repo, located [here](https://github.com/anaskhan96/soup)

### Getting started
---
```go get -u github.com/CabarrusCo/irs```

## Full example
---
```
package main

import (
	"encoding/json"
	"fmt"

	"github.com/CabarrusCo/irs"
)

func main() {
	mileageData, err := irs.GrabStandardMileageRatesByYear(2020)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n\n", mileageData)

	jsonMileage, err := json.Marshal(mileageData)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonMileage))
}
```

### API Example
---
An example of how we use this as a real time, queryable API is coming soon!
