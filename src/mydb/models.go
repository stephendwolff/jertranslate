package mydb

import "time"

/*

class MemberState(models.Model):
    name = models.CharField(max_length=50, default="")
    in_europe = models.BooleanField(default=True)

    def __str__(self):
        return self.name

*/


type MemberState struct {
	Name             	string `json:"name"`
	InEurope        	bool `json:"translate"`
}

/*

class Language(models.Model):
    name = models.CharField(max_length=50, default="")
    translate = models.BooleanField(default=True)
    iso_two_letter_code = models.CharField(max_length=2, default="")

    def __str__(self):
        return self.name

*/

type Language struct {
	ID					uint64 `json:"id"`
	Name             	string `json:"name"`
	Translate        	bool `json:"translate"`
	IsoTwoLetterCode   	string `json:"iso_two_letter_code"`
}

/*

class PoemLine(models.Model):
    line = models.TextField()
    line_number = models.PositiveSmallIntegerField()


*/


type PoemLine struct {
	ID				uint64 `json:"id"`
	Line			string `json:"line"`
	LineNumber		uint64 `json:"line_number"`
}

/*

class Translation(models.Model):
    original = models.TextField()
    translation = models.TextField(default='')
    erosion = models.TextField(default='')
    original_language = models.ForeignKey(
        'Language', null=True, on_delete=SET_NULL,
        related_name='original_language')
    translation_language = models.ForeignKey(
        'Language', null=True, on_delete=SET_NULL,
        related_name='translation_language')
    poem_line = models.ForeignKey(PoemLine, null=True, blank=True, on_delete=SET_NULL)
    collected_at = models.DateTimeField(auto_created=True)
    use = models.BooleanField(default=False)
    order = models.IntegerField(default=0)
*/




type Translation struct {
	Original				string 		`json:"original"`
	Translation				string 		`json:"translation"`
	Erosion					string 		`json:"erosion"`
	OriginalLanguage		*Language	`json:"original_language"`
	TranslationLanguage		*Language	`json:"translation_language"`
	PoemLine				*PoemLine	`json:"poem_line"`
	CollectedAt				time.Time 	`json:"collected_at"`
	Use						bool 		`json:"use"`
	Order					uint8 		`json:"order"`
}

