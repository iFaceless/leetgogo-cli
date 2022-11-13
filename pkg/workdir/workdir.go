package workdir

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/iFaceless/leetgogo-cli/pkg/leetcode/entity"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Metadata struct {
	filename             string
	SolvedProblemCount   int       `json:"solved_problem_count"`
	FavoriteProblemCount int       `json:"favorite_problem_count"`
	CreatedAt            time.Time `json:"created_at"`
	ModifiedAt           time.Time `json:"modified_at"`
}

type Workdir struct {
	dirname      string
	solutionsDir string
	readmeFile   string
	metadata     *Metadata
}

func New(dir string) *Workdir {
	return &Workdir{
		dirname:      dir,
		solutionsDir: filepath.Join(dir, "solutions"),
		readmeFile:   filepath.Join(dir, "README.md"),
		metadata:     loadMetadata(filepath.Join(dir, ".leetgogo")),
	}
}

func (wd *Workdir) Init() error {
	if wd.IsInitialized() {
		return errors.New("workdir is already initialized")
	}

	return wd.init()
}

func (wd *Workdir) init() error {
	err := os.MkdirAll(wd.solutionsDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to mkdir `solutions`: %s", err)
	}

	readmeInitialContent := strings.ReplaceAll(readme, "{{SolvedProblemsCount}}", "0")
	err = os.WriteFile(wd.readmeFile, []byte(readmeInitialContent), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to init readme file: %s", err)
	}

	return dumpMetadata(wd.metadata)
}

func (wd *Workdir) IsInitialized() bool {
	if _, err := os.Stat(wd.metadata.filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func (wd *Workdir) GenerateProblemSolutionTemplate(problem *entity.Problem) error {
	if !wd.IsInitialized() {
		return fmt.Errorf("current workdir is not initialized")
	}

	if problem == nil {
		return errors.New("invalid problem")
	}

	problemDir := filepath.Join(wd.solutionsDir, problem.ID)
	_ = os.MkdirAll(problemDir, os.ModePerm)

	mainFilename := filepath.Join(problemDir, "main.go")
	_ = os.WriteFile(mainFilename, []byte(main), os.ModePerm)

	solutionDir := filepath.Join(problemDir, "solution")
	_ = os.MkdirAll(solutionDir, os.ModePerm)

	docFilename := filepath.Join(solutionDir, "doc.go")
	_ = os.WriteFile(docFilename, []byte(doc), os.ModePerm)

	solutionFilename := filepath.Join(solutionDir, "solution.go")
	_ = os.WriteFile(solutionFilename, []byte(problem.RawCodeDefs), os.ModePerm)

	solutionTestFilename := filepath.Join(solutionDir, "solution_test.go")
	_ = os.WriteFile(solutionTestFilename, []byte("package solution"), os.ModePerm)

	return nil
}

func loadMetadata(filename string) *Metadata {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return &Metadata{
			filename:             filename,
			SolvedProblemCount:   0,
			FavoriteProblemCount: 0,
			CreatedAt:            time.Now(),
			ModifiedAt:           time.Now(),
		}
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln("failed to load existed workdir metadata file, maybe broken: " + err.Error())
	}

	var m Metadata
	err = json.Unmarshal(content, &m)
	if err != nil {
		log.Fatalln("failed to unmarshal workdir metadata content, maybe broken: " + err.Error())
	}

	m.filename = filename

	return &m
}

func dumpMetadata(meta *Metadata) error {
	if meta == nil || meta.filename == "" {
		return fmt.Errorf("invalid workdir metadata")
	}

	content, _ := json.Marshal(meta)
	err := os.WriteFile(meta.filename, content, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to dump workdir metadata")
	}

	return nil
}
