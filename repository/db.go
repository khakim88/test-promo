package repository

type DBConfiguration struct {
	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	DBOptions         string
	MaxConnection     int
	MaxIdleConnection int
}

type PubSubConfiguration struct {
	ProjectID   string
	Credentials string
}