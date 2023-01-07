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

func (c *StoreController) StoreListHandler(w http.ResponseWriter, req *http.Request) {
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

func (c *StoreController) GrpcStoreListHandler() (string, http.Handler) {
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
		// TODO: investigate what causes apperrors.BadParam.Wrap(error, string) to panic
		return nil, connect.NewError(connect.CodeInvalidArgument, nil)
		// return nil,  apperrors.GrpcError(ctx, apperrors.BadParam.Wrap(nil, "company_cd must be more than 0"))

	}

	_, err := c.service.GetStoreListService(searchQuery, int(companyCD))
	// storeList, err := c.service.GetStoreListService(searchQuery, int(companyCD))
	if err != nil {
		return nil, apperrors.GrpcError(ctx, err)
	}

	res := connect.NewResponse(&storev1.ListStoresResponse{
		// TODO: investigate how to convert []models.Store to []*storev1.ListStoresResponse_Store
		// Stores: storeList,
	})

	return res, nil
}
