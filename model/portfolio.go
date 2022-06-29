package model

type Portfolio struct {
	ID                string `json:"_id" bson:"_id"`
	Slug              string `json:"slug" bson:"slug"`
	FirstName         string `json:"first_name" bson:"first_name"`
	LastName          string `json:"last_name" bson:"last_name"`
	ProfessionalTitle string `json:"professional_title" bson:"professional_title"`

	Profiles      []Profile      `json:"profiles" bson:"profiles"`
	Contributions []Contribution `json:"contributions" bson:"contributions"`

	WorkExperiences []Experience            `json:"work_experiences" bson:"work_experiences"`
	Projects        []Project               `json:"projects" bson:"projects"`
	Educations      []Education             `json:"education" bson:"education"`
	Skills          []SkillCollection       `json:"skills" bson:"skills"`
	Certificates    []CertificateCollection `json:"certificates" bson:"certificates"`

	Interests []string `json:"interests" bson:"interests"`

	Meta Meta `json:"meta" bson:"meta"`
}

type Meta struct {
	SourceCodeURL  string `json:"source_code_url" bson:"source_code_url"`
	PdfDownloadURL string `json:"pdf_download_url" bson:"pdf_download_url"`
}

type Contribution struct {
	Name string `json:"name" bson:"name"`
	Icon string `json:"icon" bson:"icon"`
	URL  string `json:"url" bson:"url"`
}

type Education struct {
	CourseTitle string `json:"course_title" bson:"course_title"`
	Board       string `json:"board" bson:"board"`
	School      School `json:"school" bson:"school"`
	Session     string `json:"session" bson:"session"`
	StartYear   int    `json:"start_year" bson:"start_year"`
	EndYear     int    `json:"end_year,omitempty" bson:"end_year"`

	Grade        Grade    `json:"grade" bson:"grade"`
	Achievements []string `json:"achievements,omitempty" bson:"achievements"`
}

type School struct {
	Name     string `json:"name" bson:"name"`
	Logo     string `json:"logo" bson:"logo"`
	Location string `json:"location" bson:"location"`
}

type Grade struct {
	Type          string  `json:"type" bson:"type"`
	AcquiredScore float64 `json:"acquired_score" bson:"acquired_score"`
	Scale         int     `json:"scale" bson:"scale"`
}

type CertificateCollection struct {
	Type         string        `json:"type" bson:"type"`
	Certificates []Certificate `json:"certificates" bson:"certificates"`
}

type Certificate struct {
	Name     string `json:"name" bson:"name"`
	Provider string `json:"provider" bson:"provider"`
	Year     int    `json:"year" bson:"year"`
	Image    string `json:"image" bson:"image"`
}

type Profile struct {
	Name     string `json:"name" bson:"name"`
	URL      string `json:"url" bson:"url"`
	Icon     string `json:"icon" bson:"icon"`
	Username string `json:"username" bson:"username"`
}

type SkillCollection struct {
	Type   string  `json:"type" bson:"type"`
	Skills []Skill `json:"skills" bson:"skills"`
}

type Project struct {
	Title       string `json:"title" bson:"title"`
	ProjectType string `json:"project_type" bson:"project_type"`
	Logo        string `json:"logo" bson:"logo"`
	Intent      string `json:"intent" bson:"intent"`
	AppURL      string `json:"app_url,omitempty" bson:"app_url"`

	BriefDescription string          `json:"brief_description" bson:"brief_description"`
	AppliedSkills    []AppliedSkills `json:"applied_skills" bson:"applied_skills"`
	Snapshots        []Snapshot      `json:"snapshots" bson:"snapshots"`
}

type AppliedSkills struct {
	Type   string   `json:"type" bson:"type"`
	Skills []string `json:"skills" bson:"skills"`
}

type Snapshot struct {
	Caption string `json:"caption" bson:"caption"`
	URL     string `json:"url" bson:"url"`
}

type Skill struct {
	Icon string `json:"icon" bson:"icon"`
	Name string `json:"name" bson:"name"`
}

type Organization struct {
	Title            string `json:"title" bson:"title"`
	Logo             string `json:"logo" bson:"logo"`
	WebsiteURL       string `json:"website_url" bson:"website_url"`
	LinkedinURL      string `json:"linkedin_url" bson:"linkedin_url"`
	BriefDescription string `json:"brief_description" bson:"brief_description"`
	Location         string `json:"location" bson:"location"`
}

type Experience struct {
	JobTitle         string       `json:"job_title" bson:"job_title"`
	Organization     Organization `json:"organization" bson:"organization"`
	StartMonth       string       `json:"start_month" bson:"start_month"`
	StartYear        string       `json:"start_year" bson:"start_year"`
	Responsibilities []string     `json:"responsibilities" bson:"responsibilities"`
}
