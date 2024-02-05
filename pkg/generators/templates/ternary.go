package templates

// Ternary returns the first value if the last value is true, otherwise returns the second value.
func Ternary(vt any, vf any, v bool) any {
	if v {
		return vt
	}
	return vf
}
