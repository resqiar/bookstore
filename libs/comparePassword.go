package libs

import "golang.org/x/crypto/bcrypt"

func ComparePassword(hashedPassword string, plainPassword string) bool {
	// Convert the string into byte slice (bcrypt designed to work with byte slice)
	hashedPasswordBytes := []byte(hashedPassword)
	plainPasswordBytes := []byte(plainPassword)

	compareErr := bcrypt.CompareHashAndPassword(hashedPasswordBytes, plainPasswordBytes)

	// if compare process match, err should be nil.
	// otherwise it will return error.
	return compareErr == nil
}
