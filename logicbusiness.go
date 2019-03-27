package main

import (
	"net/http"
	"log"
	"banwire/dash/dashboard_banwire/db"
//	"banwire/dash/dashboard_banwire/net"
	modelito "banwire/dash/dashboard_banwire/model"
//	"time"
//	"encoding/json"
	 _ "github.com/lib/pq"   //use go get github.com/lib/pq

)


    

/////////////////////////v1

// v1ProcessDash01Grafica01  receive and handle the request from client, access DB
func v1ProcessDash01Grafica01(w http.ResponseWriter, requestData modelito.RequestDash01Grafica01) (string,string){
	defer func() {
		db.Connection.Close(nil)
	}()
    var result string
    var errorGeneral string
    var	errorGeneralNbr string

//    var resultadoDash01Grafica01 modelito.ExitoDataDash01Grafica01

    var valoresParaResponder  []modelito.Card
    
    errorGeneral=""

	////////////////////////////////////////////////validate parms
	/// START
    
    if errorGeneral==""{//continue next step
    	log.Print("CZ   STEP Validate Parms")

		/// START
	        result ="OK realizarpago"+requestData.Dash0101reference+"    :    " +requestData.Dash0101reference2+"    :    " +requestData.Dash0101Dato01+"    :    " +requestData.Dash0101Dato02+"    :    " +requestData.Dash0101Dato03+"    :    "
		    log.Print("CZ    handler Listening test realizarpago:"+result)
		    
		     log.Print("CZ   STEP Validate paramters request")
		    errorGeneral= validaReqDash01Grafica01(requestData)   //in the logicrequest.go
		
		
		/// END

    }				    
		        
    if errorGeneral!="" {
    	//prepare response with error 100
    	log.Print("CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 - Missing parameter"	+errorGeneral
		errorGeneralNbr="100"
    }

	////////////////////////////////////////////////DB	
	if errorGeneral==""{//continue next step

       	    	log.Print("CZ   STEP Consume DB")
         valoresParaResponder,errorGeneral =logicDBProcessDash01Grafica01(requestData, errorGeneral) 


    }				    
    if errorGeneral!="" && errorGeneralNbr==""{
    	//prepare response with error 310
    	log.Print("CZ    Prepare Response with 310. Error obtaining cards:"+errorGeneral)
    	errorGeneral="ERROR:310 -  Error obtaining cards -"	+errorGeneral
	    errorGeneralNbr="310"
    }

		        
	//response
    log.Print("CZ    handler DB Listening test v1ProcessDash01Grafica01  2")					

	//////////    format the response
    if errorGeneral==""{//continue next step
		log.Print("CZ   STEP Validate Parms")
			/// START
		fieldDataBytesJson,err := getJsonResponseV2(valoresParaResponder)
		
		log.Print("CZ    handler Listening test v1ProcessDash01Grafica01  3")	
		
		result ="OK get Dash0101reference: "+requestData.Dash0101reference+"resultado"
		//////////    write the response
		w.Header().Set("Content-Type", "application/json")
		 w.Write(fieldDataBytesJson)
		 
		 log.Print("CZ    handler Listening test v1ProcessDash01Grafica01  4"+"<html><body>"+ result+"</body></html>")
			         
        if err!=nil{
        	log.Print("Eror en generando response")
            errorGeneral= err.Error()
        }		
		
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


// v1ProcessDash01Grafica02  receive and handle the request from client, access DB
func v1ProcessDash01Grafica02(w http.ResponseWriter, requestData modelito.RequestDash01Grafica02) (string,string){
	defer func() {
		db.Connection.Close(nil)
	}()
    var result string
    var errorGeneral string
    var	errorGeneralNbr string

//    var resultadoDash01Grafica01 modelito.ExitoDataDash01Grafica01

    var valoresParaResponder  []modelito.Card
    
    errorGeneral=""

	////////////////////////////////////////////////validate parms
	/// START
    
    if errorGeneral==""{//continue next step
    	log.Print("CZ   STEP Validate Parms")

		/// START
	        result ="OK realizarpago"+requestData.Dash0102reference+"    :    " +requestData.Dash0102reference2+"    :    " +requestData.Dash0102Dato01+"    :    " +requestData.Dash0102Dato02+"    :    " +requestData.Dash0102Dato03+"    :    "
		    log.Print("CZ    handler Listening test realizarpago:"+result)
		    
		     log.Print("CZ   STEP Validate paramters request")
		    errorGeneral= validaReqDash01Grafica02(requestData)   //in the logicrequest.go
		
		
		/// END

    }				    
		        
    if errorGeneral!="" {
    	//prepare response with error 100
    	log.Print("CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 - Missing parameter"	+errorGeneral
		errorGeneralNbr="100"
    }

	////////////////////////////////////////////////DB	
	if errorGeneral==""{//continue next step

       	    	log.Print("CZ   STEP Consume DB")
         valoresParaResponder,errorGeneral =logicDBProcessDash01Grafica02(requestData, errorGeneral) 


    }				    
    if errorGeneral!="" && errorGeneralNbr==""{
    	//prepare response with error 310
    	log.Print("CZ    Prepare Response with 310. Error obtaining cards:"+errorGeneral)
    	errorGeneral="ERROR:310 -  Error obtaining cards -"	+errorGeneral
	    errorGeneralNbr="310"
    }

		        
	//response
    log.Print("CZ    handler DB Listening test v1ProcessDash01Grafica02  2")					

	//////////    format the response
    if errorGeneral==""{//continue next step
		log.Print("CZ   STEP Validate Parms")
			/// START
		fieldDataBytesJson,err := getJsonResponseV2(valoresParaResponder)
		
		log.Print("CZ    handler Listening test v1ProcessDash01Grafica02  3")	
		
		result ="OK get Dash0101reference: "+requestData.Dash0102reference+"resultado"
		//////////    write the response
		w.Header().Set("Content-Type", "application/json")
		 w.Write(fieldDataBytesJson)
		 
		 log.Print("CZ    handler Listening test v1ProcessDash01Grafica02  4"+"<html><body>"+ result+"</body></html>")
			         
        if err!=nil{
        	log.Print("Eror en generando response")
            errorGeneral= err.Error()
        }		
		
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

// v1ProcessDash02Grafica01  receive and handle the request from client, access DB
func v1ProcessDash02Grafica01(w http.ResponseWriter, requestData modelito.RequestDash02Grafica01) (string,string){
	defer func() {
		db.Connection.Close(nil)
	}()
    var result string
    var errorGeneral string
    var	errorGeneralNbr string

//    var resultadoDash01Grafica01 modelito.ExitoDataDash02Grafica01

    var valoresParaResponder  []modelito.Datadash
    
    errorGeneral=""

	////////////////////////////////////////////////validate parms
	/// START
    
    if errorGeneral==""{//continue next step
    	log.Print("CZ   STEP Validate Parms")

		/// START
	        result ="OK realizarpago"+requestData.Dash0201reference+"    :    " +requestData.Dash0201reference2+"    :    " +requestData.Dash0201Dato01+"    :    " +requestData.Dash0201Dato02+"    :    " +requestData.Dash0201Dato03+"    :    "
		    log.Print("CZ    handler Listening test realizarpago:"+result)
		    
		     log.Print("CZ   STEP Validate paramters request")
		    errorGeneral= validaReqDash02Grafica01(requestData)   //in the logicrequest.go
		
		
		/// END

    }				    
		        
    if errorGeneral!="" {
    	//prepare response with error 100
    	log.Print("CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 - Missing parameter"	+errorGeneral
		errorGeneralNbr="100"
    }

	////////////////////////////////////////////////DB	
	if errorGeneral==""{//continue next step

       	    	log.Print("CZ   STEP Consume DB")
                   if(Config_dbStringType=="mysql"){//mysql
                        valoresParaResponder,errorGeneral =logicDBMysqlProcessDash02Grafica01(requestData, errorGeneral)  //in logicdbmysql.go
                   }else{//postgres
                         valoresParaResponder,errorGeneral =logicDBProcessDash02Grafica01(requestData, errorGeneral)  //in logicdb.go            
                   }
         


    }				    
    if errorGeneral!="" && errorGeneralNbr==""{
    	//prepare response with error 310
    	log.Print("CZ    Prepare Response with 310. Error obtaining cards:"+errorGeneral)
    	errorGeneral="ERROR:310 -  Error obtaining cards -"	+errorGeneral
	    errorGeneralNbr="310"
    }

		        
	//response
    log.Print("CZ    handler DB Listening test v1ProcessDash02Grafica01  2")					

	//////////    format the response
    if errorGeneral==""{//continue next step
		log.Print("CZ   STEP Validate Parms")
			/// START
		fieldDataBytesJson,err := getJsonResponseDatadashV1(valoresParaResponder)
		
		log.Print("CZ    handler Listening test v1ProcessDash02Grafica01  3")	
		
		result ="OK get Dash0101reference: "+requestData.Dash0201reference+"resultado"
		//////////    write the response
		w.Header().Set("Content-Type", "application/json")
		 w.Write(fieldDataBytesJson)
		 
		 log.Print("CZ    handler Listening test v1ProcessDash02Grafica01  4"+"<html><body>"+ result+"</body></html>")
			         
        if err!=nil{
        	log.Print("Eror en generando response")
            errorGeneral= err.Error()
        }		
		
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

// v1ProcessDash02Grafica02  receive and handle the request from client, access DB
func v1ProcessDash02Grafica02(w http.ResponseWriter, requestData modelito.RequestDash02Grafica02) (string,string){
	defer func() {
		db.Connection.Close(nil)
	}()
    var result string
    var errorGeneral string
    var	errorGeneralNbr string

//    var resultadoDash01Grafica02 modelito.ExitoDataDash02Grafica02

    var valoresParaResponder  []modelito.Card
    
    errorGeneral=""

	////////////////////////////////////////////////validate parms
	/// START
    
    if errorGeneral==""{//continue next step
    	log.Print("CZ   STEP Validate Parms")

		/// START
	        result ="OK realizarpago"+requestData.Dash0202reference+"    :    " +requestData.Dash0202reference2+"    :    " +requestData.Dash0202Dato01+"    :    " +requestData.Dash0202Dato02+"    :    " +requestData.Dash0202Dato03+"    :    "
		    log.Print("CZ    handler Listening test realizarpago:"+result)
		    
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
	if errorGeneral==""{//continue next step

       	    	log.Print("CZ   STEP Consume DB")
         valoresParaResponder,errorGeneral =logicDBProcessDash02Grafica02(requestData, errorGeneral) 


    }				    
    if errorGeneral!="" && errorGeneralNbr==""{
    	//prepare response with error 310
    	log.Print("CZ    Prepare Response with 310. Error obtaining cards:"+errorGeneral)
    	errorGeneral="ERROR:310 -  Error obtaining cards -"	+errorGeneral
	    errorGeneralNbr="310"
    }

		        
	//response
    log.Print("CZ    handler DB Listening test v1ProcessDash02Grafica02  2")					

	//////////    format the response
    if errorGeneral==""{//continue next step
		log.Print("CZ   STEP Validate Parms")
			/// START
		fieldDataBytesJson,err := getJsonResponseV2(valoresParaResponder)
		
		log.Print("CZ    handler Listening test v1ProcessDash02Grafica02  3")	
		
		result ="OK get Dash0101reference: "+requestData.Dash0202reference+"resultado"
		//////////    write the response
		w.Header().Set("Content-Type", "application/json")
		 w.Write(fieldDataBytesJson)
		 
		 log.Print("CZ    handler Listening test v1ProcessDash02Grafica02  4"+"<html><body>"+ result+"</body></html>")
			         
        if err!=nil{
        	log.Print("Eror en generando response")
            errorGeneral= err.Error()
        }		
		
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

// v1ProcessDash02Grafica03  receive and handle the request from client, access DB
func v1ProcessDash02Grafica03(w http.ResponseWriter, requestData modelito.RequestDash02Grafica03) (string,string){
	defer func() {
		db.Connection.Close(nil)
	}()
    var result string
    var errorGeneral string
    var	errorGeneralNbr string

//    var resultadoDash01Grafica03 modelito.ExitoDataDash02Grafica03

    var valoresParaResponder  []modelito.Card
    
    errorGeneral=""

	////////////////////////////////////////////////validate parms
	/// START
    
    if errorGeneral==""{//continue next step
    	log.Print("CZ   STEP Validate Parms")

		/// START
	        result ="OK realizarpago"+requestData.Dash0203reference+"    :    " +requestData.Dash0203reference2+"    :    " +requestData.Dash0203Dato01+"    :    " +requestData.Dash0203Dato02+"    :    " +requestData.Dash0203Dato03+"    :    "
		    log.Print("CZ    handler Listening test realizarpago:"+result)
		    
		     log.Print("CZ   STEP Validate paramters request")
		    errorGeneral= validaReqDash02Grafica03(requestData)   //in the logicrequest.go
		
		
		/// END

    }				    
		        
    if errorGeneral!="" {
    	//prepare response with error 100
    	log.Print("CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 - Missing parameter"	+errorGeneral
		errorGeneralNbr="100"
    }

	////////////////////////////////////////////////DB	
	if errorGeneral==""{//continue next step

       	    	log.Print("CZ   STEP Consume DB")
         valoresParaResponder,errorGeneral =logicDBProcessDash02Grafica03(requestData, errorGeneral) 


    }				    
    if errorGeneral!="" && errorGeneralNbr==""{
    	//prepare response with error 310
    	log.Print("CZ    Prepare Response with 310. Error obtaining cards:"+errorGeneral)
    	errorGeneral="ERROR:310 -  Error obtaining cards -"	+errorGeneral
	    errorGeneralNbr="310"
    }

		        
	//response
    log.Print("CZ    handler DB Listening test v1ProcessDash02Grafica03  2")					

	//////////    format the response
    if errorGeneral==""{//continue next step
		log.Print("CZ   STEP Validate Parms")
			/// START
		fieldDataBytesJson,err := getJsonResponseV2(valoresParaResponder)
		
		log.Print("CZ    handler Listening test v1ProcessDash02Grafica03  3")	
		
		result ="OK get Dash0203reference: "+requestData.Dash0203reference+"resultado"
		//////////    write the response
		w.Header().Set("Content-Type", "application/json")
		 w.Write(fieldDataBytesJson)
		 
		 log.Print("CZ    handler Listening test v1ProcessDash02Grafica03  4"+"<html><body>"+ result+"</body></html>")
			         
        if err!=nil{
        	log.Print("Eror en generando response")
            errorGeneral= err.Error()
        }		
		
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

// v1ProcessDash03Grafica01  receive and handle the request from client, access DB
func v1ProcessDash03Grafica01(w http.ResponseWriter, requestData modelito.RequestDash03Grafica01) (string,string){
	defer func() {
		db.Connection.Close(nil)
	}()
    var result string
    var errorGeneral string
    var	errorGeneralNbr string

//    var resultadoDash03Grafica01 modelito.ExitoDataDash03Grafica01

    var valoresParaResponder  []modelito.Card
    
    errorGeneral=""

	////////////////////////////////////////////////validate parms
	/// START
    
    if errorGeneral==""{//continue next step
    	log.Print("CZ   STEP Validate Parms")

		/// START
	        result ="OK realizarpago"+requestData.Dash0301reference+"    :    " +requestData.Dash0301reference2+"    :    " +requestData.Dash0301Dato01+"    :    " +requestData.Dash0301Dato02+"    :    " +requestData.Dash0301Dato03+"    :    "
		    log.Print("CZ    handler Listening test realizarpago:"+result)
		    
		     log.Print("CZ   STEP Validate paramters request")
		    errorGeneral= validaReqDash03Grafica01(requestData)   //in the logicrequest.go
		
		
		/// END

    }				    
		        
    if errorGeneral!="" {
    	//prepare response with error 100
    	log.Print("CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 - Missing parameter"	+errorGeneral
		errorGeneralNbr="100"
    }

	////////////////////////////////////////////////DB	
	if errorGeneral==""{//continue next step

       	    	log.Print("CZ   STEP Consume DB")
         valoresParaResponder,errorGeneral =logicDBProcessDash03Grafica01(requestData, errorGeneral) 


    }				    
    if errorGeneral!="" && errorGeneralNbr==""{
    	//prepare response with error 310
    	log.Print("CZ    Prepare Response with 310. Error obtaining cards:"+errorGeneral)
    	errorGeneral="ERROR:310 -  Error obtaining cards -"	+errorGeneral
	    errorGeneralNbr="310"
    }

		        
	//response
    log.Print("CZ    handler DB Listening test v1ProcessDash02Grafica01  2")					

	//////////    format the response
    if errorGeneral==""{//continue next step
		log.Print("CZ   STEP Validate Parms")
			/// START
		fieldDataBytesJson,err := getJsonResponseV2(valoresParaResponder)
		
		log.Print("CZ    handler Listening test v1ProcessDash03Grafica01  3")	
		
		result ="OK get Dash0301reference: "+requestData.Dash0301reference+"resultado"
		//////////    write the response
		w.Header().Set("Content-Type", "application/json")
		 w.Write(fieldDataBytesJson)
		 
		 log.Print("CZ    handler Listening test v1ProcessDash03Grafica01  4"+"<html><body>"+ result+"</body></html>")
			         
        if err!=nil{
        	log.Print("Eror en generando response")
            errorGeneral= err.Error()
        }		
		
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

// v1ProcessDash03Grafica02  receive and handle the request from client, access DB
func v1ProcessDash03Grafica02(w http.ResponseWriter, requestData modelito.RequestDash03Grafica02) (string,string){
	defer func() {
		db.Connection.Close(nil)
	}()
    var result string
    var errorGeneral string
    var	errorGeneralNbr string

//    var resultadoDash03Grafica02 modelito.ExitoDataDash03Grafica02

    var valoresParaResponder  []modelito.Card
    
    errorGeneral=""

	////////////////////////////////////////////////validate parms
	/// START
    
    if errorGeneral==""{//continue next step
    	log.Print("CZ   STEP Validate Parms")

		/// START
	        result ="OK realizarpago"+requestData.Dash0302reference+"    :    " +requestData.Dash0302reference2+"    :    " +requestData.Dash0302Dato01+"    :    " +requestData.Dash0302Dato02+"    :    " +requestData.Dash0302Dato03+"    :    "
		    log.Print("CZ    handler Listening test realizarpago:"+result)
		    
		     log.Print("CZ   STEP Validate paramters request")
		    errorGeneral= validaReqDash03Grafica02(requestData)   //in the logicrequest.go
		
		
		/// END

    }				    
		        
    if errorGeneral!="" {
    	//prepare response with error 100
    	log.Print("CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 - Missing parameter"	+errorGeneral
		errorGeneralNbr="100"
    }

	////////////////////////////////////////////////DB	
	if errorGeneral==""{//continue next step

       	    	log.Print("CZ   STEP Consume DB")
         valoresParaResponder,errorGeneral =logicDBProcessDash03Grafica02(requestData, errorGeneral) 


    }				    
    if errorGeneral!="" && errorGeneralNbr==""{
    	//prepare response with error 310
    	log.Print("CZ    Prepare Response with 310. Error obtaining cards:"+errorGeneral)
    	errorGeneral="ERROR:310 -  Error obtaining cards -"	+errorGeneral
	    errorGeneralNbr="310"
    }

		        
	//response
    log.Print("CZ    handler DB Listening test v1ProcessDash03Grafica02  2")					

	//////////    format the response
    if errorGeneral==""{//continue next step
		log.Print("CZ   STEP Validate Parms")
			/// START
		fieldDataBytesJson,err := getJsonResponseV2(valoresParaResponder)
		
		log.Print("CZ    handler Listening test v1ProcessDash03Grafica02  3")	
		
		result ="OK get Dash0302reference: "+requestData.Dash0302reference+"resultado"
		//////////    write the response
		w.Header().Set("Content-Type", "application/json")
		 w.Write(fieldDataBytesJson)
		 
		 log.Print("CZ    handler Listening test v1ProcessDash03Grafica02  4"+"<html><body>"+ result+"</body></html>")
			         
        if err!=nil{
        	log.Print("Eror en generando response")
            errorGeneral= err.Error()
        }		
		
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

// v1ProcessDash03rafica03  receive and handle the request from client, access DB
func v1ProcessDash03Grafica03(w http.ResponseWriter, requestData modelito.RequestDash03Grafica03) (string,string){
	defer func() {
		db.Connection.Close(nil)
	}()
    var result string
    var errorGeneral string
    var	errorGeneralNbr string

//    var resultadoDash03Grafica03 modelito.ExitoDataDash03Grafica03

    var valoresParaResponder  []modelito.Card
    
    errorGeneral=""

	////////////////////////////////////////////////validate parms
	/// START
    
    if errorGeneral==""{//continue next step
    	log.Print("CZ   STEP Validate Parms")

		/// START
	        result ="OK realizarpago"+requestData.Dash0303reference+"    :    " +requestData.Dash0303reference2+"    :    " +requestData.Dash0303Dato01+"    :    " +requestData.Dash0303Dato02+"    :    " +requestData.Dash0303Dato03+"    :    "
		    log.Print("CZ    handler Listening test realizarpago:"+result)
		    
		     log.Print("CZ   STEP Validate paramters request")
		    errorGeneral= validaReqDash03Grafica03(requestData)   //in the logicrequest.go
		
		
		/// END

    }				    
		        
    if errorGeneral!="" {
    	//prepare response with error 100
    	log.Print("CZ    Prepare Response with 100. Missing parameter:"+errorGeneral)
    	errorGeneral="ERROR:100 - Missing parameter"	+errorGeneral
		errorGeneralNbr="100"
    }

	////////////////////////////////////////////////DB	
	if errorGeneral==""{//continue next step

       	    	log.Print("CZ   STEP Consume DB")
         valoresParaResponder,errorGeneral =logicDBProcessDash03Grafica03(requestData, errorGeneral) 


    }				    
    if errorGeneral!="" && errorGeneralNbr==""{
    	//prepare response with error 310
    	log.Print("CZ    Prepare Response with 310. Error obtaining cards:"+errorGeneral)
    	errorGeneral="ERROR:310 -  Error obtaining cards -"	+errorGeneral
	    errorGeneralNbr="310"
    }

		        
	//response
    log.Print("CZ    handler DB Listening test v1ProcessDash03Grafica03  2")					

	//////////    format the response
    if errorGeneral==""{//continue next step
		log.Print("CZ   STEP Validate Parms")
			/// START
		fieldDataBytesJson,err := getJsonResponseV2(valoresParaResponder)
		
		log.Print("CZ    handler Listening test v1ProcessDash03Grafica03  3")	
		
		result ="OK get Dash0303reference: "+requestData.Dash0303reference+"resultado"
		//////////    write the response
		w.Header().Set("Content-Type", "application/json")
		 w.Write(fieldDataBytesJson)
		 
		 log.Print("CZ    handler Listening test v1ProcessDash03Grafica03  4"+"<html><body>"+ result+"</body></html>")
			         
        if err!=nil{
        	log.Print("Eror en generando response")
            errorGeneral= err.Error()
        }		
		
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
