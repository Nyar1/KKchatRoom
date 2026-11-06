package dbmodel

// User 表示用户账号。
type User struct {
	KKUId string `json:"kkUId" gorm:"column:kkUId;type:varchar(64);primaryKey;comment:数据库主键"`

	KKCRId string `json:"kkcrId" gorm:"column:kkcrId;type:varchar(32);not null;uniqueIndex;comment:用户唯一账号"`

	KKCRName string `json:"kkcrName" gorm:"column:kkcrName;type:varchar(64);not null;comment:用户昵称"`

	PWD string `json:"pwd,omitempty" gorm:"column:pwd;type:varchar(255);not null;comment:明文密码"`

	PasswordHash string `json:"passwordHash,omitempty" gorm:"column:passwordHash;type:varchar(255);not null;comment:密码哈希"`

	CreatedAt int64 `json:"createdAt" gorm:"column:createdAt;autoCreateTime:milli;comment:创建时间，毫秒"`

	UpdatedAt int64 `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime:milli;comment:更新时间，毫秒"`
}
