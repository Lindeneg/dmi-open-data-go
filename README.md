**work-in-progress**

Simple [DMI API](https://confluence.govcloud.dk/display/FDAPI/Danish+Meteorological+Institute+-+Open+Data) wrapper. Heavily inspired by [this](https://github.com/LasseRegin/dmi-open-data) Python library made by [Lasse Regin](https://github.com/LasseRegin) of [Woodsense](https://en.woodsense.dk/).

I'd love to use the wrapper in Go and thus why this module is being created. 

*Also, as I'm quite new to Go, any feedback is highly welcome.*

##### Todo

- Extend API support
- Implement tests
- Implement channels
- Docs
- Create separate CLI package

##### Usage

1) Install 

`go get -u github.com/lindeneg/dmi-open-data-go/v2`

2) Get API keys [here](https://confluence.govcloud.dk/pages/viewpage.action?pageId=26476690)

3) Get some data
```Go
// main.go
package main

import (
	"fmt"
	"os"

	"github.com/lindeneg/dmi-open-data-go/v2/client"
	"github.com/lindeneg/dmi-open-data-go/v2/constants"
)

func main() {
  // also NewClimateDataClient
  c := client.NewMetObsClient(os.Getenv("DMI_METOBS_API_KEY"))

  // example method: returns unmarshaled geo-json data ... or an error
  obs, err := c.GetObservations(client.GetObservationsConfig{Limit: 100})
	
  if err != nil {
		fmt.Println(err)
		return
	}

	
  fmt.Printf("Returned %d entries\n", obs.NumberReturned)
	
  if obs.NumberReturned > 0 {
		fmt.Printf("First entry: %v\n", obs.Features[0])
	}
}
```

4) Run it!
`go run main.go`