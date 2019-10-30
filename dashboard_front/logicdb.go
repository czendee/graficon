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

                        //set the graphnbr to 31 for 21 and for 61
			//set the graphnbr to 71 for 41 and for 81
                           switch  {
                            	   case requestData.Dash0202reference2== "21"  :
                                        requestData.Dash0202reference2 ="31"
                            	   case requestData.Dash0202reference2== "61"  :
                                        requestData.Dash0202reference2="71"
                            	   case requestData.Dash0202reference2== "41"  :
                                        requestData.Dash0202reference2="31"			   
				   case requestData.Dash0202reference2== "81" :
                                        requestData.Dash0202reference2="71"
				   case requestData.Dash0202reference2== "221" :
                                        requestData.Dash0202reference2="31"
				   case requestData.Dash0202reference2== "222" :
                                        requestData.Dash0202reference2="11"	 //transacciones
				   case requestData.Dash0202reference2== "225" :
                                        requestData.Dash0202reference2="21"	 //rechazadas general 1005
				   case requestData.Dash0202reference2== "226" :
                                        requestData.Dash0202reference2="41"	 //rechazadas hards
				   case requestData.Dash0202reference2== "227" :
                                        requestData.Dash0202reference2="61"	 //rechazadas invalid merchat
				   case requestData.Dash0202reference2== "228" :
                                        requestData.Dash0202reference2="81"	 //rechazadas not honored
				   case requestData.Dash0202reference2== "223" :
                                        requestData.Dash0202reference2="41,61,81"	
				  
                                     }			
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
							var cuantosMaxVisual  int
							cuantosMaxVisual=7   //default V1
							 switch  {
							   case requestData.Dash0202reference2== "11" && requestData.Dash0202Dato01 =="99" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash01_99
								 
							   case requestData.Dash0202reference2== "11" && requestData.Dash0202Dato01 =="24" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash01_24
							   case requestData.Dash0202reference2== "11" && requestData.Dash0202Dato01 =="48" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash01_48

							   case requestData.Dash0202reference2== "21" && requestData.Dash0202Dato01 =="99" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash02_99
								 //set the graph value nbr to 31, as this is already there amounts per hour
								 //do not use the one this value had 21
								 requestData.Dash0202reference2 ="31"

							   case requestData.Dash0202reference2== "21" && requestData.Dash0202Dato01 =="7" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash02_7
								 //set the graph value nbr to 31, as this is already there amounts per hour
								 //do not use the one this value had 21
								 requestData.Dash0202reference2 ="31"

							   case requestData.Dash0202reference2== "21" && requestData.Dash0202Dato01 =="30" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash02_30
								 //set the graph value nbr to 31, as this is already there amounts per hour
								 //do not use the one this value had 21
								 requestData.Dash0202reference2 ="31"

							   case requestData.Dash0202reference2== "31" && requestData.Dash0202Dato01 =="99" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash03_99
							   case requestData.Dash0202reference2== "31" && requestData.Dash0202Dato01 =="24" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash03_24
							   case requestData.Dash0202reference2== "31" && requestData.Dash0202Dato01 =="48" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash03_48

							   case requestData.Dash0202reference2== "41" && requestData.Dash0202Dato01 =="99" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash04_99
								 //set the graph value nbr to 31, as this is already there amounts per hour
								 //do not use the one this value had 41
								 requestData.Dash0202reference2 ="31"

							   case requestData.Dash0202reference2== "41" && requestData.Dash0202Dato01 =="7" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash04_7
								 //set the graph value nbr to 31, as this is already there amounts per hour
								 //do not use the one this value had 41
								 requestData.Dash0202reference2 ="31"

							   case requestData.Dash0202reference2== "41" && requestData.Dash0202Dato01 =="30" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash04_30
								 //set the graph value nbr to 31, as this is already there amounts per hour
								 //do not use the one this value had 41
								 requestData.Dash0202reference2 ="31"


							   case requestData.Dash0202reference2== "51" && requestData.Dash0202Dato01 =="99" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash05_99
							   case requestData.Dash0202reference2== "51" && requestData.Dash0202Dato01 =="24" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash05_24
							   case requestData.Dash0202reference2== "51" && requestData.Dash0202Dato01 =="48" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash05_48

							   case requestData.Dash0202reference2== "61" && requestData.Dash0202Dato01 =="99" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash06_99
								 //set the graph value nbr to 31, as this is already there amounts per hour
								 //do not use the one this value had 61
								 requestData.Dash0202reference2 ="31"

							   case requestData.Dash0202reference2== "61" && requestData.Dash0202Dato01 =="7" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash06_7
								 //set the graph value nbr to 71, as this is already there amounts per hour denegadas
								 //do not use the one this value had 61
								 requestData.Dash0202reference ="71"

							   case requestData.Dash0202reference2== "61" && requestData.Dash0202Dato01 =="30" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash06_30
								 //set the graph value nbr to 31, as this is already there amounts per hour denegadas
								 //do not use the one this value had 61
								 requestData.Dash0202reference2 ="71"

							   case requestData.Dash0202reference2== "71" && requestData.Dash0202Dato01 =="99" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash07_99
							   case requestData.Dash0202reference2== "71" && requestData.Dash0202Dato01 =="24" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash07_24
							   case requestData.Dash0202reference2== "71" && requestData.Dash0202Dato01 =="48" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash07_48

							   case requestData.Dash0202reference2== "81" && requestData.Dash0202Dato01 =="99" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash08_99
								 //set the graph value nbr to 31, as this is already there amounts per hour denegadas
								 //do not use the one this value had 81
								 requestData.Dash0202reference2 ="71"
							   case requestData.Dash0202reference2== "81" && requestData.Dash0202Dato01 =="7" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash08_7
								 //set the graph value nbr to 31, as this is already there amounts per hour denegadas
								 //do not use the one this value had 81
								 requestData.Dash0202reference2 ="71"					 
							   case requestData.Dash0202reference2== "81" && requestData.Dash0202Dato01 =="30" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash08_30
								 //set the graph value nbr to 31, as this is already there amounts per hour denegadas
								 //do not use the one this value had 81
								 requestData.Dash0202reference2 ="71"
			// version 2 dash21 dash22 dash23
			//START
							   case requestData.Dash0202reference2== "221" && requestData.Dash0202Dato01 =="00" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_00
								 requestData.Dash0202reference2 ="31"
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "221" && requestData.Dash0202Dato01 =="01" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_01
								 requestData.Dash0202reference2 ="31"
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "221" && requestData.Dash0202Dato01 =="02" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_02
								 requestData.Dash0202reference2 ="31"
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "221" && requestData.Dash0202Dato01 =="70" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_70
								 requestData.Dash0202reference2 ="31"
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "221" && requestData.Dash0202Dato01 =="71" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_71
								 requestData.Dash0202reference2 ="31"
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "221" && requestData.Dash0202Dato01 =="72" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_72
								 requestData.Dash0202reference2 ="31"
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "221" && requestData.Dash0202Dato01 =="30" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_30
								 requestData.Dash0202reference2 ="31"
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "221" && requestData.Dash0202Dato01 =="31" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_31
								 requestData.Dash0202reference2 ="31"
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "221" && requestData.Dash0202Dato01 =="32" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_32
								 requestData.Dash0202reference2 ="31"
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "222" && requestData.Dash0202Dato01 =="00" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_00
								 requestData.Dash0202reference2 ="11" //transacciones aprobadas
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "222" && requestData.Dash0202Dato01 =="01" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_01
								 requestData.Dash0202reference2 ="11" //transacciones aprobadas
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "222" && requestData.Dash0202Dato01 =="02" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_02
								 requestData.Dash0202reference2 ="11" //transacciones aprobadas
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "222" && requestData.Dash0202Dato01 =="70" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_70
								 requestData.Dash0202reference2 ="11" //transacciones aprobadas
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "222" && requestData.Dash0202Dato01 =="71" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_71
								 requestData.Dash0202reference2 ="11" //transacciones aprobadas
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "222" && requestData.Dash0202Dato01 =="72" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_72
								 requestData.Dash0202reference2 ="11" //transacciones aprobadas
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "222" && requestData.Dash0202Dato01 =="30" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_30
								 requestData.Dash0202reference2 ="11" //transacciones aprobadas
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "222" && requestData.Dash0202Dato01 =="31" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_31
								 requestData.Dash0202reference2 ="11" //transacciones aprobadas
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "222" && requestData.Dash0202Dato01 =="32" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_32
								 requestData.Dash0202reference2 ="11" //transacciones aprobadas
								 cuantosMaxVisual=len(Config_topCustomer)
//1005
							   case requestData.Dash0202reference2== "225" && requestData.Dash0202Dato01 =="00" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_00
								 requestData.Dash0202reference2 ="21" //rechazadas general
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "225" && requestData.Dash0202Dato01 =="01" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_01
								 requestData.Dash0202reference2 ="21" //rechazadas general
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "225" && requestData.Dash0202Dato01 =="02" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_02
								 requestData.Dash0202reference2 ="21" //rechazadas general
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "225" && requestData.Dash0202Dato01 =="70" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_70
								 requestData.Dash0202reference2 ="21" //rechazadas general
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "225" && requestData.Dash0202Dato01 =="71" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_71
								 requestData.Dash0202reference2 ="21" //rechazadas general
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "225" && requestData.Dash0202Dato01 =="72" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_72
								 requestData.Dash0202reference2 ="21" //rechazadas general
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "225" && requestData.Dash0202Dato01 =="30" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_30
								 requestData.Dash0202reference2 ="21" //rechazadas general
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "225" && requestData.Dash0202Dato01 =="31" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_31
								 requestData.Dash0202reference2 ="21" //rechazadas general
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "225" && requestData.Dash0202Dato01 =="32" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_32
								 requestData.Dash0202reference2 ="21" //rechazadas general
								 cuantosMaxVisual=len(Config_topCustomer)								 
//	hards
							   case requestData.Dash0202reference2== "226" && requestData.Dash0202Dato01 =="00" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_00
								 requestData.Dash0202reference2 ="41" //rechazadas hards
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "226" && requestData.Dash0202Dato01 =="01" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_01
								 requestData.Dash0202reference2 ="41" //rechazadas hards
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "226" && requestData.Dash0202Dato01 =="02" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_02
								 requestData.Dash0202reference2 ="41" //rechazadas hards
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "226" && requestData.Dash0202Dato01 =="70" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_70
								 requestData.Dash0202reference2 ="41" //rechazadas hards
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "226" && requestData.Dash0202Dato01 =="71" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_71
								 requestData.Dash0202reference2 ="41" //rechazadas hards
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "226" && requestData.Dash0202Dato01 =="72" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_72
								 requestData.Dash0202reference2 ="41" //rechazadas hards
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "226" && requestData.Dash0202Dato01 =="30" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_30
								 requestData.Dash0202reference2 ="41" //rechazadas hards
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "226" && requestData.Dash0202Dato01 =="31" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_31
								 requestData.Dash0202reference2 ="41" //rechazadas hards
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "226" && requestData.Dash0202Dato01 =="32" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_32
								 requestData.Dash0202reference2 ="41" //rechazadas hards
								 cuantosMaxVisual=len(Config_topCustomer)									 
//rechazadas invalid merchant
							   case requestData.Dash0202reference2== "227" && requestData.Dash0202Dato01 =="00" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_00
								 requestData.Dash0202reference2 ="61" //rechazadas invalid merchant
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "227" && requestData.Dash0202Dato01 =="01" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_01
								 requestData.Dash0202reference2 ="61" //rechazadas invalid merchant
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "227" && requestData.Dash0202Dato01 =="02" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_02
								 requestData.Dash0202reference2 ="61" //rechazadas invalid merchant
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "227" && requestData.Dash0202Dato01 =="70" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_70
								 requestData.Dash0202reference2 ="61" //rechazadas invalid merchant
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "227" && requestData.Dash0202Dato01 =="71" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_71
								 requestData.Dash0202reference2 ="61" //rechazadas invalid merchant
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "227" && requestData.Dash0202Dato01 =="72" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_72
								 requestData.Dash0202reference2 ="61" //rechazadas invalid merchant
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "227" && requestData.Dash0202Dato01 =="30" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_30
								 requestData.Dash0202reference2 ="61" //rechazadas invalid merchant
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "227" && requestData.Dash0202Dato01 =="31" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_31
								 requestData.Dash0202reference2 ="61" //rechazadas invalid merchant
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "227" && requestData.Dash0202Dato01 =="32" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_32
								 requestData.Dash0202reference2 ="61" //rechazadas invalid merchant
								 cuantosMaxVisual=len(Config_topCustomer)									 
//rechazadas Not honor								 
							   case requestData.Dash0202reference2== "228" && requestData.Dash0202Dato01 =="00" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_00
								 requestData.Dash0202reference2 ="81" //rechazadas Not honor
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "228" && requestData.Dash0202Dato01 =="01" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_01
								 requestData.Dash0202reference2 ="81" //rechazadas Not honor
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "228" && requestData.Dash0202Dato01 =="02" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_02
								 requestData.Dash0202reference2 ="81" //rechazadas Not honor
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "228" && requestData.Dash0202Dato01 =="70" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_70
								 requestData.Dash0202reference2 ="81" //rechazadas Not honor
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "228" && requestData.Dash0202Dato01 =="71" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_71
								 requestData.Dash0202reference2 ="81" //rechazadas Not honor
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "228" && requestData.Dash0202Dato01 =="72" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_72
								 requestData.Dash0202reference2 ="81" //rechazadas Not honor
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "228" && requestData.Dash0202Dato01 =="30" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_30
								 requestData.Dash0202reference2 ="81" //rechazadas Not honor
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "228" && requestData.Dash0202Dato01 =="31" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_31
								 requestData.Dash0202reference2 ="81" //rechazadas Not honor
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "228" && requestData.Dash0202Dato01 =="32" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_32
								 requestData.Dash0202reference2 ="81" //rechazadas Not honor
								 cuantosMaxVisual=len(Config_topCustomer)									 
//
							   case requestData.Dash0202reference2== "223" && requestData.Dash0202Dato01 =="00" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_30
								 requestData.Dash0202reference2 ="41,61,81"
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "223" && requestData.Dash0202Dato01 =="01" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_31
								 requestData.Dash0202reference2 ="41,61,81"
								 cuantosMaxVisual=len(Config_topCustomer)
							   case requestData.Dash0202reference2== "223" && requestData.Dash0202Dato01 =="02" :
								cualConfig_comandosqlfront = Config_comandosql_frontdash21_02
								 requestData.Dash0202reference2 ="41,61,81"	
								 cuantosMaxVisual=len(Config_topCustomer)

			//END
			// version 2 dash21 dash22 dash23				 

							     }
								log.Print("manda a func  logicDBProcessDash02Grafica02 ref:\n"+requestData.Dash0202reference);                          
								log.Print(" ref2:\n"+requestData.Dash0202reference2);
							
							    resultDatadash,errCards =modelito.GetDatadash0202(db,
							       requestData.Dash0202reference,
							     requestData.Dash0202reference2,
							     requestData.Dash0202Dato01,
							     requestData.Dash0202Dato02,
							     requestData.Dash0202Dato03,
								cualConfig_comandosqlfront )

								 log.Print("regresa func  logicDBProcessDash02Grafica02 ok!\n"+cualConfig_comandosqlfront)
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
									log.Print("el registor traede la db :"+d.ID+" valor:"+d.Valoramount)
									    cuantos =cuantos+1

									//after the cuantosMaxVisualth, sum in one
									if(cuantos >cuantosMaxVisual){
									    valors,err01 := strconv.Atoi(d.Valoramount)
									    if(err01!=nil){

									    }//end if

									    acumuladoLastHr = acumuladoLastHr +valors
									}else{

									    //take the firsts cuantosMaxVisual , as they come
									    resultLastHrDatadash =append (resultLastHrDatadash,d)

									}//end else de cuantos >cuantosMaxVisual
								}//end for
							    if cuantos >cuantosMaxVisual {

								//set values for the Other element
								otherLastHrDatadash.ID ="1"

								valor2s := strconv.Itoa(acumuladoLastHr)
								otherLastHrDatadash.Valoramount=valor2s                                
								otherLastHrDatadash.Valorcolumna=valor2s
								otherLastHrDatadash.Nombrecolumna="Other"
								otherLastHrDatadash.Numsecuecial="10"

								resultLastHrDatadash =append (resultLastHrDatadash,otherLastHrDatadash)
							    } else { //not enough records for the cuantosMaxVisual lines in the screen
								     i := cuantos //cuantos is less than cuantosMaxVisual, so set the rest of the 
								    if cuantos ==0{
									    i=1  
								    }
								    for i < cuantosMaxVisual+1 {
									//set values for the Other element
									otherLastHrDatadash.ID ="1"

									otherLastHrDatadash.Valoramount="0"
									otherLastHrDatadash.Valorcolumna="0"
									otherLastHrDatadash.Nombrecolumna=" "
									otherLastHrDatadash.Numsecuecial="10"

									resultLastHrDatadash =append (resultLastHrDatadash,otherLastHrDatadash)
									i = i + 1
								    }

							    }//end else cuantos
/*
						if cuantos == 0 {
						  log.Print("DB: records not found")
						  errorGeneral="Not records found for the logicDBProcessDash02Grafica02 reference received"
						}		
*/
					    }//end else
				
				
					}
				    
				//  END fetchFromDB
   
   //END
   	  return  resultLastHrDatadash, errorGeneral

   }

