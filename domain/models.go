package domain

type ActuatorLinks struct {
	Links map[string]interface{} `json:"_links"`
}

type ActuatorEnvProperties struct {
	ActiveProfiles  []string                     `json:"activeProfiles"`
	PropertySources []ActuatorEnvPropertySources `json:"propertySources"`
}

type ActuatorEnvPropertySources struct {
	Name       string                 `json:"name"`
	Properties map[string]interface{} `json:"properties"`
}

type ActuatorHealthProperties struct {
	Status string `json:"status"`
}

type ActuatorInfoProperties struct {
	Git ActuatorInfoGitProperties `json:"git"`
}

type ActuatorInfoGitProperties struct {
	Branch string `json:"branch"`
	Build  struct {
		Host string `json:"host"`
		User struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"user"`
		Version string `json:"version"`
	} `json:"build"`
	Closest struct {
		Tag struct {
			Commit struct {
				Count string `json:"count"`
			} `json:"commit"`
			Name string `json:"name"`
		} `json:"tag"`
	} `json:"closest"`
	Commit struct {
		CommitID string `json:"id"`
		ID       struct {
			Abbrev   string `json:"abbrev"`
			Describe string `json:"describe"`
			Full     string `json:"full"`
		} `json:"id"`
		Message struct {
			Full  string `json:"full"`
			Short string `json:"short"`
		} `json:"message"`
		Time string `json:"time"`
		User struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"user"`
	} `json:"commit"`
	Dirty  string `json:"dirty"`
	Remote struct {
		Origin struct {
			URL string `json:"url"`
		} `json:"origin"`
	} `json:"remote"`
	Tags  string `json:"tags"`
	Total struct {
		Commit struct {
			Count string `json:"count"`
		} `json:"commit"`
	} `json:"total"`
}
