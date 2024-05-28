package model

const TableNameClassPlanUser = "class_plan_user"

// ClassPlanUser mapped from table <class_plan_user>
type ClassPlanUser struct {
	ID              int  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID            string `gorm:"column:uuid;not null" json:"uuid"`
	ClassPlanTypeID int  `gorm:"column:class_plan_type_id;not null" json:"class_plan_type_id"`
	UserPlanID      int  `gorm:"column:user_plan_id;not null" json:"user_plan_id"`
}

// TableName ClassPlanUser's table name
func (*ClassPlanUser) TableName() string {
	return TableNameClassPlanUser
}
