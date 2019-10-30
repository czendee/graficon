package main

import (
        "time"
        "fmt"
)


const (
	DefaultRunMode = ""
	ApiRunMode     = "api"
	BatchRunMode   = "batch"
)

func init() {
//	loaded.Add(1)
	
}



func stuff() {
    fmt.Println("doing job")
}


func dashBack01Grafica01() {      //extract the PAGADAS/ACEPTADAS amount  31
    var errorGeneral string

    var valoresParaResponder  string
    fmt.Println("doing job :dashBack01Grafica01")
    valoresParaResponder,errorGeneral=v1ProcessDashBack01Grafica01("31")   // in logicbusiness.go

    if errorGeneral==""{//continue next step
        fmt.Println("respuesta job :"+valoresParaResponder)
    }else{
        fmt.Println("error respuesta job :"+errorGeneral)
    }
    
}

func dashBack01Grafica02() {   //extract the Denegadas  amount 71
    var errorGeneral string

    var valoresParaResponder  string
    fmt.Println("doing job :dashBack01Grafica01")
    valoresParaResponder,errorGeneral=v1ProcessDashBack01Grafica01("71")   // in logicbusiness.go

    if errorGeneral==""{//continue next step
        fmt.Println("respuesta job :"+valoresParaResponder)
    }else{
        fmt.Println("error respuesta job :"+errorGeneral)
    }
    
}

func dashBack01GraficaAll() {    //extract all the pagadas  num transactions   11
    var errorGeneral string

    var valoresParaResponder  string
    fmt.Println("doing job :dashBack01Grafica01")
    valoresParaResponder,errorGeneral=v1ProcessDashBack01Grafica01("11")   // in logicbusiness.go

    if errorGeneral==""{//continue next step
        fmt.Println("respuesta job :"+valoresParaResponder)
    }else{
        fmt.Println("error respuesta job :"+errorGeneral)
    }
    
}


func dashBack05GraficaAll() {    //extract all the denegadas  num transactions 51
    var errorGeneral string

    var valoresParaResponder  string
    fmt.Println("doing job :dashBack01Grafica01")
    valoresParaResponder,errorGeneral=v1ProcessDashBack01Grafica01("51")   // in logicbusiness.go

    if errorGeneral==""{//continue next step
        fmt.Println("respuesta job :"+valoresParaResponder)
    }else{
        fmt.Println("error respuesta job :"+errorGeneral)
    }
    
}


func dashBackExtraeRechazadas( codigoData string) {    //extract Data for a Data conde
    var errorGeneral string

    var valoresParaResponder  string
    fmt.Println("doing job :dashBackExtraeRechazadas"+codigoData)
    valoresParaResponder,errorGeneral=v1ProcessDashBack01Grafica01(codigoData)   // in logicbusiness.go

    if errorGeneral==""{//continue next step
        fmt.Println("respuesta job :"+valoresParaResponder)
    }else{
        fmt.Println("error respuesta job :"+errorGeneral)
    }
    
}


func main() {
    fmt.Println("main  ticker  01")

    
	LoadConfiguration()


fmt.Println("main  ticker 02+rechazadasTambien")
    
    ticker:= time.NewTicker(time.Duration(Config_task_1periocity)*time.Second)
fmt.Println("main ticker  03")    
    go func(){
          for t:=range ticker.C{
              fmt.Println("main ticker  05:",t)
               dashBack01Grafica01() //obtains from source total amount per hour pagagas
              fmt.Println("main ticker  06:",t)               
               dashBack01Grafica02() //obtains from source total amount per hour denegadas
               dashBack01GraficaAll() //obtains from source number of transactions per hour pagagas
               dashBack05GraficaAll() //obtains from source number of transactions per hour denegadas	
		  fmt.Println("rechazadas21,41,61,81")
		   dashBackExtraeRechazadas( "21") //rechazadas- ALL 1005   - 21
		  
		  dashBackExtraeRechazadas( "41") //rechazadas- hards      -41
		  dashBackExtraeRechazadas( "61") //rechazadas- invalid merch -61
		  dashBackExtraeRechazadas( "81") //rechazadas- not honored   -81
          }
    }()

    
    time.Sleep(time.Duration(Config_task_timeframe) * time.Second)
    fmt.Println("main ticker  07")
 //to quit the goroutine
    ticker.Stop()    
    
    fmt.Println("main ticker  08")
}
