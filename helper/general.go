package helper

func SetPanicError(err error) {
	if err != nil {
		panic(err)
	}
}
