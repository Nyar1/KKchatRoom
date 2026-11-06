package dbmodel

// HubRecord 保存聊天会话信息。
const (
	HubType_PRIVATE = "PRIVATE"
	HubType_GROUP   = "GROUP"
)

type HubRecord struct {
	HubID string `json:"hubId" gorm:"column:hubId;type:varchar(64);primaryKey;comment:会话唯一标识"`

	HubType string `json:"hubType" gorm:"column:hubType;type:varchar(16);not null;comment:会话类型PRIVATE或GROUP"`

	HubName string `json:"hubName" gorm:"column:hubName;type:varchar(64);not null;default:'';comment:群聊名称，私聊可为空"`

	CreatedAt int64 `json:"createdAt" gorm:"column:createdAt;autoCreateTime:milli;comment:创建时间，毫秒"`

	UpdatedAt int64 `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime:milli;comment:更新时间，毫秒"`
}

// HubMember 保存会话成员关系。
const (
	RoleMember int8 = iota // 0：普通成员
	RoleAdmin              // 1：管理员
	RoleOwner              // 2：群主
)

type HubMember struct {
	HubID string `json:"hubId" gorm:"column:hubId;type:varchar(64);primaryKey;comment:会话ID"`

	UserID string `json:"userId" gorm:"column:userId;type:varchar(64);primaryKey;index;comment:用户ID"`

	Role int8 `json:"role" gorm:"column:role;not null;default:0;comment:角色0成员1管理员2群主"`

	Alive bool `json:"alive" gorm:"column:alive;not null;comment:是否仍为会话成员"`

	CreatedAt int64 `json:"createdAt" gorm:"column:createdAt;autoCreateTime:milli;comment:首次加入时间，毫秒"`

	UpdatedAt int64 `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime:milli;comment:最后更新时间，毫秒"`
}
