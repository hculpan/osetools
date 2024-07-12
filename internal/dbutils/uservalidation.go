package dbutils

import (
	"context"
)

// This is a stub function, just to get the basic project
// working. This should be replaced with actual code to
// retrieve the user's hashed password from the datastore
func RetrieveUserHash(username string) (string, error) {
	user, err := queries.GetUser(context.Background(), username)
	if err != nil {
		return "", err
	}
	return user.Password, nil
}
