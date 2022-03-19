// Code generated by goctl. DO NOT EDIT.
package types

type AddUserReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	NickName string `json:"nickName"`
}

type AddUserResp struct {
	Id int64 `json:"id"`
}

type QueryUserReq struct {
	Id int64 `path:"id"`
}

type QueryUserResp struct {
	Id         int64  `json:"id"`
	Account    string `json:"account"`
	NickName   string `json:"nickName"`
	CreateTime string `json:"createTime"`
}

type UpdateUserReq struct {
	Id       int64  `json:"id"`
	NickName string `json:"nickName"`
}

type DeleteUserReq struct {
	Id int64 `json:"id"`
}