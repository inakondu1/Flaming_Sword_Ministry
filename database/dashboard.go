package database

func CountUsers() (int, error) {

	var count int

	err := DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)

	return count, err
}

func CountSermons() (int, error) {

	var count int

	err := DB.QueryRow("SELECT COUNT(*) FROM sermons").Scan(&count)

	return count, err
}

func CountAnnouncements() (int, error) {

	var count int

	err := DB.QueryRow("SELECT COUNT(*) FROM announcements").Scan(&count)

	return count, err
}
