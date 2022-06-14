// Code generated by goctl. DO NOT EDIT.
package types

type RegisterReq struct {
	Username   string `json:"username"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Password   string `json:"password"`
	Gender     string `json:"gender"`
	Dob        string `json:"dob"`
}

type RegisterResp struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Gender     string `json:"gender"`
	Dob        string `json:"dob"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UpdateUserReq struct {
	Id         string  `path:"id"`
	Username   *string `json:"username"`
	First_name *string `json:"first_name"`
	Last_name  *string `json:"last_name"`
	Password   *string `json:"password"`
	Gender     *string `json:"gender"`
	Dob        *string `json:"dob"`
}

type GetUserReq struct {
	Id string `path:"id"`
}

type RefreshReq struct {
	RefreshToken string `json:"refresh_token"`
}
