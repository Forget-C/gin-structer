package schema

import "github.com/Forget-C/http-structer/pkg/schema/base"

type IDUriReq struct {
	ID string `json:"id" uri:"id" binding:"required"`
}

type GetUserReq struct {
	IDUriReq
}

type ListUsersReq struct {
	base.WSearchReq
	base.DefaultListReq
}

type CreateUserReq struct {
	Name  string `json:"name" binding:"required"`
	Title string `json:"title"`
}

type UpdateUserDataReq struct {
	Title string `json:"title"`
}

type UpdateUserReq struct {
	IDUriReq
	UpdateUserDataReq
}

type DeleteUserReq struct {
	IDUriReq
}

type GetApproveReq struct {
	IDUriReq
}

type ListApprovesReq struct {
	base.WSearchReq
	base.DefaultListReq
}

type CreateApproveReq struct {
	Desc string `json:"desc" binding:"required"`
}

type CommentApproveDataReq struct {
	Status  int    `json:"status" binding:"required;oneof=1 2"`
	Comment string `json:"comment"`
}
