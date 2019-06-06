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


func BackMysqlDatadash0201(db *sql.DB,comandosqlorigen_origin string) ([]Datadash, error) {
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


func BackMysqlInsertDatadash0201(db *sql.DB,graphnbr string, valoresToInsert []Datadash) (string, error) {
	log.Print("BackInsertDatadash0201 01!\n")
    var resultado string
    resultado ="executed not OK"
    

    
    secuencia :=0
    for _, d := range valoresToInsert {
        secuencia =secuencia +1
        statement := fmt.Sprintf("insert into banwiredash02graph02 (numsecuecial_grupoddatos, "+
         "numsecuecial,  "+
         "nombrecolumna, "+
         "created_at,     "+
         "valorcolumna,   "+
          "valoramount,   "+
           "graphnbr)     "+
        "VALUES(          "+
        "    (select max(numsecuecial_grupoddatos)+1  from banwiredash02graph01 p where p.graphnbr ='%s'),    "+
        "     %s,          "+
        "    '%s',       "+
         "    current_timestamp,   "+
         "    '%s',       "+
         "     %s)",  
              graphnbr,
              secuencia,
              d.Nombrecolumna,
              d.Valorcolumna,
              d.Valoramount    )

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
    


    log.Print("GetInsertDatadash0201 04!\n"+resultado)
    return resultado, nil
}