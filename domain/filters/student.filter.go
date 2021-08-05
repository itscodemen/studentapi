package filters

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentFilter struct {
	SortField *string `form:"sort_field"`
	SortDest  *string `form:"sort_dir"`
}

func (sf *StudentFilter) GetSortField() string {
	if sf.SortField == nil {
		return ""
	}

	return *sf.SortField
}

func (sf *StudentFilter) Scope(db *gorm.DB) *gorm.DB {
	if sf.SortField != nil {
		if sf.SortDest != nil {
			db = db.Order(fmt.Sprintf("%s %s", sf.GetSortField(), *sf.SortDest))
		}
	}
	return db
}

func NewStudentFilterFromCtx(c *gin.Context) (*StudentFilter, error) {
	var sf StudentFilter

	if err := c.ShouldBindQuery(&sf); err != nil {
		return nil, err
	}
	return &sf, nil
}
