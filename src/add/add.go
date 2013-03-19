/**
 * Created with IntelliJ IDEA.
 * User: lucas
 * Date: 3/19/13
 * Time: 3:11 PM
 */
package add

import "fmt"

func Add(a []int, res chan<- int) {
	r := 0

	for _, v := range a {
		r += v
	}

	res <- r
}

func Main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}

	n := len(a)
	ch := make(chan int)
	go Add(a[:n/2], ch)
	go Add(a[n/2:], ch)

	fmt.Println("sum:", <-ch + <-ch)
}
