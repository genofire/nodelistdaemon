package main

import (
	"flag"
	"time"

	"github.com/genofire/nodelistdaemon/runtime"
)

const DIRECTORY_URL = "https://raw.githubusercontent.com/freifunk/directory.api.freifunk.net/master/directory.json"

var url string
var filePath string

func main() {
	flag.StringVar(&url, "url", DIRECTORY_URL, "api directory")
	flag.StringVar(&filePath, "json", "/tmp/nodelist.json", "where to save json file")
	flag.Parse()
	f := runtime.NewFetcher(url, time.Minute, filePath)
	f.Start()
}
