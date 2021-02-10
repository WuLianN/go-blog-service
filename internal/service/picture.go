package service



type CountPictureRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type PictureListRequest struct {
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

func (svc *Service) CountPicture(param *CountPictureRequest) (int, error) {
	return svc.dao.CountPicture(param.Name, param.State)
}

func (svc *Service) GetPictureList(param *PictureListRequest, pager *app.Pager) ([]*model.Picture, error) {
	return svc.dao.GetPictureList(param.Name, param.State, pager.Page, pager.PageSize)
}