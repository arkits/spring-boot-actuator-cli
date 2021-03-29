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
	Git   ActuatorInfoGitProperties `json:"git"`
	Title string                    `json:"title"`
	Build map[string]interface{}    `json:"build"`
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
	Commit map[string]interface{} `json:"commit"`
	Dirty  string                 `json:"dirty"`
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

type ActuatorMetricsProperties struct {
	Names []string `json:"names"`
}
