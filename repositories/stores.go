package repositories

import (
	"context"
	"fmt"

	"github.com/MatsuoTakuro/my-template-connect-go/models"
)

func SelectStoreList(
	ctx context.Context, db Queryer, searchQuery string, companyCD int,
) ([]models.Store, error) {
	const sqlStr = `
		SELECT store_cd, company_cd, store_name, address, latitude, longitude
		FROM stores
		WHERE company_cd = $1
		AND store_name LIKE $2;
	`

	rows, err := db.QueryContext(ctx, sqlStr, companyCD, fmt.Sprint("%", searchQuery, "%"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stores := make([]models.Store, 0)
	for rows.Next() {
		var store models.Store
		err = rows.Scan(&store.StoreCD, &store.CompanyCD, &store.StoreName, &store.Address, &store.Latitude, &store.Longitude)
		if err != nil {
			return nil, err
		}

		stores = append(stores, store)
	}

	return stores, nil
}
