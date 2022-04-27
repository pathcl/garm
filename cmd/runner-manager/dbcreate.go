package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os/signal"
	"runner-manager/config"
	"runner-manager/database/sql"
	"runner-manager/params"
	"runner-manager/util"
)

var (
	conf    = flag.String("config", config.DefaultConfigFilePath, "runner-manager config file")
	version = flag.Bool("version", false, "prints version")
)

var Version string

func main() {
	flag.Parse()
	if *version {
		fmt.Println(Version)
		return
	}
	ctx, stop := signal.NotifyContext(context.Background(), signals...)
	defer stop()
	fmt.Println(ctx)

	cfg, err := config.NewConfig(*conf)
	if err != nil {
		log.Fatalf("Fetching config: %+v", err)
	}

	db, err := sql.NewSQLDatabase(ctx, cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	txt := "ana are mere prune și alune"

	enc, err := util.Aes256EncodeString(txt, "pamkotepAyksemfeghoibidEwCivbaut")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("encrypted: %d\n", len(enc))

	dec, err := util.Aes256DecodeString(enc, "pamkotepAyksemfeghoibidEwCivbaut")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dec)

	repo, err := db.CreateRepository(ctx, "gabriel-samfira", "scripts", "")
	if err != nil {
		log.Fatal(err)
	}

	pool, err := db.CreateRepositoryPool(ctx, repo.ID, params.CreatePoolParams{
		ProviderName:   "lxd_local",
		MaxRunners:     10,
		MinIdleRunners: 1,
		Image:          "ubuntu:20.04",
		Flavor:         "default",
		Tags: []string{
			"myrunner",
			"superAwesome",
		},
		OSType: config.Linux,
		OSArch: config.Amd64,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pool)

	pool2, err := db.CreateRepositoryPool(ctx, repo.ID, params.CreatePoolParams{
		ProviderName:   "lxd_local2",
		MaxRunners:     10,
		MinIdleRunners: 1,
		Image:          "ubuntu:20.04",
		Flavor:         "default",
		Tags: []string{
			"myrunner",
			"superAwesome2",
		},
		OSType: config.Linux,
		OSArch: config.Amd64,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pool2)

	pool3, err := db.FindRepositoryPoolByTags(ctx, repo.ID, []string{"myrunner", "superAwesome2"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pool3)
}
