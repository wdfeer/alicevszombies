package log

func Error(msg ...any) {
	println("ERROR:", msg)
}

func Warn(msg ...any) {
	println("WARNING:", msg)
}

func Info(msg ...any) {
	println("INFO:", msg)
}
