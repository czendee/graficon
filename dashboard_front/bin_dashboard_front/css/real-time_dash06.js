
"use strict";
var valueRotacion=20000;
var valueRotacionOriginal=20000;

var valuePasarNext=150000;
var valuePasarNextOriginal=150000;

var previousGroupNumber=0;//used to keep track of what group was displayed previously

var previousGroupNumber0202=0;//used to keep track of what group was displayed previously


function show_selected2() {
    alert("aj");
    var selector = document.getElementById('id_of_select');
    valueRotacion= selector[selector.selectedIndex].value;

    document.getElementById('display').innerHTML = valueRotacion;
}


$(function () {
//document.getElementById('btn').addEventListener('click', show_selected2);;
//document.getElementById('btn02').addEventListener('click', go_next2);;
//document.getElementById('id_of_select').addEventListener('click', show_selected3);;
//document.getElementById('id_of_selectRotacion').addEventListener('click', show_selected_rotacion3);;

var activities = document.getElementById("id_of_selectRotacion");

activities.addEventListener("change", function() {
    alert("aj id_of_selectRotacion");

    var selector = document.getElementById('id_of_selectRotacion');
    valueRotacion= selector[selector.selectedIndex].value;
    alert("aj id_of_selectRotacion:"+valueRotacion);

});

var activities = document.getElementById("id_of_select");

activities.addEventListener("change", function() {
    alert("aj id_of_select");

    var selector = document.getElementById('id_of_select');
    valuePasarNext= selector[selector.selectedIndex].value;
    alert("aj id_of_select:"+valuePasarNext);

});


	var activeUsers = 99,
			pageViewsPerSecondLowerLimit,
			pageViewsPerSecondUpperLimit,
			yValuePageViewsPerSecond,
			sumYValuePageViewsPerSecond = 0,
			numberOfSeconds = 1,
			updateChartsInterval,
			updateChartsIntervalLowerLimit = 4000, // milliseconds
			updateChartsIntervalUpperLimit = 6000, // milliseconds
			timeoutIdUpdateCharts;
	
	var pageViewsPerSecondDataPoints = [];
	var pageViewsPerMinuteDataPoints = [];
	
	// data for demo only
	var initialDataPageViewsPerSecond = [1,2,4,4,3,2,1,5,1,2,2,0,0,1,2,4,5,3,4,2,2,2,2,0,1,2,4,4,4,5,5,1,2,4,1,1,1,0,0,1,2,3,3,5,5,2,0,1,1,0,2,2,2,0,4,1,4,4,2,2];
	var initialDataPageViewsPerMinute = [110,107,122,107,128,108,100,110,118,93,95,112,108,95,96,114,107,105,124,104,131,94,109,110,108,99,104,90,104,109,89,121,118,93,109,113,106,100,101,119,76,137,112,104,98,89,104,96,120,111,108,95,93,100,101,110,98,105,107,135];
	
	// data for demo only
	var data = [
		{
			activeUsers: 99,

			categories: [
				{ name: "Men Clothing", users: 8 },
				{ name: "Women Clothing", users: 9 },
				{ name: "Gadgets", users: 10 },
				{ name: "Books", users: 5 },
				{ name: "Others", users: 23 }
			]
            
		},
		{
			activeUsers: 56,
			categories: [
				{ name: "Men Clothing", users: 9 },
				{ name: "Women Clothing", users: 9 },
				{ name: "Gadgets", users: 10 },
				{ name: "Books", users: 5 },
				{ name: "Others", users: 23 }
			]
            
		},
		{
			activeUsers: 57,
			categories: [
				{ name: "Men Clothing", users: 8 },
				{ name: "Women Clothing", users: 9 },
				{ name: "Gadgets", users: 11 },
				{ name: "Books", users: 5 },
				{ name: "Others", users: 24 }
			]
            
		},
		{
			activeUsers: 58,
			categories: [
				{ name: "Men Clothing", users: 9 },
				{ name: "Women Clothing", users: 10 },
				{ name: "Gadgets", users: 11 },
				{ name: "Books", users: 4 },
				{ name: "Others", users: 24 }
			]
            
		},
		{
			activeUsers: 59,

			categories: [
				{ name: "Men Clothing", users: 9 },
				{ name: "Women Clothing", users: 10 },
				{ name: "Gadgets", users: 11 },
				{ name: "Books", users: 5 },
				{ name: "Others", users: 24 }
			]
		},
		{
			activeUsers: 60,

			categories: [
				{ name: "Men Clothing", users: 9 },
				{ name: "Women Clothing", users: 10 },
				{ name: "Gadgets", users: 12 },
				{ name: "Books", users: 5 },
				{ name: "Others", users: 24 }
			]
		},
		{
			activeUsers: 61,

			categories: [
				{ name: "Men Clothing", users: 9 },
				{ name: "Women Clothing", users: 10 },
				{ name: "Gadgets", users: 12 },
				{ name: "Books", users: 5 },
				{ name: "Others", users: 25 }
			]
		}
	];
	
	CanvasJS.addColorSet("customColorSet", [ 
		"#393f63",
		"#e5d8B0", 
		"#ffb367", 
		"#f98461", 
		"#d9695f",
		"#e05850",
	]);

	// CanvasJS pie chart to active users by category
	var usersCategoryPieChart = new CanvasJS.Chart("users-category-pie-chart", {
		animationDuration: 800,
		animationEnabled: true,
		backgroundColor: "transparent",
		colorSet: "customColorSet",
		legend: {
			fontFamily: "calibri",
			fontSize: 14,
			horizontalAlign: "left",
			verticalAlign: "center",
			maxWidth: null,
			itemTextFormatter: function (e) {
				return e.dataPoint.name + ": " + Math.round(e.dataPoint.y / activeUsers * 100) + "%";  
			}
		},
		toolTip: {
			cornerRadius: 0,
			fontStyle: "normal",
			contentFormatter: function (e) {
				return e.entries[0].dataPoint.name + ": " + Math.round(e.entries[0].dataPoint.y / activeUsers * 100) + "% (" + e.entries[0].dataPoint.y  + ")";
			} 
		},
		data: [
			{
				legendMarkerType: "square",
				radius: "90%",
				showInLegend: true,
				startAngle: 90,
				type: "pie",
				dataPoints: [
					{ y: 8, name:"Men Clothing" },
					{ y: 9, name:"Women Clothing" },
					{ y: 10, name:"Gadgets" },
					{ y: 5, name:"Books" },
					{ y: 23, name:"Others" }
				]
			}
		]
	});
			
	// CanvasJS column chart to show live page views per minute
	var pageViewsPerMinuteColumnChart = new CanvasJS.Chart("page-views-per-minute-column-chart", {
		animationDuration: 800,
		animationEnabled: true,
		backgroundColor: "transparent",
		axisX: {
			interval: 1,
			intervalType: "minute",
			labelAutoFit: false,
			labelFontColor: "#717171",
			lineColor: "#a2a2a2",
			tickColor: "transparent",
			tickLength: 2
//            ,
//			labelFormatter: function(e) {
//				var diff, currentTime = (new Date()).getTime();
//				diff = Math.floor((e.value.getTime() - currentTime) / (1000 * 60));
//				return diff % 15 < 0 ? "" : diff + " min";
//			}
		},
		axisY: {
			includeZero: false,
			gridThickness: 0,
			labelFontColor: "#717171",
			lineColor: "#a2a2a2",
			tickColor: "#a2a2a2"
		},
		toolTip: {
			cornerRadius: 0,
			fontStyle: "normal"
		},
		data: [
			{
				color: "#424973",
				xValueFormatString: "hh:mm TT",
				type: "column",
				dataPoints : pageViewsPerMinuteDataPoints
			}
		]
	});


	//----------------------------------------------------------------------------------//
	var allCharts = [
//		usersDeviceDoughnutChart,
//		usersMediumPieChart,
		usersCategoryPieChart,
//		pageViewsPerSecondColumnChart,
		pageViewsPerMinuteColumnChart
		
	];
	
	// generate random number between given range
	function generateRandomNumber (minimum, maximum) {
		return Math.floor(Math.random() * (maximum - minimum + 1)) + minimum;
	}
	
	
	function updateUsersCategoryChart(dataIndex) {
		for (var i = 0; data[dataIndex].categories.length>= i  ; i++)
			usersCategoryPieChart.options.data[0].dataPoints[i].y = data[dataIndex].categories[i].users;
		
		usersCategoryPieChart.render();
	}
	

function updateUsersCategoryChartWithDB(dataIndex)
{
                    //use a parameter PageTitle {{.PageTitle}} set in net_v1.go with the server url
        $.getJSON("{{.PageTitle}}/v1/getDash02Grafica02?reference="+previousGroupNumber0202+"&reference2=62&dato01=99&dato02=88&dato03=77", function(data) {                  
            var linea=0;

             var exito="0";//error/not found
//             var exito=1;// found and groupNubr is greater than previousGroupNoumber0202,so get the most current data from db into javascript
//             var exito=2;// found and groupNubr is not greater than previousGroupNoumber0202, so use the current data in javascript

            $.each(data, function(key, value){
                if(linea ==0){//primera linea, viene estatus Success or Error

                }
                if(linea ==1){//second linea, viene estatus value 1 or Error
                     var statusResultado= value;
                     exito=statusResultado
                }
                if(linea ==2 && exito=="1"){//third  linea, viene nuevo group number

                   var grupoNumberResultado= value;
                   previousGroupNumber0202 =grupoNumberResultado;

                }                
                if(linea ==3 && exito=="1"){//foruth linea, viene an array with the result data
                    var arrayResultados= value;//here the array of data structres passed
                    arrayResultados.forEach(myFunction0202) //set the values in the  graph points ,below

                    //pageViewsPerMinuteDataPoints.push({ x: parseInt(value[0]),y: parseInt(value[1]) });
             
                }
                linea=linea+1;
                 console.log("json banwire recent"+value[1]); 
                 
            });	
            if(exito=="1"){ //only update data if new data group was found, more recent than the previuous group number

                 usersCategoryPieChart.render();

            }
        });
      
        usersCategoryPieChart.render();

}


function myFunction0202(item, index) {
   //  demoP.innerHTML = demoP.innerHTML + "index[" + index + "]: " + item + "<br>"; 

   usersCategoryPieChart.options.data[0].dataPoints[index].y = parseInt(item["data_valuea"]);
   usersCategoryPieChart.options.data[0].dataPoints[index].name = item["data_name"];

/*
    Object.keys(item).forEach(function(key) {
    console.log('Key : ' + key + ', Value : ' + item[key])
    })
*/

}

	// update all charts with revelant demo data, except "Page Views Per Second" and "Page Views Per Minute" charts
	function updateCharts(dataIndex) {
		activeUsers = data[dataIndex].activeUsers;
//		pageViewsPerSecondLowerLimit = data[dataIndex].pageViewsPerSecondLowerLimit;
//		pageViewsPerSecondUpperLimit = data[dataIndex].pageViewsPerSecondUpperLimit;

//dash02 no tiene		updateUsersDeviceChart(dataIndex);
//dash02 no tiene		updateUsersMediumPieChart(dataIndex);
		// usar datos demo     updateUsersCategoryChart(dataIndex);
        //usar datos de la db
        updateUsersCategoryChartWithDB(dataIndex);
		
	}
	
	function updateChartsAtRandomIntervals() {
		var dataIndex = generateRandomNumber(0, data.length - 1);
		updateChartsInterval = generateRandomNumber(updateChartsIntervalLowerLimit, updateChartsIntervalUpperLimit);
				
		updateCharts(dataIndex);
		
		if (timeoutIdUpdateCharts)
			clearTimeout(timeoutIdUpdateCharts);
		
		timeoutIdUpdateCharts = setTimeout(function () {
			updateChartsAtRandomIntervals();
		}, updateChartsInterval);
	}


	// populate "Chart 0201 " charts with initial data
	function populateChart0201() {
		
        console.log("populateChart0201 json banwire populateChart0201 "); 

//use a parameter PageTitle {{.PageTitle}} set in net_v1.go with the server url

        $.getJSON("{{.PageTitle}}/v1/getDash02Grafica02?reference=0&reference2=61&dato01=99&dato02=88&dato03=77", function(data) {  
            
            $.each(data, function(key, value){
            
                pageViewsPerMinuteDataPoints.push({ x: parseInt(value[0]),y: parseInt(value[1]) });
            
                 console.log("populateChart0201 json banwire"+value[1]); 
                 
            });	
            pageViewsPerMinuteColumnChart.render();      
        });
     pageViewsPerMinuteColumnChart.render();   
	}


	// update "updateChart0201" chart every set of seconds 
	function updateChart0201() {
			console.log("updateChart0201: json banwire recent GET");
            
//        $.getJSON("https://8070-dot-3809294-dot-devshell.appspot.com/v1/getjsondatabanwire", function(data) {  
         //send the  previousGroupNumber=0;
         // datePreviousGroupNumber=0;
         // currentPreviousGroupNumber=0;

         //the first time, previousGroupNumber will be cero, 
         // the service will decide if max available is greater than previousGroupNumber
         //  if greater   then obtain the next
         //  if equal or less, then keep the current group og data, display the message: the most current data available, and the current Date, previousDate

//use a parameter PageTitle {{.PageTitle}} set in net_v1.go with the server url
        $.getJSON("{{.PageTitle}}/v1/getDash02Grafica02?reference="+previousGroupNumber+"&reference2=61&dato01=99&dato02=88&dato03=77", function(data) {                  
            var linea=0;

             var exito="0";//error/not found
//             var exito=1;// found and groupNubr is greater than previousGroupNoumber,so get the most current data from db into javascript
//             var exito=2;// found and groupNubr is not greater than previousGroupNoumber, so use the current data in javascript

            $.each(data, function(key, value){
                if(linea ==0){//primera linea, viene estatus Success or Error

                }
                if(linea ==1){//second linea, viene estatus value 1 or Error
                     var statusResultado= value;
                     exito=statusResultado
                }
                if(linea ==2 && exito=="1"){//third  linea, viene nuevo group number

                   var grupoNumberResultado= value;
                   previousGroupNumber =grupoNumberResultado;

                }                
                if(linea ==3 && exito=="1"){//foruth linea, viene an array with the result data
                    var arrayResultados= value;//here the array of data structres passed
                    arrayResultados.forEach(myFunction0201) //set the values in the  graph points ,below

                    //pageViewsPerMinuteDataPoints.push({ x: parseInt(value[0]),y: parseInt(value[1]) });
             
                }
                linea=linea+1;
                 console.log("json banwire recent"+value[1]); 
                 
            });	
            if(exito=="1"){ //only update data if new data group was found, more recent than the previuous group number
                if (pageViewsPerMinuteDataPoints.length > 60 ) {
                        pageViewsPerMinuteDataPoints.shift();			
                }
                pageViewsPerMinuteColumnChart.render();

            }
        });
      
     pageViewsPerMinuteColumnChart.render();   
}

function myFunction0201(item, index) {
//  demoP.innerHTML = demoP.innerHTML + "index[" + index + "]: " + item + "<br>"; 

	pageViewsPerMinuteDataPoints.push({ x: parseInt(item["data_valuec"]),y: parseInt(item["data_valuea"]) });

    Object.keys(item).forEach(function(key) {
    console.log('Key : ' + key + ', Value : ' + item[key])
    //        pageViewsPerMinuteDataPoints.push({ x: parseInt(value[0]),y: parseInt(value[1]) });

    })

}

	// chart properties cutomized further based on screen width	
	function chartPropertiesCustomization(chart) {
		if ($(window).outerWidth() >= 1920) {			
			
			chart.options.legend.fontSize = 14;
			chart.options.legend.horizontalAlign = "left";
			chart.options.legend.verticalAlign = "center";
			chart.options.legend.maxWidth = null;
			
		}else if (1920 >= $(window).outerWidth()   && $(window).outerWidth() >= 1200) {
			
			chart.options.legend.fontSize = 14;
			chart.options.legend.horizontalAlign = "left";
			chart.options.legend.verticalAlign = "center";
			chart.options.legend.maxWidth = 140;
			
		} else if (1200 >= $(window).outerWidth()  && $(window).outerWidth() >= 992) {
			
			chart.options.legend.fontSize = 12;
			chart.options.legend.horizontalAlign = "center";
			chart.options.legend.verticalAlign = "top";
			chart.options.legend.maxWidth = null;
			
		} else if (992 >= $(window).outerWidth()  ) {
			
			chart.options.legend.fontSize = 14;
			chart.options.legend.horizontalAlign = "center";
			chart.options.legend.verticalAlign = "bottom";
			chart.options.legend.maxWidth = null;
			
		}
		
		chart.render();
	}
	
	function customizeCharts() {
//dash02 no tiene		chartPropertiesCustomization(usersDeviceDoughnutChart);
//dash02 no tiene		chartPropertiesCustomization(usersMediumPieChart);
		chartPropertiesCustomization(usersCategoryPieChart);
	}
	
	function renderAllCharts() {
		for (var i = 0; allCharts.length>= i  ; i++)
			allCharts[i].render();
	}
	
	function sidebarToggleOnClick() {
		$('#sidebar-toggle-button').on('click', function () {
			$('#sidebar').toggleClass('sidebar-toggle');
			$('#page-content-wrapper').toggleClass('page-content-toggle');
			renderAllCharts();
		});	
	}

	function checkRotationValues() {
        if(valueRotacionOriginal==valueRotacion){
             //continue with the same value
        }else{
            setInterval(updateChart0201, valueRotacion);//every 20 seconds
            valueRotacionOriginal=valueRotacion;
        }
         
        if(valuePasarNextOriginal==valuePasarNext){
             //continue with the same value
        }else{
            setInterval(go_next2, valuePasarNext);//every 20 seconds
            valuePasarNextOriginal=valuePasarNext;
        }         
	}

    function go_next2() {
        alert("next");
        window.location.href = "/v1/dash02";  

    }


	(function init() {
		customizeCharts();
		$(window).resize(customizeCharts);
//		populateChart0202();
        populateChart0201();
//		setInterval(updateChart0202, 1000); //every 1 second
        setInterval(updateChart0201, valueRotacionOriginal);//every 20 seconds
		setTimeout(updateChartsAtRandomIntervals, 4000); //every 4 seconds

        setInterval(checkRotationValues, valueRotacionOriginal);//every 20 seconds
        
        

        setInterval(go_next2, valuePasarNextOriginal);//every 150 seconds
        

		sidebarToggleOnClick();
	})();
	
    
});