# Standard Time Package (std/time)

This package provides unit parsing functionality for time, compatible with Go's standard `time.Duration`.

## Usage

```go
package main

import (
    "fmt"
    stdtime "github.com/armourstill/str2quantity/std/time"
)

func main() {
    // Compatible with standard library format
    d, _ := stdtime.ParseDuration("1h30m")
    fmt.Println(d) // 1h30m0s

    // Additive format support
    // "1h30m" -> 1 hour + 30 minutes

    // Precision check: minimum granularity is 1ns
    _, err := stdtime.ParseDuration("0.5ns")
    if err != nil {
        fmt.Println("Error:", err) // Error: precision loss ...
    }
}
```

## Units

The base unit is **Nanosecond (ns)** (scale = 1.0).

*   **SI Units**: `ns`, `us`/`µs`, `ms`, `s`
*   **Common Units**: `m` (minute), `h` (hour), `d` (day), `w` (week)
