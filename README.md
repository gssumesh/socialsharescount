# SocialSharesCount

Written in GO lang, this project is a wrapper API on popular social sites and returns share count of a given url. **Star it if you Like it**

## Installation
```
import "github.com/gssumesh/socialsharescount"

```

## Getting Started

To get started just get the GO package.

```go
package main

import ( 
  "github.com/gssumesh/socialsharescount"
  "fmt"
)
func main() {
  url := "http://www.google.com"
  socialCounts := make(map[string]int) // To Store Result
  
  socialCounts = socialsharescount.GetAll(url)
  fmt.Println(socialCounts)
}

```
Result:
```
map[FacebookShare:10780840 Stumbleupon:255527 Mail_ru:609 Vkontakte:373 LinkedIn:22805 Pintrest:75108 FacebookLike:21689862 Buffer:9825 Odnoklassniki:531 FacebookComment:2450901 FacebookTotal:34921603]

```

## Supported Networks
	*Facebook Share*
	*Facebook Like*
	*Facebook Comment*
	*Facebook Total*
	*LinkedIn*
	*Pinterest*
	*Stumbleupon*
	*Buffer*
	*Mail.ru*
	*Vkontakte*
	*Odnoklassniki*

## To Be Done
	*Google Plus*
	*Reddit*


## Built With

* GOlang
* Eclipse

## Authors

* **Sumesh Suvarna** - *gs.sumesh@gmail.com* - [gssumesh](https://github.com/gssumesh)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Inspired by Ruby Gem :Social Shares: https://github.com/Timrael/social_shares

## Links

* [![GoDoc](https://godoc.org/github.com/gssumesh/socialsharescount?status.svg)](https://godoc.org/github.com/gssumesh/socialsharescount)
