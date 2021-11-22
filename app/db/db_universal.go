package db

type DB struct {
	postgre *PostgreSQL
	redis *Redis
}

var dbUniversal DB

func InitDB(nameDb string) *DB{
	switch nameDb {
	case "redis":
		dbUniversal.redis, _ = InitRedis()
	case "postgre":
		dbUniversal.postgre, _ = InitPostgreSQL()
	}
	return &dbUniversal
}
