# Klik Resi Go SDK

Official Go client for the [Klik Resi](https://klikresi.com) API.
Track shipments, calculate shipping rates, and look up Indonesian locations
(provinces, cities, districts) across couriers such as JNE, J&T, Shopee
Express, SiCepat, TIKI, and more.

Full API reference: [docs.klikresi.com](https://docs.klikresi.com)

## Requirements

- Go 1.24+

## Installation

```sh
go get github.com/klikresi/go-sdk
```

## Quickstart

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/klikresi/go-sdk"
)

func main() {
	client := klikresi.NewClient("YOUR-API-KEY")
	ctx := context.Background()

	tracking, err := client.Tracking.Get(ctx, "YOUR-AWB", klikresi.CourierJNE)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tracking.Status)
	for _, h := range tracking.Histories {
		fmt.Println(h.Date, h.Message)
	}
}
```

## Usage

### Tracking

```go
// Basic tracking (charged per successful request).
t, err := client.Tracking.Get(ctx, "YOUR-AWB", klikresi.CourierJNE)

// ID Express requires an extra `number` (phone number) parameter.
// When provided, it is passed through as a query parameter.
t, err := client.Tracking.Get(ctx, "YOUR-AWB", klikresi.CourierIDExpress, klikresi.WithNumber("08123456789"))
```

The response contains a normalized `DeliveryStatus`
(`InfoReceived`, `InTransit`, `OutForDelivery`, `FailedAttempt`, `Delivered`,
`ReturnToSender`, `Exception`, `Expired`, `Pending`) plus origin, destination,
and the full event history.

### Rates

```go
// By district IDs, optionally filtered to specific couriers.
r, err := client.Rates.CalculateByID(ctx, "33.08.20", "32.09.31", 1, klikresi.CourierJNE)

// By location names.
r, err := client.Rates.CalculateByName(ctx, "Secang, Kabupaten Magelang, Jawa Tengah", "Depok, Kabupaten Cirebon, Jawa Barat", 1)

// By postal codes.
r, err := client.Rates.CalculateByPostalCode(ctx, 56195, 45155, 1)
```

### Location

All location endpoints are cursor-paginated. Request one page, or use the
`All*` helpers to fetch everything automatically.

```go
// Search locations by keyword.
page, err := client.Location.Search(ctx, "depok", klikresi.WithLimit(10))

// Single pages.
provinces, err := client.Location.Provinces(ctx)
cities, err    := client.Location.Cities(ctx, "33")
districts, err := client.Location.Districts(ctx, "33.08")

// Everything, following cursors automatically.
all, err := client.Location.AllLocations(ctx, "depok")
allProvinces, err := client.Location.AllProvinces(ctx)
```

### Me

```go
// Fetch the profile of the account that owns the API key.
me, err := client.Me.Get(ctx)
```

The response contains the account `ID`, `Name`, `Email`, and current `Balance`.

### Courier codes

Use the exported constants: `klikresi.CourierSPX`, `klikresi.CourierJNE`,
`klikresi.CourierJNT`, `klikresi.CourierSicepat`, `klikresi.CourierNinja`,
`klikresi.CourierPos`, `klikresi.CourierSAP`, `klikresi.CourierLEX`,
`klikresi.CourierLion`, `klikresi.CourierIDExpress`, `klikresi.CourierAnteraja`,
`klikresi.CourierWahana`, `klikresi.CourierTiki`.

## Errors

Any non-2xx response returns an `*klikresi.APIError` implementing `error`,
with the HTTP `StatusCode` and the API's `Message`:

```go
var apiErr *klikresi.APIError
if errors.As(err, &apiErr) {
	fmt.Println(apiErr.StatusCode, apiErr.Message)
}
```

## Configuration

The client constructor takes only your API key. The base URL defaults to
`https://klikresi.com`; it can be overridden with the `KLIKRESI_BASE_URL`
environment variable (useful for tests and proxies). Every request is bound
to the `context.Context` you pass and times out after 30 seconds.

## Examples

See the [`examples`](examples) directory for runnable samples
(`KLIKRESI_API_KEY=... go run ./examples/tracking`).

## License

MIT
