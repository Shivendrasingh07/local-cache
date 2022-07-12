package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"local-cache/config"
	"local-cache/provider"
	"local-cache/provider/configprovider"
	"local-cache/provider/localCache"
	"net/http"
	"os"
	"time"
)

const (
	defaultServerRequestTimeoutMinutes      = 2
	defaultServerReadHeaderTimeoutSeconds   = 30
	defaultServerRequestWriteTimeoutMinutes = 30
)

type Server struct {
	Router     *mux.Router
	httpServer *http.Server
	Config     provider.ConfigProvider
	LocalCache provider.LocalCacheInterface
}

func InitializeServer() *Server {
	secretConfig := configprovider.NewConfigProvider()

	localCacheProvider := localCache.NewLocalCacheProvider(secretConfig.GetString(config.ConfPath))

	return &Server{
		Config:     secretConfig,
		LocalCache: localCacheProvider,
	}
}

//func InitializeServer() {
//
//
//
//}

func (srv *Server) Start() {
	addr := GetPort()
	httpSrv := &http.Server{
		Addr:              addr,
		Handler:           srv.initializeRoutes(),
		ReadTimeout:       defaultServerRequestTimeoutMinutes * time.Minute,
		ReadHeaderTimeout: defaultServerReadHeaderTimeoutSeconds * time.Second,
		WriteTimeout:      defaultServerRequestWriteTimeoutMinutes * time.Minute,
	}
	srv.httpServer = httpSrv

	logrus.Info("Server running at PORT ", addr)
	if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Fatalf("Start %v", err)
		return
	}
	fmt.Println("test start")
}

func GetPort() string {
	var port = os.Getenv("MyPort")
	if port == "" {
		port = "5000"
	}
	return ":" + port
}
