package main

import (
//	"net/http"
	"log"
	"banwire/dash/dashboard_back/db"
//	"banwire/dash/dashboard_back/net"
//	modelito "banwire/dash/dashboard_back/model"
//	"time"
//	"encoding/json"
	 _ "github.com/lib/pq"   //use go get github.com/lib/pq

)


func v1ProcessDashBack01Grafica01() (string,string){
	defer func() {
		db.Connection.Close(nil)
	}()
    var errorGeneral string
    

    var valoresParaResponder  string

       	    	log.Print("CZ   STEP Consume DB")

                   if(Config_dbStringType=="mysql"){//mysql
                        valoresParaResponder,errorGeneral =logicDBMysqlProcessDashBack01Grafica01(errorGeneral)  //in logicdbmysql.go
                   }else{//postgres
                         valoresParaResponder,errorGeneral =logicDBProcessDashBack01Grafica01( errorGeneral)  //in logicdb.go            
                   }         
    return valoresParaResponder,errorGeneral
}
    

