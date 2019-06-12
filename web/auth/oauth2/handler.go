package oauth2

import "errors"

func passwordHandler(username, password string) (string, error) {
	if username == "fxms" && password == "fxms" {
		return "1", nil
	}
	return "", errors.New("invalid username or password")
}
