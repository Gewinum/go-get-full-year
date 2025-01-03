package main

import (
	"fmt"
	"github.com/Gewinum/go-get-full-year/fullyear"
)

func main() {
	cli := fullyear.NewGetFullYearClient()
	resp, err := cli.GetFullYear()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.YearString)
}
