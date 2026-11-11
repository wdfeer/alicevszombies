package internal

import "alicevszombies/internal/log"

func LoadUserData() {
	log.Info("Loading user data...")
	loadOptions()
	loadRunHistory()
	loadHistory()
}

func SaveUserData() {
	log.Info("Saving user data...")
	saveOptions()
	saveRunHistory()
	saveHistory()
}
