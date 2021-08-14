package helper


// StringInSlice function for checking whether string in slice
// str string searched string
// list []string slice
func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

//CheckStringNotNull check string data is not empty
func CheckStringNotNull(data string) bool {
	if data == "" {
		return false
	}

	return true
}