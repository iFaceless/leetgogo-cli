package entity

import "encoding/json"

type Problem struct {
	ID          string `json:"questionId"`
	Title       string `json:"title"`
	Difficulty  string `json:"difficulty"`
	Content     string `json:"content"`
	RawStats    string `json:"stats"`
	RawCodeDefs string `json:"codeDefinition"`
	Case        string `json:"sampleTestCase"`
	AllCases    string `json:"exampleTestcases"`
	RawMetaData string `json:"metaData"`
}

func (p *Problem) Stats() (s *Stats, err error) {
	if p.RawStats == "" {
		return nil, nil
	}
	err = json.Unmarshal([]byte(p.RawStats), &s)
	return
}

func (p *Problem) CodeDefs() (defs []CodeDefinition, err error) {
	if p.RawCodeDefs == "" {
		return nil, nil
	}
	err = json.Unmarshal([]byte(p.RawCodeDefs), &defs)
	return
}

func (p *Problem) MetaData() (m *MetaData, err error) {
	if p.RawMetaData == "" {
		return nil, err
	}
	err = json.Unmarshal([]byte(p.RawMetaData), &m)
	return
}

type Stats struct {
	TotalAccepted      string `json:"totalAccepted"`
	TotalSubmission    string `json:"totalSubmission"`
	TotalAcceptedRaw   int    `json:"totalAcceptedRaw"`
	TotalSubmissionRaw int    `json:"totalSubmissionRaw"`
	AcRate             string `json:"acRate"`
}

type CodeDefinition struct {
	Value string `json:"value"`
	Text  string `json:"text"`
	Code  string `json:"defaultCode"`
}

type MetaData struct {
	Name   string  `json:"name,omitempty"`
	Params []Param `json:"params"`
	Return Return  `json:"return"`
}

type Param struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Return struct {
	Type string `json:"type"`
}
