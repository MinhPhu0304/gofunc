![go functional](https://github.com/MinhPhu0304/go-functional/blob/master/gogo.png?raw=true "go functional header img")


# Go-Functional
[![Build Status](https://travis-ci.com/MinhPhu0304/gofunc.svg?branch=master)](https://travis-ci.com/MinhPhu0304/gofunc)  ![GitHub](https://img.shields.io/github/license/MinhPhu0304/go-functional)  ![GitHub top language](https://img.shields.io/github/languages/top/minhphu0304/go-functional)

Functional library for Golang programmers. Inspired by [Ramda JS](https://ramdajs.com) and [Lodash](https://lodash.com)

## Available functions
- ## ```AddInt```: Add all integer in parameters together
Example: 
  ```go
    import (
      gofunc "https://github.com/MinhPhu0304/gofunc"
    )
    func main() {
      input := 1
      gofunc.AddInt(input) // Expected: 1
      gofunc.AddInt(1,2,3) // Expected: 6
    }
  ```
- ## ```AddString```:  Add all integer in parameters together with a space between each  parameters
Example: 
```go
    import (
      gofunc "https://github.com/MinhPhu0304/gofunc"
    )
    func main() {
      input := "Hello"
      gofunc.AddString(input) // Expected: Hello
      gofunc.AddString(input, "world") // Expected: Hello world
    }
  ```
- ## ```All```:  Return true if all elements of the list match the predicate
Example: 
```go
    import (
      gofunc "https://github.com/MinhPhu0304/gofunc"
    )

    func main () {
      input := []int{3,3,3}
      input2 := []string{"hey","hey"}
      checkAllElementIs3 := gofunc.All(3)
      checkAllElementIs3(input) // Return true
      checkAllElementIs3(input2) // Return false
    }
  ```
- ## ```First```:  Get the first element of any array type
Example: 
```go
    import (
      gofunc "https://github.com/MinhPhu0304/gofunc"
    )
    func main() {
      input := [3]int {1,2,3}
      gofunc.First(input) // Expected: 1
      stringInput := [3]string {"Hello", "Go", "developers"}
      gofunc.First(stringInput) // Expected: Hello
    }
  ```
- ## ```Last```:  Get the last element of any array type
Example: 
```go
    import (
      gofunc "https://github.com/MinhPhu0304/gofunc"
    )
    func main() {
      input := [3]int {1,2,3}
      gofunc.Last(input) // Expected: 3
      stringInput := [3]string {"Hello", "Go", "developers"}
      gofunc.Last(stringInput) // Expected: developers
    }
  ```
- ## ```Sum```:  Add all value in an integer list
Example: 
```go
    import (
      gofunc "https://github.com/MinhPhu0304/gofunc"
    )
    func main () {
      array := []int{1, 2, 3}
      gofunc.Sum(array) // Expected: 6
    }
  ```
- ## ```Always```:  Return curried function which always return value passed in
Example: 
```go
    import (
      gofunc "https://github.com/MinhPhu0304/gofunc"
    )
    func main () {
      alwaysTrue := gofunc.Always(true)
      alwaysTrue() // return true      
    }
  ```
- ## ```Prop```:  Return field value of a struct if present. ***Note***: all field should be export from the ```struct```
Example: 
```go
    import (
      gofunc "https://github.com/MinhPhu0304/gofunc"
    )

    type Person struct {
	    Name string
    }

    func main () {
      test := Person{Name: "test"}
      name := gofunc.Prop("Name", test) // Return "test"
    }
  ```

