package models

import "blog-service/utils"

type Project struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	Image      string `json:"image"`
	CategoryId string `json:"category_id"`
	UserId     string `json:"user_id"`
}

func NewProject(name, url, image string) *Project {
	return &Project{
		ID:    utils.GenerateUUID(),
		Name:  name,
		Url:   url,
		Image: image,
	}
}
