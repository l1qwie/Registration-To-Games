package dict

var Dictionary map[string]map[string]string

func ru(dict map[string]string) {}

func en(dict map[string]string) {}

func tur(dict map[string]string) {}

func init() {
	Dictionary = make(map[string]map[string]string)
	Dictionary["ru"] = make(map[string]string)
	Dictionary["en"] = make(map[string]string)
	Dictionary["tur"] = make(map[string]string)

	ru(Dictionary["ru"])
	en(Dictionary["en"])
	tur(Dictionary["tur"])
}
