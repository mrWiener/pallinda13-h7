/**
 * Created with IntelliJ IDEA.
 * User: lucas
 * Date: 3/19/13
 * Time: 1:44 PM
 */
package ex35

import "pic"

func Pic(dx, dy int) [][]uint8 {
	out := make([][]uint8, dy)

	y := 1

	for i := range out {
		y += i
		out[i] = make([]uint8, dy)

		for j := 0; j < dy; j++ {
			y += j
			out[i][j] = uint8(dx ^ (y/dx)*y)
		}
	}

	return out
}

func Main() {
	pic.Show(Pic)
}
