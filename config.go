package leetcode

const (
	baseURL          = "https://leetcode.cn"
	graphqlURL       = "https://leetcode.cn/graphql"
	problemsURL      = "https://leetcode.cn/api/problems/$category/"
	problemDetailURL = "https://leetcode.cn/problems/%s/description/"
	executeURL       = "https://leetcode.cn/problems/%s/interpret_solution/"
	checkResultURL   = "https://leetcode.cn/submissions/detail/%s/check/"
	submitURL        = "https://leetcode.cn/problems/%s/submit/"
)

var supportedLangs = map[string]struct{}{
	"bash":       {},
	"c":          {},
	"cpp":        {},
	"csharp":     {},
	"golang":     {},
	"java":       {},
	"javascript": {},
	"kotlin":     {},
	"mysql":      {},
	"php":        {},
	"python":     {},
	"python3":    {},
	"ruby":       {},
	"rust":       {},
	"scala":      {},
	"swift":      {},
}

type ProblemCategory string

const (
	ProblemCategoryAll                         = ""
	ProblemCategoryAlgorithms  ProblemCategory = "algorithms"
	ProblemCategoryConcurrency ProblemCategory = "concurrency"
	ProblemCategoryDatabase    ProblemCategory = "database"
	ProblemCategoryShell       ProblemCategory = "shell"
)

const (
	defaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"
	defaultCsrfToken = "I42YeoigwdpI4oHHiLlYKUjcJlbuk5ZTm5udiRiehwaUuG6b3LMEXbQpQU6yyjR7"
	defaultSession   = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiNDg2ODgxIiwiX2F1dGhfdXNlcl9iYWNrZW5kIjoiZGphbmdvLmNvbnRyaWIuYXV0aC5iYWNrZW5kcy5Nb2RlbEJhY2tlbmQiLCJfYXV0aF91c2VyX2hhc2giOiIwOTlmMzFkZTkwMGE5N2VhODJlOTM4MmMwYjRlMTMxNTJmY2RkMjRlNzc1NWI4YTBkNzc2ODI4NTRkMTFkOGRkIiwiaWQiOjQ4Njg4MSwiZW1haWwiOiIiLCJ1c2VybmFtZSI6ImlmYWNlbGVzcyIsInVzZXJfc2x1ZyI6ImlmYWNlbGVzcyIsImF2YXRhciI6Imh0dHBzOi8vYXNzZXRzLmxlZXRjb2RlLmNuL2FsaXl1bi1sYy11cGxvYWQvdXNlcnMvaWZhY2VsZXNzL2F2YXRhcl8xNTU0OTk0MTQxLnBuZyIsInBob25lX3ZlcmlmaWVkIjp0cnVlLCJfdGltZXN0YW1wIjoxNjY2Nzk2MzgyLjU4NjI0MjIsImV4cGlyZWRfdGltZV8iOjE2NjkzMTY0MDAsInZlcnNpb25fa2V5XyI6MCwibGF0ZXN0X3RpbWVzdGFtcF8iOjE2NjgwMDAxNTh9.jsapqlQ7tUcAbO19YIpjchqeVfilALyK_sNzY_N2omk"
)
