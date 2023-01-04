package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MatsuoTakuro/my-template-connect-go/controllers/params"
)

type storeListParams struct {
	searchQuery string
	companyCD   string
}

func TestStoreListHandler(t *testing.T) {
	var tests = []struct {
		name       string
		query      storeListParams
		resultCode int
	}{
		{name: "number companyCD", query: storeListParams{searchQuery: "田", companyCD: "1"}, resultCode: http.StatusOK},
		{name: "string companyCD", query: storeListParams{searchQuery: "田", companyCD: "aaa"}, resultCode: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:8080/store?%s=%s&%s=%s", params.SearchQuery, tt.query.searchQuery, params.CompanyCD, tt.query.companyCD)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			res := httptest.NewRecorder()

			sCon.StoreListHandler(res, req)

			if res.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, res.Code)
			}
		})
	}
}
