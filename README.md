# github.com/jamietanna/speakeasy-example-spec

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasyapi.dev/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/advanced-setup/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/jamietanna/speakeasy-example-spec
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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

	ctx := context.Background()
	res, err := s.Pet.UpdatePet(ctx, components.Pet{
		ID:   speakeasyexamplespec.Int64(10),
		Name: "doggie",
		PhotoUrls: []string{
			"<value>",
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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Pet](docs/sdks/pet/README.md)

* [UpdatePet](docs/sdks/pet/README.md#updatepet) - Update an existing pet
* [AddPet](docs/sdks/pet/README.md#addpet) - Add a new pet to the store
* [FindPetsByStatus](docs/sdks/pet/README.md#findpetsbystatus) - Finds Pets by status
* [FindPetsByTags](docs/sdks/pet/README.md#findpetsbytags) - Finds Pets by tags
* [GetPetByID](docs/sdks/pet/README.md#getpetbyid) - Find pet by ID
* [DeletePet](docs/sdks/pet/README.md#deletepet) - Deletes a pet
* [UploadFile](docs/sdks/pet/README.md#uploadfile) - uploads an image

### [Store](docs/sdks/store/README.md)

* [GetInventory](docs/sdks/store/README.md#getinventory) - Returns pet inventories by status
* [PlaceOrder](docs/sdks/store/README.md#placeorder) - Place an order for a pet
* [GetOrderByID](docs/sdks/store/README.md#getorderbyid) - Find purchase order by ID
* [DeleteOrder](docs/sdks/store/README.md#deleteorder) - Delete purchase order by ID

### [User](docs/sdks/user/README.md)

* [CreateUser](docs/sdks/user/README.md#createuser) - Create user
* [CreateUsersWithListInput](docs/sdks/user/README.md#createuserswithlistinput) - Creates list of users with given input array
* [LoginUser](docs/sdks/user/README.md#loginuser) - Logs user into the system
* [LogoutUser](docs/sdks/user/README.md#logoutuser) - Logs out current logged in user session
* [GetUserByName](docs/sdks/user/README.md#getuserbyname) - Get user by user name
* [UpdateUser](docs/sdks/user/README.md#updateuser) - Update user
* [DeleteUser](docs/sdks/user/README.md#deleteuser) - Delete user
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.APIErrorInvalidInput | 400                            | application/json               |
| sdkerrors.APIErrorUnauthorized | 401                            | application/json               |
| sdkerrors.APIErrorNotFound     | 404                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

### Example

```go
package main

import (
	"context"
	"errors"
	speakeasyexamplespec "github.com/jamietanna/speakeasy-example-spec"
	"github.com/jamietanna/speakeasy-example-spec/models/components"
	"github.com/jamietanna/speakeasy-example-spec/models/sdkerrors"
	"log"
)

func main() {
	s := speakeasyexamplespec.New(
		speakeasyexamplespec.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Pet.UpdatePet(ctx, components.Pet{
		ID:   speakeasyexamplespec.Int64(10),
		Name: "doggie",
		PhotoUrls: []string{
			"<value>",
		},
	})
	if err != nil {

		var e *sdkerrors.APIErrorInvalidInput
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.APIErrorUnauthorized
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.APIErrorNotFound
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://{environment}.petstore.io` | `environment` (default is `prod`) |

#### Example

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
		speakeasyexamplespec.WithServerIndex(0),
		speakeasyexamplespec.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Pet.UpdatePet(ctx, components.Pet{
		ID:   speakeasyexamplespec.Int64(10),
		Name: "doggie",
		PhotoUrls: []string{
			"<value>",
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

#### Variables

Some of the server options above contain variables. If you want to set the values of those variables, the following options are provided for doing so:
 * `WithEnvironment speakeasyexamplespec.ServerEnvironment`

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
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
		speakeasyexamplespec.WithServerURL("https://{environment}.petstore.io"),
		speakeasyexamplespec.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Pet.UpdatePet(ctx, components.Pet{
		ID:   speakeasyexamplespec.Int64(10),
		Name: "doggie",
		PhotoUrls: []string{
			"<value>",
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type     | Scheme   |
| -------- | -------- | -------- |
| `APIKey` | apiKey   | API key  |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
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

	ctx := context.Background()
	res, err := s.Pet.UpdatePet(ctx, components.Pet{
		ID:   speakeasyexamplespec.Int64(10),
		Name: "doggie",
		PhotoUrls: []string{
			"<value>",
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
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
