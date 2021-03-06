package main

import (
	"fmt"
	"log"

	"github.com/nafiar/tkpd-recom-engine/common/config"
	"github.com/nafiar/tkpd-recom-engine/common/redis"
	apiDelivery "github.com/nafiar/tkpd-recom-engine/internal/app/delivery/web/api"
	redisUserInfoRepo "github.com/nafiar/tkpd-recom-engine/internal/app/repository/userinfo/redis"
	userDataInfoUC "github.com/nafiar/tkpd-recom-engine/internal/app/usecase/userdata/info"
)

func main() {
	cfg := config.GetConfig()

	err := config.InitConfig()
	if err != nil {
		err = fmt.Errorf("Could not init config. Err: %+v", err)
		log.Fatal(err)
	}

	cfg = config.GetConfig()

	redisUserInfo := redis.NewConnection(redis.ConnectionConfig{
		Address:   cfg.Redis["user_data"].Connection,
		MaxActive: cfg.Redis["user_data"].MaxActive,
		MaxIdle:   cfg.Redis["user_data"].MaxIdle,
	})

	// initialize repository, usecase & delivery
	userInfoRepo := redisUserInfoRepo.New(redisUserInfo)
	userDataInfoUC := userDataInfoUC.New(userInfoRepo)
	api := apiDelivery.New(userDataInfoUC)
	api.Serve()
}
