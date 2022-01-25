package schema

type (
	AccountNumbers struct {
		AccountNumber [9]Display
	}
	Display struct {
		DisplayRow1 [3]string
		DisplayRow2 [3]string
		DisplayRow3 [3]string
	}
	JsonRequest struct {
		Numbers string `json:"numbers" binding:"required"`
	}
)

var (
	Display0 = Display{
		DisplayRow1: [3]string{" ", "_", " "},
		DisplayRow2: [3]string{"|", " ", "|"},
		DisplayRow3: [3]string{"|", "_", "|"},
	}
	Display1 = Display{
		DisplayRow1: [3]string{" ", " ", " "},
		DisplayRow2: [3]string{" ", " ", "|"},
		DisplayRow3: [3]string{" ", " ", "|"},
	}
	Display2 = Display{
		DisplayRow1: [3]string{" ", "_", " "},
		DisplayRow2: [3]string{" ", "_", "|"},
		DisplayRow3: [3]string{"|", "_", " "},
	}
	Display3 = Display{
		DisplayRow1: [3]string{" ", "_", " "},
		DisplayRow2: [3]string{" ", "_", "|"},
		DisplayRow3: [3]string{" ", "_", "|"},
	}
	Display4 = Display{
		DisplayRow1: [3]string{" ", " ", " "},
		DisplayRow2: [3]string{"|", "_", "|"},
		DisplayRow3: [3]string{" ", " ", "|"},
	}
	Display5 = Display{
		DisplayRow1: [3]string{" ", "_", " "},
		DisplayRow2: [3]string{"|", "_", " "},
		DisplayRow3: [3]string{" ", "_", "|"},
	}
	Display6 = Display{
		DisplayRow1: [3]string{" ", "_", " "},
		DisplayRow2: [3]string{"|", "_", " "},
		DisplayRow3: [3]string{"|", "_", "|"},
	}
	Display7 = Display{
		DisplayRow1: [3]string{" ", "_", " "},
		DisplayRow2: [3]string{" ", " ", "|"},
		DisplayRow3: [3]string{" ", " ", "|"},
	}
	Display8 = Display{
		DisplayRow1: [3]string{" ", "_", " "},
		DisplayRow2: [3]string{"|", "_", "|"},
		DisplayRow3: [3]string{"|", "_", "|"},
	}
	Display9 = Display{
		DisplayRow1: [3]string{" ", "_", " "},
		DisplayRow2: [3]string{"|", "_", "|"},
		DisplayRow3: [3]string{" ", "_", "|"},
	}
	Displayx = Display{
		DisplayRow1: [3]string{" ", "_", " "},
		DisplayRow2: [3]string{" ", "_", " "},
		DisplayRow3: [3]string{" ", "_", "|"},
	}

	ACTNumber AccountNumbers
)
