package sort

func ShellSort(a []int) []int {
	gap := 1
	size := len(a)
	for size >= 3*gap {
		gap = 3*gap + 1
	}
	for gap > 0 {
		for i := gap; i < size; i++ {
			for j := i; j > gap-1; j -= gap {
				if a[j-gap] >= a[j] {
					a[j-gap], a[j] = a[j], a[j-gap]
				}
			}
		}
		gap = (gap - 1) / 3
	}
	return a
}
