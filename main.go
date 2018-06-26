package main

import "flag"
import (
	"./app/settings"
	"./app/conf"
	"./app/worker"
	"fmt"
	"os"
	"log"
)

func main()  {
	fmt.Println("start");
	settings := settings.Settings{}
	flag.IntVar(&settings.ThreadCount, "tread", 10, "thread count")
	flag.StringVar(&settings.ConfigPath, "config", "./config.json", "config path")
	flag.StringVar(&settings.ReportPath, "output", "./reports/", "report path")
	flag.StringVar(&settings.Stage, "stage", "", "stage host if need")

	flag.Parse()

	if _, err := os.Stat(settings.ConfigPath); os.IsNotExist(err) {
		log.Fatalf("ERROR: file=%s, does not exist", settings.ConfigPath)
		os.Exit(1)
	}

	config := conf.ReadConfig(settings.ConfigPath)
	worker.NewWorker(config, settings)
}