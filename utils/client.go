package utils

import (
	"go-lainh-site/types"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
)

func FetchBlogs() []types.BlogItem {
	var result types.BlogsResult
	cmsUrl := os.Getenv("CMS_URL") + "items/blogs"

	client := resty.New()
	client.SetHeader("accept", "application/json")

	_, err := client.R().SetResult(&result).Get(cmsUrl)

	if err != nil {
		log.Println(err)
	}

	return result.Data
}

func FetchFeaturedBlogs() []types.BlogItem {
	var result types.BlogsResult
	cmsUrl := os.Getenv("CMS_URL") + "items/blogs?limit=4&sort[]=-date_created"

	client := resty.New()
	client.SetHeader("accept", "application/json")

	res, err := client.R().SetResult(&result).Get(cmsUrl)

	if err != nil {
		log.Println(err)
	}

	log.Println(res.String())

	return result.Data
}

func FetchBlog(blogId string) types.BlogItem {
	var result *struct {
		Data types.BlogItem
	}

	cmsUrl := os.Getenv("CMS_URL") + "items/blogs/" + blogId

	client := resty.New()
	client.SetHeader("accept", "application/json")

	_, err := client.R().SetResult(&result).Get(cmsUrl)

	if err != nil {
		log.Println(err)
	}

	return result.Data
}

func FetchProjects() []types.ProjectItem {
	var result types.ProjectsResult
	cmsUrl := os.Getenv("CMS_URL") + "items/projects"

	client := resty.New()
	client.SetHeader("accept", "application/json")

	_, err := client.R().SetResult(&result).Get(cmsUrl)

	if err != nil {
		log.Println(err)
	}

	return result.Data
}

func FetchFeaturedProject() []types.ProjectItem {
	var result types.ProjectsResult
	cmsUrl := os.Getenv("CMS_URL") + "items/projects?limit=4&sort[]=-date_created"

	client := resty.New()
	client.SetHeader("accept", "application/json")

	_, err := client.R().SetResult(&result).Get(cmsUrl)

	if err != nil {
		log.Println(err)
	}

	return result.Data
}

func FetchProject(projectId string) types.ProjectItem {
	var result *struct {
		Data types.ProjectItem
	}

	cmsUrl := os.Getenv("CMS_URL") + "items/projects/" + projectId

	client := resty.New()
	client.SetHeader("accept", "application/json")

	_, err := client.R().SetResult(&result).Get(cmsUrl)

	if err != nil {
		log.Println(err)
	}

	return result.Data
}

func FetchAllProblem() []types.ProblemItem {
	var result types.ProblemsResult
	cmsUrl := os.Getenv("CMS_URL") + "items/problems"

	client := resty.New()
	client.SetHeader("accept", "application/json")

	_, err := client.R().SetResult(&result).Get(cmsUrl)

	if err != nil {
		log.Println(err)
	}

	return result.Data
}

func FetchProblem(problemId string) types.ProblemItem {
	var result *struct {
		Data types.ProblemItem
	}

	cmsUrl := os.Getenv("CMS_URL") + "items/problems/" + problemId

	client := resty.New()
	client.SetHeader("accept", "application/json")

	_, err := client.R().SetResult(&result).Get(cmsUrl)

	if err != nil {
		log.Println(err)
	}

	return result.Data
}
