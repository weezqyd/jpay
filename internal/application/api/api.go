package api

import (
	"github.com/weezqyd/jpay/internal/application/core"
	"log"
	"net/http"
	"strconv"
)
import "github.com/gin-gonic/gin"

type Adapter struct {
	core *core.Service
}

func NewApi(core *core.Service) *Adapter {
	return &Adapter{
		core: core,
	}
}

func (a *Adapter) GetAllCustomers(ctx *gin.Context) {
	country := ctx.Query("country")
	var data []*core.Customer
	var err error
	
	if country != "" {
		data, err = a.core.GetCustomersByCountry(country)
	} else {
		page := uint64(0)
		if pageNo := ctx.Query("page"); pageNo != "" {
			p, err := strconv.Atoi(pageNo)
			if err == nil {
				page = uint64(p)
			}
		}
		data, err = a.core.GetAllCustomers(page)
	}
	if err != nil {
		log.Printf("error while listing customers by country err: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "something went wrong",
		})
		return
	}
	
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
	return
}
