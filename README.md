<img height="200px" src="gopher.png" alt="gopher" align="right"/>

# groove

[![GoDoc](https://godoc.org/github.com/hscells/groove?status.svg)](https://godoc.org/github.com/hscells/groove)
[![Go Report Card](https://goreportcard.com/badge/github.com/hscells/groove)](https://goreportcard.com/report/github.com/hscells/groove)
[![gocover](http://gocover.io/_badge/github.com/hscells/groove)](https://gocover.io/github.com/hscells/groove)

_Query analysis pipeline framework_

groove is a library for pipeline construction for query analysis. The groove pipeline comprises a query source (the
format of the queries), a statistic source (a source for computing information retrieval statistics), preprocessing
steps, any measurements to make, and any output formats.

The groove library is primarily used in [boogie](https://github.com/hscells/boogie) which is a front-end DSL for groove.
If using groove as a Go library, refer to the simple example below which loads Medline queries and analyses them using
Elasticsearch and finally outputs the result into a JSON file.

## API Usage

```go
// Construct the pipeline.
p := NewGroovePipeline(
    query.NewTransmuteQuerySource(query.MedlineTransmutePipeline),
    stats.NewElasticsearchStatisticsSource(stats.ElasticsearchHosts("http://example.com:9200"),
                                           stats.ElasticsearchIndex("pubmed"),
                                           stats.ElasticsearchField("abstract"),
                                           stats.ElasticsearchDocumentType("doc")),
    Preprocess(preprocess.AlphaNum, preprocess.Lowercase),
    Measurement(analysis.TermCount{}, preqpp.AvgIDF{}),
    Output(output.JsonFormatter),
)

// Execute it on a directory of queries.
s, err := p.Execute("./medline")
if err != nil {
    log.Fatal(err)
}

// Print the result.
log.Println(s[0])
```

## Logo

The Go gopher was created by [Renee French](https://reneefrench.blogspot.com/), licensed under
[Creative Commons 3.0 Attributions license](https://creativecommons.org/licenses/by/3.0/).