# Distributed Document Search

### This is a small project to crawl web sources, download documents, index those documents, and allow users to search.


## Contents
1. [concierge](./concierge/) - the custodian indexers web data
2. [librarian](./librarian/) - handles user interactions and data requests

## Pre-requisites:
1. install jq:  `homebrew install jq`
### Adding Data to the Concierge server
1. start the server: `go run ./concierge/main.go`
2. Add content with a curl command: `curl -X POST -d '{"title": "Hackers: Heroes of Computer Revolution", "url": "http://www.gutenberg.org/cache/epub/729/pg729.txt"}' http://localhost:8080/api/feeder | jq`
