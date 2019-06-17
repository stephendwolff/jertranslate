// Sample translate-quickstart translates "Hello, world!" into Russian.
package main

import (
	"log"
	"mydb"
)

func main() {
	mydb.SetDSN("jerusalem:jerusalem@tcp(127.0.0.1:3306)/jerusalem?charset=utf8")

	// get list of languages
	languages, err := mydb.GetLanguages()
	if err != nil {
		log.Fatalf("Failed to get languages: %v", err)
	}
	log.Print(languages)


	// get list of poem lines
	lineObjs, err := mydb.GetPoemLines()
	if err != nil {
		log.Fatalf("Failed to get lines: %v", err)
	}

	var poemLines []string
	for _, lineObj := range lineObjs {
		poemLines = append(poemLines, lineObj.Line)
	}

	// iterate through languages, translating to and from languages for the erosion
	for i, language := range languages {
		if language.IsoTwoLetterCode != "en" {
			log.Print(i, language.IsoTwoLetterCode)

			// get translation into language
			translations, err := mydb.GetGoogleTranslations(&poemLines, "en", language.IsoTwoLetterCode)
			if err != nil {
				log.Fatalf("Failed to translate: %v", err)
			}
			log.Print(i, translations)

			// get translation back into English
			erosions, err :=  mydb.GetGoogleTranslations(&translations, language.IsoTwoLetterCode, "en")
			if err != nil {
				log.Fatalf("Failed to translate back: %v", err)
			}
			log.Print(i, erosions)

			// store
			err = mydb.StoreTranslationsAndErosions(&translations, &erosions, "en", language.IsoTwoLetterCode, &languages, &lineObjs)
			if err != nil {
				log.Fatalf("Failed to store translations: %v", err)
			}
		}
	}
}

