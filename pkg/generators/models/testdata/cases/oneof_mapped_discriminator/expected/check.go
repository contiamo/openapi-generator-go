package generatortest

func check() bool {

	foo := Error{}
	switch foo.Discriminator() {
	case ErrorDiscriminatorAuth:
		return true
	}
	return false
}
