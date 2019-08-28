package main

import (
	"fmt"
	"log"
//	"strings"
    "strconv"
//	"banwire/dash/dashboard_front/db"
	modelito "banwire/dash/dashboard_front/model"
	 "database/sql"
	 _ "github.com/lib/pq"   //use go get github.com/lib/pq     this allows postgres use
//	miu "banwire/dash/dashboard_front/util"
)

/*
    const (
        DB_USER     = "lerepagr"
        DB_PASSWORD = "Ag8q2utgSsVy2tyR7_M9cNYbzsqSvwma"
        DB_NAME     = "lerepagr"
        DB_SERVER     = "stampy.db.elephantsql.com" //"54.163.207.112"
        DB_PORT      = 5432
    )
 */


///////////////// ///////////////////////////////////////version 3.2



 
////////////////////////////////////////////////////////7
//Dash 01 Graficas


   func logicDBProcessDash01Grafica01(requestData modelito.RequestDash01Grafica01, errorGeneral string) ([]modelito.Card,string) {
	////////////////////////////////////////////////obtain parms in JSON
   //START    
var resultCards []modelito.Card
var errCards error

				//  START fetchFromDB
				    var errdb error
				    var db *sql.DB
				    // Create connection string
					connString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
						Config_DB_server,Config_DB_name, Config_DB_user, Config_DB_pass, Config_DB_port)
				
				     if (connString !="si"){

                     }
//"mysql", "root:password1@tcp(127.0.0.1:3306)/test"

					 // Create connection pool
//					db, errdb = sql.Open("postgres", connString)
//this use the values set up in the configuration.go
                  log.Print("Usando para conectar : " + Config_dbStringType)
					db, errdb = sql.Open(Config_dbStringType, Config_connString)
                    

					if errdb != nil {
						log.Print("Error creating connection pool: " + errdb.Error())
						errorGeneral=errdb.Error()
					}
					// Close the database connection pool after program executes
					 defer db.Close()
					if errdb == nil {
						log.Print("Connected!\n")
				
					
						errPing := db.Ping()
						if errPing != nil {
						  log.Print("Error: Could not establish a connection with the database:"+ errPing.Error())
							  errorGeneral=errPing.Error()
						}else{
					         log.Print("Ping ok!\n")
//					         var misCards modelito.Card
					         
					         resultCards,errCards =modelito.GetCardsByCustomer(db,requestData.Dash0101reference)
					         					         log.Print("regresa func  getCardsByCustomer ok!\n")
							if errCards != nil {
							  log.Print("Error: :"+ errCards.Error())
							  errorGeneral=errCards.Error()
							}
							var cuantos int
							cuantos = 0
				         	for _, d := range resultCards {
				         		log.Print("el registor trae:"+d.Token+" "+d.Bin)
							    cuantos =1
			         		}
							if cuantos == 0 {
							  log.Print("DB: records not found")
							  errorGeneral="Not cards found for the customer reference received"
							}		

					    }
				
				
					}
				    
				//  END fetchFromDB
   
   //END
   	  return  resultCards, errorGeneral
   }



   func logicDBProcessDash01Grafica02(requestData modelito.RequestDash01Grafica02, errorGeneral string) ([]modelito.Card,string) {
	////////////////////////////////////////////////obtain parms in JSON
   //START    
var resultCards []modelito.Card
var errCards error

				//  START fetchFromDB
				    var errdb error
				    var db *sql.DB
				    // Create connection string
					connString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
						Config_DB_server,Config_DB_name, Config_DB_user, Config_DB_pass, Config_DB_port)
				
				     if (connString !="si"){

                     }
//"mysql", "root:password1@tcp(127.0.0.1:3306)/test"

					 // Create connection pool
//					db, errdb = sql.Open("postgres", connString)
//this use the values set up in the configuration.go
                  log.Print("Usando para conectar : " + Config_dbStringType)
					db, errdb = sql.Open(Config_dbStringType, Config_connString)
                    

					if errdb != nil {
						log.Print("Error creating connection pool: " + errdb.Error())
						errorGeneral=errdb.Error()
					}
					// Close the database connection pool after program executes
					 defer db.Close()
					if errdb == nil {
						log.Print("Connected!\n")
				
					
						errPing := db.Ping()
						if errPing != nil {
						  log.Print("Error: Could not establish a connection with the database:"+ errPing.Error())
							  errorGeneral=errPing.Error()
						}else{
					         log.Print("Ping ok!\n")
//					         var misCards modelito.Card
					         
					         resultCards,errCards =modelito.GetCardsByCustomer(db,requestData.Dash0102reference)
					         					         log.Print("regresa func  getCardsByCustomer ok!\n")
							if errCards != nil {
							  log.Print("Error: :"+ errCards.Error())
							  errorGeneral=errCards.Error()
							}
							var cuantos int
							cuantos = 0
				         	for _, d := range resultCards {
				         		log.Print("el registor trae:"+d.Token+" "+d.Bin)
							    cuantos =1
			         		}
							if cuantos == 0 {
							  log.Print("DB: records not found")
							  errorGeneral="Not cards found for the customer reference received"
							}		

					    }
				
				
					}
				    
				//  END fetchFromDB
   
   //END
   	  return  resultCards, errorGeneral
   }


   func logicDBCheckMoreRecent0201(requestData modelito.RequestDash02Grafica01, errorGeneral string) (string,string) {
	////////////////////////////////////////////////obtain parms in JSON
   //START    
        var resultDatadashGroupNumber string
        var errCards error

        //  START fetchFromDB
            var errdb error
            var db *sql.DB
            // Create connection string
            connString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
                Config_DB_server,Config_DB_name, Config_DB_user, Config_DB_pass, Config_DB_port)
        
                if (connString !="si"){

                }
//"mysql", "root:password1@tcp(127.0.0.1:3306)/test"

                // Create connection pool
//					db, errdb = sql.Open("postgres", connString)
//this use the values set up in the configuration.go
            log.Print("Usando para conectar : " + Config_dbStringType)
            db, errdb = sql.Open(Config_dbStringType, Config_connString)
            

            if errdb != nil {
                log.Print("Error creating connection pool: " + errdb.Error())
                errorGeneral=errdb.Error()
            }
            // Close the database connection pool after program executes
                defer db.Close()
            if errdb == nil {
                log.Print("Connected!\n")
        
            
                errPing := db.Ping()
                if errPing != nil {
                    log.Print("Error: Could not establish a connection with the database:"+ errPing.Error())
                        errorGeneral=errPing.Error()
                }else{
                        log.Print("Ping ok in MySql :) !\n")

                        
                        resultDatadashGroupNumber,errCards =modelito.GetMaxGroupNumberDatadash0201(db,
                        requestData.Dash0201reference,
                        requestData.Dash0201reference2,
                        requestData.Dash0201Dato01,
                        requestData.Dash0201Dato02,
                        requestData.Dash0201Dato03)

                        log.Print("regresa func  logicDBMysqlCheckMoreRecent0201 ok!\n")
                    if errCards != nil {
                        log.Print("Error: :"+ errCards.Error())
                        errorGeneral=errCards.Error()
                    }

                    if resultDatadashGroupNumber =="0"{
                        //no result found , greater than the previous group
                    }else{
                       //results found , more recent
                    }


                }//end else ping
        
        
            }//end else dbconn
            
        //  END fetchFromDB
   
   //END
   	  return  resultDatadashGroupNumber, errorGeneral
   }


   func logicDBProcessDash02Grafica01(requestData modelito.RequestDash02Grafica01, errorGeneral string) ([]modelito.Datadash,string) {
	////////////////////////////////////////////////obtain parms in JSON
   //START    
var resultDatadash []modelito.Datadash
var errCards error

				//  START fetchFromDB
				    var errdb error
				    var db *sql.DB
				    // Create connection string
					connString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
						Config_DB_server,Config_DB_name, Config_DB_user, Config_DB_pass, Config_DB_port)
				
				     if (connString !="si"){

                     }
//"mysql", "root:password1@tcp(127.0.0.1:3306)/test"

					 // Create connection pool
//					db, errdb = sql.Open("postgres", connString)
//this use the values set up in the configuration.go
                  log.Print("Usando para conectar : " + Config_dbStringType)
					db, errdb = sql.Open(Config_dbStringType, Config_connString)
                    

					if errdb != nil {
						log.Print("Error creating connection pool: " + errdb.Error())
						errorGeneral=errdb.Error()
					}
					// Close the database connection pool after program executes
					 defer db.Close()
					if errdb == nil {
						log.Print("Connected!\n")
				
					
						errPing := db.Ping()
						if errPing != nil {
						  log.Print("Error: Could not establish a connection with the database:"+ errPing.Error())
							  errorGeneral=errPing.Error()
						}else{
					         log.Print("Ping ok!\n")
//					         var misCards modelito.Card
					         
					         resultDatadash,errCards =modelito.GetDatadash0201(db,
                             requestData.Dash0201reference,
                             requestData.Dash0201reference2,
                             requestData.Dash0201Dato01,
                             requestData.Dash0201Dato02,
                             requestData.Dash0201Dato03)

					         log.Print("regresa func  getCardsByCustomer ok!\n")
							if errCards != nil {
							  log.Print("Error: :"+ errCards.Error())
							  errorGeneral=errCards.Error()
							}
							var cuantos int
							cuantos = 0
				         	for _, d := range resultDatadash {
				         		log.Print("el registor trae:"+d.ID+" "+d.Valoramount)
							    cuantos =1
			         		}
							if cuantos == 0 {
							  log.Print("DB: records not found")
							  errorGeneral="Not cards found for the customer reference received"
							}		

					    }
				
				
					}
				    
				//  END fetchFromDB
   
   //END
   	  return  resultDatadash, errorGeneral
   }

   func logicDBCheckMoreRecent0202(requestData modelito.RequestDash02Grafica02, errorGeneral string) (string,string) {
	////////////////////////////////////////////////obtain parms in JSON
   //START    
        var resultDatadashGroupNumber string
        var errCards error

        //  START fetchFromDB
            var errdb error
            var db *sql.DB
            // Create connection string
            connString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
                Config_DB_server,Config_DB_name, Config_DB_user, Config_DB_pass, Config_DB_port)
        
                if (connString !="si"){

                }
//"mysql", "root:password1@tcp(127.0.0.1:3306)/test"

                // Create connection pool
//					db, errdb = sql.Open("postgres", connString)
//this use the values set up in the configuration.go
            log.Print("Usando para conectar : " + Config_dbStringType)
            db, errdb = sql.Open(Config_dbStringType, Config_connString)
            

            if errdb != nil {
                log.Print("Error creating connection pool: " + errdb.Error())
                errorGeneral=errdb.Error()
            }
            // Close the database connection pool after program executes
                defer db.Close()
            if errdb == nil {
                log.Print("Connected!\n")
        
            
                errPing := db.Ping()
                if errPing != nil {
                    log.Print("Error: Could not establish a connection with the database:"+ errPing.Error())
                        errorGeneral=errPing.Error()
                }else{
                        log.Print("Ping ok in MySql :) !\n")

                        
                        resultDatadashGroupNumber,errCards =modelito.GetMaxGroupNumberDatadash0202(db,
                        requestData.Dash0202reference,
                        requestData.Dash0202reference2,
                        requestData.Dash0202Dato01,
                        requestData.Dash0202Dato02,
                        requestData.Dash0202Dato03)

                        log.Print("regresa func  logicDBMysqlCheckMoreRecent0202 ok!\n")
                    if errCards != nil {
                        log.Print("Error: :"+ errCards.Error())
                        errorGeneral=errCards.Error()
                    }

                    if resultDatadashGroupNumber =="0"{
                        //no result found , greater than the previous group
                    }else{
                       //results found , more recent
                    }


                }//end else ping
        
        
            }//end else dbconn
            
        //  END fetchFromDB
   
   //END
   	  return  resultDatadashGroupNumber, errorGeneral
   }


   func logicDBProcessDash02Grafica02(requestData modelito.RequestDash02Grafica02, errorGeneral string) ([]modelito.Datadash,string) {
	////////////////////////////////////////////////obtain parms in JSON
   //START    
var resultDatadash []modelito.Datadash
var errCards error
resultLastHrDatadash:= []modelito.Datadash{}
				//  START fetchFromDB
				    var errdb error
				    var db *sql.DB
				    // Create connection string
					connString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
						Config_DB_server,Config_DB_name, Config_DB_user, Config_DB_pass, Config_DB_port)
				
				     if (connString !="si"){

                     }
//"mysql", "root:password1@tcp(127.0.0.1:3306)/test"

					 // Create connection pool
//					db, errdb = sql.Open("postgres", connString)
//this use the values set up in the configuration.go
                  log.Print("Usando para conectar : " + Config_dbStringType)
					db, errdb = sql.Open(Config_dbStringType, Config_connString)
                    

					if errdb != nil {
						log.Print("Error creating connection pool: " + errdb.Error())
						errorGeneral=errdb.Error()
					}
					// Close the database connection pool after program executes
					 defer db.Close()
					if errdb == nil {
						log.Print("Connected!\n")
				
					
						errPing := db.Ping()
						if errPing != nil {
						  log.Print("Error: Could not establish a connection with the database:"+ errPing.Error())
							  errorGeneral=errPing.Error()
						}else{
					         log.Print("Ping ok!\n")
//					         var misCards modelito.Card
							var cualConfig_comandosqlfront string
				 switch  {
                            	   case requestData.Dash0202reference2== "11" && requestData.Dash0202Dato01 =="99" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash01_99
                            	   case requestData.Dash0202reference2== "11" && requestData.Dash0202Dato01 =="24" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash01_24
                            	   case requestData.Dash0202reference2== "11" && requestData.Dash0202Dato01 =="48" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash01_48
					 
                            	   case requestData.Dash0202reference2== "21" && requestData.Dash0202Dato01 =="99" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash02_99
                            	   case requestData.Dash0202reference2== "21" && requestData.Dash0202Dato01 =="7" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash02_7
                            	   case requestData.Dash0202reference2== "21" && requestData.Dash0202Dato01 =="30" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash02_30
					 
                            	   case requestData.Dash0202reference2== "31" && requestData.Dash0202Dato01 =="99" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash03_99
                            	   case requestData.Dash0202reference2== "31" && requestData.Dash0202Dato01 =="24" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash03_24
                            	   case requestData.Dash0202reference2== "31" && requestData.Dash0202Dato01 =="48" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash03_48
					 
                            	   case requestData.Dash0202reference2== "41" && requestData.Dash0202Dato01 =="99" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash04_99
                            	   case requestData.Dash0202reference2== "41" && requestData.Dash0202Dato01 =="7" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash04_7
                            	   case requestData.Dash0202reference2== "41" && requestData.Dash0202Dato01 =="30" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash04_30

					 
                            	   case requestData.Dash0202reference2== "51" && requestData.Dash0202Dato01 =="99" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash05_99
                            	   case requestData.Dash0202reference2== "51" && requestData.Dash0202Dato01 =="24" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash05_24
                            	   case requestData.Dash0202reference2== "51" && requestData.Dash0202Dato01 =="48" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash05_48
					 
                            	   case requestData.Dash0202reference2== "61" && requestData.Dash0202Dato01 =="99" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash06_99
                            	   case requestData.Dash0202reference2== "61" && requestData.Dash0202Dato01 =="7" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash06_7
                            	   case requestData.Dash0202reference2== "61" && requestData.Dash0202Dato01 =="30" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash06_30

					 
                            	   case requestData.Dash0202reference2== "71" && requestData.Dash0202Dato01 =="99" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash07_99
                            	   case requestData.Dash0202reference2== "71" && requestData.Dash0202Dato01 =="24" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash07_24
                            	   case requestData.Dash0202reference2== "71" && requestData.Dash0202Dato01 =="48" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash07_48
					 
                            	   case requestData.Dash0202reference2== "81" && requestData.Dash0202Dato01 =="99" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash08_99
                            	   case requestData.Dash0202reference2== "81" && requestData.Dash0202Dato01 =="7" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash08_7
                            	   case requestData.Dash0202reference2== "81" && requestData.Dash0202Dato01 =="30" :
                                        cualConfig_comandosqlfront = Config_comandosql_frontdash08_30
					 
                                     }
					         resultDatadash,errCards =modelito.GetDatadash0202(db,
                             requestData.Dash0202reference,
                             requestData.Dash0202reference2,
                             requestData.Dash0202Dato01,
                             requestData.Dash0202Dato02,
                             requestData.Dash0202Dato03,
				cualConfig_comandosqlfront )

					         log.Print("regresa func  logicDBProcessDash02Grafica02 ok!\n")
							if errCards != nil {
							  log.Print("Error: :"+ errCards.Error())
							  errorGeneral=errCards.Error()
							}
/*							var cuantos int
							cuantos = 0
				         	for _, d := range resultDatadash {
				         		log.Print("el registor trae:"+d.ID+" "+d.Valoramount)
							    cuantos =1
			         		}
*/ 
							var cuantos int
							cuantos = 0
                            

                            otherLastHrDatadash:= modelito.Datadash{}
                            acumuladoLastHr :=0
				         	for _, d := range resultDatadash {
				         		log.Print("el registor trae:"+d.ID+" "+d.Valoramount)
							    cuantos =cuantos+1
                                
                                //after the 7th, sum in one
                                if(cuantos >7){
                                    valors,err01 := strconv.Atoi(d.Valoramount)
                                    if(err01!=nil){

                                    }
                                    
                                    acumuladoLastHr = acumuladoLastHr +valors
                                }else{
						
                                    //take the first 9 , as they come
                                    resultLastHrDatadash =append (resultLastHrDatadash,d)

                                }
			         		}
                            if cuantos >7 {
                                
                                //set values for the Other element
                                otherLastHrDatadash.ID ="1"

                                valor2s := strconv.Itoa(acumuladoLastHr)
                                otherLastHrDatadash.Valoramount=valor2s                                
                                otherLastHrDatadash.Valorcolumna=valor2s
                                otherLastHrDatadash.Nombrecolumna="Other"
                                otherLastHrDatadash.Numsecuecial="10"

                                resultLastHrDatadash =append (resultLastHrDatadash,otherLastHrDatadash)
			    } else { //not enough records for the 8 lines in the screen
				     i := cuantos //cuantos is less than 7, so set the rest of the 
				    for i <= 8 {
					//set values for the Other element
					otherLastHrDatadash.ID ="1"

					otherLastHrDatadash.Valoramount="0"
					otherLastHrDatadash.Valorcolumna="0"
					otherLastHrDatadash.Nombrecolumna=" "
					otherLastHrDatadash.Numsecuecial="10"

					resultLastHrDatadash =append (resultLastHrDatadash,otherLastHrDatadash)
					i = i + 1
				    }

			    }

							if cuantos == 0 {
							  log.Print("DB: records not found")
							  errorGeneral="Not records found for the logicDBProcessDash02Grafica02 reference received"
							}		

					    }
				
				
					}
				    
				//  END fetchFromDB
   
   //END
   	  return  resultLastHrDatadash, errorGeneral

   }

   func logicDBProcessDash02Grafica02hours24(requestData modelito.RequestDash02Grafica02, errorGeneral string) ([]modelito.Datadash,string) {
	////////////////////////////////////////////////obtain parms in JSON
   //START    
var resultDatadash []modelito.Datadash
var errCards error
result24hrsDatadash:= []modelito.Datadash{}
				//  START fetchFromDB
				    var errdb error
				    var db *sql.DB
				    // Create connection string
					connString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
						Config_DB_server,Config_DB_name, Config_DB_user, Config_DB_pass, Config_DB_port)
				
				     if (connString !="si"){

                     }
//"mysql", "root:password1@tcp(127.0.0.1:3306)/test"

					 // Create connection pool
//					db, errdb = sql.Open("postgres", connString)
//this use the values set up in the configuration.go
                  log.Print("Usando para conectar : " + Config_dbStringType)
					db, errdb = sql.Open(Config_dbStringType, Config_connString)
                    

					if errdb != nil {
						log.Print("Error creating connection pool: " + errdb.Error())
						errorGeneral=errdb.Error()
					}
					// Close the database connection pool after program executes
					 defer db.Close()
					if errdb == nil {
						log.Print("Connected!\n")
				
					
						errPing := db.Ping()
						if errPing != nil {
						  log.Print("Error: Could not establish a connection with the database:"+ errPing.Error())
							  errorGeneral=errPing.Error()
						}else{
					         log.Print("Ping ok!\n")
//					         var misCards modelito.Card
					         
					         resultDatadash,errCards =modelito.GetDatadash0202hours24(db,
                             requestData.Dash0202reference,
                             requestData.Dash0202reference2,
                             requestData.Dash0202Dato01,
                             requestData.Dash0202Dato02,
                             requestData.Dash0202Dato03)

					         log.Print("regresa func  logicDBProcessDash02Grafica02 ok!\n")
							if errCards != nil {
							  log.Print("Error: :"+ errCards.Error())
							  errorGeneral=errCards.Error()
							}
							var cuantos int
							cuantos = 0
                            

                            other24hrsDatadash:= modelito.Datadash{}
                            acumulado24hrs :=0
				         	for _, d := range resultDatadash {
				         		log.Print("el registor trae:"+d.ID+" "+d.Valoramount)
							    cuantos =cuantos+1
                                
                                //after the 7th, sum in one
                                if(cuantos >7){
                                    valors,err01 := strconv.Atoi(d.Valoramount)
                                    if(err01!=nil){

                                    }
                                    
                                    acumulado24hrs = acumulado24hrs +valors
                                }else{
                                    //take the first 9 , as they come
                                    result24hrsDatadash =append (result24hrsDatadash,d)

                                }
			         		}
                            if cuantos >7 {
                                
                                //set values for the Other element
                                other24hrsDatadash.ID ="1"

                                valor2s := strconv.Itoa(acumulado24hrs)
                                other24hrsDatadash.Valoramount=valor2s                                
                                other24hrsDatadash.Valorcolumna=valor2s
                                other24hrsDatadash.Nombrecolumna="Other"
                                other24hrsDatadash.Numsecuecial="10"

                                result24hrsDatadash =append (result24hrsDatadash,other24hrsDatadash)
                            } 
							if cuantos == 0 {
							  log.Print("DB: records not found")
							  errorGeneral="Not records found for the logicDBProcessDash02Grafica02 reference received"
							}		

					    }
				
				
					}
				    
				//  END fetchFromDB
   
   //END
   	  return  result24hrsDatadash , errorGeneral
	
   }

   func logicDBProcessDash02Grafica02hours48(requestData modelito.RequestDash02Grafica02, errorGeneral string) ([]modelito.Datadash,string) {
	////////////////////////////////////////////////obtain parms in JSON
   //START    
var resultDatadash []modelito.Datadash
var errCards error

result48hrsDatadash:= []modelito.Datadash{}
				//  START fetchFromDB
				    var errdb error
				    var db *sql.DB
				    // Create connection string
					connString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
						Config_DB_server,Config_DB_name, Config_DB_user, Config_DB_pass, Config_DB_port)
				
				     if (connString !="si"){

                     }
//"mysql", "root:password1@tcp(127.0.0.1:3306)/test"

					 // Create connection pool
//					db, errdb = sql.Open("postgres", connString)
//this use the values set up in the configuration.go
                  log.Print("Usando para conectar : " + Config_dbStringType)
					db, errdb = sql.Open(Config_dbStringType, Config_connString)
                    

					if errdb != nil {
						log.Print("Error creating connection pool: " + errdb.Error())
						errorGeneral=errdb.Error()
					}
					// Close the database connection pool after program executes
					 defer db.Close()
					if errdb == nil {
						log.Print("Connected!\n")
				
					
						errPing := db.Ping()
						if errPing != nil {
						  log.Print("Error: Could not establish a connection with the database:"+ errPing.Error())
							  errorGeneral=errPing.Error()
						}else{
					         log.Print("Ping ok!\n")
//					         var misCards modelito.Card
					         
					         resultDatadash,errCards =modelito.GetDatadash0202hours48(db,
                             requestData.Dash0202reference,
                             requestData.Dash0202reference2,
                             requestData.Dash0202Dato01,
                             requestData.Dash0202Dato02,
                             requestData.Dash0202Dato03)

					         log.Print("regresa func  logicDBProcessDash02Grafica02 ok!\n")
							if errCards != nil {
							  log.Print("Error: :"+ errCards.Error())
							  errorGeneral=errCards.Error()
							}
							var cuantos int
							cuantos = 0


                            other48hrsDatadash:= modelito.Datadash{}
                            acumulado48hrs :=0
				         	for _, d := range resultDatadash {
				         		log.Print("el registor trae:"+d.ID+" "+d.Valoramount)
							    cuantos =cuantos+1
                                
                                //after the 7th, sum in one
                                if(cuantos >7){
                                    valors04,err01 := strconv.Atoi(d.Valoramount)
                                    if(err01!=nil){

                                    }

                                    acumulado48hrs = acumulado48hrs +valors04
                                }else{
                                    //take the first 9 , as they come
                                    result48hrsDatadash =append (result48hrsDatadash,d)

                                }
			         		}
                            if cuantos >7 {
				         		log.Print("add other")                                
                                //set values for the Other element
                                other48hrsDatadash.ID ="1"
                                other48hrsDatadash.Valoramount=strconv.Itoa(acumulado48hrs)
                                other48hrsDatadash.Valorcolumna=strconv.Itoa(acumulado48hrs)
                                other48hrsDatadash.Nombrecolumna="Other"
                                other48hrsDatadash.Numsecuecial="10"

                                result48hrsDatadash =append (result48hrsDatadash,other48hrsDatadash)
                            } 
							if cuantos == 0 {
							  log.Print("DB: records not found")
							  errorGeneral="Not records found for the logicDBProcessDash02Grafica02 reference received"
							}		

					    }
				
				
					}
				    
				//  END fetchFromDB
   
   //END
   	  return  result48hrsDatadash, errorGeneral

   }


   func logicDBProcessDash02Grafica03(requestData modelito.RequestDash02Grafica03, errorGeneral string) ([]modelito.Card,string) {
	////////////////////////////////////////////////obtain parms in JSON
   //START    
var resultCards []modelito.Card
var errCards error

				//  START fetchFromDB
				    var errdb error
				    var db *sql.DB
				    // Create connection string
					connString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
						Config_DB_server,Config_DB_name, Config_DB_user, Config_DB_pass, Config_DB_port)
				
				     if (connString !="si"){

                     }
//"mysql", "root:password1@tcp(127.0.0.1:3306)/test"

					 // Create connection pool
//					db, errdb = sql.Open("postgres", connString)
//this use the values set up in the configuration.go
                  log.Print("Usando para conectar : " + Config_dbStringType)
					db, errdb = sql.Open(Config_dbStringType, Config_connString)
                    

					if errdb != nil {
						log.Print("Error creating connection pool: " + errdb.Error())
						errorGeneral=errdb.Error()
					}
					// Close the database connection pool after program executes
					 defer db.Close()
					if errdb == nil {
						log.Print("Connected!\n")
				
					
						errPing := db.Ping()
						if errPing != nil {
						  log.Print("Error: Could not establish a connection with the database:"+ errPing.Error())
							  errorGeneral=errPing.Error()
						}else{
					         log.Print("Ping ok!\n")
//					         var misCards modelito.Card
					         
					         resultCards,errCards =modelito.GetCardsByCustomer(db,requestData.Dash0203reference)
					         					         log.Print("regresa func  getCardsByCustomer ok!\n")
							if errCards != nil {
							  log.Print("Error: :"+ errCards.Error())
							  errorGeneral=errCards.Error()
							}
							var cuantos int
							cuantos = 0
				         	for _, d := range resultCards {
				         		log.Print("el registor trae:"+d.Token+" "+d.Bin)
							    cuantos =1
			         		}
							if cuantos == 0 {
							  log.Print("DB: records not found")
							  errorGeneral="Not cards found for the customer reference received"
							}		

					    }
				
				
					}
				    
				//  END fetchFromDB
   
   //END
   	  return  resultCards, errorGeneral
   }

   func logicDBProcessDash03Grafica01(requestData modelito.RequestDash03Grafica01, errorGeneral string) ([]modelito.Card,string) {
	////////////////////////////////////////////////obtain parms in JSON
   //START    
var resultCards []modelito.Card
var errCards error

				//  START fetchFromDB
				    var errdb error
				    var db *sql.DB
				    // Create connection string
					connString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
						Config_DB_server,Config_DB_name, Config_DB_user, Config_DB_pass, Config_DB_port)
				
				     if (connString !="si"){
  
                     }
//"mysql", "root:password1@tcp(127.0.0.1:3306)/test"

					 // Create connection pool
//					db, errdb = sql.Open("postgres", connString)
//this use the values set up in the configuration.go
                  log.Print("Usando para conectar : " + Config_dbStringType)
					db, errdb = sql.Open(Config_dbStringType, Config_connString)
                    

					if errdb != nil {
						log.Print("Error creating connection pool: " + errdb.Error())
						errorGeneral=errdb.Error()
					}
					// Close the database connection pool after program executes
					 defer db.Close()
					if errdb == nil {
						log.Print("Connected!\n")
				
					
						errPing := db.Ping()
						if errPing != nil {
						  log.Print("Error: Could not establish a connection with the database:"+ errPing.Error())
							  errorGeneral=errPing.Error()
						}else{
					         log.Print("Ping ok!\n")
//					         var misCards modelito.Card
					         
					         resultCards,errCards =modelito.GetCardsByCustomer(db,requestData.Dash0301reference)
					         					         log.Print("regresa func  getCardsByCustomer ok!\n")
							if errCards != nil {
							  log.Print("Error: :"+ errCards.Error())
							  errorGeneral=errCards.Error()
							}
							var cuantos int
							cuantos = 0
				         	for _, d := range resultCards {
				         		log.Print("el registor trae:"+d.Token+" "+d.Bin)
							    cuantos =1
			         		}
							if cuantos == 0 {
							  log.Print("DB: records not found")
							  errorGeneral="Not cards found for the customer reference received"
							}		

					    }
				
				
					}
				    
				//  END fetchFromDB
   
   //END
   	  return  resultCards, errorGeneral
   }

   func logicDBProcessDash03Grafica02(requestData modelito.RequestDash03Grafica02, errorGeneral string) ([]modelito.Card,string) {
	////////////////////////////////////////////////obtain parms in JSON
   //START    
var resultCards []modelito.Card
var errCards error

				//  START fetchFromDB
				    var errdb error
				    var db *sql.DB
				    // Create connection string
					connString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
						Config_DB_server,Config_DB_name, Config_DB_user, Config_DB_pass, Config_DB_port)
				
				     if (connString !="si"){

                     }
//"mysql", "root:password1@tcp(127.0.0.1:3306)/test"

					 // Create connection pool
//					db, errdb = sql.Open("postgres", connString)
//this use the values set up in the configuration.go
                  log.Print("Usando para conectar : " + Config_dbStringType)
					db, errdb = sql.Open(Config_dbStringType, Config_connString)
                    

					if errdb != nil {
						log.Print("Error creating connection pool: " + errdb.Error())
						errorGeneral=errdb.Error()
					}
					// Close the database connection pool after program executes
					 defer db.Close()
					if errdb == nil {
						log.Print("Connected!\n")
				
					
						errPing := db.Ping()
						if errPing != nil {
						  log.Print("Error: Could not establish a connection with the database:"+ errPing.Error())
							  errorGeneral=errPing.Error()
						}else{
					         log.Print("Ping ok!\n")
//					         var misCards modelito.Card
					         
					         resultCards,errCards =modelito.GetCardsByCustomer(db,requestData.Dash0302reference)
					         					         log.Print("regresa func  getCardsByCustomer ok!\n")
							if errCards != nil {
							  log.Print("Error: :"+ errCards.Error())
							  errorGeneral=errCards.Error()
							}
							var cuantos int
							cuantos = 0
				         	for _, d := range resultCards {
				         		log.Print("el registor trae:"+d.Token+" "+d.Bin)
							    cuantos =1
			         		}
							if cuantos == 0 {
							  log.Print("DB: records not found")
							  errorGeneral="Not cards found for the customer reference received"
							}		

					    }
				
				
					}
				    
				//  END fetchFromDB
   
   //END
   	  return  resultCards, errorGeneral
   }

   func logicDBProcessDash03Grafica03(requestData modelito.RequestDash03Grafica03, errorGeneral string) ([]modelito.Card,string) {
	////////////////////////////////////////////////obtain parms in JSON
   //START    
var resultCards []modelito.Card
var errCards error

				//  START fetchFromDB
				    var errdb error
				    var db *sql.DB
				    // Create connection string
					connString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
						Config_DB_server,Config_DB_name, Config_DB_user, Config_DB_pass, Config_DB_port)
				
				     if (connString !="si"){

                     }
//"mysql", "root:password1@tcp(127.0.0.1:3306)/test"

					 // Create connection pool
//					db, errdb = sql.Open("postgres", connString)
//this use the values set up in the configuration.go
                  log.Print("Usando para conectar : " + Config_dbStringType)
					db, errdb = sql.Open(Config_dbStringType, Config_connString)
                    

					if errdb != nil {
						log.Print("Error creating connection pool: " + errdb.Error())
						errorGeneral=errdb.Error()
					}
					// Close the database connection pool after program executes
					 defer db.Close()
					if errdb == nil {
						log.Print("Connected!\n")
				
					
						errPing := db.Ping()
						if errPing != nil {
						  log.Print("Error: Could not establish a connection with the database:"+ errPing.Error())
							  errorGeneral=errPing.Error()
						}else{
					         log.Print("Ping ok!\n")
//					         var misCards modelito.Card
					         
					         resultCards,errCards =modelito.GetCardsByCustomer(db,requestData.Dash0303reference)
					         					         log.Print("regresa func  getCardsByCustomer ok!\n")
							if errCards != nil {
							  log.Print("Error: :"+ errCards.Error())
							  errorGeneral=errCards.Error()
							}
							var cuantos int
							cuantos = 0
				         	for _, d := range resultCards {
				         		log.Print("el registor trae:"+d.Token+" "+d.Bin)
							    cuantos =1
			         		}
							if cuantos == 0 {
							  log.Print("DB: records not found")
							  errorGeneral="Not cards found for the customer reference received"
							}		

					    }
				
					}
				    
				//  END fetchFromDB
   //END
   	  return  resultCards, errorGeneral
   }
