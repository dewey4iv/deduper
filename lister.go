package deduper

type Lister interface {
	Add(option string) error
	List() ([]string, error)
}
