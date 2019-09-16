package main

import (
	"net/http"
	"log"
	"banwire/dash/dashboard_front/db"
//	"banwire/dash/dashboard_front/net"
	modelito "banwire/dash/dashboard_front/model"
//	"time"
//	"encoding/json"
	 _ "github.com/lib/pq"   //use go get github.com/lib/pq

)


    

/////////////////////////v1


// v1ProcessDash02Grafica02  receive and handle the request from client, access DB
func v1ProcessDash02Grafica02(w http.ResponseWriter, requestData modelito.RequestDash02Grafica02) (string,string){
    	log.Print("v1ProcessDash02Grafica02 entre")

	defer func() {
		db.Connection.Close(nil)
	}()
    var result string
    var errorGeneral string
    var	errorGeneralNbr string



    var valoresParaResponder  []modelito.Datadash
    
    errorGeneral=""

	////////////////////////////////////////////////validate parms
	/// START
    
    if errorGeneral==""{//continue next step
    	log.Print("CZ   STEP Validate Parms")

		/// START
	        result ="OK realizarpago"+requestData.Dash0202reference+"    :    " +requestData.Dash0202reference2+"    :    " +requestData.Dash0202Dato01+"    :    " +requestData.Dash0202Dato02+"    :    " +requestData.Dash0202Dato03+"    :    "
		    log.Print("CZ    handler Listening: realiza consulta dash02diagram02:"+result)
		    
		     log.Print("CZ   STEP Validate paramters request")
		    errorGeneral= validaReqDash02Grafica02(requestData)   //in the logicrequest.go
		
		
		/// END

    }				    
		        
    if errorGeneral!="" {
    	//prepare response with error 100
    	log.Print("CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 - Missing parameter"	+errorGeneral
		errorGeneralNbr="100"
    }

	////////////////////////////////////////////////DB	
//A. first check if a most recent group is available

    var existMoreRecent string
    existMoreRecent="0"
	if errorGeneral==""{//continue next step

        log.Print("CZ   STEP Consume DB")
        if(Config_dbStringType=="mysql"){//mysql            
            existMoreRecent,errorGeneral =logicDBMysqlCheckMoreRecent0202(requestData, errorGeneral)  //in logicdbmysql.go
        }else{//postgres
                existMoreRecent,errorGeneral =logicDBCheckMoreRecent0202(requestData, errorGeneral)  //in logicdb.go            
        }
         
    }				    
    if errorGeneral!="" && errorGeneralNbr==""{
    	//prepare response with error 305
    	log.Print("CZ    Prepare Response with 350. Error checking more recent group data for dash02diag02:"+errorGeneral)
    	errorGeneral="ERROR:305 -  Error checking more recent group data for dash02diag02 -"	+errorGeneral
	    errorGeneralNbr="305"
    }
 

//B. get the most recent data group
	if errorGeneral=="" && existMoreRecent !="0" {//continue next step, only if more recent data group is available

       	    	log.Print("CZ   STEP Consume DB")
                   
                   requestData.Dash0202reference = existMoreRecent  //set the group number found, as reference
                   

                   log.Print("v1ProcessDash02Grafica02 grop found"+requestData.Dash0202reference)
                   if(Config_dbStringType=="mysql"){//mysql
                        if(requestData.Dash0202Dato01 == "99"){
                            valoresParaResponder,errorGeneral =logicDBMysqlProcessDash02Grafica02(requestData, errorGeneral)  //in logicdbmysql.go
                        }else{
                             if(requestData.Dash0202Dato01 == "24"){
                                    valoresParaResponder,errorGeneral =logicDBMysqlProcessDash02Grafica02(requestData, errorGeneral)  //in logicdbmysql.go
                              }else{
                                   if(requestData.Dash0202Dato01 == "48"){
                                        valoresParaResponder,errorGeneral =logicDBMysqlProcessDash02Grafica02(requestData, errorGeneral)  //in logicdbmysql.go
                                   }else{
                                       //error
                                   }//end 48
                              }//end 24
                        }//end 99
                        
                   }else{//postgres
                          //this logic was replaced and set in logicdb.go, using the values set in configuration.go and defeined in config.json  
                        //if(requestData.Dash0202Dato01 == "99"){
                            valoresParaResponder,errorGeneral =logicDBProcessDash02Grafica02(requestData, errorGeneral)  //in logicdb.go            
                        //}else{
                        //     if(requestData.Dash0202Dato01 == "24"){
                        //            valoresParaResponder,errorGeneral =logicDBProcessDash02Grafica02hours24(requestData, errorGeneral)  //in logicdb.go            
                        //      }else{
                        //           if(requestData.Dash0202Dato01 == "48"){
                        //                valoresParaResponder,errorGeneral =logicDBProcessDash02Grafica02hours48(requestData, errorGeneral)  //in logicdb.go            
                        //           }else{
                                       //error
                        //           }//end 48
                        //      }//end 24
                        //}//end 99


                   }
         


    }				    
    if errorGeneral!="" && errorGeneralNbr==""{
    	//prepare response with error 310
    	log.Print("CZ    Prepare Response with 310. Error obtaining data dash02diag02:"+errorGeneral)
    	errorGeneral="ERROR:310 -  Error obtaining data dash02diag02 -"	+errorGeneral
	    errorGeneralNbr="310"
    }

		        
	//response
    log.Print("CZ    handler DB Listening test v1ProcessDash02Grafica02  2")					

	//////////    format the response
    //mo
    if errorGeneral==""{//continue next step
        if existMoreRecent == "0"{  //no more recnet group data available
            fieldDataBytesJson,err := getJsonResponseDatadashNoRecentDataV1()   //logicresponse.go
            result ="OK get Dash0202reference: No more recent data"
            if err!=nil{
                log.Print("Eror en generando response")
                errorGeneral= err.Error()
            }		
            //////////    write the response
            w.Header().Set("Content-Type", "application/json")
            w.Write(fieldDataBytesJson)
        }else{   //more recent data available

            log.Print("CZ   STEP Validate Parms")
                /// START
            fieldDataBytesJson,err := getJsonResponseDatadashV1(valoresParaResponder)   //logicresponse.go
            
            log.Print("CZ    handler Listening test v1ProcessDash02Grafica02  3")	
            
            result ="OK get Dash0202reference: "+requestData.Dash0202reference+"resultado"
            if err!=nil{
                log.Print("Eror en generando response")
                errorGeneral= err.Error()
            }		
            //////////    write the response
            w.Header().Set("Content-Type", "application/json")
            w.Write(fieldDataBytesJson)
        }    

        
        log.Print("CZ    handler Listening test v1ProcessDash02Grafica02  4"+"<html><body>"+ result+"</body></html>")
                    
        
        /// END



    }				    
		 
    if errorGeneral!="" && errorGeneralNbr==""{//continue next step
    	log.Print("CZ   prepare the JSON response for ERROR")

	    //  START 
	    errorGeneral="ERROR:330 -Error preparing the response"	+errorGeneral
	    errorGeneralNbr="330"
	    //  END
    }

     return errorGeneral, errorGeneralNbr

}

