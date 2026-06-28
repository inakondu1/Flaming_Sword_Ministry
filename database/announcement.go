package database

import "Flaming_Sword_Ministry/models"

func CreateAnnouncement(title, message string) error {

	_, err := DB.Exec(
		"INSERT INTO announcements (title, message) VALUES (?, ?)",
		title, message,
	)

	return err
}

func GetAnnouncements() ([]models.Announcement, error) {

	rows, err := DB.Query(`
		SELECT id, title, message
		FROM announcements
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Announcement

	for rows.Next() {
		var a models.Announcement
		rows.Scan(&a.ID, &a.Title, &a.Message)
		list = append(list, a)
	}

	return list, nil
}