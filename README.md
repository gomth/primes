# Golang prime numbers generator

*  *generate.go* - **Prime number Generator**

[![CircleCI](https://circleci.com/gh/gomth/primes/tree/main.svg?style=svg)](https://circleci.com/gh/gomth/primes/tree/main)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/687e7af393b64826814c41a77398f52d)](https://www.codacy.com/gh/gomth/primes/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=gomth/primes&amp;utm_campaign=Badge_Grade)
[![codebeat badge](https://codebeat.co/badges/0a565ed6-3bcd-47d2-988d-4d1dcf3ce05f)](https://codebeat.co/projects/github-com-gomth-primes-main)
[![Coverage Status](https://coveralls.io/repos/github/gomth/primes/badge.svg?branch=main)](https://coveralls.io/github/gomth/primes?branch=main)

##### Usage:

```golang
import "github.com/gomth/primes"

primeNumers, err := primes.Generate(1234)
if err != nil {
    return nil, err
}
```

##### Package contents:
| API | Engine | time complexity | space complexity | Details |
| ------ | ------ | ------ | ------ | ------ |
| Generate | [Sieve of Eratosthenes mod.][erat] | O(\frac{n}{2}log(log\;n)) | O(\frac{n}{2}) | mod. by skipping all even numbers |
| GenerateEratosthenes | [Sieve of Eratosthenes mod.][erat] | O(\frac{n}{2}log(log\;n)) | O(\frac{n}{2}) | mod. by skipping all even numbers |
| GenerateAtkin | [Sieve of Atkin][atkin] | O(\frac{n^{\frac{1}{2}}}{log\;n}) | O(n) | - |

##### TODO:
- automate displaying benchmarks
- use cpu and memory profile output
- fix latex in markdown

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [erat]: <https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes>
   [atkin]: <https://en.wikipedia.org/wiki/Sieve_of_Atkin>