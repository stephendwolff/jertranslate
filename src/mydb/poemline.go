package mydb

import (
	"log"
	"strconv"
)

func GetPoemLines() ([]PoemLine, error) {
	log.Println("Get Poem Lines called, ")
	rows, err := Select("SELECT * FROM `trans_poemline`")

	if err != nil {
		return nil, err
	} else {
		log.Println("Got Poem Lines: ", rows)
		if len(rows) == 0 {
			return nil, err
		}
	}

	poemLines := make([]PoemLine, len(rows))

	for i := range rows {

		//Line			string `json:"line"`
		//LineNumber	uint8 `json:"line_number"`
		idString := rows[i]["id"].(string)
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			log.Println("Language: ", rows[i])
		}

		line := rows[i]["line"].(string)
		lineNumberString := rows[i]["line_number"].(string)
		lineNumber, err := strconv.ParseUint(lineNumberString, 10, 32)
		if err != nil {
			log.Println("lineNumber: ", rows[i])
		}
		poemLines[i] = PoemLine {
			ID: id,
			Line: line,
			LineNumber: lineNumber,
		}
	}

	return poemLines, nil
}