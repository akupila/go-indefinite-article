# indefinite

Package indefinite attempts to resolve the indefinite article `a` or `an` for a
given english word. The grammar rules are not strictly defined and the result
is best-effort based on common exceptions.

```go
article := indefinite.Article("unicorn")
fmt.Println(article + " unicorn") // a unicorn
```

---

The implementation is based on the [PERL implementation][1] by Damian Conway
and [PHP port][2] by Niko Salminen.

[1]: https://metacpan.org/pod/Lingua::EN::Inflect#Selecting-indefinite-articles
[2]: https://github.com/Kaivosukeltaja/php-indefinite-article
