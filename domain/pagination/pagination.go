package pagination

import (
	"math"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	firstPage = 1
	perPage   = 5
)

type Pagination struct {
	Total   int64 `json:"total" binding:"-"`
	Pages   int   `json:"pages" binding:"-"`
	Page    int   `json:"page" form:"page" binding:"omitempty"`
	PerPage int   `json:"per_page" form:"per_page" binding:"required"`
}

func (p *Pagination) getLimit() int {
	return p.PerPage
}

func (p *Pagination) getOffset() int {
	if p.Page == firstPage {
		return 0
	}

	return p.PerPage * (p.Page - 1)
}

// CountPages prepares Page value from existing data
func (p *Pagination) CountPages() {
	if p.PerPage == 0 || p.Total == 0 {
		p.Pages = 0
		return
	}

	if p.Page == 0 {
		p.Page = firstPage
	}

	p.Pages = int(math.Ceil(float64(p.Total) / float64(p.PerPage)))
}

func (p *Pagination) Scope(db *gorm.DB) (*gorm.DB, error) {
	if p.Total == 0 {
		err := db.Count(&p.Total).Error
		if err != nil {
			return db, err
		}
	}

	db = db.Limit(p.getLimit()).Offset(p.getOffset())
	p.CountPages()

	return db, nil
}

func InitPagination() *Pagination {
	p := &Pagination{
		Page:    firstPage,
		PerPage: perPage,
	}

	return p
}

func NewPaginationFromGinCtx(c *gin.Context) (*Pagination, error) {
	p := InitPagination()
	if err := c.ShouldBindQuery(&p); err != nil {
		return nil, err
	}

	return p, nil
}
