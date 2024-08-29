package types

type BlogItem struct {
	Id           int
	Title        string
	Content      string
	Description  string
	Created_time string
}

type BlogsResult struct {
	List []BlogItem
}

type ProjectItem struct {
	Id           int
	Title        string
	Content      string
	Description  string
	Created_time string
}

type ProjectsResult struct {
	List []ProjectItem
}

type ProblemItem struct {
	Id           int
	Title        string
	Content      string
	Tag          string
	Created_time string
}

type ProblemsResult struct {
	List []ProblemItem
}
