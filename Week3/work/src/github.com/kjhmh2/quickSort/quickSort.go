package quickSort

func partition(array []int, left, right int) int {
    key := array[left]
    // loop
    for left < right {
        for (left < right && array[right] >= key) {
            right --
        }
        // change
        array[left] = array[right]
        for (left < right && array[left] <= key) {
            left ++
        }
        array[right] = array[left]
    }
    array[right] = key
    return right
}

//quicksort method implementation:
func quickSort(temp []int, left, right int) {
    if left >= right {
        return
    }
    index := partition(temp, left, right)
    quickSort(temp, left, index - 1)
    quickSort(temp, index + 1, right)
}