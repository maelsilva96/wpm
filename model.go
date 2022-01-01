package wpm

type Config struct {
	DbConnection string
	TypeRequest  string
	Path         string
	Guid         string
}

type Temporary struct {
	Data []byte
}

type ModelBind struct {
	Body       []byte
	Headers    map[string]string
	Params     map[string]string
	Queries    map[string]string
	StatusCode int
}
