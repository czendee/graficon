package model
import (
    "database/sql"
//    "errors"
	"fmt"
	"log"
//	miu "banwire/dash/dashboard_back/util"
	
)


type Datadashmysql struct {
    ID    string    `sql:"type:bigint"`
    Numsecuecialgrupoddatos  string `sql:"type:bigint`
    Numsecuecial  string `sql:"type:bigint`
    Nombrecolumna   string    `sql:"type:varchar(4)`
    Createdat   string    `sql:"type:timestamp`
    Valoramount   string    `sql:"type:bigint`
    Valorcolumna   string    `sql:"type:varchar(30)`
}

func (u *Datadashmysql) getDatadashMySql(db *sql.DB) error {
    statement := fmt.Sprintf("SELECT id_record, Numsecuecial_grupoddatos,Numsecuecial FROM banwiredash02graph01 WHERE id_record=%d", u.ID)
    return db.QueryRow(statement).Scan(&u.ID, &u.Numsecuecialgrupoddatos)
}


func BackMysqlDatadash0201(db *sql.DB) ([]Datadash, error) {
	log.Print("BackDatadash0201 01!\n")
    var resultado string
    resultado ="executed not OK"
    datadashs := []Datadash{}

    statement := fmt.Sprintf("SELECT id_record, Numsecuecial_grupoddatos,Numsecuecial,nombrecolumna,Created_at,Valorcolumna,Valoramount  FROM banwiredash02graph01  order by Numsecuecial;")
 	log.Print("BackDatadash0201 02!\n")
    rows, err := db.Query(statement)
    log.Print("BackDatadash0201 02.1!\n")
    if err != nil {
        return datadashs, err
    }
    log.Print("BackDatadash0201 02.5!\n")
    defer rows.Close()


    for rows.Next() {
        resultado ="executed OK"

    	 log.Print("BackDatadash0201 03!\n")
        var u Datadash
        if err := rows.Scan(&u.ID, &u.Numsecuecialgrupoddatos, &u.Numsecuecial, &u.Nombrecolumna,&u.Createdat,&u.Valoramount, &u.Valorcolumna); err != nil {

//        if err := rows.Scan(&u.Token, &u.Bin); err != nil {
        	log.Print("GetDatadash0201 err!\n"+err.Error())
            return datadashs, err
        }
    	 log.Print("GetDatadash0201 03.5!\n")
        datadashs = append(datadashs, u)
    }
    log.Print("GetDatadash0201 04!\n"+resultado)
    return datadashs, nil
}


func BackMysqlInsertDatadash0201(db *sql.DB) (string, error) {
	log.Print("BackInsertDatadash0201 01!\n")
    var resultado string
    resultado ="executed not OK"
    datadashs := []Datadash{}

 //for each
//{
/*	    statement := fmt.Sprintf("INSERT INTO banwirepayment( token, created_at, amount) VALUES('%s',current_timestamp,'%s')",  u.Token, u.Amount)
	    _, err := db.Exec(statement)
     	log.Print("exec ejecutado")
	    if err != nil {
     	   log.Print("exec ejecutado:error "+err.Error())
	        return err
	    }
*/
//}

    statement := fmt.Sprintf("SELECT id_record, Numsecuecial_grupoddatos,Numsecuecial,nombrecolumna,Created_at,Valorcolumna,Valoramount  FROM banwiredash02graph01  order by Numsecuecial;")
 	log.Print("BackInsertDatadash0201 02!\n")
    rows, err := db.Query(statement)
    log.Print("BackInsdertDatadash0201 02.1!\n")
    if err != nil {
        return resultado, err
    }
    log.Print("BackInsertDatadash0201 02.5!\n")
    defer rows.Close()


    for rows.Next() {
        resultado ="executed OK"

    	 log.Print("BackInsertDatadash0201 03!\n")
        var u Datadash
        if err := rows.Scan(&u.ID, &u.Numsecuecialgrupoddatos, &u.Numsecuecial, &u.Nombrecolumna,&u.Createdat,&u.Valoramount, &u.Valorcolumna); err != nil {

        	log.Print("GetInsertDatadash0201 err!\n"+err.Error())
            return resultado, err
        }
    	 log.Print("GetInsertDatadash0201 03.5!\n")
        datadashs = append(datadashs, u)
    }
    log.Print("GetInsertDatadash0201 04!\n"+resultado)
    return resultado, nil
}