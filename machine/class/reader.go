package class

type ClassReader interface {
	readClass(className string) ([]byte, error)
}
