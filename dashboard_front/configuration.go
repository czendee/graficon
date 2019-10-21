package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"fmt"
	"banwire/dash/dashboard_front/db"
	"banwire/dash/dashboard_front/path"
)

// Initializes the Config variables
var Config config

var HTTPListen string
var configFile string

var RunMode string

var Config_DB_pass string
var Config_DB_user string
var Config_DB_name string
var Config_DB_server string
var Config_DB_port int
var Config_WS_crbanwire_pass string
var Config_dbStringType string
var Config_connString string
var Config_env_server string
var Config_env_url string
var Config_env_port string

var Config_comandosql_frontdash01_99 string
var Config_comandosql_frontdash01_24 string
var Config_comandosql_frontdash01_48 string

var Config_comandosql_frontdash02_99 string
var Config_comandosql_frontdash02_7 string
var Config_comandosql_frontdash02_30 string

var Config_comandosql_frontdash03_99 string
var Config_comandosql_frontdash03_24 string
var Config_comandosql_frontdash03_48 string

var Config_comandosql_frontdash04_99 string
var Config_comandosql_frontdash04_7 string
var Config_comandosql_frontdash04_30 string

var Config_comandosql_frontdash05_99 string
var Config_comandosql_frontdash05_24 string
var Config_comandosql_frontdash05_48 string

var Config_comandosql_frontdash06_99 string
var Config_comandosql_frontdash06_7 string
var Config_comandosql_frontdash06_30 string

var Config_comandosql_frontdash07_99 string
var Config_comandosql_frontdash07_24 string
var Config_comandosql_frontdash07_48 string

var Config_comandosql_frontdash08_99 string
var Config_comandosql_frontdash08_7 string
var Config_comandosql_frontdash08_30 string
//version 2:START
var Config_comandosql_frontdash21_00 string
var Config_comandosql_frontdash21_01 string
var Config_comandosql_frontdash21_02 string

var Config_comandosql_frontdash21_70 string
var Config_comandosql_frontdash21_71 string
var Config_comandosql_frontdash21_72 string

var Config_comandosql_frontdash21_30 string
var Config_comandosql_frontdash21_31 string
var Config_comandosql_frontdash21_32 string

var Config_topCustomer[8] string
//version 2:END
/*    const (
        DB_USER     = "lerepagr"        
        DB_PASSWORD = "Ag8q2utgSsVy2tyR7_M9cNYbzsqSvwma"
        DB_NAME     = "lerepagr"
        DB_SERVER     = "stampy.db.elephantsql.com" //"54.163.207.112"
        DB_PORT      = 5432

        CR_BANWIRE_USER      = "pruebasbw" //this was mentioned by charly Jan 2019
    )
*/

func init() {
	flag.StringVar(&RunMode, "mode", "api", "Service mode run (Options: api, batch)")
//to be set later	flag.StringVar(&HTTPListen, "http", ":8077", "Path where HTTP Server will listening")
	flag.StringVar(&configFile, "config", "./conf/config.json", "Path of configuration file")
//	flag.StringVar(&configFile, "config", "config.json", "Path of configuration file")

	configFile = path.RelativePath(configFile)
}


// loadConfiguration loads the configuration file
func LoadConfiguration() {
	log.Print("Loading configuration...v3.5")

	if d, e := ioutil.ReadFile(configFile); e == nil {
		e := json.Unmarshal(d, &Config)
		if e != nil {
//			log.Panicf("Error in unmarshalling the configuration file: %s", e.Error())
            log.Print("Configuration was not was loaded!")
		}else{
			log.Print("Configuration was YESS was loaded!")
		}

        
		log.Print("Configuration was loaded!(check previo")
	} else {
		//log.Panicf("Error in opening the configuration file: %s", e)
			log.Print("Error in opening the configuration file %s", e)
	}

    //reset the port
    	flag.StringVar(&HTTPListen, "http", Config_env_port, "Path where HTTP Server will listening")
}

// config is the configuration structure object
type config struct {
//	Database configDatabase `json:"database"`
	Database configDatabase `json:"othersources"`
}

// configDatabase is the database structure object for configuration
type configDatabase struct{}

// UnmarshalJSON handles desearialization of configDatabase
// and loads the othersources: database connections and webservices connections
func (c *configDatabase) UnmarshalJSON(data []byte) error {
 log.Print("UnmarshalJSON 00!:othersources")
	var cc = []struct {
		Type string                   `json:"type"`
		Nodes []map[string]interface{} `json:"nodes"`
	}{}
			log.Print("UnmarshalJSON 01!")
	err := json.Unmarshal(data, &cc)
	if err != nil {
		return err
	}
			log.Print("UnmarshalJSON 02!")
	for _, d := range cc {
					log.Print("UnmarshalJSON 03.!"+d.Type)
		switch d.Type {
		case "crbanwire":
			log.Print("UnmarshalJSON 2.04.2!")
			for _, n := range d.Nodes {
							log.Print("UnmarshalJSON 2.05!")
				if active, _ := n["active"].(bool); active {
                    passcrban,_:=n["passwordcrbanwire"].(string)
                    
                    log.Print("---- The value  was loaded"+passcrban)

                     Config_WS_crbanwire_pass = passcrban
                    log.Print("---- The crbanwire value  was assigned es "+Config_WS_crbanwire_pass)
				}
				log.Print("UnmarshalJSON 2.06!")
			}

		case "envserver":
			log.Print("UnmarshalJSON 2.04.2! envserver")
			for _, n := range d.Nodes {
							log.Print("UnmarshalJSON 2.05!")
				if active, _ := n["active"].(bool); active {
                    valenv,_:=n["envlevel"].(string)
                    valurl,_:=n["envurl"].(string)
                    valport,_:=n["envport"].(string)
                    log.Print("---- The value  was loaded"+valenv)

                     Config_env_server = valenv
                     Config_env_url = valurl
                     Config_env_port = valport
                    log.Print("---- The env level value  was assigned es "+Config_env_server)
                    log.Print("---- The env level port  was assigned es "+Config_env_port)
				}
				log.Print("UnmarshalJSON 2.06! envserver")
			}

		case "postgresql":
			log.Print("UnmarshalJSON 04!")
			for _, n := range d.Nodes {
							log.Print("UnmarshalJSON 05!")
				if active, _ := n["active"].(bool); active {
					host, _ := n["host"].(string)
					port, _ := n["port"].(float64)
					_db, _ := n["db"].(string)
					user, _ := n["user"].(string)
					pass, _ := n["password"].(string)
                    
                    
					                   
                    commandSqldash01_99, _ := n["comandosqlorigendash01_99"].(string)
                    commandSqldash01_24, _ := n["comandosqlorigendash01_24"].(string)
                    commandSqldash01_48, _ := n["comandosqlorigendash01_48"].(string)
                    commandSqldash02_99, _ := n["comandosqlorigendash02_99"].(string)
                    commandSqldash02_7, _ := n["comandosqlorigendash02_7"].(string)
                    commandSqldash02_30, _ := n["comandosqlorigendash02_30"].(string)
                    commandSqldash03_99, _ := n["comandosqlorigendash03_99"].(string)
                    commandSqldash03_24, _ := n["comandosqlorigendash03_24"].(string)
                    commandSqldash03_48, _ := n["comandosqlorigendash03_48"].(string)
                    commandSqldash04_99, _ := n["comandosqlorigendash04_99"].(string)
                    commandSqldash04_7, _ := n["comandosqlorigendash04_7"].(string)
                    commandSqldash04_30, _ := n["comandosqlorigendash04_30"].(string)
                    commandSqldash05_99, _ := n["comandosqlorigendash05_99"].(string)
                    commandSqldash05_24, _ := n["comandosqlorigendash05_24"].(string)
                    commandSqldash05_48, _ := n["comandosqlorigendash05_48"].(string)
                    commandSqldash06_99, _ := n["comandosqlorigendash06_99"].(string)
                    commandSqldash06_7, _ := n["comandosqlorigendash06_7"].(string)
                    commandSqldash06_30, _ := n["comandosqlorigendash06_30"].(string)
                    commandSqldash07_99, _ := n["comandosqlorigendash07_99"].(string)
                    commandSqldash07_24, _ := n["comandosqlorigendash07_24"].(string)
                    commandSqldash07_48, _ := n["comandosqlorigendash07_48"].(string)
                    commandSqldash08_99, _ := n["comandosqlorigendash08_99"].(string)
                    commandSqldash08_7, _ := n["comandosqlorigendash08_7"].(string)
                    commandSqldash08_30, _ := n["comandosqlorigendash08_30"].(string)
					
	//new version 2: start					
                    commandSqldash21_00, _ := n["comandosqlorigendash21_00"].(string)
                    commandSqldash21_01, _ := n["comandosqlorigendash21_01"].(string)
                    commandSqldash21_02, _ := n["comandosqlorigendash21_02"].(string)
                    commandSqldash21_70, _ := n["comandosqlorigendash21_70"].(string)
                    commandSqldash21_71, _ := n["comandosqlorigendash21_71"].(string)
                    commandSqldash21_72, _ := n["comandosqlorigendash21_72"].(string)					
                    commandSqldash21_30, _ := n["comandosqlorigendash21_30"].(string)
                    commandSqldash21_31, _ := n["comandosqlorigendash21_31"].(string)
                    commandSqldash21_32, _ := n["comandosqlorigendash21_32"].(string)					
					
                    topCustomer01, _       := n["topCustomer_01"].(string)	
		    topCustomer02, _       := n["topCustomer_02"].(string)						
		    topCustomer03, _       := n["topCustomer_03"].(string)						
		    topCustomer04, _       := n["topCustomer_04"].(string)						
		    topCustomer05, _       := n["topCustomer_05"].(string)						
		    topCustomer06, _       := n["topCustomer_06"].(string)						
		    topCustomer07, _       := n["topCustomer_07"].(string)						
		    topCustomer08, _       := n["topCustomer_08"].(string)											
	//new version 2: end                    

				
					

                    Config_DB_pass =pass
                    Config_DB_user =user
                    Config_DB_name =_db
                    Config_DB_server =host
                    Config_DB_port =int(port)            
                    
                    Config_dbStringType="postgres"
                    Config_connString = fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
						Config_DB_server,Config_DB_name, Config_DB_user, Config_DB_pass, Config_DB_port)
				

                    Config_comandosql_frontdash01_99 =commandSqldash01_99
                    Config_comandosql_frontdash01_24 =commandSqldash01_24
                    Config_comandosql_frontdash01_48 =commandSqldash01_48					
                    Config_comandosql_frontdash02_99 =commandSqldash02_99
                    Config_comandosql_frontdash02_7 =commandSqldash02_7
                    Config_comandosql_frontdash02_30 =commandSqldash02_30					
                    Config_comandosql_frontdash03_99 =commandSqldash03_99
                    Config_comandosql_frontdash03_24 =commandSqldash03_24
                    Config_comandosql_frontdash03_48 =commandSqldash03_48					
                    Config_comandosql_frontdash04_99 =commandSqldash04_99
                    Config_comandosql_frontdash04_7 =commandSqldash04_7
                    Config_comandosql_frontdash04_30 =commandSqldash04_30
					
                    Config_comandosql_frontdash05_99 =commandSqldash05_99
                    Config_comandosql_frontdash05_24 =commandSqldash05_24
                    Config_comandosql_frontdash05_48 =commandSqldash05_48					
                    Config_comandosql_frontdash06_99 =commandSqldash06_99
                    Config_comandosql_frontdash06_7 =commandSqldash06_7
                    Config_comandosql_frontdash06_30 =commandSqldash06_30					
                    Config_comandosql_frontdash07_99 =commandSqldash07_99
                    Config_comandosql_frontdash07_24 =commandSqldash07_24
                    Config_comandosql_frontdash07_48 =commandSqldash07_48					
                    Config_comandosql_frontdash08_99 =commandSqldash08_99
                    Config_comandosql_frontdash08_7 =commandSqldash08_7
                    Config_comandosql_frontdash08_30 =commandSqldash08_30
//new version 2   START					
                    Config_comandosql_frontdash21_00 =commandSqldash21_00
                    Config_comandosql_frontdash21_01 =commandSqldash21_01
                    Config_comandosql_frontdash21_02 =commandSqldash21_02
                    Config_comandosql_frontdash21_70 =commandSqldash21_70
                    Config_comandosql_frontdash21_71 =commandSqldash21_71
                    Config_comandosql_frontdash21_72 =commandSqldash21_72					
                    Config_comandosql_frontdash21_30 =commandSqldash21_30
                    Config_comandosql_frontdash21_31 =commandSqldash21_31
                    Config_comandosql_frontdash21_32 =commandSqldash21_32
		    Config_topCustomer[0] = topCustomer01
		    Config_topCustomer[1] = topCustomer02
		    Config_topCustomer[2] = topCustomer03
		    Config_topCustomer[3] = topCustomer04
		    Config_topCustomer[4] = topCustomer05
		    Config_topCustomer[5] = topCustomer06
		    Config_topCustomer[6] = topCustomer07
		    Config_topCustomer[7] = topCustomer08
//new version 2    END					
                    log.Print("---- The DB values  was assigned "+Config_DB_server)                    
                    log.Print("---- The DB values  was assigned "+Config_DB_user)
                    log.Print("---- The DB values  was assigned "+Config_DB_pass)
                    log.Print("---- The DB values  was assigned "+Config_DB_name)
                    log.Print("---- ne sql 03 99 "+Config_comandosql_frontdash03_99)
                    log.Print("---- ne sql 02 7 "+Config_comandosql_frontdash02_7)		
                    log.Print("---- ne sql 08 30 "+Config_comandosql_frontdash08_30)	
//new version 2   START					
                    log.Print("---- ne sql 21 00 "+Config_comandosql_frontdash21_00)					
                    log.Print("---- ne sql 21 01 "+Config_comandosql_frontdash21_01)					
		    log.Print("---- ne sql 21 02 "+Config_comandosql_frontdash21_02)
                    log.Print("---- ne sql 21 00 "+Config_comandosql_frontdash21_70)					
                    log.Print("---- ne sql 21 01 "+Config_comandosql_frontdash21_31)					
		    log.Print("---- ne sql 21 02 "+Config_comandosql_frontdash21_32)					
//new version 2   END					
					if e := db.Connection.Set(db.NewPgDb(host, int(port), _db, user, pass)); e == nil {
						log.Print("---- The postgresql database was loaded"+host)
						log.Print("---- The postgresql database was loaded"+_db)
					} else {
						return e
					}
				}
							log.Print("UnmarshalJSON 06!")
			}
		case "mysql":
			log.Print("UnmarshalJSON 04! mysql")
			for _, n := range d.Nodes {
							log.Print("UnmarshalJSON 05! mysql")
				if active, _ := n["active"].(bool); active {
					host, _ := n["host"].(string)
					port, _ := n["port"].(float64)
					_db, _ := n["db"].(string)
					user, _ := n["user"].(string)
					pass, _ := n["password"].(string)
                    
                    

                    Config_DB_pass =pass
                    Config_DB_user =user
                    Config_DB_name =_db
                    Config_DB_server =host
                    Config_DB_port =int(port)            
                    


                    Config_dbStringType="mysql"
                    Config_connString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
//                    Config_connString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
						Config_DB_user, Config_DB_pass,Config_DB_server,Config_DB_port,Config_DB_name )
				
				
//"mysql", "root:password1@tcp(127.0.0.1:3306)/test"

 

                    log.Print("---- The DB values  was assigned "+Config_DB_server)                    
                    log.Print("---- The DB values  was assigned "+Config_DB_user)
                    log.Print("---- The DB values  was assigned "+Config_DB_pass)
                    log.Print("---- The DB values  was assigned "+Config_DB_name)
					if e := db.Connection.Set(db.NewPgDb(host, int(port), _db, user, pass)); e == nil {
						log.Print("---- The postgresql database was loaded"+host)
						log.Print("---- The postgresql database was loaded"+_db)
					} else {
						return e
					}
				}
							log.Print("UnmarshalJSON 06!")
			}

			break
		}
			log.Print("UnmarshalJSON 07!")
	}

	return nil
}
