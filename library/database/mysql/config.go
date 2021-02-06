package mysql

type Config struct {
	Host      string
	Port      int
	Username  string
	Password  string
	Database  string
	Verbose   bool
	SingTable bool
}
