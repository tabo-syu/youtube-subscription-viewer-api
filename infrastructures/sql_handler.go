package infrastructures

type SqlHandler struct{}

func NewSqlHandler(config *Config) *SqlHandler {
	return &SqlHandler{}
}
