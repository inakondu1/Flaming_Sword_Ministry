package database

import (
	"Flaming_Sword_Ministry/models"
)

// Save a sermon
func CreateSermon(sermon models.Sermon) error {

	query := `
	INSERT INTO sermons (
		title,
		bible_verse,
		scripture_references,
		content,
		category,
		date,
		created_by
	)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err := DB.Exec(
		query,
		sermon.Title,
		sermon.BibleVerse,
		sermon.References,
		sermon.Content,
		sermon.Category,
		sermon.Date,
		sermon.CreatedBy,
	)

	return err
}

// Get all sermons
func GetAllSermons() ([]models.Sermon, error) {

	rows, err := DB.Query(`
		SELECT
			id,
			title,
			bible_verse,
			scripture_references,
			content,
			category,
			date,
			created_by
		FROM sermons
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sermons []models.Sermon

	for rows.Next() {

		var sermon models.Sermon

		err := rows.Scan(
			&sermon.ID,
			&sermon.Title,
			&sermon.BibleVerse,
			&sermon.References,
			&sermon.Content,
			&sermon.Category,
			&sermon.Date,
			&sermon.CreatedBy,
		)

		if err != nil {
			return nil, err
		}

		sermons = append(sermons, sermon)
	}

	return sermons, nil
}
