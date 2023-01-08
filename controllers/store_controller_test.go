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

func TestHttpStoreListHandler(t *testing.T) {
	tests := map[string]struct {
		query      storeListParams
		resultCode int
	}{
		"number companyCD": {
			query: storeListParams{
				searchQuery: "田",
				companyCD:   "1",
			},
			resultCode: http.StatusOK,
		},
		"string companyCD": {
			query: storeListParams{
				searchQuery: "田",
				companyCD:   "aaa",
			},
			resultCode: http.StatusBadRequest,
		},
	}

	for name, sub := range tests {
		t.Run(name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:8080/store?%s=%s&%s=%s", params.SearchQuery, sub.query.searchQuery, params.CompanyCD, sub.query.companyCD)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			res := httptest.NewRecorder()

			sCon.HttpStoreListHandler(res, req)

			if res.Code != sub.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", sub.resultCode, res.Code)
			}
		})
	}
}

// TODO: implement test code for StoreListHandler on grpc
func TestStoreListHandler(t *testing.T) {}
