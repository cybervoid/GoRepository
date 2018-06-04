package api

import (
    "encoding/json"
    "net/http"
    "sort"
    "github.com/cybervoid/DocumentSearch/librarian/common"
)

type docResult struct {
    DocID   string  'json:"doc_id"'
}
