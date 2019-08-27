package main

import (
	"fmt"
	"log"
//	"strings"
//	"banwire/dash/dashboard_back/db"
	modelito "banwire/dash/dashboard_back/model"
	 "database/sql"
	 _ "github.com/go-sql-driver/mysql"   //use go get -u github.com/go-sql-driver/mysql   this allows mysql use
//	miu "banwire/dash/dashboard_back/util"
)



///////////////// ///////////////////////////////////////version 1


////////////////////////////////////////////////////////7
//Dash 01 Graficas

func logicDBMysqlProcessDashBackGetDataFor01Grafica01( errorGeneral string,graphnbr string) ([]modelito.Datadash,string) {
	////////////////////////////////////////////////obtain parms in JSON
   //START    
var resultDatadash []modelito.Datadash
var errCards error

				//  START fetchFromDB
				    var errdb error
				    var db *sql.DB
				    // Create connection string
				//	connString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
				//		Config_DB_server,Config_DB_name, Config_DB_user, Config_DB_pass, Config_DB_port)
					connString03 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
						Config_DB_origin_user,Config_DB_origin_pass, Config_DB_origin_server, Config_DB_origin_port, Config_DB_origin_name)
				
				     if (connString03 !="si"){

                     }
//this use the values set up in the configuration.go
                  log.Print("Usando para conectar : " + Config_dbStringType_origin)
                  log.Print("Usando para conectar origin : " + Config_connString_origin)
                  log.Print("Usando para conectar origin real: " +connString03 )
					//db, errdb = sql.Open(Config_dbStringType_origin, Config_connString_origin)
                                   db, errdb = sql.Open(Config_dbStringType_origin, connString03)                    

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
                                    var cualConfig_comandosqlorigen_origin string

					         		switch graphnbr {
                            		case "31":
                                        cualConfig_comandosqlorigen_origin = Config_comandosqlorigen_origindash03
                                    case "11":
                                        cualConfig_comandosqlorigen_origin = Config_comandosqlorigen_origindash01                                    
                                    case "51":
                                        cualConfig_comandosqlorigen_origin = Config_comandosqlorigen_origindash05
                                    case "71":
                                         cualConfig_comandosqlorigen_origin = Config_comandosqlorigen_origindash07
                                    case "72":
                                          cualConfig_comandosqlorigen_origin = Config_comandosqlorigen_origindash07
                                    case "73":
                                          cualConfig_comandosqlorigen_origin = Config_comandosqlorigen_origindash07
                                     }		
							log.Print("Commnad to execute!\n"+cualConfig_comandosqlorigen_origin)
					         resultDatadash,errCards =modelito.BackMysqlDatadash0201(db,cualConfig_comandosqlorigen_origin) //defined in the config.json and set in the configuration.go

							if errCards != nil {
							  log.Print("Error: :"+ errCards.Error())
							  errorGeneral=errCards.Error()
							}	

					    }
				
				
					}
				    
				//  END fetchFromDB
   
   //END
   	  return  resultDatadash, errorGeneral
   }


func logicDBMysqlProcessDashBackInsertData01Grafica01( errorGeneral string,graphnbr string, valoresToInsert []modelito.Datadash) (string,string) {
	////////////////////////////////////////////////obtain parms in JSON
   //START    
var resultDatadash string
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
					         
					         resultDatadash,errCards =modelito.BackMysqlInsertDatadash0201(db,graphnbr,valoresToInsert)

							if errCards != nil {
							  log.Print("Error: :"+ errCards.Error())
							  errorGeneral=errCards.Error()
							}	

					    }
				
				
					}
				    
				//  END fetchFromDB
   
   //END
   	  return  resultDatadash, errorGeneral
   }


