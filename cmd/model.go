package model

type User struct {
	Id         int     `:"Id" json:"Id"`
	Mobile     int     `:"Mobile" json:"Mobile"`         // 手机号
	Email      string  `:"Email" json:"Email"`           // 邮箱
	Password   string  `:"Password" json:"Password"`     // 密码
	NickName   string  `:"NickName" json:"NickName"`     // 昵称
	Money      float64 `:"Money" json:"Money"`           // 账户余额
	Frozen     float64 `:"Frozen" json:"Frozen"`         // 冻结金额
	PayAccount string  `:"PayAccount" json:"PayAccount"` // 支付宝账号
	RealName   string  `:"RealName" json:"RealName"`     // 真实姓名
	Pid        int     `:"Pid" json:"Pid"`               // 推荐人 id
	CreatedAt  string  `:"CreatedAt" json:"CreatedAt"`
}
