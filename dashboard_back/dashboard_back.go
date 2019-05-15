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


func dashBack01Grafica01() {
    var errorGeneral string

    var valoresParaResponder  string
    fmt.Println("doing job :dashBack01Grafica01")
    valoresParaResponder,errorGeneral=v1ProcessDashBack01Grafica01()   // in logicbusiness.go

    if errorGeneral==""{//continue next step
        fmt.Println("respuesta job :"+valoresParaResponder)
    }else{
        fmt.Println("error respuesta job :"+errorGeneral)
    }
    
}





func main() {
    fmt.Println("main  ticker  01")

    
	LoadConfiguration()


fmt.Println("main  ticker 02")
    
    ticker:= time.NewTicker(time.Duration(Config_task_1periocity)*time.Second)
fmt.Println("main ticker  03")    
    go func(){
          for t:=range ticker.C{
              fmt.Println("main ticker  05:",t)
               dashBack01Grafica01()
          }
    }()

    
    time.Sleep(time.Duration(Config_task_timeframe) * time.Second)
    fmt.Println("main ticker  06")
 //to quit the goroutine
    ticker.Stop()    
    
    fmt.Println("main ticker  07")
}