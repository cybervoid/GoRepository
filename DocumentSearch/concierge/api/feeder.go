
package api

import (
    "crypto/sha1"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "time"
    "../common"
)

type payload struct {
    URL     string 'json:"url"'
    Title   string 'json:"title"'
}
type document struct {
    Doc   string 'json:"-"'
    Title string 'json:"title"'
    DocID string 'json:"DocID"'
}
type token struct {
    Line   string 'json:"-"'
    Token  string 'json:"token"'
    Title  string 'json:"title"'
    DocID  string 'json:"doc_id"'
    LIndex int    'json:"line_index"'
    Index  int    'json:"token_index"'
}
type dMsg struct {
    DocID string
    Ch chan document
}
type lMsg struct {
    LIndex  int
    DocID   string
}
type lMeta struct {
    LIndex  int
    DocID   string
    Line    string
}
type dAllMsg struct {
    Ch chan []document
}

// done signals all listening goroutines to stop
var done chan bool
// dGetCh is used to retrieve a single document from store.
var dGetCh chan dMsg
// lGetCh is used to retrieve a single line from store.
var lGetCh chan lMsg
// lStoreCh is used to put a line into store
var lStoreCh chan lMeta

// iAddCh is used to add token to index (librarian).
var iAddCh chan token
// dStoreCh is used to put a document into store.
var dStoreCh chan document
// dProcessCh is used to process a document and convert it to tokens
var dProcessCh chan document
// dGetAllCh is used to retrieve all documents in store.
var dGetAllCh chan dAllMsg
// pProcessCh is used to process the /feeder's payload and start the indexing process.
var pProcessCh chan payload

// StartFeederSystem() initializes all channels and starts all goroutines
// we are using a standard function instead of 'init()'
// because we don't want the channels & goroutines to be initialized during testing
// Unless explicitly required by a particular test.
func StartFeederSystem() {
    done = make(chan bool)

    dGetCh = make(chan dMsg, 8)
    dGetAllCh = make(chan dAllMsg)

    iAddCh = make(chan payload, 8)
    pProcessCh = make(chan payload, 8)
    dStoreCh = make(chan document, 8)
    dProcessCh = make(chan document, 8)
    lGetCh = make(chan lMSg)
    lStoreCh = make(chan lMeta, 8)

    for i := 0; i < 4; i++ {
        go indexAdder(iAddCh, done)
        go docProcessor(pProcessCh, dStoreCh, dProcessCh, done)
        go indexProcess(dProcessCh, lStoreCh, iAddCh, done)
        go docStore(dStoreCh, dGetCh, dGetAllCh, done)
        go lineStore(lStoreCh, lGetCh, done)
    }
}

// indexAdder adds token to index (Librarian).
func indexAdder(ch chan token, done chan bool) {
    for {
        select {
        case tok := <-ch:
            fmt.Println("adding to librarian:", tok.Token)
        case <- done:
            common.Log("Exiting indexAdder.")
            return
        }
    }
}
// lineStore maintains a catalog of all lines for all documents being indexed.
func lineStore(ch chan lMeta, callback chan lMsg, done chan bool) {
    store := map[string]string{}
    for {
        select {
        case line := <-ch:
            id := fmt.Sprintf("%s-%d", line.DocID, line.LIndex)
            store[id] = line.Line

        case ch := <-callback:
            line := ""
            id := fmt.Sprintf("%s-%d", ch.DocID, ch.LIndex)
            if l, exists := store[id]; exists {
                line = l
            }
            ch.Ch <- line
        case <-done:
            common.Log("Exiting docStore.")
            return
        }
    }
}
