package model

type User struct {
	ID       uint
	Name     string `gorm:"type:varchar(250)"`
	Email    string `gorm:"type:varchar(250);unique;not null"`
	Password string `gorm:"type:varchar(250);not null"`
	Mobile   int    `gorm:"not null"`
	Active   int    `gorm:"type:tinyint(10);default:1"`
	Date     `gorm:"embedded"`
}

func (User) TableName() string {
	return "users"
}
func FindUserByEmail(email_id string) (int64, error) {
	var user User
	var count int64

	if result := db.Model(&user).Where("email=? AND active=?", email_id, 1).Count(&count); result.Error != nil {
		return 0, result.Error
	}
	return count, nil

}
