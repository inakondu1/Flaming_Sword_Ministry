package database

import "Flaming_Sword_Ministry/models"

// Save a prayer request
func CreatePrayer(prayer models.Prayer) error {

	_, err := DB.Exec(`
		INSERT INTO prayer_requests
		(name, phone, request, status)
		VALUES (?, ?, ?, ?)
	`,
		prayer.Name,
		prayer.Phone,
		prayer.Request,
		prayer.Status,
	)

	return err
}

// Get all prayer requests
func GetAllPrayers() ([]models.Prayer, error) {

	rows, err := DB.Query(`
		SELECT id, name, phone, request, status, created_at
		FROM prayer_requests
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prayers []models.Prayer

	for rows.Next() {

		var p models.Prayer

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Phone,
			&p.Request,
			&p.Status,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		prayers = append(prayers, p)
	}

	return prayers, nil
}
