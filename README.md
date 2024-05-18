# base26

`base26` implements a simple base26 encoding and decoding utility.

Base26 is useful for encoding numbers into a string of letters, and decoding with short strings.

For example, user can use base26 string for numeric id like Docker object hash, but more short hand way. Shorter string is human-readable and easy to type, less mistake.

## Usage

```go
package main

import (
    "fmt"
    "github.com/hitoshi44/base26"
)

func main() {

	fmt.Println(base26.EncodeUint(1234567890))
    // >> dzxprwk

	fmt.Println(base26.MustDecodeUint("dzxprwk"))
    // >> 1234567890
}
```

## TODO

- [ ] Add Bytes to Bytes enc/dec functions

## License

See [LICENSE](./LICENSE)
