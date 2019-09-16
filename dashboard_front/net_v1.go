
package main

import (
	"net/http"
	"fmt"
	"log"
	"banwire/dash/dashboard_front/db"
	"banwire/dash/dashboard_front/net"
	modelito "banwire/dash/dashboard_front/model"

    "banwire/dash/dashboard_front/model"
	"banwire/dash/dashboard_front/model/pgsql"
    "encoding/json"

    //start: para dashboard
	"html/template"
	"io/ioutil"
//	"log"
//	"net/http"
//	"regexp"
   //end: para dashboard
    
//	"time"
//	"encoding/json"
//	 "database/sql"
	 _ "github.com/lib/pq"   //use go get github.com/lib/pq
     
)

    
// init loads the routes for version 1
func init() {
//	var _r = net.GetRouter()
//	var r = _r.PathPrefix("/v1").Subrouter()

    var r = net.GetRouter()
	//route for test
	    log.Print("cz  init net_v1")

	r.Handle("/v1/testdbcharly", netHandle(handleTestV1, nil)).Methods("GET")     //logicbusiness.go


/////////start Para Dashboard


    r.Handle("/v1/getdash01grafica01", netHandle(v1handleGETDash01Grafica01, nil)).Methods("GET")     //logicbusiness.go
    r.Handle("/v1/getdash01grafica02", netHandle(v1handleGETDash01Grafica02, nil)).Methods("GET")     //logicbusiness.go

    r.Handle("/v1/getdash02grafica01", netHandle(v1handleGETDash02Grafica01, nil)).Methods("GET")     //logicbusiness.go
    r.Handle("/v1/getdash02grafica02", netHandle(v1handleGETDash02Grafica02, nil)).Methods("GET")     //logicbusiness.go
    r.Handle("/v1/getdash02grafica03", netHandle(v1handleGETDash02Grafica03, nil)).Methods("GET")     //logicbusiness.go

    r.Handle("/v1/getdash03grafica01", netHandle(v1handleGETDash03Grafica01, nil)).Methods("GET")     //logicbusiness.go
    r.Handle("/v1/getdash03grafica02", netHandle(v1handleGETDash03Grafica02, nil)).Methods("GET")     //logicbusiness.go
    r.Handle("/v1/getdash03grafica03", netHandle(v1handleGETDash03Grafica03, nil)).Methods("GET")     //logicbusiness.go


//    r.Handle("/v1/getjsondatabanwire", netHandle(getJsonDataBanwire, nil)).Methods("GET")     //logicbusiness.go
    r.Handle("/v1/getDash02Grafica01", netHandle(v1handleGETDash02Grafica01, nil)).Methods("GET")     //logicbusiness.go

    r.Handle("/v1/getDash02Grafica02", netHandle(v1handleGETDash02Grafica02, nil)).Methods("GET")     //logicbusiness.go    


    r.Handle("/v1/dash01", netHandle(dash01Handler, nil)).Methods("GET")     
    r.Handle("/v1/dash02", netHandle(dash02Handler, nil)).Methods("GET")     
    r.Handle("/v1/dash03", netHandle(dash03Handler, nil)).Methods("GET") 
    r.Handle("/v1/dash04", netHandle(dash04Handler, nil)).Methods("GET")
    r.Handle("/v1/dash05", netHandle(dash05Handler, nil)).Methods("GET") 
    r.Handle("/v1/dash06", netHandle(dash06Handler, nil)).Methods("GET") 
    r.Handle("/v1/dash07", netHandle(dash07Handler, nil)).Methods("GET") 
    r.Handle("/v1/dash08", netHandle(dash08Handler, nil)).Methods("GET")     
    r.Handle("/v1/ejemplo", netHandle(ejemploHandler, nil)).Methods("GET")     

    r.Handle("/font-roboto.css", netHandle(serveCss01, nil)).Methods("GET")     
    r.Handle("/bootstrap4-alpha3.min.css", netHandle(serveCss02, nil)).Methods("GET")     
    r.Handle("/font-awesome.min.css", netHandle(serveCss03, nil)).Methods("GET")     
    r.Handle("/style.css", netHandle(serveCss04, nil)).Methods("GET")     
    r.Handle("/jquery-3.1.0.min.js", netHandle(serveJs01, nil)).Methods("GET")     
    r.Handle("/tether.min.js", netHandle(serveJs02, nil)).Methods("GET")     
    r.Handle("/bootstrap4-alpha3.min.js", netHandle(serveJs03, nil)).Methods("GET")     
    r.Handle("/canvasjs.min.js", netHandle(serveJs04, nil)).Methods("GET")     
    r.Handle("/real-time.js", netHandle(serveJs05, nil)).Methods("GET")     
    r.Handle("/real-time_dash01.js", netHandle(serveJs05dash01, nil)).Methods("GET")  
    r.Handle("/real-time_dash02.js", netHandle(serveJs05dash02, nil)).Methods("GET")  
    r.Handle("/real-time_dash03.js", netHandle(serveJs05dash03, nil)).Methods("GET")
    r.Handle("/real-time_dash04.js", netHandle(serveJs05dash04, nil)).Methods("GET")
    r.Handle("/real-time_dash05.js", netHandle(serveJs05dash05, nil)).Methods("GET")
    r.Handle("/real-time_dash06.js", netHandle(serveJs05dash06, nil)).Methods("GET")
    r.Handle("/real-time_dash07.js", netHandle(serveJs05dash07, nil)).Methods("GET")
    r.Handle("/real-time_dash08.js", netHandle(serveJs05dash08, nil)).Methods("GET")  




/////////end Para Dashboard

}

// handleTest is an example for receive and handle the request from client
func handleTestV1(w http.ResponseWriter, r *http.Request) {
	var response []byte
	var err error

	defer func() {
		db.Connection.Close(r)
		if err != nil {
			err = model.UnmarshalJSONError(err)
			response, _ = json.Marshal(err)
		}
		if response != nil {
			rw := net.ResponseWriterJSON(w)
			rw.Write(response)
		}
	}()

    log.Print("cz  handleTestV1")

	fmt.Fprint(w,"youtochi   antes de  probar la db")
	
	_pg := pgsql.Db{
		db.Connection.GetPgDb(r, false),
	}

	fmt.Fprint(w,"youtochi   establece conn - probar la db")

    //llama una func que acceda la db
    var miCardResult model.Card
    fmt.Fprint(w,"     youtochi   establece conn - 2   ")
    err, miCardResult= _pg.GetCardByTokenAndCust("crd.2Gm9cob47gMEet9hjcMbVS3oJwmf","14")
    


	fmt.Fprint(w,"   youtochi   ver si ok  - probar la db       ")
    log.Println(" 3  youtochi   ver si ok  - probar la db!      \n")
    


	if err != nil {
		var e = model.NewErrorGroup()
		if _e, ok := err.(*pgsql.Error); ok {

			switch _e.Type {
			case pgsql.ErrNotFound:
//				e.Push(model.InvalidValueError("id", error_not_found))
                 log.Print("ERROR carloitos Not Found!\n")
				break
//			case pgsql.ErrNotFoundStatusPayment:
//				e.Push(model.InvalidValueError("id_status_payment", "The status payment id does not exists"))
//				break
			default:
				log.Panic(_e.Type)
				break
			}

		} else {
			log.Panic(err)
		}

		if e.HasError() {
			err = e
			return
		}
	}else{
       log.Println("        youtochi   quesque si  - probar la db!      \n")
            
        log.Print("       youtochi   ver si ok  - ver el dato!\n"+miCardResult.Token)
    }
        
	fmt.Fprint(w,"youtochi   establece conn - probar la db")
    
	rw := net.ResponseWriterJSON(w)
	rw.Write([]byte(`{"readyCarlitos":true}`))
}



//////////////////////////////////start Para Dashboard

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func getJsonDataBanwire(w http.ResponseWriter, r *http.Request) {
    /// START
  //var uno string
  //uno="[[1,11],[2,6],[3,8],[4,5],[5,4],[6,6],[7,7],[8,6],[9,9],[10,12],[11,14],[12,14],[13,16],[14,12],[15,9],[16,6],[17,5],[18,6],[19,8],[20,10],[21,11],[22,11],[23,15],[24,16],[25,13],[26,9],[27,4],[28,2],[29,-1],[30,4],[31,2],[32,5],[33,0],[34,-5],[35,0],[36,0],[37,2],[38,3],[39,8],[40,8],[41,6],[42,7],[43,12],[44,17],[45,20],[46,16],[47,13],[48,10],[49,13],[50,17],[51,21],[52,19],[53,18],[54,21],[55,26],[56,28],[57,33],[58,28],[59,33],[60,29],[61,24],[62,21],[63,26],[64,21],[65,18],[66,23],[67,24],[68,29],[69,29],[70,29],[71,28],[72,32],[73,28],[74,27],[75,31],[76,35],[77,36],[78,31],[79,26],[80,24],[81,29],[82,33],[83,35],[84,34],[85,36],[86,38],[87,33],[88,34],[89,36],[90,31],[91,34],[92,37],[93,35],[94,37],[95,40],[96,41],[97,42],[98,41],[99,42],[100,38]]"
	w.Header().Set("Content-Type", "application/json")
//	w.Write( []byte("Hello") )
    w.Write( []byte("[[1,11],[2,6],[3,8],[4,5],[5,4],[6,6],[7,7],[8,6],[9,9],[10,12],[11,14],[12,14],[13,16],[14,12],[15,9],[16,6],[17,5],[18,6],[19,8],[20,10],[21,11],[22,11],[23,15],[24,16],[25,13],[26,9],[27,4],[28,2],[29,-1],[30,4] ]") )

//	rw := net.ResponseWriterJSON(w)
//	rw.Write([]byte(`{"readyCarlitos":true}`))   
}

///////////////////////////////////////////////////////////////////////////////////////////////////////77
///Dashs Graficas

// v1handleDBDash01Grafica01  receive and handle the request from client, mthod POST
func v1handlePOSTDash01Grafica01 (w http.ResponseWriter, r *http.Request) {
	defer func() {
		db.Connection.Close(nil)
	}()
    var errorGeneral string
    var errorGeneralNbr string
    var requestData modelito.RequestDash01Grafica01
//    var requestData modelito.RequestPayment
    
    errorGeneral=""
requestData,errorGeneral =obtainPostParmsProcessDash01Grafica01(r,errorGeneral)  //logicrequest_post.go

	////////////////////////////////////////////////validate parms
	/// START
	////////////////////////////////////////////////validate parms
	/// START
    
    if errorGeneral=="" {

		errorGeneral,errorGeneralNbr= v1ProcessDash01Grafica01(w , requestData)    //logicbusiness.go 
	}

    if errorGeneral!=""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		log.Print("CZ   STEP Get the ERROR response JSON ready")
		
			/// START
		fieldDataBytesJson,err := getJsonResponseError(errorGeneral, errorGeneralNbr)
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		if(err!=nil){
			
		}
	
    } 
					
}

// v1handleGETDash01Grafica01  receive and handle the request from client, method GET
func v1handleGETDash01Grafica01 (w http.ResponseWriter, r *http.Request) {
	defer func() {
		db.Connection.Close(nil)
	}()

					
}

// v1handleGETDash01Grafica02  receive and handle the request from client, method GET
func v1handleGETDash01Grafica02 (w http.ResponseWriter, r *http.Request) {
	defer func() {
		db.Connection.Close(nil)
	}()

					
}


// v1handleGETDash02Grafica01  receive and handle the request from client, method GET
func v1handleGETDash02Grafica01 (w http.ResponseWriter, r *http.Request) {
	defer func() {
		db.Connection.Close(nil)
	}()

					
}


// v1handleGETDash02Grafica02  receive and handle the request from client, method GET
func v1handleGETDash02Grafica02 (w http.ResponseWriter, r *http.Request) {
	defer func() {
		db.Connection.Close(nil)
	}()

    var errorGeneral string
    var	errorGeneralNbr string
    
    var requestData modelito.RequestDash02Grafica02
    errorGeneral=""
requestData,errorGeneral =obtainParmsProcessDash02Grafica02(r,errorGeneral) //logicrequest.go

	////////////////////////////////////////////////validate parms
	/// START
    
    if errorGeneral=="" {

		errorGeneral,errorGeneralNbr= v1ProcessDash02Grafica02(w , requestData)    //logicbusiness.go 
	}

    if errorGeneral!=""{
    	//send error response if any
    	//prepare an error JSON Response, if any
		log.Print("CZ   STEP Get the ERROR response JSON ready")
		
			/// START
		fieldDataBytesJson,err := getJsonResponseError(errorGeneral, errorGeneralNbr)
		//////////    write the response (ERROR)
		w.Header().Set("Content-Type", "application/json")
		w.Write(fieldDataBytesJson)	
		if(err!=nil){
			
		}
	
    } 
					
}



// v1handleGETDash02Grafica03  receive and handle the request from client, method GET
func v1handleGETDash02Grafica03 (w http.ResponseWriter, r *http.Request) {
	defer func() {
		db.Connection.Close(nil)
	}()

					
}
// v1handleGETDash03Grafica01  receive and handle the request from client, method GET
func v1handleGETDash03Grafica01 (w http.ResponseWriter, r *http.Request) {
	defer func() {
		db.Connection.Close(nil)
	}()

					
}


// v1handleGETDash03Grafica02  receive and handle the request from client, method GET
func v1handleGETDash03Grafica02 (w http.ResponseWriter, r *http.Request) {
	defer func() {
		db.Connection.Close(nil)
	}()

					
}



// v1handleGETDash03Grafica03  receive and handle the request from client, method GET
func v1handleGETDash03Grafica03 (w http.ResponseWriter, r *http.Request) {
	defer func() {
		db.Connection.Close(nil)
	}()
					
}


////////////////////////////////////////////////////////////////////////////77

func serveCss01(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/font-roboto.css")
}


func serveCss02(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/bootstrap4-alpha3.min.css")
}

func serveCss03(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/font-awesome.min.css")
}

func serveCss04(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/style.css")
}

func serveJs01(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/jquery-3.1.0.min.js")
}
func serveJs02(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/tether.min.js")
}
func serveJs03(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/bootstrap4-alpha3.min.js")
}
func serveJs04(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/canvasjs.min.js")
}

func serveJs05(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "css/real-time.js")
}

func serveJs05dash01(w http.ResponseWriter, r *http.Request) {

    log.Print("cz  serveJs05dash01:"+ "css/real-time_dash01.js")
//     http.ServeFile(w, r, "css/real-time_dash01.js")  //define in configuration.go


data := TodoPageData{
			PageTitle: Config_env_url,
}
tmpl := template.Must(template.ParseFiles("css/real-time_dash01.js"))
tmpl.Execute(w, data)


}

func serveJs05dash02(w http.ResponseWriter, r *http.Request) {
    log.Print("cz  serveJs05dash01:"+ "css/real-time_dash02.js")
//    http.ServeFile(w, r, "css/"+Config_env_server+"real-time_dash02.js")  //define in configuration.go


data := TodoPageData{
			PageTitle: Config_env_url,
}
tmpl := template.Must(template.ParseFiles("css/real-time_dash02.js"))
tmpl.Execute(w, data)

}

func serveJs05dash03(w http.ResponseWriter, r *http.Request) {
    log.Print("cz  serveJs05dash03:"+ "css/real-time_dash03.js")
//    http.ServeFile(w, r, "css/"+Config_env_server+"real-time_dash03.js")  //define in configuration.go

    
data := TodoPageData{
			PageTitle: Config_env_url,
}
tmpl := template.Must(template.ParseFiles("css/real-time_dash03.js"))
tmpl.Execute(w, data)

}

func serveJs05dash04(w http.ResponseWriter, r *http.Request) {
    log.Print("cz  serveJs05dash04:"+ "css/real-time_dash04.js")
//    http.ServeFile(w, r, "css/"+Config_env_server+"real-time_dash04.js")  //define in configuration.go

    
data := TodoPageData{
			PageTitle: Config_env_url,
}
tmpl := template.Must(template.ParseFiles("css/real-time_dash04.js"))
tmpl.Execute(w, data)

}

func serveJs05dash05(w http.ResponseWriter, r *http.Request) {
    log.Print("cz  serveJs05dash05:"+ "css/real-time_dash05.js")
//    http.ServeFile(w, r, "css/"+Config_env_server+"real-time_dash05.js")  //define in configuration.go

    
data := TodoPageData{
			PageTitle: Config_env_url,
}
tmpl := template.Must(template.ParseFiles("css/real-time_dash05.js"))
tmpl.Execute(w, data)

}

func serveJs05dash06(w http.ResponseWriter, r *http.Request) {
    log.Print("cz  serveJs05dash06:"+ "css/real-time_dash06.js")
//    http.ServeFile(w, r, "css/"+Config_env_server+"real-time_dash06.js")  //define in configuration.go

    
data := TodoPageData{
			PageTitle: Config_env_url,
}
tmpl := template.Must(template.ParseFiles("css/real-time_dash06.js"))
tmpl.Execute(w, data)

}

func serveJs05dash07(w http.ResponseWriter, r *http.Request) {
    log.Print("cz  serveJs05dash07:"+ "css/real-time_dash07.js")
//    http.ServeFile(w, r, "css/"+Config_env_server+"real-time_dash07.js")  //define in configuration.go

    
data := TodoPageData{
			PageTitle: Config_env_url,
}
tmpl := template.Must(template.ParseFiles("css/real-time_dash07.js"))
tmpl.Execute(w, data)

}

func serveJs05dash08(w http.ResponseWriter, r *http.Request) {
    log.Print("cz  serveJs05dash08:"+ "css/real-time_dash08.js")
//    http.ServeFile(w, r, "css/"+Config_env_server+"real-time_dash08.js")  //define in configuration.go

    
data := TodoPageData{
			PageTitle: Config_env_url,
}
tmpl := template.Must(template.ParseFiles("css/real-time_dash08.js"))
tmpl.Execute(w, data)

}

func ejemploHandler(w http.ResponseWriter, r *http.Request) {

log.Print("cz  ejemploHandler")

http.ServeFile(w,r,"ejemplo_copy.html")

log.Print("CZ   STEP ejemploHandler 01")
/*
    
//    title := r.URL.Path[len("/v1/ejemplo/"):]
    log.Print("CZ   STEP ejemploHandler 02")
    p, _ := loadPage(title)
    t, _ := template.ParseFiles("ejemplo.html")
    t.Execute(w, p)
*/    
}

type TodoPageData struct {
	PageTitle string
}
func dash01Handler(w http.ResponseWriter, r *http.Request) {

    log.Print("cz  dash01Handler with param"+Config_env_url)

//    http.ServeFile(w,r,"dash01.html")

data := TodoPageData{
			PageTitle: Config_env_server,
}
tmpl := template.Must(template.ParseFiles("dash01.html"))
tmpl.Execute(w, data)

    log.Print("CZ   STEP dash01Handler 01")
    
}

func dash02Handler(w http.ResponseWriter, r *http.Request) {

    log.Print("cz  dash02Handler")

    //http.ServeFile(w,r,"dash02.html")

    data := TodoPageData{
			PageTitle: Config_env_server,
    }
    tmpl := template.Must(template.ParseFiles("dash02.html"))
    tmpl.Execute(w, data)

    log.Print("CZ   STEP dash02Handler 01")
    
}

func dash03Handler(w http.ResponseWriter, r *http.Request) {

    log.Print("cz  dash03Handler")

    //http.ServeFile(w,r,"dash03.html")

    data := TodoPageData{
			PageTitle: Config_env_server,
    }
    tmpl := template.Must(template.ParseFiles("dash03.html"))
    tmpl.Execute(w, data)

    log.Print("CZ   STEP dash03Handler 01")
    
}

func dash04Handler(w http.ResponseWriter, r *http.Request) {

    log.Print("cz  dash04Handler")

    //http.ServeFile(w,r,"dash04.html")

    data := TodoPageData{
			PageTitle: Config_env_server,
    }
    tmpl := template.Must(template.ParseFiles("dash04.html"))
    tmpl.Execute(w, data)

    log.Print("CZ   STEP dash04Handler 01")
    
}

func dash05Handler(w http.ResponseWriter, r *http.Request) {

    log.Print("cz  dash05Handler")

    //http.ServeFile(w,r,"dash05.html") 

    data := TodoPageData{
			PageTitle: Config_env_server,
    }
    tmpl := template.Must(template.ParseFiles("dash05.html"))
    tmpl.Execute(w, data)

    log.Print("CZ   STEP dash05Handler 01")
    
}

func dash06Handler(w http.ResponseWriter, r *http.Request) {

    log.Print("cz  dash06Handler")

    //http.ServeFile(w,r,"dash06.html")

    data := TodoPageData{
			PageTitle: Config_env_server,
    }
    tmpl := template.Must(template.ParseFiles("dash06.html"))
    tmpl.Execute(w, data)

    log.Print("CZ   STEP dash06Handler 01")
    
}

func dash07Handler(w http.ResponseWriter, r *http.Request) {

    log.Print("cz  dash07Handler")

//    http.ServeFile(w,r,"dash07.html")


data := TodoPageData{
			PageTitle: Config_env_server,
}
tmpl := template.Must(template.ParseFiles("dash07.html"))
tmpl.Execute(w, data)

    log.Print("CZ   STEP dash07Handler 01")
    
}

func dash08Handler(w http.ResponseWriter, r *http.Request) {

    log.Print("cz  dash08Handler")

    //http.ServeFile(w,r,"dash08.html")

    data := TodoPageData{
			PageTitle: Config_env_server,
    }
    tmpl := template.Must(template.ParseFiles("dash08.html"))
    tmpl.Execute(w, data)

    log.Print("CZ   STEP dash08Handler 01")
    
}

/////////////////////////END para dashboard
