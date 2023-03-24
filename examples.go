package comparator
/* MAIN EXAMPLE

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

*/