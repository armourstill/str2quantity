# Standard Storage Package (std/storage)

This package provides unit parsing for digital storage quantities. It supports `ParseBytes` (general use) and `ParseBits` (high precision).

## Usage

```go
package main

import (
    "fmt"
    stdstorage "github.com/armourstill/str2quantity/std/storage"
)

func main() {
    // Scenario A: General capacity config (supports large numbers and decimals)
    // Even "1bit" will result in 0.125 Bytes
    bytes, _ := stdstorage.ParseBytes("1.5GB")
    fmt.Printf("%.2f Bytes\n", bytes)

    // Scenario B: Network packet counting/Hardware counting (Integer bits)
    // Rejects "0.5 bit", supports up to ~1.15 EiB
    bits, err := stdstorage.ParseBits("100Mb")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%d bits\n", bits)
}
```

## Units

The base unit is **Bit (b)** (scale = 1.0).

*   **Base Units**: `b`/`bit`/`bits`, `B`/`Byte`/`Bytes` (1B = 8b)
*   **IEC Standard Prefixes** (1024-based): `Ki`, `Mi`, `Gi`, `Ti`, `Pi`, `Ei`
*   **JEDEC/Binary Prefixes** (1024-based by default in this package): `k`/`K`, `m`/`M`, `g`/`G`, `t`/`T`, `p`/`P`, `e`/`E`
