package main

import (
	"fmt"
	"log"
//	"strings"
//	"banwire/dash/dashboard_back/db"
	modelito "banwire/dash/dashboard_back/model"
	 "database/sql"
	 _ "github.com/lib/pq"   //use go get github.com/lib/pq     this allows postgres use
//	miu "banwire/dash/dashboard_back/util"
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

func logicDBProcessDashBack01Grafica01( errorGeneral string) (string,string) {
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
					         
					         resultDatadash,errCards =modelito.BackDatadash0201(db)

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


