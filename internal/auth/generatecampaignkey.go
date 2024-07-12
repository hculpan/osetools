package auth

import "math/rand/v2"

const KEY_SIZE int = 24

func GenerateCampaignKey() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	result := make([]byte, KEY_SIZE)
	for i := range result {
		result[i] = letters[rand.IntN(len(letters))]
	}
	return string(result)
}
