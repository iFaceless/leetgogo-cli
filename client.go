package leetcode

import (
	"errors"
	"fmt"
	"github.com/iFaceless/leetgogo-cli/entity"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

type ExecuteType int

const (
	ExecuteTypeTest ExecuteType = iota
	ExecuteTypeSubmit
)

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

func (cli *Client) TestProblemSolution(solution *entity.Solution, testCases []string) error {
	return cli.executeProblemSolution(ExecuteTypeTest, solution, testCases)
}

func (cli *Client) SubmitProblemSolution(solution *entity.Solution) error {
	return cli.executeProblemSolution(ExecuteTypeSubmit, solution, nil)
}

func (cli *Client) executeProblemSolution(executeType ExecuteType, solution *entity.Solution, testCases []string) error {
	return nil
}
