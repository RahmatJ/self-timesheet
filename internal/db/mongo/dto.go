package mongo

type (
	config struct {
		hostName string
		port     int64
		userName string
		password string
		dbName   string
		timeout  int64
	}
)
