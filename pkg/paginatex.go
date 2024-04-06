package pkg

import (
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"reflect"
	"strings"
)

const (
	DefaultPage     = 1
	DefaultPageSize = 10
)

// Param 分页参数
type Param struct {
	DB       *gorm.DB
	Page     int64
	Limit    int64
	OrderBy  []string
	Select   string
	Unsecure bool
}

// Pagination 分页响应数据
type Pagination struct {
	Count       int64 `json:"count"`
	Total       int64 `json:"total"`
	TotalPages  int64 `json:"totalPages"`
	Offset      int64 `json:"offset"`
	PerPage     int64 `json:"perPage"`
	CurrentPage int64 `json:"currentPage"`
	PrevPage    int64 `json:"prevPage"`
	NextPage    int64 `json:"nextPage"`
}

func Paginate(p *Param, result interface{}) *Pagination {
	db := p.DB

	var count int64
	db.Count(&count)

	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}
	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			if !p.Unsecure {
				db = db.Order(o)
				continue
			}
			tempStr := strings.Split(o, ",")
			for _, v := range tempStr {
				strS := strings.Split(v, " ")
				if len(strS) < 2 {
					continue
				}
				if len(strS) == 2 {
					desc := false
					if strings.ToLower(strS[1]) == "desc" {
						desc = true
					}
					db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: strS[0]}, Desc: desc})
				} else if len(strS) > 2 {
					// 以最后一个空格分割   // 防止分割错误 例如 "CONVERT(name USING GBK) ASC"
					lastSpaceIndex := strings.LastIndex(v, " ")
					by := strings.ToLower(v[lastSpaceIndex+1:])
					if by != "desc" && by != "asc" {
						continue
					}
					desc := false
					if by == "desc" {
						desc = true
					}
					// 防止SQL注入
					db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: v[:lastSpaceIndex]}, Desc: desc})
				}
			}
		}
	}

	var paginator Pagination
	var offset int64

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	if p.Select != "" {
		db = db.Select(p.Select)
	}

	db.Limit(cast.ToInt(p.Limit)).Offset(cast.ToInt(offset)).Find(result)

	paginator.Total = count
	paginator.CurrentPage = p.Page

	slice := reflect.ValueOf(result)

	if slice.Kind() == reflect.Slice {
		paginator.Count = int64(slice.Len())
	}
	if slice.Kind() == reflect.Ptr {
		paginator.Count = int64(slice.Elem().Len())
	}

	paginator.Offset = offset
	paginator.PerPage = p.Limit
	paginator.TotalPages = int64(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		paginator.PrevPage = p.Page - 1
	} else {
		paginator.PrevPage = p.Page
	}

	if p.Page == paginator.TotalPages {
		paginator.NextPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}
	return &paginator
}
