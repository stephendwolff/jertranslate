SELECT l.name, p.line, .t.original, t.translation, t.erosion
FROM trans_translation t
	LEFT JOIN trans_language l ON l.`id`=t.`translation_language_id`
	LEFT JOIN trans_poemline p ON p.`id`=t.`poem_line_id`
ORDER BY l.name, poem_line_id;