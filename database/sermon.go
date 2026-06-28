package database

import (
	"Flaming_Sword_Ministry/models"
)

// CreateSermon saves a sermon into the database.
func CreateSermon(sermon models.Sermon) error {

	query := `
	INSERT INTO sermons (
		title,
		bible_verse,
		references,
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

// GetAllSermons returns every sermon in the database.
func GetAllSermons() ([]models.Sermon, error) {

	query := `
	SELECT
		id,
		title,
		bible_verse,
		references,
		content,
		category,
		date,
		created_by
	FROM sermons
	ORDER BY id DESC
	`

	rows, err := DB.Query(query)
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
