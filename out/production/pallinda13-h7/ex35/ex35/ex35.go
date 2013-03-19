/**
 * Created with IntelliJ IDEA.
 * User: lucas
 * Date: 3/19/13
 * Time: 1:44 PM
 */
package ex35

import "pic"

func Pic(dx, dy int) [][]uint8 {
	out := make([][]uint8,0, dy)

	for i := 0; i < dy; i++ {
		out[i] = make([]uint8,0, dy)

		for j := 0; j < dy; j++ {
			out[i][j] = uint8(i+j)
		}
	}

	return out
}

func Ex35() {
	pic.Show(Pic)
}
