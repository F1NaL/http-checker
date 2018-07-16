package conf

type Result struct {
	Status   bool
	GotValue int
	Task     Task
	UserAgent UserAgent
	TestResults []TestResult
}

type TestResult struct {
	Status bool
	Title string
	Task string
	ExpectValue string
	GotValue string
}
