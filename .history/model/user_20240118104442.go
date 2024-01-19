package model

// type UserInfo struct {
// 	UserName string `json:"userName"`
// 	PassWord string `json:"password"`
// }
type Login struct {
	UserName string `json:"userName" binding:"required"`
	Passward string `json:"passward" binding:"required"`
}
