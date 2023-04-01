package random

import "errors"

func CheckName(name string) bool {
	if name != "Dan" && name != "Jeff"{
		return false
	}
	
	return true
}

func GetFullName(firstName string, lastName string) (string, error) {
	if firstName == "" && lastName == "" {
		return "", errors.New("missing firstName and lastName")
	} else if firstName == "" && lastName != ""{
		return "", errors.New("missing firstName")
	} else if lastName == "" && firstName != ""{
		return "", errors.New("missing lastName")
	} else {
		return firstName + " " + lastName, nil
	}
}
