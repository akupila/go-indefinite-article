# go-indefinite-article

[![Build Status](https://github.com/akupila/go-indefinite-article/workflows/Go/badge.svg)][build]
[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)][godev]

Package `indefinite` resolves the indefinite article `a` or `an` for a given english word.

## Install

```
go get github.com/akupila/go-indefinite-article
```

## Example

```go
article := indefinite.Article("unicorn")
fmt.Println(article + " unicorn") // a unicorn
```
See [indefinite_test.go][test] for more.

---

The implementation is based on the [PERL implementation][perl] by Damian Conway
and [PHP port][php] by Niko Salminen.

[perl]: https://metacpan.org/pod/Lingua::EN::Inflect#Selecting-indefinite-articles
[php]: https://github.com/Kaivosukeltaja/php-indefinite-article
[test]: ./indefinite_test.go
[build]: https://github.com/akupila/go-indefinite-article/actions?query=workflow%3AGo
[godev]: https://pkg.go.dev/github.com/akupila/go-indefinite-article
