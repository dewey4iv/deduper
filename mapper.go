package deduper

type Mapper interface {
	Add(key string, option string)
	Map() (map[string][]string, error)
}
