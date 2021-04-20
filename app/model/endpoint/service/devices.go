package model

type DeviceBindReq struct {
	Point    string `json:"point"`    // Point 班级号
	Password string `json:"password"` // Password 密码散列值 (SHA512)
}

type DeviceResp struct {
	HasSetting bool   `json:"has_setting"` //标识是否存在远程配置文件（布尔值）
	Token      string `json:"token"`       // JWT Token
}
