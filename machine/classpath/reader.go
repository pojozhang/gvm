package classpath

type ClassReader interface {
	readClass(className string) ([]byte, error)
}
