package model
import (
    "database/sql"
//    "errors"
	"fmt"
	"log"
//	miu "banwire/dash/dashboard_banwire/util"
	
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

func GetDatadash0201(db *sql.DB, reference string,reference02 string,dato01 string,dato02 string,dato03 string) ([]Datadash, error) {
	log.Print("GetDatadash0201 01!\n")

//	statement := fmt.Sprintf("SELECT id_record, Numsecuecial_grupoddatos,nombrecolumna,Created_at,Valorcolumna,Valoramount  FROM banwiredash02graph01   WHERE  reference= '%s'  order by Created_at ASC;", reference)
//	statement := fmt.Sprintf("SELECT id_record, Numsecuecial_grupoddatos,nombrecolumna,Created_at,Valorcolumna,Valoramount  FROM banwiredash02graph01   WHERE  Numsecuecial_grupoddatos= %s  order by Created_at ASC;", reference)
//    statement := fmt.Sprintf("SELECT id_record, Numsecuecial_grupoddatos,Numsecuecial,nombrecolumna,Created_at,Valorcolumna,Valoramount  FROM banwiredash02graph01    order by Created_at ASC;")
    statement := fmt.Sprintf("SELECT id_record, Numsecuecial_grupoddatos,Numsecuecial,nombrecolumna,Created_at,Valorcolumna,Valoramount  FROM banwiredash02graph01  WHERE  Numsecuecial_grupoddatos= %s  order by Numsecuecial;",reference)
 	log.Print("GetDatadash0201 02!\n")
    rows, err := db.Query(statement)
    log.Print("GetDatadash0201 02.1!\n")
    if err != nil {
        return nil, err
    }
    log.Print("GetDatadash0201 02.5!\n")
    defer rows.Close()
    datadashs := []Datadash{}
    for rows.Next() {
    	 log.Print("GetDatadash0201 03!\n")
        var u Datadash
        if err := rows.Scan(&u.ID, &u.Numsecuecialgrupoddatos, &u.Numsecuecial, &u.Nombrecolumna,&u.Createdat,&u.Valoramount, &u.Valorcolumna); err != nil {

//        if err := rows.Scan(&u.Token, &u.Bin); err != nil {
        	log.Print("GetDatadash0201 err!\n"+err.Error())
            return nil, err
        }
    	 log.Print("GetDatadash0201 03.5!\n")
        datadashs = append(datadashs, u)
    }
    log.Print("GetDatadash0201 04!\n")
    return datadashs, nil
}
