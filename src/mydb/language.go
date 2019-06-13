package mydb

import (
	"log"
)

func GetLanguages() (interface{}, error) {
	log.Println("Get Languages called, ")
	rows, err := Select("SELECT * FROM `trans_language`")

	if err != nil {
		return nil, err
	} else {
		log.Println("Languages got ", rows)
		if len(rows) == 0 {
			return nil, err
		}
	}

	languages := make(map[string]interface{})

	for i := range rows {

		//Name             	string `json:"name"`
		//Translate        	bool `json:"translate"`
		//IsoTwoLetterCode   	string `json:"iso_two_letter_code"`

		name := rows[i]["name"].(string)
		translate := rows[i]["translate"].(string)
		isoTwoLetterCode := rows[i]["iso_two_letter_code"].(string)
		languages[isoTwoLetterCode] = map[string]interface{} {
			"name": name,
			"translate": translate,
		}
	}

	return languages, nil
}