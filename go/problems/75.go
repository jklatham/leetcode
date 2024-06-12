package problems

/*
Given an array nums with n objects...

...sort them in-place so that objects of the same color are adjacent.

int 0 = red
int 1 = white
int 2 = blue

Constraints:
- Do not use sort function.
- n == nums.length
- 1 <= n <= 300
- nums[i] is either 0, 1, or 2.

*/

func SortColors(num []int) {

	for j := range num {
		for k := range num {
			if num[j] <= num[k] {
				num[j], num[k] = num[k], num[j]
			}
		}
	}
}
