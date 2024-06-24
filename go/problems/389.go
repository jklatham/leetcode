package problems

func FindTheDifference(s string, t string) byte {

	var count int

	for _, i := range s + t {
		// sum mod 2
		count ^= int(i)
	}

	return byte(count)
}
