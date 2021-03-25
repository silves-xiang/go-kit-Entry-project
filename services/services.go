package services

type Services interface {
	Concat(a,b string) (c string)
	Diff(a,b string) (c string)
	Health() (a bool)
}
type ServiceA struct {}
func (s ServiceA) Concat (a , b string) (c string) {
	return a+b
}
func (s ServiceA) Diff (a , b string) (c string) {
	if a == b {
		return "the same"
	}
	return  "not same"
}
func (s ServiceA) Health () (a bool) {
	return true
}