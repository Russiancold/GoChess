package main

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	TypeShift  uint8 = 1
	TypeLength uint8 = 3
	TypeColorMask uint8 = 15
	MovedShift uint8 = 4
	ColorMask  uint8 = 1
	TypeMask   uint8 = 7 << 1
	MovedMask  uint8 = 1 << 4
)

const (
	King = uint8((iota + 1) << 1)
	Queen
	Rook
	Bishop
	Knight
	Pawn
)

const (
	WPawn = Pawn + 1
	BPawn = Pawn
)

var (
	enPassantFile         = uint8(8)
	availaibleSquareBuf   [50]uint8
	availaibleSquareCount uint8
	turn = uint8(1)
)

func toggleTurn() {
	turn ^= 1
}

func main() {
	generateBoard()
	printBoard()
	var sq1 uint8
	var sq2 uint8
	for {
		fmt.Scanln(&sq1)
		fmt.Scanln(&sq2)
		if turn == board[sq1]&ColorMask {
			if err := Move(&sq1, &sq2); err != nil {
				fmt.Println(err)
			}
			toggleTurn()
			printBoard()
		} else {
			fmt.Println("It's not your turn")
		}
	}
}

var board [64]uint8

func printBoard() {
	for i := 7; i >= 0; i-- {
		for j := 0; j < 8; j ++ {
			fmt.Print(board[i*8+j]&TypeMask>>1, " ")
		}
		fmt.Println()
	}
}

func generateBoard() {
	for i := 8; i < 16; i++ {
		setWhite(&board[i])
		setType(&board[i], Pawn)

		setBlack(&board[40+i])
		setType(&board[40+i], Pawn)
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

func availaibleContains(a *uint8) bool {
	for i := uint8(0); i < availaibleSquareCount; i++ {
		if availaibleSquareBuf[i] == *a {
			return true
		}
	}
	return false
}

func Move(square1, square2 *uint8) error {
	fmt.Println(enPassantFile)
	fmt.Println(strconv.FormatInt(int64(board[*square1]), 2))
	getAvailableSquares(square1)
	if availaibleContains(square2) {
		if board[*square1]&TypeColorMask == WPawn {
			if getRank(square1) == 1 {
				board[*square1] |= MovedMask
				if *square2-*square1 == 16 {
					enPassantFile = getFile(square1)
					board[*square2], board[*square1] = board[*square1], 0
					return nil
				} else {
					enPassantFile = 8
					board[*square2], board[*square1] = board[*square1], 0
					return nil
				}
			}
			if getFile(square2) != getFile(square1) && board[*square2] == 0 {
				board[*square2], board[*square1] = board[*square1], 0
				board[*square2 - 8] = 0
				enPassantFile = 8
				return nil
			}
		}
		if board[*square1]&TypeColorMask == BPawn {
			if getRank(square1) == 6 {
				board[*square1] |= MovedMask
				if *square1-*square2 == 16 {
					enPassantFile = getFile(square1)
					board[*square2], board[*square1] = board[*square1], 0
					return nil
				} else {
					enPassantFile = 8
					board[*square2], board[*square1] = board[*square1], 0
					return nil
				}
			}
			if getFile(square2) != getFile(square1) && board[*square2] == 0 {
				board[*square2], board[*square1] = board[*square1], 0
				board[*square2 + 8] = 0
				enPassantFile = 8
				return nil
			}
		}
		board[*square2], board[*square1] = board[*square1], 0
		enPassantFile = 8
		return nil
	}
	return errors.New("Invalid move")
}

func getAvailableSquares(square *uint8) {
	piece := board[*square]
	switch piece & TypeMask {
	case Pawn:
		getPawnAvailableSquares(square)
		return
	}
	return
}

//TODO: En passant
func getPawnAvailableSquares(square *uint8) {
	if board[*square]&ColorMask == 1 {
		if getFile(square) != 0 {
			if board[*square+7] != 0 && board[*square+7]&ColorMask == 0 {
				availaibleSquareBuf[availaibleSquareCount] = *square + 7
				availaibleSquareCount++
			}
		}
		if getFile(square) != 7 {
			if board[*square+9] != 0 && board[*square+9]&ColorMask == 0 {
				availaibleSquareBuf[availaibleSquareCount] = *square + 9
				availaibleSquareCount++
			}
		}
		if board[*square+8] == 0 {
			availaibleSquareBuf[availaibleSquareCount] = *square + 8
			availaibleSquareCount++
			if board[*square+16] == 0 && board[*square]&MovedMask>>MovedShift == 0 {
				availaibleSquareBuf[availaibleSquareCount] = *square + 16
				availaibleSquareCount++
			}
		}
		if enPassantFile != 8 {
			if enPassantFile-getFile(square) == 1 && getRank(square) == 4 {
				availaibleSquareBuf[availaibleSquareCount] = *square + 9
				availaibleSquareCount++
			}
			if getFile(square)-enPassantFile == 1 && getRank(square) == 4 {
				availaibleSquareBuf[availaibleSquareCount] = *square + 7
				availaibleSquareCount++
			}
		}
		return
	}
	if getFile(square) != 0 {
		if board[*square-9] != 0 && board[*square-9]&ColorMask == 1 {
			availaibleSquareBuf[availaibleSquareCount] = *square - 9
			availaibleSquareCount++
		}
	}
	if getFile(square) != 7 {
		if board[*square-7] != 0 && board[*square+9]&ColorMask == 1 {
			availaibleSquareBuf[availaibleSquareCount] = *square -7
			availaibleSquareCount++
		}
	}
	if board[*square-8] == 0 {
		availaibleSquareBuf[availaibleSquareCount] = *square - 8
		availaibleSquareCount++
		if board[*square-16] == 0 && board[*square]&MovedMask>>MovedShift == 0 {
			availaibleSquareBuf[availaibleSquareCount] = *square - 16
			availaibleSquareCount++
		}
	}
	if enPassantFile != 8 {
		if enPassantFile-getFile(square) == 1 && getRank(square) == 3 {
			availaibleSquareBuf[availaibleSquareCount] = *square - 7
			availaibleSquareCount++
		}
		if getFile(square)-enPassantFile == 1 && getRank(square) == 3 {
			availaibleSquareBuf[availaibleSquareCount] = *square - 9
			availaibleSquareCount++
		}
	}
	return

}

func getIndex(rank uint8, file uint8) uint8 {
	return (rank-1)*8 + (file - 1)
}

func getFile(square *uint8) uint8 {
	return *square % 8
}

func getRank(i *uint8) uint8 {
	return *i / 8
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
