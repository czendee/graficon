package main

import (
//	"net/http"
	"log"
	"banwire/dash/dashboard_back/db"
//	"banwire/dash/dashboard_back/net"
	modelito "banwire/dash/dashboard_back/model"
//	"time"
//	"encoding/json"
	 _ "github.com/lib/pq"   //use go get github.com/lib/pq

)


func v1ProcessDashBack01Grafica01(graphnbr string) (string,string){
	defer func() {
		db.Connection.Close(nil)
	}()
    var errorGeneral string
    
    var dataObtainedFromSource []modelito.Datadash
    var valoresParaResponder  string

       	    	log.Print("CZ   STEP Consume DB")

                   if(Config_dbStringType_origin=="mysql"){//mysql
                        dataObtainedFromSource,errorGeneral =logicDBMysqlProcessDashBackGetDataFor01Grafica01(errorGeneral,graphnbr )  //in logicdbmysql.go
                   }else{//postgres
                         dataObtainedFromSource,errorGeneral =logicDBProcessDashBackGetDataFor01Grafica01( errorGeneral,graphnbr)  //in logicdb.go postgres           
                   }        
                   if(errorGeneral!=""){//handle error

                   }else{
                       //values found in the source db
                       //now, insert them in the dest db
                        if(Config_dbStringType=="mysql"){//mysql
                            valoresParaResponder,errorGeneral =logicDBMysqlProcessDashBackInsertData01Grafica01(errorGeneral,graphnbr,dataObtainedFromSource)  //in logicdbmysql.go
                        }else{//postgres
                            valoresParaResponder,errorGeneral =logicDBProcessDashBackInsertData01Grafica01( errorGeneral,graphnbr,dataObtainedFromSource)  //in logicdb.go  postgres          
                        }         

                   }
    return valoresParaResponder,errorGeneral
}
    

