package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    var count int
	for _, square := range cb[file] {
        if square {
            count++
        }
    }
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    var count int
	
    if rank < 1 || rank > 8 {
        return count
    }
    
	for _, v := range cb {
        if v[rank-1] {
            count++
        }
    }
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    var count int
    for _, squares := range cb {
        count += len(squares)
    }
    return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    var count int
    for _, squares := range cb {
        for _, sq := range squares {
            if sq {
                count++
            }
        }
    }
    return count
}
