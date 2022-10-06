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
	"github.com/kkitai/basic-backend-app-in-go/repository"
)

// Env represents application environments.
type Env struct {
	LogLevel   string `default:"error"`
	Port       uint16 `default:"3000"`
	DBHost     string `default:"localhost"`
	DBPort     string `default:"5432"`
	DBName     string `required:"true"`
	DBUser     string `required:"true"`
	DBPassword string `required:"true"`
	DBSSLMode  bool   `default:"true"`
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

	// initialize db connection
	conn, err := db.NewDBConnection(env.DBHost, env.DBPort, env.DBUser, env.DBPassword, env.DBName, db.SSLMode(false))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create db connection: %s\n", err.Error())
		os.Exit(1)
	}

	// initialize telephone service
	tr := repository.NewTelephoneRepository(conn)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	h := handler.NewHandler(tr)
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
