package repositories

import (
	"database/sql"

	"github.com/MatsuoTakuro/my-template-connect-go/models"
)

func SelectStoreList(db *sql.DB, searchQuery string, companyCD int) ([]models.Store, error) {
	const sqlStr = `
		SELECT store_cd, company_cd, store_name, address, latitude, longitude
		FROM stores
		WHERE company_cd = ?
		AND store_name LIKE CONCAT('%', ?, '%');
	`

	rows, err := db.Query(sqlStr, companyCD, searchQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	storeArray := make([]models.Store, 0)
	for rows.Next() {
		var store models.Store
		err = rows.Scan(&store.StoreCD, &store.CompanyCD, &store.StoreName, &store.Address, &store.Latitude, &store.Longitude)
		if err != nil {
			return nil, err
		}

		storeArray = append(storeArray, store)
	}

	return storeArray, nil
}
