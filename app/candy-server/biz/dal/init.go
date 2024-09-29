package dal

import (
	"github.com/mokaz111/candy-server/biz/dal/mongo"
)

func Init() {
	//redis.Init()
	//mysql.Init()
	mongo.Init()
}
