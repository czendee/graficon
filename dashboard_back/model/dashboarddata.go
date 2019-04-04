package model
import (
    "database/sql"
//    "errors"
	"fmt"
	"log"
//	miu "banwire/dash/dashboard_back/util"
	
)
type Datadash struct {
    ID    string    `sql:"type:bigint"`
    Numsecuecialgrupoddatos  string `sql:"type:bigint`
    Numsecuecial  string `sql:"type:bigint`
    Nombrecolumna   string    `sql:"type:varchar(4)`
    Createdat   string    `sql:"type:timestamp`
    Valoramount   string    `sql:"type:bigint`
    Valorcolumna   string    `sql:"type:varchar(30)`
}
func (u *Datadash) getDatadash(db *sql.DB) error {
    statement := fmt.Sprintf("SELECT id_record, Numsecuecial_grupoddatos,Numsecuecial FROM banwiredash02graph01 WHERE id_record=%d", u.ID)
    return db.QueryRow(statement).Scan(&u.ID, &u.Numsecuecialgrupoddatos)
}


 func BackDatadash0201(db *sql.DB) (string, error) {
	log.Print("BackDatadash0201 01!\n")
    var resultado string
    resultado ="executed not OK"

    statement := fmt.Sprintf("SELECT id_record, Numsecuecial_grupoddatos,Numsecuecial,nombrecolumna,Created_at,Valorcolumna,Valoramount  FROM banwiredash02graph01  order by Numsecuecial;")
 	log.Print("BackDatadash0201 02!\n")
    rows, err := db.Query(statement)
    log.Print("BackDatadash0201 02.1!\n")
    if err != nil {
        return resultado, err
    }
    log.Print("BackDatadash0201 02.5!\n")
    defer rows.Close()
    datadashs := []Datadash{}

    
    for rows.Next() {
        resultado ="executed OK"

    	 log.Print("BackDatadash0201 03!\n")
        var u Datadash
        if err := rows.Scan(&u.ID, &u.Numsecuecialgrupoddatos, &u.Numsecuecial, &u.Nombrecolumna,&u.Createdat,&u.Valoramount, &u.Valorcolumna); err != nil {

//        if err := rows.Scan(&u.Token, &u.Bin); err != nil {
        	log.Print("GetDatadash0201 err!\n"+err.Error())
            return resultado, err
        }
    	 log.Print("GetDatadash0201 03.5!\n")
        datadashs = append(datadashs, u)
    }
    log.Print("GetDatadash0201 04!\n")
    return resultado, nil
}

