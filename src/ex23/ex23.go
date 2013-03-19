/**
 * Created with IntelliJ IDEA.
 * User: lucas
 * Date: 3/19/13
 * Time: 1:10 PM
 */
package ex23

import "fmt"

func Sqrt(x, limit float64) float64 {
	z := x
	delta := limit

	for delta >= limit {
		delta = z
		z = z - ((z*z - x)/(2.0*z))
		delta -= z
	}

	return z
}

func Main() {
	fmt.Println(Sqrt(4, 0.0001));
}
