package structures

type Category struct {
	title string
}

func (c *Category) CreateCategory(title string) {
	c.title = title
}
