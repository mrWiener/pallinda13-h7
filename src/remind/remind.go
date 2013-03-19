/**
 * Created with IntelliJ IDEA.
 * User: lucas
 * Date: 3/19/13
 * Time: 2:51 PM
 * To change this template use File | Settings | File Templates.
 */
package remind

import (
	"time"
	"fmt"
)

func Remind(text string, paus time.Duration) {
	go func() {
		for {
			time.Sleep(paus);
			fmt.Printf("Klockan är %s: %s\n", time.Now().Format("15.04"), text)
		}
	}()
}

func Main() {
	Remind("Dags att äta", time.Second*3)
	Remind("Dags att arbeta", time.Second*8)
	Remind("Dags att sova", time.Second*24)
}
