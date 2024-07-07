// 2582 - Pass the Pillow

package problems

import "fmt"

func PassThePillow(n int, time int) int {
	pillow := 1
	dir := 1

	for time > 0 {

		if pillow == n {
			dir = -1
		}
		if pillow == 1 {
			dir = 1
		}
		pillow += dir
		time--
		fmt.Println("Pillow:", pillow, "Time:", time)
	}
	return pillow
}
