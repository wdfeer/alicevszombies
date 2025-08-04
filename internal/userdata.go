package internal

func LoadUserData() {
	loadOptions()
	loadRunHistory()
	loadHistory()
}

func SaveUserData() {
	saveOptions()
	saveRunHistory()
	saveHistory()
}
