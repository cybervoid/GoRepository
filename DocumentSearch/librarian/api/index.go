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
}
