package svc

import (
	"database/sql"
	"fmt"
	"ticket-booking-platform/services/user-service/internal/config"
	"ticket-booking-platform/services/user-service/logic"
	"ticket-booking-platform/services/user-service/repository"

	_ "github.com/lib/pq"
)

type ServiceContext struct {
	Config         config.Config
	DB             *sql.DB
	UserRepo       repository.UserRepository
	RegisterLogic  *logic.RegisterLogic
	LoginLogic     *logic.LoginLogic
	ProfileLogic   *logic.ProfileLogic
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host, c.Database.Port, c.Database.User,
		c.Database.Password, c.Database.DBName, c.Database.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)

	return &ServiceContext{
		Config:        c,
		DB:            db,
		UserRepo:      userRepo,
		RegisterLogic: logic.NewRegisterLogic(userRepo),
		LoginLogic:    logic.NewLoginLogic(userRepo, c.JWT.Secret),
		ProfileLogic:  logic.NewProfileLogic(userRepo),
	}
}
