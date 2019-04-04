package main

import (
//	"banwire/dash/dashboard_back/net"
//	conf "banwire/dash/dashboard_back/conf"
//	"flag"
//	"log"
//	"sync"
    "github.com/jasonlvhit/gocron" // Use here: go get github.com/jasonlvhit/gocron
        "time"
        "fmt"
)

//var loaded sync.WaitGroup

const (
	DefaultRunMode = ""
	ApiRunMode     = "api"
	BatchRunMode   = "batch"
)

func init() {
//	loaded.Add(1)
	
}

/*
func main() {
//	flag.Parse()
//	LoadConfiguration()

//	loaded.Done()

	log.Print("Jobs is ready for run...")

	// Do jobs with params
//	gocron.Every(3).Second().Do(taskWithParams, 1, "hello")

	// Do jobs without params
//	gocron.Every(5).Second().Do(task)
	gocron.Every(1).Seconds().Do(task)
    log.Print("Jobs is running")
      time.Sleep(10 * time.Second)
      log.Print("Jobs is running after sleep")


	switch RunMode {
	case BatchRunMode:
		BatchTest()
		break
	case DefaultRunMode, ApiRunMode:
		startServer()
		break
	default:
		log.Print("Run mode is unknown")
		break
	}

}

*/

/*

func startServer() {
	log.Printf("HTTP Listening: %s", HTTPListen)
	net.HTTPListener(HTTPListen)
}
*/

/*

func task() {
		log.Print("I am runnning task.")
}

func taskWithParams(a int, b string) {
		log.Print(a, b)
}

*/
func jobs(quit <-chan bool) {
    for {
        //cron jobs
        g := gocron.NewScheduler()

        g.Every(uint64(Config_task_1periocity)).Seconds().Do(dashBack01Grafica01)
        g.Every(uint64(Config_task_2periocity)).Second().Do(stuff)

//        g := gocron.NewScheduler()
//        g.Every(1).Second().Do(stuff)


        select {
        case <-quit:
            // here we receive from quit and exit
            // if `g` has started, we may want to `g.Clear()`
            // or the scheduled jobs will continue, we will just not be waiting for them to finish.
            return
        case <-g.Start():
            // here we know all the jobs are done, and go round the loop again
        }
    }
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
    fmt.Println("main   01")

    //	flag.Parse()
	LoadConfiguration()

//	loaded.Done()
fmt.Println("main   02")
    q := make(chan bool)
    go jobs(q)
    
    time.Sleep(time.Duration(Config_task_timeframe) * time.Second)

 //to quit the goroutine
    q <- true    
    close(q)
    fmt.Println("main   03")
}