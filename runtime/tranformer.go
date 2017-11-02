package runtime

var transformers []func([]byte, string, *Fetcher) error = []func([]byte, string, *Fetcher) error{
	transformNodelist, transformMeshviewerV2, transformMeshviewerV1, transformMeshviewerFFRGB, transformYanic,
}

func transform(body []byte, site_code string, f *Fetcher) {
	for _, trans := range transformers {
		err := trans(body, site_code, f)
		if err == nil {
			return
		}
	}
}
