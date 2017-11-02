package runtime

import (
	"github.com/FreifunkBremen/yanic/jsontime"

	yanicMeshviewer "github.com/FreifunkBremen/yanic/output/meshviewer"
	yanicMeshviewerFFRGB "github.com/FreifunkBremen/yanic/output/meshviewer-ffrgb"
	yanicNodelist "github.com/FreifunkBremen/yanic/output/nodelist"
)

func transformNodelistNode(n *yanicNodelist.Node, sitecode string) *yanicMeshviewerFFRGB.Node {
	node := &yanicMeshviewerFFRGB.Node{
		NodeID:    n.ID,
		Hostname:  n.Name,
		SiteCode:  sitecode,
		Firstseen: jsontime.Now(),
		Lastseen:  n.Status.LastContact,
		IsOnline:  n.Status.Online,
		Clients:   n.Status.Clients,
	}
	if pos := n.Position; pos != nil {
		node.Location = &yanicMeshviewerFFRGB.Location{
			Latitude:   pos.Lat,
			Longtitude: pos.Long,
		}
	}
	return node
}

func transformMeshviewerNode(n *yanicMeshviewer.Node, sitecode string) *yanicMeshviewerFFRGB.Node {
	if nodeinfo := n.Nodeinfo; nodeinfo != nil {
		node := &yanicMeshviewerFFRGB.Node{
			NodeID:    nodeinfo.NodeID,
			Hostname:  nodeinfo.Hostname,
			SiteCode:  sitecode,
			Firstseen: n.Firstseen,
			Lastseen:  n.Lastseen,
			IsOnline:  n.Flags.Online,
			Clients:   n.Statistics.Clients,
		}
		if pos := nodeinfo.Location; pos != nil {
			node.Location = &yanicMeshviewerFFRGB.Location{
				Latitude:   pos.Latitude,
				Longtitude: pos.Longtitude,
			}
		}
		return node
	}
	return nil
}
