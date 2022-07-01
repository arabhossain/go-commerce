package configs

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	dbUser     string
	dbPswd     string
	dbProtocol string
	dbHost     string
	dbPort     string
	dbName     string
	testDBHost string
	testDBName string
	apiPort    string
	migrate    string
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.dbUser, "dbuser", os.Getenv("MYSQL_USER"), "DB user name")
	flag.StringVar(&conf.dbPswd, "dbpswd", os.Getenv("MYSQL_PASSWORD"), "DB pass")
	flag.StringVar(&conf.dbPort, "dbport", os.Getenv("MYSQL_PORT"), "DB port")
	flag.StringVar(&conf.dbProtocol, "dbProtocol", os.Getenv("MYSQL_PROTOCOL"), "DB Protocol")
	flag.StringVar(&conf.dbHost, "dbhost", os.Getenv("MYSQL_HOST"), "DB host")
	flag.StringVar(&conf.dbName, "dbname", os.Getenv("MYSQL_DB"), "DB name")
	flag.StringVar(&conf.testDBHost, "testdbhost", os.Getenv("TEST_DB_HOST"), "test database host")
	flag.StringVar(&conf.testDBName, "testdbname", os.Getenv("TEST_DB_NAME"), "test database name")
	flag.StringVar(&conf.apiPort, "apiPort", os.Getenv("API_PORT"), "API Port")
	flag.StringVar(&conf.migrate, "migrate", os.Getenv("DB_MIGRATE_TYPE"), "specify if we should be migrating DB 'up' or 'down'")

	flag.Parse()

	return conf
}

func (c *Config) GetDBConnStr() string {
	return c.getDBConnStr(c.dbName)
}

func (c *Config) GetTestDBConnStr() string {
	return c.getDBConnStr(c.testDBName)
}

func (c *Config) getDBConnStr(dbname string) string {
	return fmt.Sprintf(
		"%s:%s@/%s?multiStatements=true",//"%s:%s@%s(%s:%s)/%s?multiStatements=true",
		c.dbUser,
		c.dbPswd,
		//c.dbProtocol,
		//c.dbHost,
		//c.dbPort,
		dbname,
	)
}

func (c *Config) GetAPIPort() string {
	return ":" + c.apiPort
}

func (c *Config) GetMigration() string {
	return c.migrate
}
