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
	ctx := context.Background()

	s := speakeasyexamplespec.New(
		speakeasyexamplespec.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Pet.UpdatePet(ctx, components.Pet{
		ID:   speakeasyexamplespec.Pointer[int64](10),
		Name: "doggie",
		Category: &components.Category{
			ID:   speakeasyexamplespec.Pointer[int64](1),
			Name: speakeasyexamplespec.Pointer("Dogs"),
		},
		PhotoUrls: []string{
			"<value 1>",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Pet != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->