package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/petme/api/cmd/serverd/router"
	"github.com/petme/api/pkg/db/pg"
	"github.com/petme/api/pkg/httpserv"
	pkgerrors "github.com/pkg/errors"
)

const projectDirName = "petme"

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Exiting...")
}

func run(ctx context.Context) error {
	loadEnv()
	log.Println("Starting app initialization")
	dbOpenConnection, err := strconv.Atoi(os.Getenv("DB_POOL_MAX_OPEN_CONNS"))
	if err != nil {
		return pkgerrors.WithStack(fmt.Errorf("invalid db pool max open conns: %w", err))
	}
	dbIdleConnection, err := strconv.Atoi(os.Getenv("DB_POOL_MAX_IDLE_CONNS"))
	if err != nil {
		return pkgerrors.WithStack(fmt.Errorf("invalid db pool max idle conns: %w", err))
	}

	conn, err := pg.NewPool(os.Getenv("DB_URL"), dbOpenConnection, dbIdleConnection)
	if err != nil {
		return err
	}

	defer conn.Close()

	rtr, err := initRouter(ctx, conn)

	log.Println("App initialization completed")

	httpserv.NewServer(rtr.Handler()).Start(ctx)

	return nil
}

func initRouter(
	ctx context.Context,
	dbConn pg.BeginnerExecutor) (router.Router, error) {
	return router.New(
		ctx,
		nil,
		nil,
	), nil
}
