package main

import (
	"fmt"
	"log"
//	"strings"
//	"banwire/dash/dashboard_banwire/db"
	modelito "banwire/dash/dashboard_banwire/model"
	 "database/sql"
	 _ "github.com/go-sql-driver/mysql"   //use go get -u github.com/go-sql-driver/mysql   this allows mysql use
//	miu "banwire/dash/dashboard_banwire/util"
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


///////////////// ///////////////////////////////////////version 1


////////////////////////////////////////////////////////7
//Dash 01 Graficas


   func logicDBMysqlProcessDash01Grafica01(requestData modelito.RequestDash01Grafica01, errorGeneral string) ([]modelito.Card,string) {
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



   func logicDBMysqlProcessDash01Grafica02(requestData modelito.RequestDash01Grafica02, errorGeneral string) ([]modelito.Card,string) {
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


   func logicDBMysqlProcessDash02Grafica01(requestData modelito.RequestDash02Grafica01, errorGeneral string) ([]modelito.Datadash,string) {
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
					         log.Print("Ping ok in MySql :) !\n")
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

   func logicDBMysqlProcessDash02Grafica02(requestData modelito.RequestDash02Grafica02, errorGeneral string) ([]modelito.Card,string) {
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
					         
					         resultCards,errCards =modelito.GetCardsByCustomer(db,requestData.Dash0202reference)
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

   func logicDBMysqlProcessDash02Grafica03(requestData modelito.RequestDash02Grafica03, errorGeneral string) ([]modelito.Card,string) {
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

   func logicDBMysqlProcessDash03Grafica01(requestData modelito.RequestDash03Grafica01, errorGeneral string) ([]modelito.Card,string) {
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

   func logicDBMysqlProcessDash03Grafica02(requestData modelito.RequestDash03Grafica02, errorGeneral string) ([]modelito.Card,string) {
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

   func logicDBMysqlProcessDash03Grafica03(requestData modelito.RequestDash03Grafica03, errorGeneral string) ([]modelito.Card,string) {
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
