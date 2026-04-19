package task_scheduler

func leastInterval(tasks []byte, n int) int {
	freq := make([]int, 26)
	maxNumber := 0
	maxCount := 0

	for _, task := range tasks {
		freq[task-'A']++
		if maxNumber == freq[task-'A'] {
			maxCount++
		} else if maxNumber < freq[task-'A'] {
			maxNumber = freq[task-'A']
			maxCount = 1
		}
	}

	gapCount := maxNumber - 1
	gapLength := n - (maxCount - 1)
	empty := gapCount * gapLength
	availableTasks := len(tasks) - maxNumber*maxCount
	idles := max(0, empty-availableTasks)

	return len(tasks) + idles
}
