package globals

// InArray 类似于php的in_array
func StringInArray(f string, arr []string) bool {
	for _, i := range arr {
		if i == f {
			return true
		}
	}
	return false
}
