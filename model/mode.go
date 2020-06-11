package model


type GithubUser struct {
	Id         int    `json:"id"`
	Login      string `json:"login"`
	NodeId     string `json:"node_id"`
	AvatarUrl  string `json:"avatar_url"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	CreateTime int    `json:"create_time"`
	UpdateTime int    `json:"update_time"`
}