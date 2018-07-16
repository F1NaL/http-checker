package report

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"time"
)

func MakeReport(data interface{}, outputPath string) {
	fmt.Println("report", data)
	tmpl := template.Must(template.ParseFiles("./app/report/report.html"))
	jsondata, eee := json.Marshal(data)
	tplData := ReportData{string(jsondata)}
	if eee != nil {
		fmt.Println(eee)
	}
	var tpl bytes.Buffer
	e := tmpl.Execute(&tpl, tplData)
	if e != nil {
		fmt.Println(e)
	}
	err := ioutil.WriteFile(outputPath+"report_"+time.Now().Format(time.RFC3339)+".html", tpl.Bytes(), 0644)
	if err != nil {
		fmt.Println(err)
	}
}

type ReportData struct {
	Reports string
}
