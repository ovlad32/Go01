package SD

import "log"

type Currency struct {
	IsoCode string `json:"isoCode"`
	Name string `json:"name"`
	FullName string  `json:"fullName"`
}

func getAllCurrencies () ([]Currency){
	result := make([]Currency,0,20);
	rows, err := Db.Query("select IsoCode, Name from nsi.currency order by country nulls last, is_metal desc");
	if err != nil {
		log.Fatal("allCurrencies",err)
	}
	defer rows.Close();
	for rows.Next() {
		var cur Currency;
		if err := rows.Scan(&cur.IsoCode,&cur.Name);err!=nil {
			if err != nil {
				log.Fatal("allCurrencies",err)
			}
		};
		result = append(result,cur);
	}
	if err:=rows.Err(); err!=nil {
		if err != nil {
			log.Fatal("allCurrencies",err)
		}
	}

	return result;

}