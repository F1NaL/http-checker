package main

import "flag"
import (
	"github.com/F1NaL/http-checker/app/conf"
	"github.com/F1NaL/http-checker/app/settings"
	"github.com/F1NaL/http-checker/app/worker"
	"fmt"
	"log"
	"os"
)

var (
	version string
	build   string
	commit  string
	docs    string
)

func main() {
	fmt.Println("start")
	settings := settings.Settings{}
	flag.IntVar(&settings.ThreadCount, "tread", 10, "thread count")
	flag.StringVar(&settings.ConfigPath, "config", "github.com/F1NaL/http-checker/config.json", "config path")
	flag.StringVar(&settings.ReportPath, "output", "github.com/F1NaL/http-checker/reports/", "report path")
	flag.StringVar(&settings.Stage, "stage", "", "stage host if need")

	flag.Parse()

	if _, err := os.Stat(settings.ConfigPath); os.IsNotExist(err) {
		log.Fatalf("ERROR: file=%s, does not exist", settings.ConfigPath)
		os.Exit(1)
	}

	printVersion()

	config := conf.ReadConfig(settings.ConfigPath)
	os.Exit(worker.NewWorker(config, settings))
}

//program build data
func printVersion() {
	fmt.Print(`
 __          __      __                                __                  __                           
/  |        /  |    /  |                              /  |                /  |                          
$$ |____   _$$ |_  _$$ |_     ______          _______ $$ |____    ______  $$ |   __   ______    ______  
$$      \ / $$   |/ $$   |   /      \        /       |$$      \  /      \ $$ |  /  | /      \  /      \ 
$$$$$$$  |$$$$$$/ $$$$$$/   /$$$$$$  |      /$$$$$$$/ $$$$$$$  | $$$$$$  |$$ |_/$$/ /$$$$$$  |/$$$$$$  |
$$ |  $$ |  $$ | __ $$ | __ $$ |  $$ |      $$ |      $$ |  $$ | /    $$ |$$   $$<  $$    $$ |$$ |  $$/ 
$$ |  $$ |  $$ |/  |$$ |/  |$$ |__$$ |      $$ \_____ $$ |  $$ |/$$$$$$$ |$$$$$$  \ $$$$$$$$/ $$ |      
$$ |  $$ |  $$  $$/ $$  $$/ $$    $$/       $$       |$$ |  $$ |$$    $$ |$$ | $$  |$$       |$$ |      
$$/   $$/    $$$$/   $$$$/  $$$$$$$/         $$$$$$$/ $$/   $$/  $$$$$$$/ $$/   $$/  $$$$$$$/ $$/       
                            $$ |                                                                        
                            $$ |                                                                        
                            $$/                                                                        
`)
	fmt.Printf("Version: %s\nBuild Time: %s\nGit Commit Hash: %s\nDocs: %s\n\n\n", version, build, commit, docs)
}
