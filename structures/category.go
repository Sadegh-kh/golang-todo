package structures

type Category struct {
	title string
}

var categoryStorage []Category

func (c *Category) CreateCategory(title string) {
	c.title = title
	categoryStorage = append(categoryStorage, *c)
}

func GetCategoryList() []Category {
	return categoryStorage
}
