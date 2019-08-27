package model
import (
    "database/sql"
//    "errors"
	"fmt"
	"log"
//	miu "banwire/dash/dashboard_front/util"
	
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
    statement := fmt.Sprintf("SELECT id_record, Numsecuecial_grupoddatos,Numsecuecial,nombrecolumna,Created_at,Valorcolumna,Valoramount  FROM banwiredash02graph01  WHERE  Numsecuecial_grupoddatos= %s  and graphnbr = '%s' order by Numsecuecial;",reference, reference02)
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


func GetMaxGroupNumberDatadash0201(db *sql.DB, reference string,reference02 string,dato01 string,dato02 string,dato03 string) (string, error) {
    var valueresult string
    valueresult ="0"
	log.Print("GetMaxGroupNumberDatadash0201 01!\n")

    statement := fmt.Sprintf("select max(numsecuecial_grupoddatos)  as gruponumero from banwiredash02graph01 where numsecuecial_grupoddatos>%s and graphnbr = '%s' group by numsecuecial_grupoddatos;",reference,reference02)
 	log.Print("GetMaxGroupNumberDatadash0201 02!\n")
    rows, err := db.Query(statement)
    log.Print("GetMaxGroupNumberDatadash0201 02.1!\n")
    if err != nil {
        return "NOK", err
    }
    log.Print("GetMaxGroupNumberDatadash0201 02.5!\n")
    defer rows.Close()
//    datadashs := []Datadash{}
    for rows.Next() {
    	 log.Print("GetMaxGroupNumberDatadash0201 03!\n")
        var u Datadash
        if err := rows.Scan( &u.Numsecuecialgrupoddatos); err != nil {

//        if err := rows.Scan(&u.Token, &u.Bin); err != nil {
        	log.Print("GetMaxGroupNumberDatadash0201 err!\n"+err.Error())
            return "NOK", err
        }
        valueresult = u.Numsecuecialgrupoddatos
    	 log.Print("GetMaxGroupNumberDatadash0201 03.5!\n"+valueresult)
    }
    log.Print("GetMaxGroupNumberDatadash0201 04!\n")
    return valueresult, nil
}

func GetMaxGroupNumberDatadash0202(db *sql.DB, reference string,reference02 string,dato01 string,dato02 string,dato03 string) (string, error) {
    var valueresult string
    valueresult ="0"
	log.Print("GetMaxGroupNumberDatadash0202 01!\n")

    statement := fmt.Sprintf("select max(numsecuecial_grupoddatos)  as gruponumero from banwiredash02graph02 where numsecuecial_grupoddatos>%s  and graphnbr = '%s' group by numsecuecial_grupoddatos;",reference,reference02)
 	log.Print("GetMaxGroupNumberDatadash0202 02!\n")
    rows, err := db.Query(statement)
    log.Print("GetMaxGroupNumberDatadash0202 02.1!\n")
    if err != nil {
        return "NOK", err
    }
    log.Print("GetMaxGroupNumberDatadash0202 02.5!\n")
    defer rows.Close()
//    datadashs := []Datadash{}
    for rows.Next() {
    	 log.Print("GetMaxGroupNumberDatadash0202 03!\n")
        var u Datadash
        if err := rows.Scan( &u.Numsecuecialgrupoddatos); err != nil {

//        if err := rows.Scan(&u.Token, &u.Bin); err != nil {
        	log.Print("GetMaxGroupNumberDatadash0202 err!\n"+err.Error())
            return "NOK", err
        }
        valueresult = u.Numsecuecialgrupoddatos
    	 log.Print("GetMaxGroupNumberDatadash0202 03.5!\n"+valueresult)
    }
    log.Print("GetMaxGroupNumberDatadash0202 04!\n")
    return valueresult, nil
}


func GetDatadash0202(db *sql.DB, reference string,reference02 string,dato01 string,dato02 string,dato03 string, comnadosql string) ([]Datadash, error) {
	log.Print("GetDatadash0201 01!\n")

//previo    statement := fmt.Sprintf("SELECT id_record, Numsecuecial_grupoddatos,Numsecuecial,nombrecolumna,Created_at,Valorcolumna,Valoramount  FROM banwiredash02graph02  WHERE  Numsecuecial_grupoddatos= %s and graphnbr = '%s'  order by Numsecuecial;",reference, reference02)
	
	    statement := fmt.Sprintf(comnadosql,reference02) //this value is decided in the logicdb.go, set in the Configuration.go and defined in the config.json
	
	
 	log.Print("GetDatadash0202 02!\n")
    rows, err := db.Query(statement)
    log.Print("GetDatadash0202 02.1!\n")
    if err != nil {
        return nil, err
    }
    log.Print("GetDatadash0202 02.5!\n")
    defer rows.Close()
    datadashs := []Datadash{}
    for rows.Next() {
    	 log.Print("GetDatadash0202 03!\n")
        var u Datadash
        if err := rows.Scan(&u.ID, &u.Numsecuecialgrupoddatos, &u.Numsecuecial, &u.Nombrecolumna,&u.Createdat,&u.Valoramount, &u.Valorcolumna); err != nil {

//        if err := rows.Scan(&u.Token, &u.Bin); err != nil {
        	log.Print("GetDatadash0202 err!\n"+err.Error())
            return nil, err
        }
    	 log.Print("GetDatadash0202 03.5!\n")
        datadashs = append(datadashs, u)
    }
    log.Print("GetDatadash0202 04!\n")
    return datadashs, nil
}

func GetDatadash0202original(db *sql.DB, reference string,reference02 string,dato01 string,dato02 string,dato03 string) ([]Datadash, error) {
	log.Print("GetDatadash0201 01!\n")

    statement := fmt.Sprintf("SELECT id_record, Numsecuecial_grupoddatos,Numsecuecial,nombrecolumna,Created_at,Valorcolumna,Valoramount  FROM banwiredash02graph02  WHERE  Numsecuecial_grupoddatos= %s and graphnbr = '%s'  order by Numsecuecial;",reference, reference02)
	
	
	
 	log.Print("GetDatadash0202 02!\n")
    rows, err := db.Query(statement)
    log.Print("GetDatadash0202 02.1!\n")
    if err != nil {
        return nil, err
    }
    log.Print("GetDatadash0202 02.5!\n")
    defer rows.Close()
    datadashs := []Datadash{}
    for rows.Next() {
    	 log.Print("GetDatadash0202 03!\n")
        var u Datadash
        if err := rows.Scan(&u.ID, &u.Numsecuecialgrupoddatos, &u.Numsecuecial, &u.Nombrecolumna,&u.Createdat,&u.Valoramount, &u.Valorcolumna); err != nil {

//        if err := rows.Scan(&u.Token, &u.Bin); err != nil {
        	log.Print("GetDatadash0202 err!\n"+err.Error())
            return nil, err
        }
    	 log.Print("GetDatadash0202 03.5!\n")
        datadashs = append(datadashs, u)
    }
    log.Print("GetDatadash0202 04!\n")
    return datadashs, nil
}


func GetDatadash0202hours24(db *sql.DB, reference string,reference02 string,dato01 string,dato02 string,dato03 string) ([]Datadash, error) {
	log.Print("GetDatadash0202hours24 01!\n")

    statement := fmt.Sprintf("SELECT 1 as id_record, 2 as Numsecuecial_grupoddatos,3 as Numsecuecial,nombrecolumna,current_date as Created_at,total as Valorcolumna,total as Valoramount  from (SELECT nombrecolumna,sum(Valoramount) as total  FROM banwiredash02graph02  WHERE  graphnbr = '%s' and created_At > current_date-1 group by nombrecolumna order by total desc) as buen;", reference02)
 	log.Print("GetDatadash0202hours24 02!\n")
    rows, err := db.Query(statement)
    log.Print("GetDatadash0202hours24 02.1!\n")
    if err != nil {
        return nil, err
    }
    log.Print("GetDatadash0202hours24 02.5!\n")
    defer rows.Close()
    datadashs := []Datadash{}
    for rows.Next() {
    	 log.Print("GetDatadash0202hours24 03!\n")
        var u Datadash
        if err := rows.Scan(&u.ID, &u.Numsecuecialgrupoddatos, &u.Numsecuecial, &u.Nombrecolumna,&u.Createdat,&u.Valoramount, &u.Valorcolumna); err != nil {

//        if err := rows.Scan(&u.Token, &u.Bin); err != nil {
        	log.Print("GetDatadash0202 err!\n"+err.Error())
            return nil, err
        }
    	 log.Print("GetDatadash0202hours24 03.5!\n")
        datadashs = append(datadashs, u)
    }
    log.Print("GetDatadash0202hours24 04!\n")
    return datadashs, nil
}

func GetDatadash0202hours48(db *sql.DB, reference string,reference02 string,dato01 string,dato02 string,dato03 string) ([]Datadash, error) {
	log.Print("GetDatadash0202hours48 01!\n")

    statement := fmt.Sprintf("SELECT 1 as id_record, 2 as Numsecuecial_grupoddatos,3 as Numsecuecial,nombrecolumna,current_date as Created_at,total as Valorcolumna,total as Valoramount  from (SELECT nombrecolumna,sum(Valoramount) as total  FROM banwiredash02graph02  WHERE  graphnbr = '%s' and created_At > current_date-2 group by nombrecolumna order by total desc) as buen;", reference02)
 	log.Print("GetDatadash0202hours48 02!\n")
    rows, err := db.Query(statement)
    log.Print("GetDatadash0202hours48 02.1!\n")
    if err != nil {
        return nil, err
    }
    log.Print("GetDatadash0202hours48 02.5!\n")
    defer rows.Close()
    datadashs := []Datadash{}
    for rows.Next() {
    	 log.Print("GetDatadash0202hours48 03!\n")
        var u Datadash
        if err := rows.Scan(&u.ID, &u.Numsecuecialgrupoddatos, &u.Numsecuecial, &u.Nombrecolumna,&u.Createdat,&u.Valoramount, &u.Valorcolumna); err != nil {

//        if err := rows.Scan(&u.Token, &u.Bin); err != nil {
        	log.Print("GetDatadash0202hours48 err!\n"+err.Error())
            return nil, err
        }
    	 log.Print("GetDatadash0202hours48 03.5!\n")
        datadashs = append(datadashs, u)
    }
    log.Print("GetDatadash0202hours48 04!\n")
    return datadashs, nil
}
