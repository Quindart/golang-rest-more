package models

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

func NewCategory(id, name, path string) *Category {
	return &Category{
		Id:   id,
		Name: name,
		Path: path,
	}
}
