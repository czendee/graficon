package model
import (
    "database/sql"
//    "errors"
	"fmt"
	"log"
    "strconv"
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


 func BackDatadash0201(db *sql.DB, comandosqlorigen_origin string) ([]Datadash, error) {
	log.Print("BackDatadash0201 01!\n")
    var resultado string
    resultado ="executed not OK"
    datadashs := []Datadash{}

    statement := fmt.Sprintf(comandosqlorigen_origin)  //defined in the config.json and set in the configuration.go
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
        if err := rows.Scan( &u.Nombrecolumna,&u.Valoramount); err != nil {

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

 func BackNextGroupNumber(db *sql.DB,grphnbr string) ( string) {
	log.Print("BackNextGroupNumber 01!\n")
    var resultado string
    resultado ="executed not OK"
    
    statement := fmt.Sprintf("select max(numsecuecial_grupoddatos)+1 as nextnumero from banwiredash02graph02 p where p.graphnbr ='%s'",grphnbr)  //defined in the config.json and set in the configuration.go
                     //
 	log.Print("BackNextGroupNumber 02!\n")
    rows, err := db.Query(statement)
    log.Print("BackNextGroupNumber 02.1!\n")
    if err != nil {
        return resultado
    }
    log.Print("BackNextGroupNumber 02.5!\n")
    defer rows.Close()
    
    for rows.Next() {
        resultado ="executed OK"

    	 log.Print("BackNextGroupNumber 03!\n")
        var u Datadash
        if err := rows.Scan( &u.Nombrecolumna); err != nil {
        	log.Print("BackNextGroupNumber err!\n"+err.Error())
            return resultado
        }
    	log.Print("BackNextGroupNumber 03.5!\n")
        resultado=u.Nombrecolumna
    }
    log.Print("BackNextGroupNumber 04!\n"+resultado)
    return resultado
}


func BackInsertDatadash0201(db *sql.DB,graphnbr string, valoresToInsert []Datadash, proxGroupNumber string) (string, error) {
	log.Print("BackInsertDatadash0201 01!\n")
    var resultado string
    resultado ="executed not OK"
    
    secuencia :=0
    for _, d := range valoresToInsert {
        secuencia =secuencia +1

         var strSecuencia string
         strSecuencia = strconv.Itoa(secuencia)
        d.Valorcolumna = d.Valoramount

        statement := fmt.Sprintf("insert into banwiredash02graph02 (numsecuecial_grupoddatos,"+
         "numsecuecial, "+
         "nombrecolumna, "+
         "created_at,  "+
         "valorcolumna, "+
         " valoramount, "+
           "graphnbr)   "+
        "VALUES(    "+
//            " (select max(numsecuecial_grupoddatos)+1  from banwiredash02graph02 p where p.graphnbr ='%s'),     "+
            " '%s' ,     "+
            " %s,     "+
             " '%s',     "+
             " current_timestamp,    "+
             " '%s',    "+
              " %s," +
              " '%s' )",  
              proxGroupNumber,
              strSecuencia,
              d.Nombrecolumna,
              d.Valorcolumna,
              d.Valoramount,
              graphnbr   )
         log.Print("exec ejecutado comando:"+statement)
	    _, err := db.Exec(statement)
     	log.Print("exec ejecutado row:"+d.Nombrecolumna)
	    if err != nil {
     	   log.Print("exec ejecutado:error "+err.Error())
	        return resultado,err
	    }
        

    }//end for each

    if(secuencia>0){
        resultado ="executed OK"
    }else{
        resultado ="executed not OK"
    }
    

 	log.Print("BackInsertDatadash0201 02!\n")
    log.Print("InsertDatadash0201 04!\n"+resultado)
    return resultado, nil
}

 func BackDatadash0202sp(db *sql.DB, graphnbr string) (string, error) {
	log.Print("BackDatadash0202sp 01!\n")
    var resultado string
    resultado ="executed not OK"

    statement := fmt.Sprintf("    SELECT getdatadashboard01(5,2,'"+graphnbr+"',"+graphnbr+");")
 	log.Print("BackDatadash0202sp 02!\n")
    rows, err := db.Query(statement)
    log.Print("BackDatadash0202sp 02.1!\n")
    if err != nil {
        return resultado, err
    }
    log.Print("BackDatadash0202sp 02.5!\n")
    defer rows.Close()

    for rows.Next() {
        resultado ="executed OK"
    }

/*    datadashs := []Datadash{}

    for rows.Next() {
        resultado ="executed OK"

    	 log.Print("BackDatadash0202sp 03!\n")
        var u Datadash
        if err := rows.Scan(&u.ID, &u.Numsecuecialgrupoddatos, &u.Numsecuecial, &u.Nombrecolumna,&u.Createdat,&u.Valoramount, &u.Valorcolumna); err != nil {

//        if err := rows.Scan(&u.Token, &u.Bin); err != nil {
        	log.Print("BackDatadash0202sp err!\n"+err.Error())
            return resultado, err
        }
    	 log.Print("BackDatadash0202sp 03.5!\n")
        datadashs = append(datadashs, u)
    }
    */
    log.Print("BackDatadash0202sp 04!\n")
    return resultado, nil
}