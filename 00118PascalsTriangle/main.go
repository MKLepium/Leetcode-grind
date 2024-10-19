package main

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{{1}}
	}

	triangle := make([][]int, numRows)
	triangle[0] = []int{1}
	for i := 1; i < numRows; i++ {
		triangle[i] = generateNextLine(triangle[i-1])
	}
	return triangle
} 

func generateNextLine(prevLine []int) []int {
	nextLine := make([]int, len(prevLine)+1)
	nextLine[0] = 1
	nextLine[len(nextLine)-1] = 1

	for i := 1; i < len(nextLine)-1; i++ {
		nextLine[i] = prevLine[i-1] + prevLine[i]
	}

	return nextLine
}
