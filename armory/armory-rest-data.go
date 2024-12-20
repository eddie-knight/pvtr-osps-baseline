package armory

type RestData struct {
	Id   int
	Name string
}

func (r *ArmoryData) Rest() RestData {
	if r.rest.Name == "" {
		r.loadRestData()
	}
	return r.rest
}

func (r *ArmoryData) loadRestData() error {
	return nil
}
