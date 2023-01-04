package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/MatsuoTakuro/my-template-connect-go/apperrors"
	"github.com/MatsuoTakuro/my-template-connect-go/controllers/params"
	"github.com/MatsuoTakuro/my-template-connect-go/controllers/services"
)

type StoreController struct {
	service services.StoreServicer
}

func NewStoreController(s services.StoreServicer) *StoreController {
	return &StoreController{service: s}
}

func (c *StoreController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `{"message": "Hello, world! by store"}`)
}

func (c *StoreController) StoreListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var searchQuery string
	if sq, ok := queryMap[params.SearchQuery]; ok && len(sq) > 0 {
		var err error
		searchQuery = sq[0]
		if err != nil {
			err = apperrors.BadParam.Wrap(err, "queryparam must be number")
			apperrors.ErrorHandler(w, req, err)
			return
		}
	} else {
		searchQuery = ""
	}
	var companyCD int
	if cd, ok := queryMap[params.CompanyCD]; ok && len(cd) > 0 {
		var err error
		companyCD, err = strconv.Atoi(cd[0])
		if err != nil {
			err = apperrors.BadParam.Wrap(err, "queryparam must be number")
			apperrors.ErrorHandler(w, req, err)
			return
		}
	} else {
		companyCD = 1
	}

	storeList, err := c.service.GetStoreListService(searchQuery, companyCD)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	if err = json.NewEncoder(w).Encode(storeList); err != nil {
		err = apperrors.ResBodyEncodeFailed.Wrap(err, "fail to encode response body")
		apperrors.ErrorHandler(w, req, err)
		return
	}
}
