package commands

var Collections map[string]string

func init() {
	Collections = make(map[string]string)

	Collections["ping"] = "ping"
}
