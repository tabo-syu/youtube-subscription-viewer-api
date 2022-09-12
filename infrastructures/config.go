package infrastructures

type DB struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	TimeZone string
}

type Youtube struct{}

type Config struct {
	DB      DB
	Youtube Youtube
}
