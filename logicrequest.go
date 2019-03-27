package main

import (
	"net/http"
	"log"
	"encoding/json"
	modelito "banwire/dash/dashboard_banwire/model"
	 _ "github.com/lib/pq"   //use go get github.com/lib/pq
	 
)

///////////////7//get


   func obtainParmsGettokenizedcards(r *http.Request, errorGeneral string )(modelito.RequestTokenizedCards, string){
   	var requestData modelito.RequestTokenizedCards
	////////////////////////////////////////////////obtain parms in JSON
   //START    
    log.Print("cz  handleDBGettokenizedcards")

	    log.Print("CZ    handlerDB Listening test obtienetarjetastokenizadas")
	    
    	err := r.ParseForm()
		if err != nil {
	    	//prepare response with error 100
	    	log.Print("CZ    Prepare Response with 380. Missing parameter:"+errorGeneral)
	    	errorGeneral="ERROR:380 -"	+err.Error()
		}
		v := r.Form
		requestData.Cardreference = v.Get("cardreference")

    //END
   	
   	 return  requestData, errorGeneral
   }



   func obtainParmsProcessPayment(r *http.Request, errorGeneral string) (modelito.RequestPayment,string){
   	 var requestData modelito.RequestPayment
	////////////////////////////////////////////////obtain parms in JSON
   //START    
    log.Print("cz  handleProcesspayment")
 		 log.Print("CZ    handler Listening test realizarpago")
    		    
    	err := r.ParseForm()
		if err != nil {
	    	//prepare response with error 100
	    	log.Print("CZ    Prepare Response with 180. Missing parameter:"+errorGeneral)
	    	errorGeneral="ERROR:180 -"	+err.Error()
		}
		v := r.Form
		requestData.Clientreference = v.Get("clientreference")
		requestData.Paymentreference = v.Get("paymentreference")
		requestData.Token = v.Get("token")
		requestData.Cvv = v.Get("cvv")
		requestData.Amount = v.Get("amount")

   //END
   	 
   	 return requestData,errorGeneral
   }

   func obtainParmsGeneratetokenized(r *http.Request, errorGeneral string) (modelito.RequestTokenized,string) {
	////////////////////////////////////////////////obtain parms in JSON
   //START    
     var requestData modelito.RequestTokenized
    log.Print("cz  handleGeneratetokenized")
	    log.Print("CZ    handler Listening test handleGeneratetokenized")
    		    
    	err := r.ParseForm()
		if err != nil {
	    	//prepare response with error 280
	    	log.Print("CZ    Prepare Response with 280. Error in JSON Request:"+errorGeneral)
	    	errorGeneral="ERROR :280 -Error in JSON Request-"	+err.Error()
		}
		v := r.Form
		requestData.Clientreference = v.Get("clientreference")
		requestData.Paymentreference = v.Get("paymentreference")
		requestData.Card = v.Get("card")
		requestData.Exp = v.Get("exp")
		requestData.Cvv = v.Get("cvv")

   //END
   	  return  requestData, errorGeneral
   }

////////////////////////Post



   func obtainPostParmsGettokenizedcards(r *http.Request, errorGeneral string )(modelito.RequestTokenizedCards, string){
   	var requestData modelito.RequestTokenizedCards
	////////////////////////////////////////////////obtain parms in JSON
   //START    
    log.Print("cz  handleDBGettokenizedcards")

	    log.Print("CZ    handlerDB Listening test obtienetarjetastokenizadas")
	 
 			decoder := json.NewDecoder(r.Body)
		
			err := decoder.Decode(&requestData)
			if err != nil {
		    	log.Print("CZ    Prepare Response with 380. JSON format/Missing parameter:"+errorGeneral)
		    	errorGeneral="ERROR:380 -Input JSON format/Missing parameter"	+err.Error()

			}
		
			//post   cardreference := requestData.Cardreference

    //END
   	
   	 return  requestData, errorGeneral
   }



   func obtainPostParmsProcessPayment(r *http.Request, errorGeneral string) (modelito.RequestPayment,string){
   	 var requestData modelito.RequestPayment
	////////////////////////////////////////////////obtain parms in JSON
   //START    
    log.Print("cz  handleProcesspayment")
 		 log.Print("CZ    handler Listening test realizarpago")
 			decoder := json.NewDecoder(r.Body)
		
			err := decoder.Decode(&requestData)
			if err != nil {
		    	log.Print("CZ    Prepare Response with 180. JSON format/Missing parameter:"+errorGeneral)
		    	errorGeneral="ERROR:180 -Input JSON format/Missing parameter"	+err.Error()

			}

   //END
   	 
   	 return requestData,errorGeneral
   }

   func obtainPostParmsGeneratetokenized(r *http.Request, errorGeneral string) (modelito.RequestTokenized,string) {
	////////////////////////////////////////////////obtain parms in JSON
   //START    
     var requestData modelito.RequestTokenized
    log.Print("cz  handleGeneratetokenized")
	    log.Print("CZ    handler Listening test handleGeneratetokenized")
 			decoder := json.NewDecoder(r.Body)
		
			err := decoder.Decode(&requestData)
			if err != nil {
		    	log.Print("CZ    Prepare Response with 280. JSON format/Missing parameter:"+errorGeneral)
		    	errorGeneral="ERROR:280 -Input JSON format/Missing parameter"	+err.Error()

			}

   //END
   	  return  requestData, errorGeneral
   }


////////////////////////validate input params


 	    func validaReqDash01Grafica01( parRequestData modelito.RequestDash01Grafica01) string {
            var resultado string
            
            	if parRequestData.Dash0101reference != "" {

				}else{
					resultado=" reference is required"
		        }
            	if parRequestData.Dash0101reference2 != "" {

				}else{
					resultado=" reference2 is required"
		        }

            	if parRequestData.Dash0101Dato01 != "" {

				}else{
					resultado=" Dato01 is required"
		        }

            	if parRequestData.Dash0101Dato02 != "" {

				}else{
					resultado=" Dato02 is required"
		        }

            	if parRequestData.Dash0101Dato03 != "" {

				}else{
					resultado=" Dato03 is required"
		        }

			/// END

            return resultado
	    }




 	    func validaReqDash01Grafica02( parRequestData modelito.RequestDash01Grafica02) string {
            var resultado string
            
            	if parRequestData.Dash0102reference != "" {

				}else{
					resultado=" reference is required"
		        }
            	if parRequestData.Dash0102reference2 != "" {

				}else{
					resultado=" reference2 is required"
		        }

            	if parRequestData.Dash0102Dato01 != "" {

				}else{
					resultado=" Dato01 is required"
		        }

            	if parRequestData.Dash0102Dato02 != "" {

				}else{
					resultado=" Dato02 is required"
		        }

            	if parRequestData.Dash0102Dato03 != "" {

				}else{
					resultado=" Dato03 is required"
		        }

			/// END

            return resultado
	    }

 	    func validaReqDash02Grafica01( parRequestData modelito.RequestDash02Grafica01) string {
            var resultado string
            
            	if parRequestData.Dash0201reference != "" {

				}else{
					resultado=" reference is required"
		        }
            	if parRequestData.Dash0201reference2 != "" {

				}else{
					resultado=" reference2 is required"
		        }

            	if parRequestData.Dash0201Dato01 != "" {

				}else{
					resultado=" Dato01 is required"
		        }

            	if parRequestData.Dash0201Dato02 != "" {

				}else{
					resultado=" Dato02 is required"
		        }

            	if parRequestData.Dash0201Dato03 != "" {

				}else{
					resultado=" Dato03 is required"
		        }

			/// END

            return resultado
	    }


 	    func validaReqDash02Grafica02( parRequestData modelito.RequestDash02Grafica02) string {
            var resultado string
            
            	if parRequestData.Dash0202reference != "" {

				}else{
					resultado=" reference is required"
		        }
            	if parRequestData.Dash0202reference2 != "" {

				}else{
					resultado=" reference2 is required"
		        }

            	if parRequestData.Dash0202Dato01 != "" {

				}else{
					resultado=" Dato01 is required"
		        }

            	if parRequestData.Dash0202Dato02 != "" {

				}else{
					resultado=" Dato02 is required"
		        }

            	if parRequestData.Dash0202Dato03 != "" {

				}else{
					resultado=" Dato03 is required"
		        }

			/// END

            return resultado
	    }

 	    func validaReqDash02Grafica03( parRequestData modelito.RequestDash02Grafica03) string {
            var resultado string
            
            	if parRequestData.Dash0203reference != "" {

				}else{
					resultado=" reference is required"
		        }
            	if parRequestData.Dash0203reference2 != "" {

				}else{
					resultado=" reference2 is required"
		        }

            	if parRequestData.Dash0203Dato01 != "" {

				}else{
					resultado=" Dato01 is required"
		        }

            	if parRequestData.Dash0203Dato02 != "" {

				}else{
					resultado=" Dato02 is required"
		        }

            	if parRequestData.Dash0203Dato03 != "" {

				}else{
					resultado=" Dato03 is required"
		        }

			/// END

            return resultado
	    }


 	    func validaReqDash03Grafica01( parRequestData modelito.RequestDash03Grafica01) string {
            var resultado string
            
            	if parRequestData.Dash0301reference != "" {

				}else{
					resultado=" reference is required"
		        }
            	if parRequestData.Dash0301reference2 != "" {

				}else{
					resultado=" reference2 is required"
		        }

            	if parRequestData.Dash0301Dato01 != "" {

				}else{
					resultado=" Dato01 is required"
		        }

            	if parRequestData.Dash0301Dato02 != "" {

				}else{
					resultado=" Dato02 is required"
		        }

            	if parRequestData.Dash0301Dato03 != "" {

				}else{
					resultado=" Dato03 is required"
		        }

			/// END

            return resultado
	    }


 	    func validaReqDash03Grafica02( parRequestData modelito.RequestDash03Grafica02) string {
            var resultado string
            
            	if parRequestData.Dash0302reference != "" {

				}else{
					resultado=" reference is required"
		        }
            	if parRequestData.Dash0302reference2 != "" {

				}else{
					resultado=" reference2 is required"
		        }

            	if parRequestData.Dash0302Dato01 != "" {

				}else{
					resultado=" Dato01 is required"
		        }

            	if parRequestData.Dash0302Dato02 != "" {

				}else{
					resultado=" Dato02 is required"
		        }

            	if parRequestData.Dash0302Dato03 != "" {

				}else{
					resultado=" Dato03 is required"
		        }

			/// END

            return resultado
	    }

 	    func validaReqDash03Grafica03( parRequestData modelito.RequestDash03Grafica03) string {
            var resultado string
            
            	if parRequestData.Dash0303reference != "" {

				}else{
					resultado=" reference is required"
		        }
            	if parRequestData.Dash0303reference2 != "" {

				}else{
					resultado=" reference2 is required"
		        }

            	if parRequestData.Dash0303Dato01 != "" {

				}else{
					resultado=" Dato01 is required"
		        }

            	if parRequestData.Dash0303Dato02 != "" {

				}else{
					resultado=" Dato02 is required"
		        }

            	if parRequestData.Dash0303Dato03 != "" {

				}else{
					resultado=" Dato03 is required"
		        }

			/// END

            return resultado
	    }




///////////////////////////////////////7

   func obtainParmsProcessDash01Grafica01(r *http.Request, errorGeneral string) (modelito.RequestDash01Grafica01,string){
   	 var requestData modelito.RequestDash01Grafica01
	////////////////////////////////////////////////obtain parms in JSON
   //START    
    log.Print("cz  obtainParmsProcessDash01Grafica01")
 		 log.Print("CZ    handler Listening test obtainParmsProcessDash01Grafica01")
    		    
    	err := r.ParseForm()
		if err != nil {
	    	//prepare response with error 100
	    	log.Print("CZ    Prepare Response with 180. Missing parameter:"+errorGeneral)
	    	errorGeneral="ERROR:180 -"	+err.Error()
		}
		v := r.Form
		requestData.Dash0101reference = v.Get("reference")
		requestData.Dash0101reference2 = v.Get("reference2")
		requestData.Dash0101Dato01 = v.Get("dato01")
		requestData.Dash0101Dato02 = v.Get("dato02")
		requestData.Dash0101Dato03 = v.Get("dato03")

   //END
   	 
   	 return requestData,errorGeneral
   }



   func obtainPostParmsProcessDash01Grafica01(r *http.Request, errorGeneral string) (modelito.RequestDash01Grafica01,string){
   	 var requestData modelito.RequestDash01Grafica01
	////////////////////////////////////////////////obtain parms in JSON
   //START    
    log.Print("cz  obtainPostParmsProcessDash01Grafica01")
 		 log.Print("CZ    handler Listening test obtainPostParmsProcessDash01Grafica01")
 			decoder := json.NewDecoder(r.Body)
		
			err := decoder.Decode(&requestData)
			if err != nil {
		    	log.Print("CZ    Prepare Response with 180. JSON format/Missing parameter:"+errorGeneral)
		    	errorGeneral="ERROR:180 -Input JSON format/Missing parameter"	+err.Error()

			}

   //END
   	 
   	 return requestData,errorGeneral
   }




   func obtainParmsProcessDash01Grafica02(r *http.Request, errorGeneral string) (modelito.RequestDash01Grafica02,string){
   	 var requestData modelito.RequestDash01Grafica02
	////////////////////////////////////////////////obtain parms in JSON
   //START    
    log.Print("cz  obtainParmsProcessDash01Grafica02")
 		 log.Print("CZ    handler Listening test obtainParmsProcessDash01Grafica02")
    		    
    	err := r.ParseForm()
		if err != nil {
	    	//prepare response with error 100
	    	log.Print("CZ    Prepare Response with 180. Missing parameter:"+errorGeneral)
	    	errorGeneral="ERROR:180 -"	+err.Error()
		}
		v := r.Form
		requestData.Dash0102reference = v.Get("reference")
		requestData.Dash0102reference2 = v.Get("reference2")
		requestData.Dash0102Dato01 = v.Get("dato01")
		requestData.Dash0102Dato02 = v.Get("dato02")
		requestData.Dash0102Dato03 = v.Get("dato03")

   //END
   	 
   	 return requestData,errorGeneral
   }



   func obtainPostParmsProcessDash01Grafica02(r *http.Request, errorGeneral string) (modelito.RequestDash01Grafica02,string){
   	 var requestData modelito.RequestDash01Grafica02
	////////////////////////////////////////////////obtain parms in JSON
   //START    
    log.Print("cz  obtainPostParmsProcessDash01Grafica02")
 		 log.Print("CZ    handler Listening test obtainPostParmsProcessDash01Grafica02")
 			decoder := json.NewDecoder(r.Body)
		
			err := decoder.Decode(&requestData)
			if err != nil {
		    	log.Print("CZ    Prepare Response with 180. JSON format/Missing parameter:"+errorGeneral)
		    	errorGeneral="ERROR:180 -Input JSON format/Missing parameter"	+err.Error()

			}

   //END
   	 
   	 return requestData,errorGeneral
   }

   func obtainParmsProcessDash02Grafica01(r *http.Request, errorGeneral string) (modelito.RequestDash02Grafica01,string){
   	 var requestData modelito.RequestDash02Grafica01
	////////////////////////////////////////////////obtain parms in JSON
   //START    
    log.Print("cz  obtainParmsProcessDash02Grafica01")
 		 log.Print("CZ    handler Listening test obtainParmsProcessDash02Grafica01")
    		    
    	err := r.ParseForm()
		if err != nil {
	    	//prepare response with error 100
	    	log.Print("CZ    Prepare Response with 180. Missing parameter:"+errorGeneral)
	    	errorGeneral="ERROR:180 -"	+err.Error()
		}
		v := r.Form
		requestData.Dash0201reference = v.Get("reference")
		requestData.Dash0201reference2 = v.Get("reference2")
		requestData.Dash0201Dato01 = v.Get("dato01")
		requestData.Dash0201Dato02 = v.Get("dato02")
		requestData.Dash0201Dato03 = v.Get("dato03")

   //END
   	 
   	 return requestData,errorGeneral
   }

   func obtainParmsProcessDash02Grafica02(r *http.Request, errorGeneral string) (modelito.RequestDash02Grafica02,string){
   	 var requestData modelito.RequestDash02Grafica02
	////////////////////////////////////////////////obtain parms in JSON
   //START    
    log.Print("cz  obtainParmsProcessDash02Grafica02")
 		 log.Print("CZ    handler Listening test obtainParmsProcessDash02Grafica02")
    		    
    	err := r.ParseForm()
		if err != nil {
	    	//prepare response with error 100
	    	log.Print("CZ    Prepare Response with 180. Missing parameter:"+errorGeneral)
	    	errorGeneral="ERROR:180 -"	+err.Error()
		}
		v := r.Form
		requestData.Dash0202reference = v.Get("reference")
		requestData.Dash0202reference2 = v.Get("reference2")
		requestData.Dash0202Dato01 = v.Get("dato01")
		requestData.Dash0202Dato02 = v.Get("dato02")
		requestData.Dash0202Dato03 = v.Get("dato03")

   //END
   	 
   	 return requestData,errorGeneral
   }

   func obtainParmsProcessDash02Grafica03(r *http.Request, errorGeneral string) (modelito.RequestDash02Grafica03,string){
   	 var requestData modelito.RequestDash02Grafica03
	////////////////////////////////////////////////obtain parms in JSON
   //START    
    log.Print("cz  obtainParmsProcessDash02Grafica03")
 		 log.Print("CZ    handler Listening test obtainParmsProcessDash02Grafica03")
    		    
    	err := r.ParseForm()
		if err != nil {
	    	//prepare response with error 100
	    	log.Print("CZ    Prepare Response with 180. Missing parameter:"+errorGeneral)
	    	errorGeneral="ERROR:180 -"	+err.Error()
		}
		v := r.Form
		requestData.Dash0203reference = v.Get("reference")
		requestData.Dash0203reference2 = v.Get("reference2")
		requestData.Dash0203Dato01 = v.Get("dato01")
		requestData.Dash0203Dato02 = v.Get("dato02")
		requestData.Dash0203Dato03 = v.Get("dato03")

   //END
   	 
   	 return requestData,errorGeneral
   }

   func obtainParmsProcessDash03Grafica01(r *http.Request, errorGeneral string) (modelito.RequestDash03Grafica01,string){
   	 var requestData modelito.RequestDash03Grafica01
	////////////////////////////////////////////////obtain parms in JSON
   //START    
    log.Print("cz  obtainParmsProcessDash03Grafica01")
 		 log.Print("CZ    handler Listening test obtainParmsProcessDash03Grafica01")
    		    
    	err := r.ParseForm()
		if err != nil {
	    	//prepare response with error 100
	    	log.Print("CZ    Prepare Response with 180. Missing parameter:"+errorGeneral)
	    	errorGeneral="ERROR:180 -"	+err.Error()
		}
		v := r.Form
		requestData.Dash0301reference = v.Get("reference")
		requestData.Dash0301reference2 = v.Get("reference2")
		requestData.Dash0301Dato01 = v.Get("dato01")
		requestData.Dash0301Dato02 = v.Get("dato02")
		requestData.Dash0301Dato03 = v.Get("dato03")

   //END
   	 
   	 return requestData,errorGeneral
   }

   func obtainParmsProcessDash03Grafica02(r *http.Request, errorGeneral string) (modelito.RequestDash03Grafica02,string){
   	 var requestData modelito.RequestDash03Grafica02
	////////////////////////////////////////////////obtain parms in JSON
   //START    
    log.Print("cz  obtainParmsProcessDash03Grafica02")
 		 log.Print("CZ    handler Listening test obtainParmsProcessDash03Grafica02")
    		    
    	err := r.ParseForm()
		if err != nil {
	    	//prepare response with error 100
	    	log.Print("CZ    Prepare Response with 180. Missing parameter:"+errorGeneral)
	    	errorGeneral="ERROR:180 -"	+err.Error()
		}
		v := r.Form
		requestData.Dash0302reference = v.Get("reference")
		requestData.Dash0302reference2 = v.Get("reference2")
		requestData.Dash0302Dato01 = v.Get("dato01")
		requestData.Dash0302Dato02 = v.Get("dato02")
		requestData.Dash0302Dato03 = v.Get("dato03")

   //END
   	 
   	 return requestData,errorGeneral
   }

   func obtainParmsProcessDash03Grafica03(r *http.Request, errorGeneral string) (modelito.RequestDash03Grafica03,string){
   	 var requestData modelito.RequestDash03Grafica03
	////////////////////////////////////////////////obtain parms in JSON
   //START    
    log.Print("cz  obtainParmsProcessDash03Grafica03")
 		 log.Print("CZ    handler Listening test obtainParmsProcessDash03Grafica03")
    		    
    	err := r.ParseForm()
		if err != nil {
	    	//prepare response with error 100
	    	log.Print("CZ    Prepare Response with 180. Missing parameter:"+errorGeneral)
	    	errorGeneral="ERROR:180 -"	+err.Error()
		}
		v := r.Form
		requestData.Dash0303reference = v.Get("reference")
		requestData.Dash0303reference2 = v.Get("reference2")
		requestData.Dash0303Dato01 = v.Get("dato01")
		requestData.Dash0303Dato02 = v.Get("dato02")
		requestData.Dash0303Dato03 = v.Get("dato03")

   //END
   	 
   	 return requestData,errorGeneral
   }
