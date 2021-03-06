package module

// A Resp 项目统一返回的结构
// swagger:response Resp
type Resp struct {
	// 返回体状态码
	MgsCode int `json:"mgsCode" binding:"required"`
	// 返回体信息
	MgsText string `json:"mgsText" binding:"required"`
}

type ResponsesData struct {
	MgsCode int `json:"mgs_code" binding:"required"`  // 
}

type HttpErrs struct {
	MsgCode int `json:"msg_code" binding:"required"` // 状态码
	MgsText string `json:"mgs_text" binding:"required"` // 报错信息
}