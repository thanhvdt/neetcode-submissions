func isValidSudoku(board [][]byte) bool {
	// create 3 arrays of 9 row/col/box boolean array,
	// iterate through the whole board and add each element into each suitable array to check if its valid sudoku 
	rowChecks := [9][9]bool{}
	colChecks := [9][9]bool{}
	subBoxesChecks := [3][3][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			
			// box is byte
			box := board[i][j] - '1'
			
			if rowChecks[i][box] {
				return false
			}
			rowChecks[i][box] = true

			if colChecks[j][box] {
				return false
			}
			colChecks[j][box] = true

			subBoxesRow := i/3
			subBoxesCol := j/3
			if subBoxesChecks[subBoxesRow][subBoxesCol][box] {
				return false
			}
			subBoxesChecks[subBoxesRow][subBoxesCol][box] = true
		}
	}

	return true
}
