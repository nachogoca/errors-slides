
func foo(int) (string, error)

type error interface {
	Error() string
}
