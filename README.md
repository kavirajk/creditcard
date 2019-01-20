# creditcard
A simple offline creditcard validator written in Go.
It uses [Luhn Algorithm](https://en.wikipedia.org/wiki/Luhn_algorithm) to check the creditcard number.

It contains both CLI and library.

# Install
```bash
go get -u -v github.com/kavirajk/creditcard
```

# Usage
## As CLI - Takes input from Stdin
```bash
bash-4.3$ cat cards.txt
4111-1111-1111-1111
5500-0000-0000-0004
3400-0000-0000-009
3000-0000-0000-04
3011-0008-0000-04

bash-4.3$ cat cards.txt | creditcard
3011-0008-0000-04         <INVALID>
3400-0000-0000-009        American-Express
3000-0000-0000-04         DinersClub-International
5500-0000-0000-0004       MasterCard
4111-1111-1111-1111       Visa
```

## As go package.
```go
package main

import (
	"fmt"

	"github.com/kavirajk/creditcard/pkg/creditcard"
)

func main() {
	card, err := creditcard.Parse(creditcard.Normalize("4111111111111111")) // Normalize is optional
	if err != nil {
		panic(err) // Invalid
	}
	fmt.Println("Detail:", creditcard.Pretty(card.Number()), card.Type())
}
```

## LICENSE
MIT License

Copyright (c) 2019 Kaviraj
