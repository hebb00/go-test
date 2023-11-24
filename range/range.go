package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var c int
	for k, x := range cb {
		if k == file {
			for _, i := range x {
				if i == true {
					c++
				}

			}
		}

	}
	return c

}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var c int
	if rank >= 1 && rank <= 8 {
		for _, x := range cb {
			for rr, i := range x {
				if rr == rank {
					if i == true {
						c++
					}
				}

			}
		}
	}
	return c
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var c int
	for _, x := range cb {
		c = c + len(x)

	}
	return c
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var c int
	for _, x := range cb {
			for _, i := range x {
				if i == true {
					c++
				}

			}
	

	}
	return c
}
