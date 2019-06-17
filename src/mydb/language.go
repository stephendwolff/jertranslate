package mydb

import (
	"log"
	"strconv"
)

func GetLanguages() ([]Language, error) {

	rows, err := Select("SELECT * FROM `trans_language`")

	if err != nil {
		return nil, err
	} else {
		// log.Println("Got Languages: ", rows)
		if len(rows) == 0 {
			return nil, err
		}
	}
	languages := make([]Language, len(rows))

	for i := range rows {

		//Name             	string `json:"name"`
		//Translate        	bool `json:"translate"`
		//IsoTwoLetterCode   	string `json:"iso_two_letter_code"`
		log.Println("Language: ", rows[i])
		idString := rows[i]["id"].(string)
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			log.Println("Language: ", rows[i])
		}

		name := rows[i]["name"].(string)
		translateStr := rows[i]["translate"].(string)
		translate := translateStr == "1"
		isoTwoLetterCode := rows[i]["iso_two_letter_code"].(string)
		languages[i] = Language{
			ID: id,
			IsoTwoLetterCode: isoTwoLetterCode,
			Name: name,
			Translate: translate,
		}
	}

	return languages, nil
}
