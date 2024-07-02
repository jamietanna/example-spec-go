<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	speakeasyexamplespec "github.com/jamietanna/speakeasy-example-spec"
	"github.com/jamietanna/speakeasy-example-spec/models/components"
	"log"
)

func main() {
	s := speakeasyexamplespec.New(
		speakeasyexamplespec.WithSecurity("<YOUR_API_KEY_HERE>"),
	)
	request := components.Pet{
		ID:   speakeasyexamplespec.Int64(10),
		Name: "doggie",
		Category: &components.Category{
			ID:   speakeasyexamplespec.Int64(1),
			Name: speakeasyexamplespec.String("Dogs"),
		},
		PhotoUrls: []string{
			"<value>",
		},
	}
	ctx := context.Background()
	res, err := s.Pet.UpdatePet(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.Pet != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->