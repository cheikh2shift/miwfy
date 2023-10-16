package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

func main() {

	p := message.NewPrinter(language.Hindi)

	fmt.Println(
		p.Sprintf("Total %v", number.Decimal(1_00_000)),
	)

}
