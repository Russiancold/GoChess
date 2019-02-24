package main

const (
	TypeShift  uint8 = 1
	TypeLength uint8 = 3
	MovedShift uint8 = 4
	TypeMask   uint8 = 7 << 1
)

const (
	King = (iota + 1) << 1
	Queen
	Rook
	Bishop
	Knight
	Pawn
)

func main() {
	generateBoard()

}

var board [64]uint8

func generateBoard() {
	for i := 8; i < 16; i++ {
		setWhite(&board[i])
		setType(&board[i], Pawn)

		setBlack(&board[63 - i])
		setType(&board[63 - i], Pawn)
	}
	setWhite(&board[0])
	setType(&board[0], Rook)
	setWhite(&board[1])
	setType(&board[1], Knight)
	setWhite(&board[2])
	setType(&board[2], Bishop)
	setWhite(&board[3])
	setType(&board[3], Queen)
	setWhite(&board[4])
	setType(&board[4], King)
	setWhite(&board[5])
	setType(&board[5], Bishop)
	setWhite(&board[6])
	setType(&board[6], Knight)
	setWhite(&board[7])
	setType(&board[7], Rook)
	setBlack(&board[56])
	setType(&board[56], Rook)
	setBlack(&board[57])
	setType(&board[57], Knight)
	setBlack(&board[58])
	setType(&board[58], Bishop)
	setBlack(&board[59])
	setType(&board[59], Queen)
	setBlack(&board[60])
	setType(&board[60], King)
	setBlack(&board[61])
	setType(&board[61], Bishop)
	setBlack(&board[62])
	setType(&board[62], Knight)
	setBlack(&board[63])
	setType(&board[63], Rook)
}

func getIndex(rank uint8, file uint8) uint8 {
	return (rank - 1) * 8 + (file - 1)
}

func getRank(i uint8) uint8 {
	return i / 8 + 1
}

func setWhite(piece *uint8) {
	*piece |= 1
}

func setBlack(piece *uint8) {
	*piece &= 0
}

func setType(piece *uint8, mask uint8) {
	*piece &= ^TypeMask // Clear lower 4 bits. Note: ~0xf == 0xfffffff0
	*piece |= mask & TypeMask
}
