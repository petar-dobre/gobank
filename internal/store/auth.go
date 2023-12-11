package store

func (s *PostgresStore) GetHashedPassword(email string) (string, error) {
	var hashedPassword string

	err := s.db.QueryRow(`select password from account where email = $1`, email).Scan(&hashedPassword)
	if err != nil {
		return "", err
	}

	return hashedPassword, nil
}
