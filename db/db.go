package db

import (
	"github.com/go-xorm/xorm"
	"github.com/juliotorresmoreno/teravision/config"
	"github.com/lib/pq"
)

// NewConn return new db conection
func NewConn() (*xorm.Engine, error) {
	conf := config.GetConf()
	dsn, err := parseDSN(conf.DriverName, conf.DatabaseDSN)
	if err != nil {
		return nil, err
	}
	return xorm.NewEngine(conf.DriverName, dsn)
}

func parseDSN(driver, dsn string) (string, error) {
	switch driver {
	case "postgres":
		return pq.ParseURL(dsn)
	}
	return dsn, nil
}
