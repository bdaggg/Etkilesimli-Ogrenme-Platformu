package model

var (
	GroupMap = map[int]string{
		1: "Anasayfa",
		// Diğerleri eklenecek
	}
)

type SystemModule struct {
	ID       int                  `gorm:"primarykey;autoIncrement" json:"id"`
	Path     string               `json:"-"`
	Name     string               `json:"name"`
	GroupNum int                  `json:"group_num" gorm:"default:0"`
	Actions  []SystemModuleAction `gorm:"foreignKey:ModuleID" json:"actions"`
}

func (SystemModule) TableName() string {
	return "core.module"
}

func (SystemModule) Comment() string {
	return "COMMENT ON TABLE core.module IS 'Sistem modüllerinin tutulduğu tablo'; "
}

type SystemModuleAction struct {
	ID            int           `gorm:"primarykey;autoIncrement" json:"id"`
	ModuleID      int           `json:"module_id"`
	RelatedModule *SystemModule `gorm:"foreignKey:ModuleID" json:"related_module,omitempty"`
	Name          string        `json:"name"`
	Key           string        `json:"key"`
}

func (SystemModuleAction) TableName() string {
	return "core.module_actions"
}

func (SystemModuleAction) Comment() string {
	return "COMMENT ON TABLE core.module_actions IS 'Modüllerin Yetkilerinin tutulduğu tablo'; "
}

type SystemRole struct {
	ID         int                  `gorm:"primarykey;autoIncrement" json:"id"`
	Name       string               `gorm:"index:,unique" json:"name"`
	Privileges []SystemModuleAction `gorm:"many2many:core.role_module_actions;" json:"privs"`
}

func (SystemRole) TableName() string {
	return "core.roles"
}

func (SystemRole) Comment() string {
	return "COMMENT ON TABLE core.roles IS 'Sistem rollerinin tutulduğu tablo'; "
}
