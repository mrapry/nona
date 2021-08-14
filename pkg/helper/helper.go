package helper

import "encoding/json"

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

// ToBytes convert all types to bytes
func ToBytes(i interface{}) (b []byte) {
	switch t := i.(type) {
	case []byte:
		b = t
	case string:
		b = []byte(t)
	default:
		b, _ = json.Marshal(i)
	}
	return
}