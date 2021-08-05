package sort

func BubbleSort(values []int) []int {
	for j := 0; j < len(values)-1; j++ {
		for i := 0; i < len(values)-j-1; i++ {
			if values[i] > values[i+1] {
				values[i], values[i+1] = values[i+1], values[i]
			}
		}
	}
	return values
}
