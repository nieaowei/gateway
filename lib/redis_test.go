package lib

func initT() {
	InitRedisConf("../conf/dev")
	InitRedis()
}
