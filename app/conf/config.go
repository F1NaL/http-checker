package conf

type Config struct {
	Tasks []Task `json:"tasks"`
	UserAgents []UserAgent `json:"useragents"`
}

type UserAgent struct {
	Title string `json:"title"`
	Agent string `json:"agent"`
}
