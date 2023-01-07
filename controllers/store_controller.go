package controllers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/MatsuoTakuro/my-template-connect-go/apperrors"
	"github.com/MatsuoTakuro/my-template-connect-go/controllers/params"
	"github.com/MatsuoTakuro/my-template-connect-go/controllers/services"
	storev1 "github.com/MatsuoTakuro/my-template-connect-go/gen/templateconnectgo/v1"
	"github.com/MatsuoTakuro/my-template-connect-go/gen/templateconnectgo/v1/storev1connect"
	"github.com/bufbuild/connect-go"
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

func (c *StoreController) HttpStoreListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()
	var searchQuery string
	if sq, ok := queryMap[params.SearchQuery]; ok && len(sq) > 0 {
		searchQuery = sq[0]
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

func (c *StoreController) StoreListHandler() (string, http.Handler) {
	return storev1connect.NewStoreServiceHandler(c)
}

func (c *StoreController) ListStores(
	ctx context.Context, req *connect.Request[storev1.ListStoresRequest],
) (*connect.Response[storev1.ListStoresResponse], error) {

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	searchQuery := req.Msg.SearchQuery

	companyCD := req.Msg.CompanyCd
	if companyCD == 0 {
		return nil, apperrors.GrpcError(ctx, apperrors.NewAppError(apperrors.BadParam, "company_cd must be more than 0"))

	}

	storeList, err := c.service.GetStoreListService(searchQuery, int(companyCD))
	if err != nil {
		return nil, apperrors.GrpcError(ctx, err)
	}

	var stores []*storev1.ListStoresResponse_Store
	for _, s := range storeList {
		storeRes := &storev1.ListStoresResponse_Store{
			StoreCd:   int32(s.StoreCD),
			CompanyCd: int32(s.CompanyCD),
			StoreName: s.StoreName,
			Address:   s.Address,
			Latitude:  s.Latitude,
			Longitude: s.Longitude,
		}
		stores = append(stores, storeRes)
	}

	res := connect.NewResponse(&storev1.ListStoresResponse{
		Stores: stores,
	})

	return res, nil
}
