package helper

func EmptyStringToNull(str *string) *string {
	if str == nil || *str == "" {
		return nil
	}
	return str
}
