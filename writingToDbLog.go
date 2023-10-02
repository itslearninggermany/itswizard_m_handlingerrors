package itswizard_m_handlingerrors

import "github.com/jinzhu/gorm"

/*
The content will be stored in the Log. Errors must be stored over the ErrorLog!
Examples: "User logged in"
*/
func WritingToDbLog(db *gorm.DB, user string, content string) {
	db.Save(&Log{
		Username: user,
		Content:  content,
	})
}
