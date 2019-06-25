package mydb

import (
	"cloud.google.com/go/translate"
	"context"
	"golang.org/x/text/language"
	"log"
	"time"
)

func GetGoogleTranslations(lines *[]string, sourceLanguageCode string, targetLanguageCode string) ([]string, error) {

	ctx := context.Background()

	// Creates a client.
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}

	// Set the source language.
	source, err := language.Parse(sourceLanguageCode)
	if err != nil {
		log.Fatalf("Failed to parse original language: %v", err)
		return nil, err
	}

	// Sets the target language.
	target, err := language.Parse(targetLanguageCode)
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
		return nil, err
	}

	// Translates the text into target language.
	translationsObjs, err := client.Translate(ctx, *lines, target, &translate.Options{
		Source: source,
	})

	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
	}

	var translations []string
	for _, obj := range translationsObjs {
		translations = append(translations, obj.Text)
	}

	return translations, nil
}


func StoreTranslationsAndErosions(translations *[]string,
								  erosions *[]string,
								  erosionLanguageCode string,
								  targetLanguageCode string,
								  languages *[]Language,
								  poemLines *[]PoemLine,
								  originals *[]string) error {



	var erosionLanguageID uint64
	var targetLanguageID uint64
	var poemLineID uint64

	for _, lang := range *languages {
		if lang.IsoTwoLetterCode == targetLanguageCode{
			targetLanguageID = lang.ID
		}
		if lang.IsoTwoLetterCode == erosionLanguageCode {
			erosionLanguageID = lang.ID
		}
	}

	query := "INSERT INTO trans_translation (`original`, `translation`, `erosion`, `original_language_id`, `translation_language_id`, `order`, `use`, `poem_line_id`, `collected_at`) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)"

	for i, _ := range *poemLines {

		// big assumption here, translations are in poem line order, so poemLine.ID is what we use for the insert
		poemLineID = uint64(i + 1)
		_, _, err := Exec(query,
			(*originals)[i], (*translations)[i], (*erosions)[i], erosionLanguageID, targetLanguageID, 0, true, poemLineID,
			time.Now().UTC().Format("2006-01-02 15:04:05.000"))
//		(*originals)[i], (*translations)[i], (*erosions)[i], erosionLanguageID, targetLanguageID, 0, true, poemLineID,
//			time.Now().UTC().Format("2006-01-02 15:04:05.000"))
		if err != nil {
			log.Print(err)
			return err
		}
	}
	return nil

}
