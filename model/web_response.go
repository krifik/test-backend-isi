package model

type WebResponse struct {
	Code   int    `json:"code"`
	Remark string `json:"remark"`
	Data   any    `json:"data"`
}
