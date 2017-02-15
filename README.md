# HomeFacts [![CircleCI](https://circleci.com/gh/orcaman/homefacts.svg?style=svg&circle-token=513c6145094d0cf7331aded74c53ea3a1e131b34)](https://circleci.com/gh/orcaman/homefacts)

The following is an unoffical go client for homefacts.com. HomeFacts is a website that provides stats about addresses.
This unoffical software is not affiliated with homefacts.com in any way. Visit Homefacts [here](http://www.homefacts.com/).

## Usage Example

Initialize a client, and use the `GetFacts` function, passing a string address. The response object looks like so

```json
 {
 	"crimeRate": {
 		"Class": "successIcon",
 		"Text": "Low"
 	},
 	"schoolRating": {
 		"Class": "errorIcon",
 		"Text": "B+ for 22280 S 209th Way"
 	},
 	"registeredOffenders": {
 		"Class": "warningIcon",
 		"Text": "1  within 1 mile"
 	},
 	"avgHomePrice": {
 		"Class": "warningIcon",
 		"Text": "$203,000 within 1 mile"
 	},
 	"foreclosures": {
 		"Class": "errorIcon",
 		"Text": "51  within 1 mile"
 	},
 	"naturalHazards": {
 		"Class": "successIcon",
 		"Text": "Earthquake, Hurricane "
 	},
 	"environmentalHazards": {
 		"Class": "",
 		"Text": "9  within 1 mile"
 	}
 }
```

```go
package main

import (
	"flag"
	"log"

	"encoding/json"

	"github.com/orcaman/homefacts"
)

var (
	address = flag.String("address", "22280 S 209th Way, Queen Creek, AZ 85142", "address")
)

func init() {
	flag.Parse()
}

func main() {
	c := homefacts.New()

	resp, err := c.GetFacts(&homefacts.Request{
		Address: *address,
	})

	if err != nil {
		log.Fatalln(err)
	}

	j, err := json.MarshalIndent(resp.Result, " ", "\t")

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("response: %s", j)
	// prints:
	//  {
	//  	"crimeRate": {
	//  		"Class": "successIcon",
	//  		"Text": "Low"
	//  	},
	//  	"schoolRating": {
	//  		"Class": "errorIcon",
	//  		"Text": "B+ for 22280 S 209th Way"
	//  	},
	//  	"registeredOffenders": {
	//  		"Class": "warningIcon",
	//  		"Text": "1  within 1 mile"
	//  	},
	//  	"avgHomePrice": {
	//  		"Class": "warningIcon",
	//  		"Text": "$203,000 within 1 mile"
	//  	},
	//  	"foreclosures": {
	//  		"Class": "errorIcon",
	//  		"Text": "51  within 1 mile"
	//  	},
	//  	"naturalHazards": {
	//  		"Class": "successIcon",
	//  		"Text": "Earthquake, Hurricane "
	//  	},
	//  	"environmentalHazards": {
	//  		"Class": "",
	//  		"Text": "9  within 1 mile"
	//  	}
	//  }
}
```