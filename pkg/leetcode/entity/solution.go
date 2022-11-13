package entity

type Solution struct {
	ProblemSlug string `json:"problemSlug"`
	Lang        string `json:"lang"`
	Code        string `json:"code"`
}

type StateRetriever interface {
	RetrieveState() string
}

type ExecuteSolutionResult struct {
	StatusCode             int           `json:"status_code"`
	Lang                   string        `json:"lang"`
	RunSuccess             bool          `json:"run_success"`
	StatusRuntime          string        `json:"status_runtime"`
	Memory                 int           `json:"memory"`
	CodeAnswer             []string      `json:"code_answer"`
	CodeOutput             []interface{} `json:"code_output"`
	StdOutputList          []string      `json:"std_output_list"`
	ElapsedTime            int           `json:"elapsed_time"`
	TaskFinishTime         int64         `json:"task_finish_time"`
	TaskName               string        `json:"task_name"`
	ExpectedStatusCode     int           `json:"expected_status_code"`
	ExpectedLang           string        `json:"expected_lang"`
	ExpectedRunSuccess     bool          `json:"expected_run_success"`
	ExpectedStatusRuntime  string        `json:"expected_status_runtime"`
	ExpectedMemory         int           `json:"expected_memory"`
	ExpectedCodeAnswer     []string      `json:"expected_code_answer"`
	ExpectedCodeOutput     []interface{} `json:"expected_code_output"`
	ExpectedStdOutputList  []string      `json:"expected_std_output_list"`
	ExpectedElapsedTime    int           `json:"expected_elapsed_time"`
	ExpectedTaskFinishTime int64         `json:"expected_task_finish_time"`
	ExpectedTaskName       string        `json:"expected_task_name"`
	CorrectAnswer          bool          `json:"correct_answer"`
	CompareResult          string        `json:"compare_result"`
	StatusMsg              string        `json:"status_msg"`
	State                  string        `json:"state"`
	FastSubmit             bool          `json:"fast_submit"`
	TotalCorrect           int           `json:"total_correct"`
	TotalTestcases         int           `json:"total_testcases"`
	SubmissionID           string        `json:"submission_id"`
	RuntimePercentile      interface{}   `json:"runtime_percentile"`
	StatusMemory           string        `json:"status_memory"`
	MemoryPercentile       interface{}   `json:"memory_percentile"`
	PrettyLang             string        `json:"pretty_lang"`
}

func (e *ExecuteSolutionResult) RetrieveState() string {
	return e.State
}

type SubmitSolutionResult struct {
	StatusCode        int         `json:"status_code"`
	Lang              string      `json:"lang"`
	RunSuccess        bool        `json:"run_success"`
	StatusRuntime     string      `json:"status_runtime"`
	Memory            int         `json:"memory"`
	QuestionID        string      `json:"question_id"`
	ElapsedTime       int         `json:"elapsed_time"`
	CompareResult     string      `json:"compare_result"`
	CodeOutput        string      `json:"code_output"`
	StdOutput         string      `json:"std_output"`
	LastTestcase      string      `json:"last_testcase"`
	ExpectedOutput    string      `json:"expected_output"`
	TaskFinishTime    int64       `json:"task_finish_time"`
	TaskName          string      `json:"task_name"`
	Finished          bool        `json:"finished"`
	StatusMsg         string      `json:"status_msg"`
	State             string      `json:"state"`
	FastSubmit        bool        `json:"fast_submit"`
	TotalCorrect      int         `json:"total_correct"`
	TotalTestcases    int         `json:"total_testcases"`
	SubmissionID      string      `json:"submission_id"`
	RuntimePercentile interface{} `json:"runtime_percentile"`
	StatusMemory      string      `json:"status_memory"`
	MemoryPercentile  interface{} `json:"memory_percentile"`
	PrettyLang        string      `json:"pretty_lang"`
	InputFormatted    string      `json:"input_formatted"`
	Input             string      `json:"input"`
}

func (e *SubmitSolutionResult) RetrieveState() string {
	return e.State
}
