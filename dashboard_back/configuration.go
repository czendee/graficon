package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"fmt"
	"banwire/dash/dashboard_back/db"
	"banwire/dash/dashboard_back/path"
)

// Initializes the Config variables
var Config config

var HTTPListen string
var configFile string

var RunMode string

//configuration for the dest/dash DB 
var Config_DB_pass string
var Config_DB_user string
var Config_DB_name string
var Config_DB_server string
var Config_DB_port int
var Config_WS_crbanwire_pass string
var Config_dbStringType string
var Config_connString string

var Config_task_1periocity int
var Config_task_2periocity int
var Config_task_timeframe int

//configuration for the source DB 
var Config_DB_origin_pass string
var Config_DB_origin_user string
var Config_DB_origin_name string
var Config_DB_origin_server string
var Config_DB_origin_port int
var Config_dbStringType_origin string
var Config_connString_origin string
var Config_comandosqlorigen_origindash01 string
var Config_comandosqlorigen_origindash02 string

var Config_comandosqlorigen_origindash03 string
var Config_comandosqlorigen_origindash04 string
var Config_comandosqlorigen_origindash05 string
var Config_comandosqlorigen_origindash06 string
var Config_comandosqlorigen_origindash07 string
var Config_comandosqlorigen_origindash08 string

var Config_comandosqlorigen_origindash20 string    //all rechazadas
var Config_comandosqlorigen_origindash21 string    //rechazadas- hards
var Config_comandosqlorigen_origindash22 string    //rechazadas- invalid merch
var Config_comandosqlorigen_origindash23 string    //rechazadas- hnot honored
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
	flag.StringVar(&HTTPListen, "http", ":8070", "Path where HTTP Server will listening")
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
		case "tasks":
			log.Print("tasks data 04")
			for _, n := range d.Nodes {
							log.Print("tasks data 04")
				if active, _ := n["active"].(bool); active {
					task1, _ := n["task1periocity"].(float64)
					task2, _ := n["task2periocity"].(float64)
					timeframe, _ := n["timeframe"].(float64)
                  
                    
                    Config_task_1periocity =int(task1)            
                    Config_task_2periocity =int(task2)            
                    Config_task_timeframe =int(timeframe)            

//                    log.Print("---- The Task value periocity for task 1  was assigned "+Config_task_1periocity)                    
//                    log.Print("---- The Task value periocity for task 2  was assigned "+Config_task_2periocity)
//                    log.Print("---- The Task value  timeframe was assigned "+Config_task_timeframe)
				}//end active
							log.Print("UnmarshalJSON 06!")
			}

		case "sourcedatadb":
			log.Print("UnmarshalJSON 04!")
			for _, n := range d.Nodes {
							log.Print("UnmarshalJSON 05!")
				if active, _ := n["active"].(bool); active {
					host, _ := n["host"].(string)
					port, _ := n["port"].(float64)
					_db, _ := n["db"].(string)
					user, _ := n["user"].(string)
					pass, _ := n["password"].(string)
                    typedb, _ := n["type"].(string)
                    
                    commandSqldash01, _ := n["comandosqlorigendash01"].(string)
                    commandSqldash02, _ := n["comandosqlorigendash02"].(string)
                    commandSqldash03, _ := n["comandosqlorigendash03"].(string)
                    commandSqldash04, _ := n["comandosqlorigendash04"].(string)
                    commandSqldash05, _ := n["comandosqlorigendash05"].(string)
                    commandSqldash06, _ := n["comandosqlorigendash06"].(string)
                    commandSqldash07, _ := n["comandosqlorigendash07"].(string)
                    commandSqldash08, _ := n["comandosqlorigendash08"].(string)
                    commandSqldash20, _ := n["comandosqlorigendash20"].(string)
                    commandSqldash21, _ := n["comandosqlorigendash21"].(string)
                    commandSqldash22, _ := n["comandosqlorigendash22"].(string)
                    commandSqldash23, _ := n["comandosqlorigendash23"].(string)
                    

                    Config_DB_origin_pass =pass
                    Config_DB_origin_user =user
                    Config_DB_origin_name =_db
                    Config_DB_origin_server =host
                    Config_DB_origin_port =int(port)            
                    
                    Config_dbStringType_origin=typedb   
                    Config_connString_origin = fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
						Config_DB_origin_server,Config_DB_origin_name, Config_DB_origin_user, Config_DB_origin_pass, Config_DB_origin_port)
				
                    Config_comandosqlorigen_origindash01 = commandSqldash01
                    Config_comandosqlorigen_origindash02 = commandSqldash02
                    Config_comandosqlorigen_origindash03 = commandSqldash03
                    Config_comandosqlorigen_origindash04 = commandSqldash04
                    Config_comandosqlorigen_origindash05 = commandSqldash05
                    Config_comandosqlorigen_origindash06 = commandSqldash06
                    Config_comandosqlorigen_origindash07 = commandSqldash07
		    Config_comandosqlorigen_origindash08 = commandSqldash08
                    Config_comandosqlorigen_origindash20 = commandSqldash20   //rechazadas- ALL 1005
                    Config_comandosqlorigen_origindash21 = commandSqldash21   //rechazadas- hards
                    Config_comandosqlorigen_origindash22 = commandSqldash22   //rechazadas- invalid merch
                    Config_comandosqlorigen_origindash23 = commandSqldash23   //rechazadas- not honored
                    
					
                    log.Print("---- The DB values  was assigned "+Config_DB_origin_server)                    
                    log.Print("---- The DB values  was assigned "+Config_DB_origin_user)
                    log.Print("---- The DB values  was assigned "+Config_DB_origin_pass)
                    log.Print("---- The DB values  was assigned "+Config_DB_origin_name)
					if e := db.Connection.Set(db.NewPgDb(host, int(port), _db, user, pass)); e == nil {
						log.Print("---- The postgresql database was loaded"+host)
						log.Print("---- The postgresql database was loaded"+_db)
					} else {
						return e
					}
				}
							log.Print("UnmarshalJSON 06!")
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
                    
                    

                    Config_DB_pass =pass
                    Config_DB_user =user
                    Config_DB_name =_db
                    Config_DB_server =host
                    Config_DB_port =int(port)            
                    
                    Config_dbStringType="postgres"
                    Config_connString = fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
						Config_DB_server,Config_DB_name, Config_DB_user, Config_DB_pass, Config_DB_port)
				

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
