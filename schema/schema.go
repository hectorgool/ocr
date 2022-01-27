package schema

import (
	"encoding/json"

	"github.com/google/uuid"
)

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
	LogDB struct {
		ID        uuid.UUID       `json:"id" db:"id" binding:"required"`
		EndPoint  string          `db:"endpoint"`
		Method    string          `db:"method"`
		Input     json.RawMessage `db:"json_input"`
		Output    json.RawMessage `db:"json_output"`
		CreatedOn string          `db:"created_on"`
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

	Schema = `
		CREATE TABLE IF NOT EXISTS log_db (
			id VARCHAR(36) PRIMARY KEY,
    		endpoint TEXT,
    		method TEXT,
    		json_input JSON,
			json_output JSON,
			created_on DATETIME NOT NULL DEFAULT NOW()
		);
	`
)
