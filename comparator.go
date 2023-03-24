/*Package comparator implements 
a simple library for creating/using custom comparator

1. Before using the comparator, it must be initialized. 
	Example: c := comparator.New(alphabet []string),
	or just, c := comparator.New(), it use default alphabet
	
2. The alphabet is represented as a slice of a string. You can also set it.
	Example: c.SetABC([]string{"A","B","C","abcde"}),

	The lower the slice index, the higher the priority of the value.
	You can restore the default alphabet - c.SetDefaultABC() ----
	implNote!	The alphabet elements can also take a string of two or more characters. 
				This means they have the same priority.

3. After that, methods are available to you:
	
	3.1. Compare:
			first interface{}
			second interface{}
		return:
			(int) 1 - first > second
			(int) -1 - first < second
			(int) 0 - first == second

		WARRNING: only [int,float32,float64,string] types supported yet.
		Also your arguments must have the same type.

	3.2. GetIndex:
			char string (it must be ONE symbol)
		 return:
			(int) 1 - first > second
			(int) -1 - first < second
			(int) 0 - first == second


4. For the convenience of using comparator you can use the following methods:

	4.1. Less:
			first interface{}
			second interface{}
		return:
			{true}, if first < second, else {false}
	4.2. More:
			first interface{}
			second interface{}
		return:
			{true}, if first > second, else {false} 
	4.3. Equal:
			first interface{}
			second interface{}
		return:
			{true}, if first == second, else {false} 
	4.4. LessEqual:
			first interface{}
			second interface{}
		return:
			{true}, if first <= second, else {false} 
	4.5. MoreEqual:
			first interface{}
			second interface{}
		return:
			{true}, if first >= second, else {false} 
	

And finnaly. 
	You can use this package to create generic functions or methods 
	where you need to use your own alphabet for strings. See examples if you interested.

	I hope you enjoy this library and help you avoid writing unnecessary code.
	Have fun <3


Author * Toor3-14 [chi76];
Version 0.1B - Initial;
License: MIT License;
*/

package comparator

var defaultABC = []string{"A","B","C","D","E","F","G",
                   "H","I","J","K","L","M","N",
                   "O","P","Q","R","S","T","U",
                   "V","W","X","Y","Z",

                   "a","b","c","d","e","f","g",
                   "h","i","j","k","l","m","n",
                   "o","p","q","r","s","t","u",
                   "v","w","x","y","z"}


//______________Struct Section

type comparator struct {
	abc []string
}

func New(ABC ...[]string) *comparator {
	var set []string
	if len(ABC) == 0 {
		set = defaultABC
	}else {
		set = ABC[0]
	}
	return &comparator{abc: set}
}

func (c *comparator) SetABC(ABC []string) *comparator {
	c.abc = ABC
	return c
}
func (c *comparator) SetDefaultABC() *comparator {
	c.abc = defaultABC
	return c
}
//______________Supporting Methods (used in main method/func)

func (c *comparator) GetIndex(char string) int {
	for i := 0; i < len(c.abc); i++  {

	   if(len(c.abc[i]) > 1) {
		for j := 0; j < len(c.abc[i]); j++  {
		  if(char == c.abc[i][j:j+1]) {
			return i
		  }
		}
	  }else {
		if(char == c.abc[i][0:1]) {
		  return i
		}
	  }
	
	}
	
	return len(c.abc)
}

func checkType(elem interface{}) int {
	switch elem.(type) {
	case string:
		return 1
	case float32:
		return 2
	case float64:
		return 3
	case int:
		return 4
	default:
		return 0
	}
}

//______________Main method
//              Supported types: string, float32, float64, int


func (c *comparator) Compare(first, second interface{}) int {
	
	// result - default value - 0, 
	// if it is not changed, then the values are equivalent
	var result int = 0

	if checkType(first) != checkType(second) {
		panic("Different types of variables in arguments method - comparator.Compare()")
	}
	if checkType(first) == 0 || checkType(second) == 0 {
		panic("Incorrect type(s) of arguments in method - comparator.Compare()")
	} 

	switch first.(type) {
	case string:
		first,_:= first.(string)
		second,_:= second.(string)

		if c.GetIndex(first[0:1]) > c.GetIndex(second[0:1]) {
			result = 1
		}else if c.GetIndex(first[0:1]) < c.GetIndex(second[0:1]) {
			result = -1
		}else if c.GetIndex(first[0:1]) == c.GetIndex(second[0:1]) {

			var min int = 0
			lf,ls := len(first), len(second)
			if lf > ls {
				min = ls
			}else {
				min = lf
			}

			for i:=1; i < min; i++ {
				if c.GetIndex(first[i:i+1]) > c.GetIndex(second[i:i+1]) {
					result = 1
					break
				}else if c.GetIndex(first[i:i+1]) < c.GetIndex(second[i:i+1]) {
					result = -1
					break
				}else if c.GetIndex(first[i:i+1]) == c.GetIndex(second[i:i+1]) {
					continue
				}
			}

		}
	case float32:
		f,_ := first.(float32)
		s,_ := second.(float32)
		if f > s {
			result = 1
		}else if f < s {
			result = -1
		}
	case float64:
		f,_ := first.(float64)
		s,_ := second.(float64)
		if f > s {
			result = 1
		}else if f < s {
			result = -1
		}
	default:
		f,_ := first.(int)
		s,_ := second.(int)
		if f > s {
			result = 1
		}else if f < s {
			result = -1
		}
	}
	return result
}

//______________Auxiliary Methods
//              For the usual check through [<,>, ==, <=, >=] operators

func (c *comparator) Less(first, second interface{}) bool {
	if c.Compare(first, second) == -1 {
		return true
	}else {
		return false
	}
}
func (c *comparator) More(first, second interface{}) bool {
	if c.Compare(first, second) == 1 {
		return true
	}else {
		return false
	}
}
func (c *comparator) Equal(first, second interface{}) bool {
	if c.Compare(first, second) == 0 {
		return true
	}else {
		return false
	}
}

func (c *comparator) LessEqual(first, second interface{}) bool {
	if c.Compare(first, second) <= 0 {
		return true
	}else {
		return false
	}
}
func (c *comparator) MoreEqual(first, second interface{}) bool {
	if c.Compare(first, second) >= 0 {
		return true
	}else {
		return false
	}
}

/*
	Todo:
		make better customize alphabet,
		add patterns: [a-z]/[A-Z]/[a-zA-Z]/[1-9]
*/
