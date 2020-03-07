![go functional](https://github.com/MinhPhu0304/go-functional/blob/master/gogo.png?raw=true "go functional header img")


# Go-Functional
[![Build Status](https://travis-ci.com/MinhPhu0304/go-functional.svg?branch=master)](https://travis-ci.com/MinhPhu0304/go-functional)   ![GitHub](https://img.shields.io/github/license/MinhPhu0304/go-functional)  ![GitHub top language](https://img.shields.io/github/languages/top/minhphu0304/go-functional)

Functional library for Golang programmers. Inspired by [Ramda JS](https://ramdajs.com) and [Lodash](https://lodash.com)

## Available functions
- ## ```AddInt```: Add all integer in parameters together
Example: 
  ```go
    import (
      gofunc "https://github.com/MinhPhu0304/go-functional"
    )
    input := 1
    gofunc.AddInt(input) // Expected: 1
    gofunc.AddInt(1,2,3) // Expected: 6
  ```
- ## ```AddString```:  Add all integer in parameters together with a space between each  parameters
Example: 
```go
    import (
      gofunc "https://github.com/MinhPhu0304/go-functional"
    )
    input := "Hello"
    gofunc.AddString(input) // Expected: Hello
    gofunc.AddString(input, "world") // Expected: Hello world
  ```
- ## ```First```:  Get the first element of any array type
Example: 
```go
    import (
      gofunc "https://github.com/MinhPhu0304/go-functional"
    )
    input := [3]int {1,2,3}
    gofunc.First(input) // Expected: 1
    stringInput := [3]string {"Hello", "Go", "developers"}
    gofunc.First(stringInput) // Expected: Hello
  ```
- ## ```Last```:  Get the last element of any array type
Example: 
```go
    import (
      gofunc "https://github.com/MinhPhu0304/go-functional"
    )
    input := [3]int {1,2,3}
    gofunc.Last(input) // Expected: 3
    stringInput := [3]string {"Hello", "Go", "developers"}
    gofunc.Last(stringInput) // Expected: developers
  ```
