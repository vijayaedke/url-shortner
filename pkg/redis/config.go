package redis

type RedisConfig struct {
	hostname string
	port     int
	database int
	password string
}
