## Arrays - Arrays, Slices and Maps

Arrays are a special data structure in Go that allow us to allocate contiguous blocks of fixed size memory. Arrays have some special features in Go related to how they are declared and viewed as types.

## Notes

* Arrays are fixed length data structures that can't change.
* Arrays of different sizes are considered to be of different types.
* Memory is allocated as a contigous block.

## Code Review

[Declare, initalize and iterate](example1/example1.go) ([Go Playground](http://play.golang.org/p/irrA08aCkm))

[Different type arrays](example2/example2.go) ([Go Playground](http://play.golang.org/p/LVD43cYBNS))

[Contigous memory allocations](example3/example3.go) ([Go Playground](http://play.golang.org/p/-Hg3nUdO5p))

## Exercises

### Exercise 1

Declare an array of 5 strings with each element initialized to its zero value. Declare a second array of 5 strings and initialize this array with literal string values. Assign the second array to the first and display the results of the first array.

[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/PvjKR2uU2H))
