package lead

import "gorm.io/gorm"

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   int
}
