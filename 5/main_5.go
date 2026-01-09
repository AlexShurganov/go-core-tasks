package main

func CheckInsertion(slice1, slice2 []int) (bool, []int) {
	seen := make(map[int]bool)
	var result []int

	for _, v := range slice2 {
		seen[v] = true
	}

	for _, v := range slice1 {
		if seen[v] {
			result = append(result, v)
			delete(seen, v)
		}
	}

	return len(result) > 0, result
}
