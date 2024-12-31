package armory

type SecurityInsights struct {
	ByteContent []byte   `json:"content"`
	SHA         string   `json:"sha"`
	Header      SIHeader `yaml:"header"`
}

type SIHeader struct {
	SchemaVersion string `yaml:"schema-version"`
	ChangeLogURL  string `yaml:"changelog"`
	LicenseURL    string `yaml:"license"`
}

func (s *SecurityInsights) GetData(owner, repo string) {
	return
	// err = json.Unmarshal(response, s)
	// if err != nil {
	// 	return
	// }
	// err = yaml.Unmarshal(s.ByteContent, s)
	// if err != nil {
	// 	return
	// }
	// return
}
