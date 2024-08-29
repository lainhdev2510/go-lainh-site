package utils

import (
	"go-lainh-site/types"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
)

func FetchBlogs() []types.BlogItem {
	var result types.BlogsResult

	cmsUrl := os.Getenv("CMS_URL") + "tables/mvv2llfu53r4xyr/records"
	cmsToken := os.Getenv("CMS_TOKEN")

	client := resty.New()
	client.SetHeader("xc-token", cmsToken)
	client.SetHeader("accept", "application/json")

	_, err := client.R().SetQueryParams(map[string]string{
		"limit":   "25",
		"shuffle": "0",
		"offset":  "0",
		"sort":    "-created_time",
	}).SetResult(&result).Get(cmsUrl)

	if err != nil {
		log.Println(err)
	}

	return result.List
}

func FetchFeaturedProject() []types.ProjectItem {
	var result types.ProjectsResult

	cmsUrl := os.Getenv("CMS_URL") + "tables/mgatebgdir8d6zl/records"
	cmsToken := os.Getenv("CMS_TOKEN")

	client := resty.New()
	client.SetHeader("xc-token", cmsToken)
	client.SetHeader("accept", "application/json")

	_, err := client.R().SetQueryParams(map[string]string{
		"limit":   "4",
		"shuffle": "0",
		"offset":  "0",
		"sort":    "-created_time",
	}).SetResult(&result).Get(cmsUrl)

	if err != nil {
		log.Println(err)
	}

	return result.List
}

func FetchProjects() []types.ProjectItem {
	var result types.ProjectsResult

	cmsUrl := os.Getenv("CMS_URL") + "tables/mgatebgdir8d6zl/records"
	cmsToken := os.Getenv("CMS_TOKEN")

	client := resty.New()
	client.SetHeader("xc-token", cmsToken)
	client.SetHeader("accept", "application/json")

	_, err := client.R().SetQueryParams(map[string]string{
		"limit":   "25",
		"shuffle": "0",
		"offset":  "0",
		"sort":    "-created_time",
	}).SetResult(&result).Get(cmsUrl)

	if err != nil {
		log.Println(err)
	}

	return result.List
}

func FetchFeaturedBlogs() []types.BlogItem {
	var result types.BlogsResult

	cmsUrl := os.Getenv("CMS_URL") + "tables/mvv2llfu53r4xyr/records"
	cmsToken := os.Getenv("CMS_TOKEN")

	client := resty.New()
	client.SetHeader("xc-token", cmsToken)
	client.SetHeader("accept", "application/json")

	_, err := client.R().SetQueryParams(map[string]string{
		"limit":   "4",
		"shuffle": "0",
		"offset":  "0",
		"sort":    "-created_time",
	}).SetResult(&result).Get(cmsUrl)

	if err != nil {
		log.Println(err)
	}

	return result.List
}

func FetchBlog(blogId string) types.BlogItem {
	var result types.BlogItem

	cmsUrl := os.Getenv("CMS_URL") + "tables/mvv2llfu53r4xyr/records/" + blogId
	cmsToken := os.Getenv("CMS_TOKEN")

	client := resty.New()
	client.SetHeader("xc-token", cmsToken)
	client.SetHeader("accept", "application/json")

	_, err := client.R().SetResult(&result).Get(cmsUrl)

	if err != nil {
		log.Println(err)
	}

	return result
}

func FetchProject(projectId string) types.ProjectItem {
	var result types.ProjectItem

	cmsUrl := os.Getenv("CMS_URL") + "tables/mgatebgdir8d6zl/records/" + projectId
	cmsToken := os.Getenv("CMS_TOKEN")

	client := resty.New()
	client.SetHeader("xc-token", cmsToken)
	client.SetHeader("accept", "application/json")

	_, err := client.R().SetResult(&result).Get(cmsUrl)

	if err != nil {
		log.Println(err)
	}

	return result
}

func FetchAllProblem() []types.ProblemItem {
	var result types.ProblemsResult

	cmsUrl := os.Getenv("CMS_URL") + "tables/mqx1rzwesoj8t17/records"
	cmsToken := os.Getenv("CMS_TOKEN")

	client := resty.New()
	client.SetHeader("xc-token", cmsToken)
	client.SetHeader("accept", "application/json")

	_, err := client.R().SetQueryParams(map[string]string{
		"limit":   "1000",
		"shuffle": "0",
		"offset":  "0",
		"sort":    "title",
	}).SetResult(&result).Get(cmsUrl)

	if err != nil {
		log.Println(err)
	}

	return result.List
}

func FetchProblem(projectId string) types.ProblemItem {
	var result types.ProblemItem

	cmsUrl := os.Getenv("CMS_URL") + "tables/mqx1rzwesoj8t17/records/" + projectId
	cmsToken := os.Getenv("CMS_TOKEN")

	client := resty.New()
	client.SetHeader("xc-token", cmsToken)
	client.SetHeader("accept", "application/json")

	_, err := client.R().SetResult(&result).Get(cmsUrl)

	if err != nil {
		log.Println(err)
	}

	return result
}
