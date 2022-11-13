package leetcode

import (
	"errors"
	"fmt"
	"github.com/iFaceless/leetgogo-cli/leetcode/entity"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

type actionType int

const (
	actionTypeExecute actionType = iota
	actionTypeSubmit
)

var actionURLMapping = map[actionType]string{
	actionTypeSubmit:  submitURL,
	actionTypeExecute: executeURL,
}

type Client struct {
	restCli *req.Client
}

func NewClient() *Client {
	cli := req.C().
		SetUserAgent(defaultUserAgent).
		SetTimeout(15 * time.Second).
		SetCommonHeaders(map[string]string{
			"x-requested-with": "XMLHttpRequest",
			"Origin":           baseURL,
			"Cookie": strings.Join([]string{
				"csrftoken=" + defaultCsrfToken,
				"LEETCODE_SESSION=" + defaultSession,
			}, ";"),
			"x-csrftoken": defaultCsrfToken,
		}).
		DevMode()

	return &Client{
		restCli: cli,
	}
}

func (cli *Client) RandomProblem(category ProblemCategory) (*entity.Problem, error) {
	refer := "https://leetcode.cn/problemset/all/"
	body := map[string]string{
		"query": `
query problemsetRandomFilteredQuestion($categorySlug: String!, $filters: QuestionListFilterInput) {
	problemsetRandomFilteredQuestion(categorySlug: $categorySlug, filters: $filters)
}`,
		"variables": fmt.Sprintf(`{"categorySlug":"%s","filters":{}}`, category),
	}

	type Result struct {
		Data struct {
			ProblemsetRandomFilteredQuestion string `json:"problemsetRandomFilteredQuestion"`
		}
	}
	var result Result
	resp := cli.restCli.
		Post(graphqlURL).
		SetHeader("Referer", refer).
		SetBodyJsonMarshal(body).
		SetResult(&result).
		EnableDump().
		Do()
	if resp.Err != nil {
		return nil, resp.Err
	}

	return cli.ProblemBySlug(result.Data.ProblemsetRandomFilteredQuestion)
}

func (cli *Client) ProblemBySlug(slug string) (*entity.Problem, error) {
	refer := fmt.Sprintf(problemDetailURL, slug)
	body := map[string]string{
		"query": `
query getQuestionDetail($titleSlug: String!) {,
  question(titleSlug: $titleSlug) {,
    content,
    stats,
	questionId,
	title,
	difficulty,
    codeDefinition,
    sampleTestCase,
    exampleTestcases,
    enableRunCode,
    metaData,
    translatedContent,
  },
}
`,
		"variables": fmt.Sprintf(`{"titleSlug": "%s"}`, slug),
	}

	type Result struct {
		Data struct {
			Problem entity.Problem `json:"question"`
		}
	}

	var result Result

	resp := cli.restCli.
		Post(graphqlURL).
		SetHeader("Referer", refer).
		SetBodyJsonMarshal(body).
		SetResult(&result).
		EnableDump().
		Do()

	if resp.Err != nil {
		return nil, resp.Err
	}

	return &result.Data.Problem, nil
}

func (cli *Client) ProblemIDBySlug(slug string) (string, error) {
	refer := fmt.Sprintf(problemDetailURL, slug)
	body := map[string]string{
		"query": `
query getQuestionDetail($titleSlug: String!) {,
  question(titleSlug: $titleSlug) {,
	questionId,
	title,
  },
}
`,
		"variables": fmt.Sprintf(`{"titleSlug": "%s"}`, slug),
	}

	type Result struct {
		Data struct {
			Problem entity.Problem `json:"question"`
		}
	}

	var result Result

	resp := cli.restCli.
		Post(graphqlURL).
		SetHeader("Referer", refer).
		SetBodyJsonMarshal(body).
		SetResult(&result).
		EnableDump().
		Do()

	if resp.Err != nil {
		return "", resp.Err
	}

	if result.Data.Problem.ID == "" {
		return "", fmt.Errorf("problem '%s' not found", slug)
	}

	return result.Data.Problem.ID, nil
}

// FavoriteProblem 收藏到指定的收藏夹
func (cli *Client) FavoriteProblem(slug string, favoriteNames ...string) error {
	problemID, err := cli.ProblemIDBySlug(slug)
	if err != nil {
		return fmt.Errorf("failed to get problem id by slug: %s", err)
	}

	favoriteIDHashList, err := cli.favoriteIDHashList(favoriteNames...)
	if err != nil {
		return err
	}

	refer := fmt.Sprintf(problemDetailURL, slug)
	do := func(problemId, favoriteIDHash string) error {
		body := map[string]string{
			"query": `
mutation addQuestionToFavorite($favoriteIdHash: String!, $questionId: String!) {
  addQuestionToFavorite(favoriteIdHash: $favoriteIdHash, questionId: $questionId) {
    ok
    error
    favoriteIdHash
    questionId
  }
}
`,
			"variables": fmt.Sprintf(`{"favoriteIdHash": "%s", "questionId": "%s"}`, favoriteIDHash, problemId),
		}
		type Result struct {
			Data struct {
				AddQuestionToFavorite struct {
					Ok    bool   `json:"ok"`
					Error string `json:"error"`
				} `json:"addQuestionToFavorite"`
			} `json:"data"`
		}

		var result Result

		resp := cli.restCli.
			Post(graphqlURL).
			SetHeader("Referer", refer).
			SetBodyJsonMarshal(body).
			SetResult(&result).
			EnableDump().
			Do()

		if resp.Err != nil {
			return resp.Err
		}

		if result.Data.AddQuestionToFavorite.Ok {
			return nil
		}

		if result.Data.AddQuestionToFavorite.Error != "" {
			return errors.New(result.Data.AddQuestionToFavorite.Error)
		}

		return errors.New(resp.String())
	}

	for _, idHash := range favoriteIDHashList {
		err := do(problemID, idHash)
		if err != nil {
			return fmt.Errorf("failed to favorite (%s-%s): %s", problemID, idHash, err)
		}
	}

	return nil
}

// UnfavoriteProblem 从指定的收藏夹中删除收藏
func (cli *Client) UnfavoriteProblem(slug string, favoriteNames ...string) error {
	problemID, err := cli.ProblemIDBySlug(slug)
	if err != nil {
		return fmt.Errorf("failed to get problem id by slug: %s", err)
	}

	favoriteIDHashList, err := cli.favoriteIDHashList(favoriteNames...)
	if err != nil {
		return err
	}

	refer := fmt.Sprintf(problemDetailURL, slug)
	do := func(problemId, favoriteIDHash string) error {
		body := map[string]string{
			"query": `
mutation removeQuestionFromFavorite($favoriteIdHash: String!, $questionId: String!) {
  removeQuestionFromFavorite(
    favoriteIdHash: $favoriteIdHash
    questionId: $questionId
  ) {
    ok
    error
    favoriteIdHash
    questionId
  }
}
`,
			"variables": fmt.Sprintf(`{"favoriteIdHash": "%s", "questionId": "%s"}`, favoriteIDHash, problemId),
		}
		type Result struct {
			Data struct {
				RemoveQuestionFromFavorite struct {
					Ok    bool   `json:"ok"`
					Error string `json:"error"`
				} `json:"removeQuestionFromFavorite"`
			} `json:"data"`
		}

		var result Result

		resp := cli.restCli.
			Post(graphqlURL).
			SetHeader("Referer", refer).
			SetBodyJsonMarshal(body).
			SetResult(&result).
			EnableDump().
			Do()

		if resp.Err != nil {
			return resp.Err
		}

		if result.Data.RemoveQuestionFromFavorite.Ok {
			return nil
		}

		if result.Data.RemoveQuestionFromFavorite.Error != "" {
			return errors.New(result.Data.RemoveQuestionFromFavorite.Error)
		}

		return errors.New(resp.String())
	}

	for _, idHash := range favoriteIDHashList {
		err := do(problemID, idHash)
		if err != nil {
			return fmt.Errorf("failed to unfavorite (%s-%s): %s", problemID, idHash, err)
		}
	}

	return nil
}

func (cli *Client) favoriteIDHashList(favoriteNames ...string) ([]string, error) {
	if len(favoriteNames) == 0 {
		return nil, errors.New("empty favorite names")
	}

	favorites, err := cli.UserFavorites()
	if err != nil {
		return nil, fmt.Errorf("failed to get user favorites: %s", err)
	}

	findFavoriteIDHash := func(name string) (string, bool) {
		for _, x := range favorites {
			if name == x.Name {
				return x.IDHash, true
			}
		}
		return "", false
	}

	var idHashList []string
	for _, name := range favoriteNames {
		targetID, ok := findFavoriteIDHash(name)
		if !ok {
			return nil, fmt.Errorf("favorite name not found for current user: %s", name)
		}
		idHashList = append(idHashList, targetID)
	}

	return idHashList, nil
}

func (cli *Client) UserFavorites() ([]*entity.Favorite, error) {
	refer := "https://leetcode.cn/problemset/all/"
	body := map[string]string{
		"query": `
query userFavorites {
	favoritesLists {
		allFavorites {
			idHash
			name
			isPublicFavorite
		}
	}
}
`,
		"variables": "{}",
	}

	type Result struct {
		Data struct {
			FavoritesLists struct {
				AllFavorites []*entity.Favorite `json:"allFavorites"`
			} `json:"favoritesLists"`
		} `json:"data"`
	}

	var result Result

	resp := cli.restCli.
		Post(graphqlURL).
		SetHeader("Referer", refer).
		SetBodyJsonMarshal(body).
		SetResult(&result).
		EnableDump().
		Do()

	if resp.Err != nil {
		return nil, resp.Err
	}

	return result.Data.FavoritesLists.AllFavorites, nil
}

func (cli *Client) ExecuteProblemSolution(solution *entity.Solution, testCases []string) (*entity.ExecuteSolutionResult, error) {
	var result entity.ExecuteSolutionResult
	err := cli.handleProblemSolution(actionTypeExecute, solution, testCases, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *Client) SubmitProblemSolution(solution *entity.Solution) (*entity.SubmitSolutionResult, error) {
	var result entity.SubmitSolutionResult
	err := cli.handleProblemSolution(actionTypeSubmit, solution, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

var errNeedRetry = errors.New("state not success")

func (cli *Client) handleProblemSolution(actionType actionType, solution *entity.Solution, testCases []string, result interface{}) error {
	refer := fmt.Sprintf(problemDetailURL, solution.ProblemSlug)
	problemID, err := cli.ProblemIDBySlug(solution.ProblemSlug)
	if err != nil {
		return fmt.Errorf("failed to get problem id by slug: %s", err)
	}

	_, ok := supportedLangs[solution.Lang]
	if !ok {
		return fmt.Errorf("unsupported lang %s", solution.Lang)
	}

	body := map[string]string{
		"data_input":  strings.Join(testCases, "\n"),
		"lang":        solution.Lang,
		"question_id": problemID,
		"typed_code":  solution.Code,
	}

	type Result struct {
		InterpretExpectedID string `json:"interpret_expected_id"`
		InterpretID         string `json:"interpret_id"`
		TestCases           string `json:"test_case"`
		SubmissionID        int    `json:"submission_id"`
	}

	var uploadResult Result
	requestURL := fmt.Sprintf(actionURLMapping[actionType], solution.ProblemSlug)
	uploadResp := cli.restCli.
		Post(requestURL).
		SetHeader("Referer", refer).
		SetBodyJsonMarshal(body).
		SetResult(&uploadResult).
		EnableDump().
		Do()

	if uploadResp.Err != nil {
		return uploadResp.Err
	}

	var checkID string
	switch actionType {
	case actionTypeExecute:
		checkID = uploadResult.InterpretID
	case actionTypeSubmit:
		checkID = fmt.Sprintf("%d", uploadResult.SubmissionID)
	default:
		return fmt.Errorf("unsupported action type: %d", actionType)
	}
	if checkID == "" {
		return fmt.Errorf("failed to execute or submit solution: %s", uploadResp.String())
	}

	check := func() error {
		resp := cli.restCli.
			Get(fmt.Sprintf(checkResultURL, checkID)).
			SetHeader("Referer", refer).
			SetResult(&result).
			EnableDump().
			Do()
		if resp.Err != nil {
			return err
		}

		retriever, ok := result.(entity.StateRetriever)
		if !ok {
			return errors.New("cannot get state from response content")
		}
		state := retriever.RetrieveState()

		if state == "" {
			return fmt.Errorf("failed to check, state is empty, raw response: %s", resp.String())
		}

		if state != "SUCCESS" {
			return errNeedRetry
		}

		return nil
	}

	for i := 0; i < 30; i++ {
		time.Sleep(500 * time.Millisecond)
		err := check()
		if err != nil {
			if errors.Is(err, errNeedRetry) {
				continue
			}
			return err
		}
		return nil
	}

	return errors.New("timeout to check result")
}
