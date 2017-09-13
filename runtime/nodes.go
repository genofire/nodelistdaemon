package runtime

import "github.com/FreifunkBremen/yanic/jsontime"

type Node struct {
	Firstseen jsontime.Time `json:"firstseen"`
	Lastseen  jsontime.Time `json:"lastseen"`
	IsOnline  bool          `json:"is_online"`
	NodeID    string        `json:"node_id"`
	Hostname  string        `json:"hostname"`
	SiteCode  string        `json:"community"`
	Location  *Location     `json:"location"`
	Clients   uint32        `json:"clients"`
}

type Location struct {
	Lon float64 `json:"long"`
	Lat float64 `json:"lat"`
}
