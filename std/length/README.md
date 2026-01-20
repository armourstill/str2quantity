# Standard Length Package (std/length)

This package provides unit parsing for length. The base unit is **Meter (m)** using `float64` for high precision and large scale support.

## Usage

```go
package main

import (
    "fmt"
    "github.com/armourstill/str2quantity/std/length"
)

func main() {
    // SI Units
    l1, _ := length.ParseLength("1.5km")
    fmt.Printf("1.5km = %.2f meters\n", l1) // 1500.00 meters

    // Multi-part string support
    l2, _ := length.ParseLength("1m 50cm")
    fmt.Printf("1m 50cm = %.2f meters\n", l2) // 1.50 meters
}
```

## Units

The base unit is **Meter (m)** (scale = 1.0).

*   **Base Unit**: `m`
*   **SI Prefixes**: `nm`, `µm`/`um`, `mm`, `cm`, `km`
