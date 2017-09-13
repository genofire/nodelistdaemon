package runtime

import (
	"github.com/FreifunkBremen/yanic/jsontime"

	yanicMeshviewer "github.com/FreifunkBremen/yanic/output/meshviewer"
	yanicNodelist "github.com/FreifunkBremen/yanic/output/nodelist"
	yanicRuntime "github.com/FreifunkBremen/yanic/runtime"
)

func transformNodelistNode(n *yanicNodelist.Node, sitecode string) *Node {
	node := &Node{
		NodeID:    n.ID,
		Hostname:  n.Name,
		SiteCode:  sitecode,
		Firstseen: jsontime.Now(),
		Lastseen:  n.Status.LastContact,
		IsOnline:  n.Status.Online,
		Clients:   n.Status.Clients,
	}
	if pos := n.Position; pos != nil {
		node.Location = &Location{
			Lat: pos.Lat,
			Lon: pos.Long,
		}
	}
	return node
}

func transformMeshviewerNode(n *yanicMeshviewer.Node, sitecode string) *Node {
	if nodeinfo := n.Nodeinfo; nodeinfo != nil {
		node := &Node{
			NodeID:    nodeinfo.NodeID,
			Hostname:  nodeinfo.Hostname,
			SiteCode:  sitecode,
			Firstseen: n.Firstseen,
			Lastseen:  n.Lastseen,
			IsOnline:  n.Flags.Online,
			Clients:   n.Statistics.Clients,
		}
		if pos := nodeinfo.Location; pos != nil {
			node.Location = &Location{
				Lat: pos.Latitude,
				Lon: pos.Longtitude,
			}
		}
		return node
	}
	return nil
}

func transformYanicNode(n *yanicRuntime.Node, sitecode string) *Node {
	if nodeinfo := n.Nodeinfo; nodeinfo != nil {
		node := &Node{
			NodeID:    nodeinfo.NodeID,
			Hostname:  nodeinfo.Hostname,
			SiteCode:  sitecode,
			Firstseen: n.Firstseen,
			Lastseen:  n.Lastseen,
			IsOnline:  n.Online,
			Clients:   n.Statistics.Clients.Total,
		}
		if pos := nodeinfo.Location; pos != nil {
			node.Location = &Location{
				Lat: pos.Latitude,
				Lon: pos.Longtitude,
			}
		}
		return node
	}
	return nil
}
