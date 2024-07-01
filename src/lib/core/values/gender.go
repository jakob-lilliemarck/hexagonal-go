package values

import "errors"

// Define a custom type and constants
type Gender string

const (
	female Gender = "female"
	male   Gender = "male"
)

// Provide a function to safely convert strings to the custom type
func GenderFromString(repr string) (Gender, error) {
	switch repr {
	case "female":
		return female, nil
	case "male":
		return male, nil
	default:
		return "", errors.New("unknown gender")
	}
}

// Function to check if two genders can make offspring
func CanMakeOffspring(lhs, rhs Gender) bool {
	return (lhs == female && rhs == male) || (lhs == male && rhs == female)
}
