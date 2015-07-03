package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
)

var Engine *xorm.Engine

func init() {
	revel.OnAppStart(Init)
}

func Init() {
	driver := revel.Config.StringDefault("db.driver", "mysql")
	host := revel.Config.StringDefault("db.host", "db.niukathon.io")
	port := revel.Config.IntDefault("db.port", 3306)
	database := revel.Config.StringDefault("db.database", "niukathon")
	user := revel.Config.StringDefault("db.user", "niukathon")
	password := revel.Config.StringDefault("db.pass", "niukathon")

	params := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", user, password, host, port, database)

	var err error
	Engine, err = xorm.NewEngine(driver, params)
	// defer Engine.Close()
	if err != nil {
		panic(err)
	}

	Engine.ShowSQL = true

	// err = Engine.Sync(
	// 	new(Team),
	// 	new(Event),
	// 	new(User),
	// 	new(Post),
	// )

	// if err != nil {
	// 	panic(err)
	// }

}
