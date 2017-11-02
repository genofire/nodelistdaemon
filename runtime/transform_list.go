package runtime

import (
	"encoding/json"

	yanicMeshviewer "github.com/FreifunkBremen/yanic/output/meshviewer"
	yanicMeshviewerFFRGB "github.com/FreifunkBremen/yanic/output/meshviewer-ffrgb"
	yanicNodelist "github.com/FreifunkBremen/yanic/output/nodelist"
	yanicRuntime "github.com/FreifunkBremen/yanic/runtime"
)

func transformNodelist(body []byte, sitecode string, f *Fetcher) error {
	var nodes yanicNodelist.NodeList
	err := json.Unmarshal(body, &nodes)
	if err == nil {
		for _, node := range nodes.List {
			nodeEntry := transformNodelistNode(node, sitecode)
			f.AddNode(nodeEntry)
		}
		return nil
	}
	return err
}

func transformMeshviewerV2(body []byte, sitecode string, f *Fetcher) error {
	var nodes yanicMeshviewer.NodesV2
	err := json.Unmarshal(body, &nodes)
	if err == nil {
		for _, node := range nodes.List {
			nodeEntry := transformMeshviewerNode(node, sitecode)
			f.AddNode(nodeEntry)
		}
		return nil
	}
	return err
}

func transformMeshviewerV1(body []byte, sitecode string, f *Fetcher) error {
	var nodes yanicMeshviewer.NodesV1
	err := json.Unmarshal(body, &nodes)
	if err == nil {
		for _, node := range nodes.List {
			nodeEntry := transformMeshviewerNode(node, sitecode)
			f.AddNode(nodeEntry)
		}
		return nil
	}
	return err
}

func transformMeshviewerFFRGB(body []byte, sitecode string, f *Fetcher) error {
	var nodes yanicMeshviewerFFRGB.Meshviewer
	err := json.Unmarshal(body, &nodes)
	if err == nil {
		for _, node := range nodes.Nodes {
			f.AddNode(node)
		}
		return nil
	}
	return err
}

func transformYanic(body []byte, sitecode string, f *Fetcher) error {
	var nodes yanicRuntime.Nodes
	err := json.Unmarshal(body, &nodes)
	if err == nil {
		for _, node := range nodes.List {
			nodeEntry := yanicMeshviewerFFRGB.NewNode(nil, node)
			f.AddNode(nodeEntry)
		}
		return nil
	}
	return err
}
