package models

import "blog-service/utils"

type Blog struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	AuthorId   string `json:"author"`
	ImgUrl     string `json:"img_url"`
	CategoryId string `json:"category_id"`
}

func NewBlog(title, content, authorId, imgUrl, categoryId string) *Blog {
	return &Blog{
		Id:         utils.GenerateUUID(),
		Title:      title,
		Content:    content,
		AuthorId:   authorId,
		ImgUrl:     imgUrl,
		CategoryId: categoryId,
	}
}
