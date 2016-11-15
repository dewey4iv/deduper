package deduper

type Mapper interface {
	Add(key string, option string) error
	Map() (map[string][]string, error)
}
