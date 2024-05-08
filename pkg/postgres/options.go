package postgres

type ConfigDB struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}
