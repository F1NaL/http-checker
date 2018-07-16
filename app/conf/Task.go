package conf

type Task struct {
	Url    string `json:"url"`
	Status int    `json:"status"`
	Tests []TaskTest `json:"tests"`
}

type TaskTest struct {
	Title    string `json:"title"`
	Q    string `json:"q"`
	A    string `json:"a"`
	Match    string `json:"match"`
}

const TaskTestContains  = "contains"