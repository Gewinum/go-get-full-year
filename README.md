# 🌟 go-get-full-year 🌟

Welcome to `go-get-full-year`! This is a Go client library designed to fetch the current year and additional information from the awesome [GetFullYear](https://getfullyear.com) API. 🚀

## 📦 Installation

To install the library, simply run:

```sh
go get github.com/Gewinum/go-get-full-year
```
## 🚀 Quick Start

Here's a quick example to get you started with `go-get-full-year`:

```go
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

```

## 📚 API Documentation

### `NewGetFullYearClient() *GetFullYearClient`

Creates a new instance of `GetFullYearClient`.

### `GetFullYear() (*FullYearClientResponse, error)`

Fetches the current year and additional information from the GetFullYear API.

#### FullYearClientResponse

- `Year` (int): The current year.
- `YearString` (string): The current year as a string.
- `SponsoredBy` (string): The sponsor of the current year.

## 🤝 Contributing

We welcome contributions! Feel free to open an issue or submit a pull request. Let's make this project even better together! 💪

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## 🌟 Acknowledgements

Special thanks to the [GetFullYear](https://getfullyear.com) team for providing such a cool API! 🎉

---

Made with ❤️ by [Gewinum](https://github.com/Gewinum)
