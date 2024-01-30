package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"index/suffixarray"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	searcher := Searcher{}
	err := searcher.Load("completeworks.txt")
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/search", handleSearch(searcher))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Printf("shakesearch available at http://localhost:%s...", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Searcher struct {
	CompleteWorks string
	SuffixArray   *suffixarray.Index
}

func handleSearch(searcher Searcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		limit := 20
		offset := 0
		query, ok := r.URL.Query()["q"]
		queryObj := r.URL.Query()
		limitParam := queryObj.Get("limit")
		offsetParam := queryObj.Get("offset")
		if len(limitParam) > 0 {
			convertedInt, err := strconv.Atoi(limitParam)
			if err == nil {
				limit = convertedInt
			}
		}

		if len(offsetParam) > 0 {
			convertedInt, err := strconv.Atoi(offsetParam)
			if err == nil {
				offset = convertedInt
			}
		}

		if !ok || len(query[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}
		results := searcher.Search(query[0], limit, offset)
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err := enc.Encode(results)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("encoding failure"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf.Bytes())
	}
}

func (s *Searcher) Load(filename string) error {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Load: %w", err)
	}
	s.CompleteWorks = string(dat)
	s.SuffixArray = suffixarray.New(dat)
	return nil
}

func (s *Searcher) Search(query string, limit int, offset int) []string {
	initialQuery := query
	restOfQuery := ""
	// Handling for multi-word query
	if strings.Contains(query, " ") {
		substrings := strings.SplitAfterN(query, " ", 2)
		initialQuery = strings.TrimSpace(substrings[0])
		restOfQuery = strings.TrimSpace(substrings[1])
	}
	// Find the first word of the query in uppercase, lowercase, and capitalized forms
	// since those are likely the only valid forms
	maxNumberOfResults := limit + offset
	idxs := s.SuffixArray.Lookup([]byte(strings.ToUpper(initialQuery)), -1)
	idxs = append(idxs, s.SuffixArray.Lookup([]byte(strings.ToLower(initialQuery)), -1)...)
	// Find capitalized form of query
	r := []rune(initialQuery)
	r[0] = unicode.ToUpper(r[0])
	capitalizedQuery := string(r)
	idxs = append(idxs, s.SuffixArray.Lookup([]byte(capitalizedQuery), -1)...)

	sort.Ints(idxs)
	results := []string{}
	for _, idx := range idxs {
		fullQuery := initialQuery
		if len(restOfQuery) > 0 {
			queryParts := []string{initialQuery, restOfQuery}
			fullQuery = strings.Join(queryParts, " ")
		}
		result := s.CompleteWorks[idx:(idx + len(fullQuery))]
		if strings.Contains(strings.ToLower(result), strings.ToLower(fullQuery)) {
			beginningIndex := idx - 250
			endingIndex := idx + 250
			if beginningIndex < 0 {
				beginningIndex = 0
			}
			if endingIndex > len(s.CompleteWorks) {
				endingIndex = len(s.CompleteWorks) - 1
			}
			results = append(results, s.CompleteWorks[beginningIndex:endingIndex])
		}
		if len(results) == maxNumberOfResults {
			break
		}
	}
	return results
}
