package piscine

func SuccessCase() {
	puts("=======SUCCESS CASE========")
	/// Pawn
	{
		board := []string{
			"..K",
			".P.",
			"...",
		}
		Checkmate(board)
	}

	{
		board := []string{
			"K;QR",
			"\\\\\\\\",
			"QQQQ",
			"QQQQ",
		}
		Checkmate(board)
	}

	{
		board := []string{
			"K..",
			".P.",
			"...",
		}
		Checkmate(board)
	}

	/// Rook
	{
		board := []string{
			"R\nK",
			"...",
			"...",
		}
		Checkmate(board)
	}
	{
		board := []string{
			"KあR",
			"...",
			"...",
		}
		Checkmate(board)
	}
	{
		board := []string{
			"..K",
			"..\t",
			"..R",
		}
		Checkmate(board)
	}
	{
		board := []string{
			"  R",
			"   ",
			"  K",
		}
		Checkmate(board)
	}

	/// Bishop
	{
		board := []string{
			"..K",
			".b.",
			"B..",
		}
		Checkmate(board)
	}
	{
		board := []string{
			"..B",
			".B.",
			"K..",
		}
		Checkmate(board)
	}
	{
		board := []string{
			"B..",
			"...",
			"..K",
		}
		Checkmate(board)
	}
	{
		board := []string{
			"K..",
			"...",
			"..B",
		}
		Checkmate(board)
	}

	/// Queen
	{
		board := []string{
			"RQK",
			"...",
			"...",
		}
		Checkmate(board)
	}

	{
		board := []string{
			"K.QR",
			"....",
			"....",
			"....",
		}
		Checkmate(board)
	}

	{
		board := []string{
			"K.QR",
			"....",
			"....",
			"....",
		}
		Checkmate(board)
	}

	{
		board := []string{
			"..K",
			"...",
			"..Q",
		}
		Checkmate(board)
	}
	{
		board := []string{
			"..Q",
			"...",
			"..K",
		}
		Checkmate(board)
	}

	/// Bishop
	{
		board := []string{
			"..K",
			"...",
			"Q..",
		}
		Checkmate(board)
	}
	{
		board := []string{
			"..Q",
			"...",
			"K..",
		}
		Checkmate(board)
	}
	{
		board := []string{
			"Q..",
			"...",
			"..K",
		}
		Checkmate(board)
	}
	{
		board := []string{
			"K..",
			"...",
			"..Q",
		}
		Checkmate(board)
	}
	/// Rook
	{
		board := []string{
			"R.K",
			"...",
			"...",
		}
		Checkmate(board)
	}
	{
		board := []string{
			"K.R",
			"...",
			"...",
		}
		Checkmate(board)
	}
	{
		board := []string{
			"..K",
			"...",
			"..R",
		}
		Checkmate(board)
	}
	{
		board := []string{
			"..R",
			"...",
			"..K",
		}
		Checkmate(board)
	}

	/// Bishop
	{
		board := []string{
			"..K",
			"...",
			"Q..",
		}
		Checkmate(board)
	}
	{
		board := []string{
			"..Q",
			"...",
			"K..",
		}
		Checkmate(board)
	}
	{
		board := []string{
			"Q..",
			"...",
			"..K",
		}
		Checkmate(board)
	}
	{
		board := []string{
			"K..",
			"...",
			"..Q",
		}
		Checkmate(board)
	}
	{
		board := []string{
			"...",
			".Q.",
			"..K",
		}
		Checkmate(board)
	}
}

func FailCase() {
	puts("=======FAIL CASE========")

	// pawn

	{
		board := []string{
			"..........",
			"..........",
			"..R......Q",
			"..Q.......",
			"..........",
			"..........",
			"..........",
			".........P",
			"......P..K",
			"..........",
		}
		Checkmate(board)
	}

	{
		board := []string{
			"あああ",
			"あKあ",
			"あPあ",
		}
		Checkmate(board)
	}

	{
		board := []string{
			"PPP",
			"PKP",
			"\nP\n",
		}
		Checkmate(board)
	}

	{
		board := []string{
			"   ",
			"\nKあ",
			"\tPあ",
		}
		Checkmate(board)
	}

	// rook
	{
		board := []string{
			"..........",
			"..........",
			"..R......Q",
			"..Q.......",
			"..........",
			"..........",
			"..........",
			".........P",
			"......P..K",
			"..........",
		}
		Checkmate(board)
	}

	{
		board := []string{
			"あああ",
			"あKあ",
			"Rああ",
		}
		Checkmate(board)
	}

	// bishop
	{
		board := []string{
			".....R....",
			"..........",
			"....Q....K",
			"....Q.....",
			"..........",
			".........",
			"..........",
			"..........",
			"......P.B.",
			".....P....",
		}
		Checkmate(board)
	}

	{
		board := []string{
			"あああ",
			"あKあ",
			"あBあ",
		}
		Checkmate(board)
	}

	// queen

	{
		board := []string{
			"K..",
			"..Q",
			"...",
		}
		Checkmate(board)
	}

	{
		board := []string{
			"K..",
			".R.",
			"..Q",
		}
		Checkmate(board)
	}

	{
		board := []string{
			"KPQQQ",
			"PRQQQ",
			"QQQQQ",
			"QQQQQ",
			"QQQQQ",
		}
		Checkmate(board)
	}

	// non-board
	{
		board := []string{
			"ああああ",
			"あKああ",
			"Bあああ",
		}
		Checkmate(board)
	}

	{
		board := []string{
			"...",
			".K.",
			"B..",
			"...",
		}
		Checkmate(board)
	}

	{
		board := []string{
			"...",
			".K.",
			"B..",
			" ",
		}
		Checkmate(board)
	}

	{
		board := []string{
			"...",
			".K.",
			"B..",
			"",
		}
		Checkmate(board)
	}

	{
		board := []string{
			"K",
			"R",
		}
		Checkmate(board)
	}

	{
		board := []string{
			"K\n",
			"R",
		}
		Checkmate(board)
	}

	//non-only_K

	{
		board := []string{
			"K.",
			"KQ",
		}
		Checkmate(board)
	}

	{
		board := []string{
			"KKK",
			"KQK",
			"KKK",
		}
		Checkmate(board)
	}

	{
		board := []string{
			"K..",
			".Q.",
			"..K",
		}
		Checkmate(board)
	}
}
