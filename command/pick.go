package command

import (
	"github.com/iFaceless/leetgogo-cli/leetcode"
	"github.com/iFaceless/leetgogo-cli/leetcode/entity"
	"github.com/iFaceless/leetgogo-cli/workdir"
	"github.com/spf13/cobra"
	"log"
	"net/url"
	"strings"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:     "pick [problem-slug | problem-url]",
		Short:   "pick a leetcode problem by slug or url with slug, default to a random one",
		Example: "leetgogo pick two-sum",
		Args:    cobra.MaximumNArgs(1),
		Run:     handlePickCommand,
	})
}

func handlePickCommand(_ *cobra.Command, args []string) {
	var problem *entity.Problem
	if len(args) == 0 {
		_problem, err := leetcode.NewClient().RandomProblem("algorithms")
		if err != nil {
			log.Fatalf("failed to pick a random problem: %s", err)
		}
		problem = _problem
	} else if len(args) == 1 {
		problemSlug := extractProblemSlug(args[0])
		log.Println("got problem slug: " + problemSlug)
		if problemSlug == "" {
			log.Println("invalid problem slug or url")
		}

		_problem, err := leetcode.NewClient().ProblemBySlug(problemSlug)
		if err != nil {
			log.Fatalf("failed to pick problem by slug: %s", err)
		}
		problem = _problem
	}

	wd := workdir.New("/home/chris/Projects/Go/leetgogogo")
	err := wd.GenerateProblemSolutionTemplate(problem)
	if err != nil {
		log.Fatalf("failed to generate problem solution template: %s", err)
	}
}

func extractProblemSlug(in string) string {
	if in == "" {
		return ""
	}

	result, err := url.Parse(in)
	if err != nil {
		log.Fatalf("failed to parse problem slug: %s", err)
	}

	result.Path = strings.TrimRight(result.Path, "/")
	parts := strings.Split(result.Path, "/")
	if len(parts) == 0 {
		return ""
	}

	return parts[len(parts)-1]
}
