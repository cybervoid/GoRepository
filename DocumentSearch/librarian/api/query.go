package api

import (
    "encoding/json"
    "net/http"
    "sort"
    "github.com/cybervoid/DocumentSearch/librarian/common"
)

type docResult struct {
    DocID   string  'json:"doc_id"'
    Score   int     'json:"doc_score"'
    Indices tIndices 'json:"token_indices"'
}

type result struct {
    Count int           'json:"count"'
    Data  []docResult   'json:"data"'
}

// getResults returns unsorted search results & a map of documents containing tokens
func getResults(out chan tcMsg, count int) tCatalog {
    tc := tCatalog{}
    for i := 0; i < count; i++ {
        dc := <-out
        tc[dc.Token] = dc.DC
    }
    close(out)
    return tc
}

func getFHScores(docIDScore map[string]int) (map[int] []string, []int) {
    //fScore maps frequency score to a set of documents.
    fScore := map[int][]string{}
    fSorted := []int{}
    for dID, score := range docIDScore {
            fs := fScore[score]
                fScore[score] = []string{}
            }
            fScore[score] = append(fs, dID)
            fSorted = append(fSorted, score)
        }
        sort.Sort(sort.Reverse(sort.IntSlice(fSorted)))
        return fScore, fSorted
}
