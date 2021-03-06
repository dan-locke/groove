package query

import (
	"github.com/hscells/cqr"
	"github.com/hscells/groove"
	"io/ioutil"
	"strconv"
)

type KeywordQuerySource struct {
	fields []string
}

// Load takes a directory of queries and parses them "as is".
func (kw KeywordQuerySource) Load(directory string) ([]groove.PipelineQuery, error) {
	// First, get a list of files in the directory.
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return []groove.PipelineQuery{}, err
	}

	// Next, load each query into a CQR keyword query.
	queries := []groove.PipelineQuery{}
	for _, f := range files {
		source, err := ioutil.ReadFile(directory + "/" + f.Name())
		if err != nil {
			return []groove.PipelineQuery{}, err
		}

		cqrQuery := cqr.Keyword{QueryString: string(source), Fields: kw.fields}

		topic, err := strconv.Atoi(f.Name())

		queries = append(queries, *groove.NewPipelineQuery(f.Name(), int64(topic), cqrQuery))
	}

	// Finally, return the queries.
	return queries, nil
}

func NewKeywordQuerySource(fields ...string) KeywordQuerySource {
	return KeywordQuerySource{fields: fields}
}
