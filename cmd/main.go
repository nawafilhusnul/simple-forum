package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nawafilhusnul/forum/internal/configs"
	"github.com/nawafilhusnul/forum/internal/handler/memberships"
	"github.com/nawafilhusnul/forum/internal/handler/posts"
	membershipsRepo "github.com/nawafilhusnul/forum/internal/repository/memberships"
	postsRepo "github.com/nawafilhusnul/forum/internal/repository/posts"
	membershipsSvc "github.com/nawafilhusnul/forum/internal/service/memberships"
	postsSvc "github.com/nawafilhusnul/forum/internal/service/posts"
	"github.com/nawafilhusnul/forum/pkg/internalsql"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolders([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatalf("failed to initialize config: %v\n", err)
	}

	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database: %+v\n", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipsRepo := membershipsRepo.NewRepository(db)
	membershipsSvc := membershipsSvc.NewService(membershipsRepo, cfg)
	membershipsHandler := memberships.NewHandler(r, membershipsSvc)
	membershipsHandler.RegisterRoutes()

	postsRepo := postsRepo.NewRepository(db)
	postsSvc := postsSvc.NewService(postsRepo, cfg)
	postsHandler := posts.NewHandler(r, postsSvc)
	postsHandler.RegisterRoutes()

	r.Run(cfg.Service.Port)
}
