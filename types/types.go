package types

type BlogItem struct {
	Id           int
	Title        string
	Content      string
	Description  string
	Date_created string
}

type BlogsResult struct {
	Data []BlogItem
}

type ProjectItem struct {
	Id           int
	Title        string
	Content      string
	Description  string
	Date_created string
}

type ProjectsResult struct {
	Data []ProjectItem
}

type ProblemItem struct {
	Id           int
	Title        string
	Content      string
	Tags         []string `json:"tags"`
	Date_created string
}

type ProblemsResult struct {
	Data []ProblemItem
}
