# IRS Mileage Scraper

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

Our team is responsible for software development and support for the [County](https://www.cabarruscounty.us/departments/information-technology). We work under the direction of the Chief Information Officer.

### About
---
At Cabarrus County we use the IRS mileage for travel rates and reimbursement. The problem we kept encountering is that the IRS has no API for this data. In the past we've had to manually enter this data somewhere at a central location. Each year the IRS updates the new mileage rate. This go package scrapes the data located [here](https://www.irs.gov/tax-professionals/standard-mileage-rates). and returns the inofmration

### How it Works
---
This package uses Soup to scrape the IRS table and finds the nearest indexes for the nearest year the end user passes. For more information on the Soup package, please check out it's Github Repo, located [here](https://github.com/anaskhan96/soup)

### Getting started
---
```go get -u github.com/CabarrusCo/irsmileagescraper```
```

### Getting started
---
```go get -u github.com/CabarrusCo/irsmileagescraper```
```
