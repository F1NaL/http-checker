package worker

import (
	"github.com/F1NaL/http-checker/app/conf"
	"github.com/F1NaL/http-checker/app/report"
	"github.com/F1NaL/http-checker/app/settings"
	"fmt"
	"net/http"
	"sync"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

type Worker struct {
	Task conf.Task
}

func NewWorker(config conf.Config, settings settings.Settings) int {
	fmt.Println("settings", settings)
	tasksCount := len(config.Tasks)

	jobs := make(chan conf.Task, tasksCount)
	results := []conf.Result{}

	wg := new(sync.WaitGroup)
	wg.Add(tasksCount)

	for w := 1; w <= settings.ThreadCount; w++ {
		go process(jobs, &results, wg, settings, config)
	}

	for _, task := range config.Tasks {
		jobs <- task
	}
	wg.Wait()
	report.MakeReport(results, settings.ReportPath)
	close(jobs)
	i := 0
	for _, r := range results {
		if !r.Status {
			i++
		}
	}
	return i
}

func process(jobs <-chan conf.Task, results *[]conf.Result, wg *sync.WaitGroup, settings settings.Settings, config conf.Config) {
	for j := range jobs {
		if len(config.UserAgents) > 0 {
			for _, agent := range config.UserAgents {
				*results = append(*results, checkTask(j, settings, agent))
			}
		} else {
			*results = append(*results, checkTask(j, settings, conf.UserAgent{"BaseAgent", "http-checker-client"}))
		}
		wg.Done()
	}
}

func checkTask(task conf.Task, settings settings.Settings, userAgent conf.UserAgent) conf.Result {
	url := task.Url
	if len(settings.Stage) > 0 {
		url = settings.Stage + task.Url
	}
	fmt.Println("checkTask", url, userAgent)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", userAgent.Agent)
	transport := http.Transport{}
	resp, e := transport.RoundTrip(req)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("finish tasks")
	result := resp.StatusCode == task.Status
	testsResult := make([]conf.TestResult, 0)
	if len(task.Tests) > 0 && result {
		defer resp.Body.Close()
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		for _, test := range task.Tests {
			r := checkTest(doc, test)
			if !r.Status {
				result = false
			}
			testsResult = append(testsResult, r)
		}
	}

	return conf.Result{result, resp.StatusCode, task, userAgent, testsResult}
}

func checkTest(document *goquery.Document, test conf.TaskTest) conf.TestResult {
	result := conf.TestResult{}
	result.Title = test.Title
	result.Task = test.Q + " " + test.Match + " " + test.A
	result.ExpectValue = test.A
	text := document.Find(test.Q).Text()
	result.GotValue = text
	if test.Match == conf.TaskTestContains {
		fmt.Println(conf.TaskTestContains, strings.Contains(text, test.A), text, test.A)
		result.Status = strings.Contains(text, test.A)
		return result
	}
	result.Status = (text == test.A)
	return result
}
