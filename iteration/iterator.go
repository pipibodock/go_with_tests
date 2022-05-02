package iteration

const valueRepeated = 5
func Repeated(param string) (result string) {
	for i := 0; i < valueRepeated; i++ {
		result += param
	}
	return
}