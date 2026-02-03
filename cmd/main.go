package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kautsarhasby/situs-forum/internal/configs"
	"github.com/kautsarhasby/situs-forum/internal/handlers/memberships"
	membershipRepo "github.com/kautsarhasby/situs-forum/internal/repository/memberships"
	membershipSvc "github.com/kautsarhasby/situs-forum/internal/service/memberships"
	"github.com/kautsarhasby/situs-forum/pkg/internalsql"
)

func main() {
  r := gin.Default()

  var (
    cfg *configs.Config
  )

  err:=configs.Init(
    configs.WithConfigFolder(
      []string{"./internal/configs"},
    ),
    configs.WithConfigFile("config"),
    configs.WithConfigType("yaml"),
  )

  if err != nil {
    log.Fatal("Gagal Inisiasi Config",err)
  }

  cfg = configs.Get()
  log.Println("config",cfg)

  db,err := internalsql.Connect(cfg.Database.DataSourceName)
  if err != nil {
    log.Fatal("Gagal  iniiasi database", err)
  }

  membershipRepository := membershipRepo.NewRepository(db) 
  membershipsService :=membershipSvc.NewService(membershipRepository,cfg)

  membershipHandler := memberships.NewHandler(r,membershipsService)
  membershipHandler.RegisterRoute()

  if err := r.Run(cfg.Service.Port); 
  err != nil {
    log.Fatalf("failed to run server: %v", err)
  }
}