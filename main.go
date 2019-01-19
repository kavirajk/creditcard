package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kavirajk/creditcard/pkg/creditcard"
)

func main() {
	cards := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cards = append(cards, scanner.Text())
	}

	outChan := make(chan out)
	defer close(outChan)

	// Scatter
	for i := 0; i < len(cards); i++ {
		go func(card string) {
			var o out
			cc, err := creditcard.Parse(card)
			if err != nil {
				o.err = err
			}
			o.c = cc
			o.num = creditcard.Normalize(card)
			outChan <- o
		}(cards[i])
	}

	// Gatter
	for i := 0; i < len(cards); i++ {
		fmt.Println(format(<-outChan))
	}
}

type out struct {
	c   *creditcard.CreditCard
	err error
	num string
}

func format(o out) string {
	if o.err != nil {
		return fmt.Sprintf("%-25s %s", creditcard.Pretty(o.num), o.err.Error())
	}
	return fmt.Sprintf("%-25s %s", creditcard.Pretty(o.c.Number()), o.c.Type())
}
