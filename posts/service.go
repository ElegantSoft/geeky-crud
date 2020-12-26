package posts

import "github.com/ElegantSoft/geeky-crud/models"

type PostService interface {
	Create(post *models.Post) error
	Find(post *models.Post) error
}

//type postServiceImpl struct {
//
//}
//
//func (p postServiceImpl) Create(post *models.Post) error {
//	panic("implement me")
//}
//
//func (p postServiceImpl) Find(post *models.Post) error {
//	panic("implement me")
//}
//
//func NewService() PostService {
//	return &postServiceImpl{}
//}