package runtime

type DirectoryAPI struct {
	Name     string     `json:"name"`
	NodeMaps []*NodeMap `json:"nodeMaps"`
}
type NodeMap struct {
	URL           string `json:"url"`
	Interval      string `json:"interval"`
	TechnicalType string `json:"technicalType"`
	MapType       string `json:"mapType`
}
