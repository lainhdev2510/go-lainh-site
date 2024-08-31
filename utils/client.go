package utils

import (
	"go-lainh-site/types"
	"log"

	"github.com/go-resty/resty/v2"
)

func FetchBlogs() []types.BlogItem {
	var result types.BlogsResult
	cmsUrl := "https://directus.lainh.site/items/blogs"

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
	cmsUrl := "https://directus.lainh.site/items/blogs?limit=4&sort[]=-date_created"

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

	cmsUrl := "https://directus.lainh.site/items/blogs/" + blogId

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
	cmsUrl := "https://directus.lainh.site/items/projects"

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
	cmsUrl := "https://directus.lainh.site/items/projects?limit=4&sort[]=-date_created"

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

	cmsUrl := "https://directus.lainh.site/items/projects/" + projectId

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
	cmsUrl := "https://directus.lainh.site/items/problems"

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

	cmsUrl := "https://directus.lainh.site/items/problems/" + problemId

	client := resty.New()
	client.SetHeader("accept", "application/json")

	_, err := client.R().SetResult(&result).Get(cmsUrl)

	if err != nil {
		log.Println(err)
	}

	return result.Data
}
