package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/kkitai/basic-backend-app-in-go/db"
	handler "github.com/kkitai/basic-backend-app-in-go/http"
)

type Env struct {
	// TODO: implement passing a logger to handler
	LogLevel   string `default:"error"`
	Port       uint16 `default:"3000"`
	DBHost     string
	DBPort     string `default:"5432"`
	DBName     string
	DBUser     string
	DBPassword string
}

// @title          Basic Back-end REST APP in go
// @version        1.0
// @description    sample implementation of back-end rest api written in go.
// @contact.name   kkitai
// @contact.url    https://github.com/kkitai
// @contact.email  nmgys043@gmail.com
// @license.name   MIT
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html
// @host           localhost:3000
func main() {
	var env Env
	if err := envconfig.Process("myapp", &env); err != nil {
		fmt.Fprintf(os.Stderr, "failed to load environment variables: %s\n", err.Error())
		os.Exit(1)
	}

	db, err := db.NewDB(env.DBHost, env.DBPort, env.DBUser, env.DBPassword, env.DBName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create db connection: %s\n", err.Error())
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	h := handler.NewHandler(db)
	server := &http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(env.Port), 10),
		Handler: h,
	}

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()

	fmt.Fprintf(os.Stdout, "start receiving at :%d\n", env.Port)
	fmt.Fprintln(os.Stderr, server.ListenAndServe())
}