package app

//func (e *App) CreateUser(u uuid.UUID, login string, password string, avatar []byte, name string) error {
//	if !checkPermissionForUUID(e.gate, u, core.FN_CreateUser) {
//		return core.Status_AccessDenied
//	}
//
//	var c core.ContainerCreateUser
//	var err error
//
//	c.Request.Login = login
//	c.Request.Avatar = avatar
//	c.Request.Name = name
//	c.Request.PassHash, err = bcrypt.GenerateFromPassword([]byte(password+e.app.Salt), e.app.HashCost)
//	if err != nil {
//		return err
//	}
//
//	err = e.data.CreateUser(&c)
//	return err
//}
//
//func (e *App) Login(login string, password string) (uuid.UUID, error) {
//	var c core.ContainerGetUserPassHash
//	var err error
//
//	c.Request.Login = login
//	err = e.data.GetPassHash(&c)
//	if err != nil {
//		return uuid.New(), core.Status_AccessDenied
//	}
//
//	err = bcrypt.CompareHashAndPassword(c.Response.PassHash, []byte(password))
//	if err != nil {
//		return uuid.New(), core.Status_AccessDenied
//	}
//
//	return uuid.New(), nil
//}

//func (e *User) CreateUser(ctx context.Context, r *utils.ContentRequestCreateUser) (utils.ContentResponseCreateUser, error) {
//	panic("implement me")
//}
//
//func (e *User) GetUserInfo(ctx context.Context, r *utils.ContentRequestGetUserInfo) (utils.ContentResponseGetUserInfo, error) {
//	panic("implement me")
//}
//
//func (e *User) UpdateUserInfo(ctx context.Context, r *utils.ContentRequestUpdateUserInfo) (utils.ContentResponseUpdateUserInfo, error) {
//	panic("implement me")
//}
//
//func (e *User) GetUserList(ctx context.Context, r *utils.ContentRequestGetUserList) (utils.ContentResponseGetUserList, error) {
//	panic("implement me")
//}
