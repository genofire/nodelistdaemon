package runtime

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/FreifunkBremen/yanic/jsontime"
	yanicMeshviewerFFRGB "github.com/FreifunkBremen/yanic/output/meshviewer-ffrgb"
)

type Fetcher struct {
	url           string
	filePath      string
	repeat        time.Duration
	directoryList map[string]string
	Timestemp     jsontime.Time                `json:"timestamp"`
	List          []*yanicMeshviewerFFRGB.Node `json:"nodes"`
	nodesMutex    sync.Mutex
	closed        bool
}

func NewFetcher(url string, repeat time.Duration, filePath string) *Fetcher {
	return &Fetcher{
		url:        url,
		repeat:     repeat,
		filePath:   filePath,
		nodesMutex: sync.Mutex{},
	}
}

func (f *Fetcher) Start() {
	timer := time.NewTimer(f.repeat)
	for !f.closed {
		select {
		case <-timer.C:
			f.Timestemp = jsontime.Now()
			f.work()
			f.save()
			f.List = []*yanicMeshviewerFFRGB.Node{}
			timer.Reset(f.repeat)
		}
	}
	timer.Stop()
}

func (f *Fetcher) Stop() {
	f.closed = true
}

func (f *Fetcher) AddNode(n *yanicMeshviewerFFRGB.Node) {
	if n != nil {
		f.nodesMutex.Lock()
		f.List = append(f.List, n)
		f.nodesMutex.Unlock()
	}
}

func (f *Fetcher) work() {
	err := JSONRequest(f.url, &f.directoryList)
	if err != nil {
		log.Fatal(err)
	}

	wgDirectory := sync.WaitGroup{}
	wgNodes := sync.WaitGroup{}
	count := 0

	for community, url := range f.directoryList {
		wgDirectory.Add(1)
		go func(community, url string) {
			defer wgDirectory.Done()
			var directory DirectoryAPI

			err := JSONRequest(url, &directory)
			if err != nil {
				return
			}
			for _, mapEntry := range directory.NodeMaps {
				if mapEntry.TechnicalType == "nodelist" || mapEntry.MapType == "list/status" {
					wgNodes.Add(1)
					count++
					go f.updateNodes(&wgNodes, community, mapEntry.URL)
				}
			}
		}(community, url)
	}

	log.Println("found", len(f.directoryList), "communities")
	wgDirectory.Wait()
	log.Println("wait for", count, "request for nodes")
	wgNodes.Wait()
	log.Println("found", len(f.List), "nodes")
}

func (f *Fetcher) updateNodes(wg *sync.WaitGroup, community string, url string) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	transform(body, community, f)
}

func (f *Fetcher) save() {
	f.nodesMutex.Lock()
	defer f.nodesMutex.Unlock()

	if f.filePath != "" {
		tmpFile := f.filePath + ".tmp"

		file, err := os.OpenFile(tmpFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Panic(err)
		}

		err = json.NewEncoder(file).Encode(f)
		if err != nil {
			log.Panic(err)
		}

		file.Close()
		if err := os.Rename(tmpFile, f.filePath); err != nil {
			log.Panic(err)
		}
	}
}
