package worker

import (
	"../conf"
	"../settings"
	"../report"
	"fmt"
	"sync"
	"net/http"
)

type Worker struct {
	Task conf.Task
}


func NewWorker(config conf.Config, settings settings.Settings)  {
	fmt.Println("settings", settings)
	tasksCount := len(config.Tasks)

	jobs := make(chan conf.Task, tasksCount)
	results := []conf.Result{}

	wg := new(sync.WaitGroup)
	wg.Add(tasksCount)

	for w := 1; w <= settings.ThreadCount; w++ {
		go process(jobs, &results, wg, settings)
	}

	for _, task := range config.Tasks {
		jobs <- task
	}
	wg.Wait()
	report.MakeReport(results, settings.ReportPath)
	close(jobs)
}

func process(jobs <-chan conf.Task, results *[]conf.Result, wg *sync.WaitGroup, settings settings.Settings)  {
	for j := range jobs {
		fmt.Println("run", j.Url)
		*results = append(*results, checkTask(j, settings))
		wg.Done()
	}
}

func checkTask(task conf.Task, settings settings.Settings) conf.Result {
	url := task.Url
	if len(settings.Stage) > 0 {
		url = settings.Stage + task.Url
	}
	fmt.Println("checkTask", url)
	req, _ := http.NewRequest("GET", url, nil)
	transport := http.Transport{}
	resp, _ := transport.RoundTrip(req)
	result := resp.StatusCode == task.Status
	return conf.Result{result,resp.StatusCode, task}
}