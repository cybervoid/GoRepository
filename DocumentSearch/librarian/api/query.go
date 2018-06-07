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

func getFScores(docIDScore map[string]int) (map[int] []string, []int) {
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

func getDocMaps(tc tCatalog) (map[string]int, map[string]tIndices) {
    // docIDScore maps DocIDs to occurrences of all tokens.
    // key: DocID
    // value: sum of all occurrences of tokens so far.
    docIDScore := map[string]int{}
    docIndices := map[string]tIndices{}
    // for each token's catalog
    for _, dc := range tc {
        // for each document registered under the token
        for dId, doc := range dc {
            // add to docID score
            var tokIndices tIndices
            for _, tList := range doc.Indices {
                tokIndices = append(tokIndices, tList...)
            }
            docIDScore[dID] += doc.Count
            dti := docIndices[dID]
            docIndices[dID] = append(dti, tokIndices...)
        }
    }
    return docIDScore, docIndices
}

func sortResults(tc tCatalog) []docResult {
    docIDScore, docIndices := getDocMaps(tc)
    fScore, fSorted := gFScores(docIDScore)

    results := []docResult{}
    addedDocs := map[string]bool{}
    for _, score := range fSorted {
        for _, docID := range fScore[score] {
            if _, exists := addedDocs[docID]; exists {
                continue
            }
            results = append(results, docResult {
                DocID:  docID,
                Score:  score,
                Indices:    docIndices[docID],
            })
            addedDoc[docID] = false
        }
    }
    return results
}

// getSearchResults returns a list of documents
// they are listed in the descending order of occurences.
func getSearchResults(sts []string) []docResult {
    callback := make(chan tcMsg)

    for _, st := range sts {
        go func(term string) {
            tcGet <- tcCallback {
                Token:  term,
                Ch: callback,
            }
        }(st)
    }
}
