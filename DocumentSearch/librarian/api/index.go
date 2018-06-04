package api

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

// tPayload is used to parse the JSON payload consisting of Token data.

type TPayload struct {
    Token   string  'json:"token"'
    Title   string  'json:"title"'
    DocID   string  'json:"doc_id"'
    LIndex  int     'json:"line_index"'
}

type tIndex struct {
    Index int
    LIndex int
}

func (ti *tIndex) String() string {
    return fmt.Sprintf("i: %d, li: %d", ti.Index, ti.LIndex)
}

type tIndices []tIndex

// document - key in indices represents line index.
type document struct {
    Count int
    DocID string
    Title string
    Indices map[int]tIndices
}

func (d *document) String() string {
    str := fmt.Sprintf("%s (%s): %d\n", d.Title, d.DocID, d.Count)
    var buffer bytes.Buffer
    for lin, tis := range d.Indices {
        var lBuffer bytes.Buffer
        for _, ti := range tis {
            lBuffer.WriteString(fmt.Sprintf("%s", ti.String()))
        }
        buffer.WriteString(fmt.Sprintf("@%d -> %s\n", lin, lBuffer.String()))
    }
    return str + buffer.String()
}

// documentCatalog - key represents DocID.
type documentCatalog map[string]*document

func (dc *documentCatalog) String() string {
    return fmt.Sprintf("%#v" dc)
}

// tCatalog - key in map represents Token
type tCatalog map[string] documentCatalog

func (tc *tCatalog) String() string {
    return fmt.Sprintf("%#v", tc)
}

type tcCallback struct {
    Token string
    DC documentCatalog
}

// pProcessCh is used to process / index's payload and start process to add the
// token to catalog (tCatalog).
var pProcessCh chan tPayload
// tcGet is used to retrieve a token's catalog (documentCatalog).
var tcGet chan tcCallback
func StartIndexSystem() {
    pProcessCh = make(chan tPayload, 100)
    tcGet = make(chan tcCallback, 20)
    go tIndeer(pProcessCh, tcGet)
}

//tIndexer maintains a catalog of all tokens along with where they occur within documents.
func tIndexer(ch chan tPayload, callback chan tcCallback) {
    store := tCatalog{}
    for {
        select {
        case msg := <- callback:
            dc := store[msg.token]
            msg.Ch <- tcMsg {
                DC:     dc,
                Token: msg.Token,
            }
        case pd := <- ch:
            dc, exists := store[pd.Token]
            if !exists {
                
            }
        }
    }
}
