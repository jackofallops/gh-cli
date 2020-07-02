package utils

func DerefStringSafely(input *string) string {
	if input == nil {
		return ""
	}
	return *input
}
