package service

// Check if the password is correct
func VerifyPassword(userPassword, inputPassword string) bool {
	if ComparePassword(userPassword, inputPassword) != nil {
		return false
	}

	return true
}
