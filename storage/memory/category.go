package memory

import "todo/entity"

type Category struct {
	categories []entity.Category
}

func NewCategoryStorage()*Category{
	return &Category{
		categories: []entity.Category{},
	}
}

func (c *Category)CheckCategoryByUserID(userID,categoryID int)bool{

	for _,value := range c.categories{
		if value.UserID == userID && value.ID == categoryID{
			return true
		}
	}
	return false

}