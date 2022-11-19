# Array
* Array's size should be defined at compile time that's why array are not usually used in Go lang.


    var x = [...]int{1, 2, 3}


# Slices

Go is a call by value language 
that means every time you pass parameter in function, 
Go will make a copy of that variable

* declare
  * var x = []int{1, 2, 3}
  * var data []int
    * this is nil slice
  * var x = []int{}
    * non-nil slice
* append
  * x = append(x, 10)
  * x = append(x, y...) //  y is another slice
* length
  * len(x)
* capacity
  * cap(x)
* make
  * x := make([]int, length, capacity)
* slicing
  * x[:2] | x[:exclusive]
  * x[1:] | x[inclusive:]
  * x[1:3] | x[inclusive:exclusive]
  * x[:]
* copy
  * number_of_element_copied := copy(destination_slice, source_slice)
* Maps
  * var nilMap map[string]int
  * totalWins := map[string]int{}
