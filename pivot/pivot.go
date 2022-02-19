package pivot

import "log"

// FindPivotPoint : finds the pivot of an array with o(n) complexity
func FindPivotPoint(arr []int, printPasses bool) int {
	const notFound = -1
	loopCount := 0
	vectorSize := len(arr)

	if printPasses {
		defer func() {
			log.Printf("Vector size %d - total rounds to find pivot %d", vectorSize, loopCount)
		}()
	}

	rightIndex := vectorSize - 1
	if rightIndex < 2 {
		return notFound
	}

	leftIndex := 0

	leftSum := arr[leftIndex]
	rightSum := arr[rightIndex]

	for {
		loopCount++

		if leftSum < rightSum {
			leftIndex++
			leftSum += arr[leftIndex]
		}

		if leftSum > rightSum {
			rightIndex--
			rightSum += arr[rightIndex]
		}

		switch rightIndex - leftIndex {
		case 2:
			if leftSum == rightSum {
				return leftIndex + 1
			}
			fallthrough
		case 1:
			return notFound
		default:
			if leftSum == rightSum {
				leftIndex++
				rightIndex--
				leftSum += arr[leftIndex]
				rightSum += arr[rightIndex]
			}
		}
	}
}
