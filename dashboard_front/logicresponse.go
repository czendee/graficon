package main

import (
	"log"
	modelito "banwire/dash/dashboard_front/model"
	 _ "github.com/lib/pq"   //use go get github.com/lib/pq
	 "time"
	 "encoding/json"
)



///////////////////////////////////////////////////////////////7/////version 2

///fetchDataDashDiagram

func getJsonResponseDatadashV1( datadashFound []modelito.Datadash)([]byte, error) {
	
	if(datadashFound !=nil){
        var cuantos=0;
     	for _, d := range datadashFound {
	     if(cuantos ==0){
		mainStruct = modelito.AutoGeneratedDash{StatusMessage: "Success" ,Status:"1",GroupNumber:d.Numsecuecialgrupoddatos}    //  model/request.go 
	     }
		log.Print("el registor trae:"+d.ID+" "+d.Nombrecolumna)
		w := modelito.ResponseDatadash{time.Now(),d.Numsecuecial,d.Nombrecolumna,d.Createdat,d.Valoramount,d.Valorcolumna}   //request.go
		mainStruct.Datadash = append(mainStruct.Datadash, w)     //request.go
		cuantos=cuantos+1;

 		}
		if(cuantos ==0){
		     mainStruct = modelito.AutoGeneratedDash{StatusMessage: "Error" ,Status:"0",GroupNumber:"0"}    //  model/request.go 
		}

 		return json.MarshalIndent(mainStruct, "", "  ")

	}


	return nil,nil
}

func getJsonResponseDatadashV2( datadashFound []modelito.Datadash)([]byte, error) {
	
	if(datadashFound !=nil){
		
		mainStruct := modelito.AutoGeneratedDash{StatusMessage: "Success" ,Status:"1"}    //  model/request.go
		var cuantos=0;
		var contadorTop =0;
		var primero=0;	
		var cualTopEs=-1;
		for i := 0; i < len(Config_topCustomer); i++ {

			var esteTopTraeDatos=0;
			for _, d := range datadashFound {
				if(primero ==0){
				     mainStruct = modelito.AutoGeneratedDash{StatusMessage: "Success" ,Status:"1",GroupNumber:d.Numsecuecialgrupoddatos}    //  model/request.go 
				     primero=1     
				}
                             
				if(Config_topCustomer[i] == d.Nombrecolumna){

						log.Print("el registor trae:"+d.ID+" "+d.Nombrecolumna)
						w := modelito.ResponseDatadash{time.Now(),d.Numsecuecial,d.Nombrecolumna,d.Createdat,d.Valoramount,d.Valorcolumna}   //request.go
						mainStruct.Datadash = append(mainStruct.Datadash, w)     //request.go
					    cuantos=cuantos+1;
					esteTopTraeDatos=1;

				}

			}//end for
			if(esteTopTraeDatos==0)
					//agrega uno vacio 
					w := modelito.ResponseDatadash{time.Now(),"0",Config_topCustomer[i],time.Now(),0,"0"}   //request.go
					mainStruct.Datadash = append(mainStruct.Datadash, w)     //request.go

			}//end else
	       }//end for

		if(cuantos ==0){
		     mainStruct = modelito.AutoGeneratedDash{StatusMessage: "Error" ,Status:"0",GroupNumber:"0"}    //  model/request.go 
		}

 		return json.MarshalIndent(mainStruct, "", "  ")

	}


	return nil,nil
}

func getJsonResponseDatadashNoRecentDataV1()([]byte, error) {

		mainStruct := modelito.AutoGeneratedDash{StatusMessage: "Error" ,Status:"0",GroupNumber:"0"}    //  model/request.go 

 		return json.MarshalIndent(mainStruct, "", "  ")

}



///////////////////////////////////////////////////////////////7/////version 2

///fetchCards
func getJsonResponseV2( cardsFound []modelito.Card)([]byte, error) {
	
	if(cardsFound !=nil){
		
		mainStruct := modelito.AutoGenerated{StatusMessage: "Success" ,Status:"1"}
     	for _, d := range cardsFound {
     		log.Print("el registor trae:"+d.Token+" "+d.Bin)
			w := modelito.CardData{time.Now(), d.Token,d.Last,d.Brand,d.Valid,d.Bin,d.Score,d.Type}   //request.go
			mainStruct.Cards = append(mainStruct.Cards, w)
 		}
 		return json.MarshalIndent(mainStruct, "", "  ")

/* 		
		for i := 0; i < 5; i++ {
			w := modelito.CardData{time.Now(), "G#$$%ytoywteouwytr","1234","VISA","2501","781234","1"}
			mainStruct.Cards = append(mainStruct.Cards, w)
		}

*/
	}


	return nil,nil
}

//processpayment
func getJsonResponsePaymentV2(datoPayment modelito.ExitoData)([]byte, error) {
	
									log.Print(" getJsonResponsePaymentV2 token:!\n"+datoPayment.Token)
									log.Print(" getJsonResponsePaymentV2 bin:!\n"+datoPayment.Bin)
									log.Print(" getJsonResponsePaymentV2 last:!\n"+datoPayment.LastDigits)
									
	mainStruct := modelito.ResponsePayment{StatusMessage: "Success",Status:"1"}        //response.go
    w := modelito.ExitoData{ datoPayment.Token,datoPayment.PaymentReference,datoPayment.Authcode,datoPayment.Idtransaction,datoPayment.Marca,datoPayment.Bin,datoPayment.LastDigits,datoPayment.Type}   //response.go
    mainStruct.SucessData=w
	return json.MarshalIndent(mainStruct, "", "  ")
}



/////////response for tokenize


func getJsonResponseTokenizeV2(cardTokenized modelito.Card)([]byte, error) {
	
	mainStruct := modelito.ResponseTokenize{StatusMessage: "Success",Status:"1"}
    w := modelito.CardData{time.Now(), cardTokenized.Token,cardTokenized.Last, cardTokenized.Brand,cardTokenized.Valid,cardTokenized.Bin,cardTokenized.Score,cardTokenized.Type}  
    mainStruct.Card = w


	return json.MarshalIndent(mainStruct, "", "  ")
}


////////////////////////ERROR response


func getJsonResponseError(errorMsg, errorNumber string )([]byte, error) {
	
	mainStruct := modelito.ResponseError{StatusMessage: errorMsg,Status:errorNumber}

	return json.MarshalIndent(mainStruct, "", "  ")
}
