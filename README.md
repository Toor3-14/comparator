<p align="center">
  <a href="https://github.com/Toor3-14/comparator" rel="noopener" target="_blank"><img width="150" src="logo.png" alt="comparator logo"></a>
</p>
<h1 align="center">Go Compare 😋</h1>


[![Gmail](https://img.shields.io/badge/-dmitrykeof@gmail.com-F9DB60?style=flat-square&logo=Yandex&logoColor=FF3333)](mailto:dmitrykeof@gmail.com) 
[![Telegram](https://img.shields.io/badge/Telegram-blue?style=flat-square&logo=Telegram)](https://t.me/redltoor)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Comparator is alibrary for easy creating/using custom comparator.
This library help you with creating generic functions/methods where you need use your own alphabet.
It will help the developer save time on developing their projects and make their code more beautiful.
Also will help the user to adjust your methods more flexibly to their needs.


The main goal of the project - is to standardize work with methods using different data types and needing their own alphabets

## Contents

0. [Installation](#Installation)
1. [Usage](#Usage)
    1. [Main Example](#Main-Example)
    2. [Before Using](#Before-Using)
    3. [Alphabet](#Alphabet)
    4. [Methods](#Methods)
        1. [GetIndex](#GetIndex)
        2. [Compare](#Compare)
        3. [Less](#Less)
        4. [More](#More)
        5. [Equal](#Equal)
        6. [LessEqual](#LessEqual)
        7. [MoreEqual](#MoreEqual)
2. [Contributing](#Contributing)
3. [Authors](#Authors)
4. [Current Version](#Current-Version)

## Installation

```cmd
$ go get github.com/Toor3-14/comparator
```
____
## Usage

### Main Example
```go
import (
	"fmt"
	"github.com/Toor3-14/comparator"
)
var c = comparator.New([]string{"!","A","B","C","D","E","F","G"})
func ChangeABC(abc []string) {
	c.SetABC(abc)
}
func BubleSort[T any](array []T) {
  for i := 0; i < len(array)-1; i++ {
     if (c.More(array[i],array[i+1])) {
       array[i], array[i+1] = array[i+1], array[i]
       i=-1;
     }
   }
}
func BinarySearch[T any](array []T, value T) int {
	result := -1 // as default - not found
	start := 0
	end := len(array) - 1
	for start <= end {
		mid := (start + end) / 2
		if c.Equal(array[mid], value) {
			result = mid // found
			break
		} else if c.Less(array[mid], value) {
			start = mid + 1
		} else if c.More(array[mid], value) {
			end = mid - 1
		}
	}
	
	return result
}
func main() {
	arrStr := []string{"Bin","Cosume","Alex","Dog","Cat"}
        arrInt := []int{0,4,6,7,3,2,1,9,8,5}
	
        BubleSort(arrStr)
        BubleSort(arrInt)
	
        fmt.Println(arrStr)
        fmt.Println(arrInt)
	
        fmt.Println(BinarySearch(arrStr,"Cat"))
        fmt.Println(BinarySearch(arrInt,9))
}
```
____
### Before Using
Before using the comparator, it must be initialized.
Example:
```go
c := comparator.New(alphabet []string)
```
or just:
```go
c := comparator.New(), it use default alphabet
```
____
### Alphabet
The alphabet is represented as a slice of a string. You can also set it.
Example:
```go
c.SetABC([]string{"A","B","C","abcde"})
```
The lower the slice index, the higher the priority of the value.
You can restore the default alphabet:
```go
c.SetDefaultABC()
```
IMPORTANT:
The alphabet elements can also take a string of two or more characters. 
This means they have the same priority.
____
### Methods

#### GetIndex
Takes a string type character (length must be 1) and returns its index from the alphabet
Example:
```go
c.GetIndex("A") //With the standard alphabet , it will return 0
```
____
#### Compare
Compare compares two variables of type (int, float32, float64, string) and returns the result of the comparison
Also your arguments must have the same type.

Comparison scheme:
```
x > y = 1
x < y = -1
x == y = 0
```
Example:
```go
if c.Compare("A","B") == -1 {
      ....
}
```
Attention! Don't forget that A is less than B with the standard alphabet

For a more familiar comparison , there are the following methods:
____
#### Less
This compares two variables of type (int,float32,float64,string) and returns true if the first argument is less than the second
Example:
```go
if c.Less("A", "B") {  // "A" < "B"
    ....
}
```
____
#### More
This compares two variables of type (int,float32,float64,string) and returns true if the first argument is greater than the second
Example:
```go
if c.More("B", "A") {  // "B" > "A"
    ....
}
```
____
#### Equal
This compares two variables of type (int,float32,float64,string) and returns true if the first argument is equal than the second
Example:
```go
if c.Equal("C", "C") {  // "C" == "C"
    ....
}
```
____
#### LessEqual
This compares two variables of type (int,float32,float64,string) and returns true if the first argument is less or equal than the second
Example:
```go
if c.LessEqual("C", "C") {  // "C" <= "C"
    ....
}
```
___
#### MoreEqual 
This compares two variables of type (int,float32,float64,string) and returns true if the first argument is greater or equal than the second
Example:
```go
if c.MoreEqual("B", "A") {  // "B" >= "A"
    ....
}
```
___
## Contributing

* Fork the [project](https://github.com/Toor3-14/comparator)
* Fix [project issues](https://github.com/Toor3-14/comparator/issues) or request new features

Don't hesitate ;)
___

## Authors
 * Toor3-14 / chi76

## Current Version
 0.1B
