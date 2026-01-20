# Go Quantity Parser (str2quantity)

This is a general-purpose dimension/unit parsing library designed based on Dimensional Analysis.

It supports custom unit systems and introduces **Generic Parsing** and **Precision Control**, helping to avoid precision loss when handling nanosecond-level time or bit-level storage.

## Core Features

*   **Generic Architecture (`Parse[N]`)**: Supports parsing into any numeric type (`int64`, `float64`, `uint`, `time.Duration`, etc.).
*   **Precision Control**:
    *   When the target is an integer type (e.g., `int64`), the library checks for precision loss due to unit conversion (e.g., inputting `0.5ns` or `0.5bit` will return an error).
    *   Built-in tolerance of `1e-12` to balance floating-point calculation noise and numerical checks.
*   **Physical Base Design**:
    *   **Time**: Uses `ns` as the integer base (1.0).
    *   **Storage**: Uses `bit` as the integer base (1.0).
    *   **Length**: Uses `m` (meter) as the float base (1.0).
*   **Flexible Unit System (`unit.System`)**:
    *   **Multi-part Accumulation**: Supports formats like `1h30m`.
    *   **Prefix Binding**: Supports SI/IEC prefixes (kB, KiB) and context-sensitive parsing (e.g., `k=1024` in storage vs 1000).
    *   **Priority Matching**: Resolves unit conflicts.
*   **Safety**: Built-in Dimensional Checking to prevent illegal operations like `1h + 1kg`.

## Standard Packages

The library comes with pre-defined unit systems for common use cases:

### 1. [Time (std/time)](std/time/README.md)
*   **Basic Usage**: `stdtime.ParseDuration("1h30m")`

### 2. [Storage (std/storage)](std/storage/README.md)
*   **Basic Usage**: `stdstorage.ParseBytes("1.5GB")`

### 3. [Length (std/length)](std/length/README.md)
*   **Basic Usage**: `length.ParseLength("1km 500m")`

## Advanced Usage: Custom Unit System

Use generic capabilities to build your own system.

For example: build an SI-compliant (1KB=1000B) system using `std/storage`'s Clone feature.

```go
package main

import (
    "fmt"
    "github.com/armourstill/str2quantity/parser"
    "github.com/armourstill/str2quantity/std/storage"
)

func main() {
    // 1. Clone Standard System (Clone)
    // std/storage defaults to JEDEC/Binary standard (1KB = 1024 Bytes)
    // We clone it to modify it into an "SI System" (SI Standard)
    siSys := storage.System.Clone()

    // 2. Overwrite Prefix Definitions (OverwritePrefix)
    // Change K, M, G etc. to 1000-based
    siPrefixes := []struct {
        sym string
        val float64
    }{
        {"k", 1e3}, {"K", 1e3},
        {"m", 1e6}, {"M", 1e6},
        {"g", 1e9}, {"G", 1e9},
        {"t", 1e12}, {"T", 1e12},
    }

    for _, p := range siPrefixes {
        // OverwritePrefix can directly modify the scale of existing prefixes
        // Note: Must modify all variants (e.g., both k and K)
        if err := siSys.OverwritePrefix(p.sym, p.val); err != nil {
            panic(err)
        }
    }

    // 3. Verification
    // Custom: 1KB = 1000 Bytes = 8000 bits
    val, _, _ := parser.Parse[float64]("1KB", siSys)
    fmt.Printf("SI System 1KB = %.0f bits (Expect 8000)\n", val)
    // Original: 1KB = 1024 Bytes = 8192 bits
    valStd, _, _ := parser.Parse[float64]("1KB", storage.System)
    fmt.Printf("Standard System 1KB = %.0f bits (Expect 8192)\n", valStd)
}
```

## Installation

```bash
go get github.com/armourstill/str2quantity
```

## Technical Details: Precision and Trade-offs

### Float64 vs Int64
This library allows developers to choose the underlying numeric type based on the scenario:

*   **Float64 (Default Recommended)**: Suitable for most human-readable configurations (e.g., config files). Has a large numeric range but is limited by floating-point precision (approx. 15 significant digits).
*   **Int64**: Suitable for scenarios requiring integer precision (e.g., billing, hardware counting). By setting the base unit (e.g., `bit`, `ns`) to 1.0, combined with the library's validation logic, it helps avoid implicit fractional truncation.

### Floating Point Noise Elimination
During parsing, the library internally uses a tolerance of `1e-12` to automatically handle tiny noise from floating-point operations (e.g., `29.999999...`), ensuring that integer unit conversions (e.g., `1m = 60s`) yield correct integer results when using generic int parsing.

## Roadmap

1. Standardized implementation of other international base units.
2. Support for dimensional arithmetic, e.g., `N=kg·m/s²`.
