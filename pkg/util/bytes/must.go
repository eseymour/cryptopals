package bytes

// Must is a helper that wraps calls returning ([]byte, error), panicking if the
// error is non-nil.
func Must(data []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return data
}
