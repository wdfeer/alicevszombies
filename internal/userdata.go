package internal

func LoadUserData() {
	println("INFO: Loading user data...")
	loadOptions()
	loadRunHistory()
	loadHistory()
}

func SaveUserData() {
	println("INFO: Saving user data...")
	saveOptions()
	saveRunHistory()
	saveHistory()
}
