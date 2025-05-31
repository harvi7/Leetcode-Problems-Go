package array

func transformArray(arr []int) []int {
	var isChanged bool = true
	for isChanged {
		isChanged = false
		prev := arr[0]
		for i := 1; i < len(arr)-1; i++ {
			p := arr[i]
			if arr[i] > prev && arr[i] > arr[i+1] {
				arr[i]--
				isChanged = true
			}
			if arr[i] < prev && arr[i] < arr[i+1] {
				arr[i]++
				isChanged = true
			}
			prev = p
		}
	}
	return arr
}
