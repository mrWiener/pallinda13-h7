/**
 * Created with IntelliJ IDEA.
 * User: lucas
 * Date: 3/19/13
 * Time: 2:32 PM
 * To change this template use File | Settings | File Templates.
 */
package ex43

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func () int {
	flastlast := 0
	flast := 1
	return func() int {
		f := flast + flastlast
		flastlast = flast
		flast = f
		return f
	}
}

func Main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
