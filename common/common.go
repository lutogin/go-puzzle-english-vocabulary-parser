package common

func CriticalErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
