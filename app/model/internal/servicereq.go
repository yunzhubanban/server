package internal

type DeviceBindReq struct {
	Point    string `json:"point"`    // Point 班级号
	Password string `json:"password"` // Password 密码散列值 (SHA512)
}
