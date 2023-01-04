package testdata

import "github.com/MatsuoTakuro/my-template-connect-go/models"

var storeTestData = []models.Store{
	{
		StoreCD:   4,
		CompanyCD: 1,
		StoreName: "スーパーマーケット 田村店",
		Address:   "福岡県福岡市早良区田村1-15-5",
		Latitude:  33.5446,
		Longitude: 130.3259,
	},
	{
		StoreCD:   10,
		CompanyCD: 1,
		StoreName: "スーパーセンター 石田店",
		Address:   "福岡県北九州市小倉南区八重洲町5-15",
		Latitude:  33.8400,
		Longitude: 130.8910,
	},
}
