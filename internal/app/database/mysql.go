package database

import "gitlab.com/beabys/quetzal"

type Mysql struct {
	*quetzal.Mysql
}

type ResultQuery struct {
	Error        error
	RowsAffected int64
}

func NewMysql(c *quetzal.MysqlConfig) *Mysql {
	database := quetzal.NewMysql().SetConfigs(c)
	return &Mysql{database}
}
