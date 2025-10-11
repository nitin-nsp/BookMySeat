package svc

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"ticket-booking-platform/services/user-service/internal/config"
	"ticket-booking-platform/services/user-service/internal/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel *model.UserModel
	Redis     *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initDB(c)
	rds := initRedis(c)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(db),
		Redis:     rds,
	}
}

func initDB(c config.Config) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Postgres.Host, c.Postgres.Port, c.Postgres.User, c.Postgres.Password, c.Postgres.Database)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(c.Postgres.MaxOpenConns)
	db.SetMaxIdleConns(c.Postgres.MaxIdleConns)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}

func initRedis(c config.Config) *redis.Redis {
	return redis.MustNewRedis(redis.RedisConf{
		Host: c.Redis.Host,
		Type: c.Redis.Type,
		Pass: c.Redis.Pass,
	})
}
