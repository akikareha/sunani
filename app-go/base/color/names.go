package color

var table = []struct {
	name  string
	color Color
}{
	{"clear", New(0, 0, 0, 0)},
	{"black", New(0, 0, 0, 255)},
	{"white", New(255, 255, 255, 255)},
	{"red", New(255, 0, 0, 255)},
	{"green", New(0, 255, 0, 255)},
	{"blue", New(0, 0, 255, 255)},
	{"cyan", New(0, 255, 255, 255)},
	{"magenta", New(255, 0, 255, 255)},
	{"yellow", New(255, 255, 0, 255)},
}

var byName = func() map[string]Color {
	m := make(map[string]Color, len(table))
	for _, e := range table {
		m[e.name] = e.color
	}
	return m
}()

func Names() []string {
	names := make([]string, 0, len(table))
	for _, e := range table {
		names = append(names, e.name)
	}
	return names
}

func ByName(name string) (Color, bool) {
	e, ok := byName[name]
	return e, ok
}
