package class

type ClassLoader interface {
	readClass(path string) ([]byte, error)
}
