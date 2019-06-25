// Sample translate-quickstart translates "Hello, world!" into Russian.
package main

import (
	"log"
	"mydb"
	"time"
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

	var erosionPoemLines []string
	for _, lineObj := range lineObjs {
		erosionPoemLines = append(erosionPoemLines, lineObj.Line)
	}

	// Take 3: iterate through languages, translating to and from languages for the erosion, feeding translation back in *and* not the erosion
	var firstLangCode = "en"
	for i, language := range languages {
		if language.IsoTwoLetterCode != "en" {

			log.Printf("Translate to from %s (%s)", language.IsoTwoLetterCode, language.Name)

			// get translation into language
			translations, err := mydb.GetGoogleTranslations(&poemLines, firstLangCode, language.IsoTwoLetterCode)
			if err != nil {
				log.Fatalf("Failed to translate: %v", err)
			}
			log.Print(i, translations)

			// *always* get translation back into English
			erosionPoemLines, err :=  mydb.GetGoogleTranslations(&translations, language.IsoTwoLetterCode, "en")
			if err != nil {
				log.Fatalf("Failed to translate back: %v", err)
			}
			log.Print(i, erosionPoemLines)

			duration := time.Second
			time.Sleep(duration)

			// store
			err = mydb.StoreTranslationsAndErosions(&translations, &erosionPoemLines, firstLangCode, language.IsoTwoLetterCode, &languages, &lineObjs, &poemLines)
			if err != nil {
				log.Fatalf("Failed to store translations: %v", err)
			}
			firstLangCode = language.IsoTwoLetterCode
		}
	}


	//	var apiCallCount = 0
	// Take 1: iterate through languages, translating to and from languages for the erosion
	/*	for i, language := range languages {
			if language.IsoTwoLetterCode != "en" {
	//			log.Fo("Translate to from %s (%s), language.IsoTwoLetterCode)
				log.Printf("Translate to from %s (%s)", language.IsoTwoLetterCode, language.Name)

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
				err = mydb.StoreTranslationsAndErosions(&translations, &erosions, "en", language.IsoTwoLetterCode, &languages, &lineObjs, &poemLines)
				if err != nil {
					log.Fatalf("Failed to store translations: %v", err)
				}

				duration := time.Second
				//Millisecond * 250
				time.Sleep(duration)
			}
		}
	*/


	/*
		// Take 2: iterate through languages, translating to and from languages for the erosion, feeding erosion back in
		for i, language := range languages {
			if language.IsoTwoLetterCode != "en" {
				//			log.Fo("Translate to from %s (%s), language.IsoTwoLetterCode)
				log.Printf("Translate to from %s (%s)", language.IsoTwoLetterCode, language.Name)

				// get translation into language
				translations, err := mydb.GetGoogleTranslations(&poemLines, "en", language.IsoTwoLetterCode)
				if err != nil {
					log.Fatalf("Failed to translate: %v", err)
				}
				log.Print(i, translations)

				// get translation back into English
				poemLines, err :=  mydb.GetGoogleTranslations(&translations, language.IsoTwoLetterCode, "en")
				if err != nil {
					log.Fatalf("Failed to translate back: %v", err)
				}
				log.Print(i, poemLines)

				duration := time.Second
				time.Sleep(duration)

				// store
				err = mydb.StoreTranslationsAndErosions(&translations, &poemLines, "en", language.IsoTwoLetterCode, &languages, &lineObjs, &poemLines)
				if err != nil {
					log.Fatalf("Failed to store translations: %v", err)
				}
			}
		}

	*/

}

