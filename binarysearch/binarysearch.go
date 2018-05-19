package main

func BinarySearch(array []int, target int) (int, error) {
    startPos := 0
    endPos := len(array) - 1

    for true {
        midPos := (endPos - startPos) / 2
        midVal := array[midPos]

        if midVal > target {
            startPos = midPos
        }else if midVal < target {
            endPos = midPos
        }else {
            return midPos, nil
        }
    }

    return -1, nil
}

