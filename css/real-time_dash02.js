
"use strict";
var valueRotacion=20000;
var valueRotacionOriginal=20000;

var valuePasarNext=150000;
var valuePasarNextOriginal=150000;


function show_selected2() {
    alert("aj");
    var selector = document.getElementById('id_of_select');
    valueRotacion= selector[selector.selectedIndex].value;

    document.getElementById('display').innerHTML = valueRotacion;
}



$(function () {
document.getElementById('btn').addEventListener('click', show_selected2);;
document.getElementById('btn02').addEventListener('click', go_next2);;

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
			pageViewsPerSecondLowerLimit: 0,
			pageViewsPerSecondUpperLimit: 3,

			device: [
				{ name: "PC", users: 39 },
				{ name: "Comercio", users: 5 },
				{ name: "Mobile", users: 11 }
			],
			trafficMedium: [
				{ name: "Organic", users: 38 },
				{ name: "Direct", users: 8 },
				{ name: "Paid", users: 5 },
				{ name: "Referral", users: 4 }
			],
			categories: [
				{ name: "Men Clothing", users: 8 },
				{ name: "Women Clothing", users: 9 },
				{ name: "Gadgets", users: 10 },
				{ name: "Books", users: 5 },
				{ name: "Others", users: 23 }
			],
			states: [
				{ name: "Others", users: 16 },
				{ name: "Pennsylvania", users: 4 },
				{ name: "Florida", users: 5 },
				{ name: "Texas", users: 7 },
				{ name: "New York", users: 11 },
				{ name: "California", users: 12 },
			]
		},
		{
			activeUsers: 56,
			pageViewsPerSecondLowerLimit: 0,
			pageViewsPerSecondUpperLimit: 3,

			device: [
				{ name: "PC", users: 40 },
				{ name: "Comercio", users: 5 },
				{ name: "Mobile", users: 11 }
			],
			trafficMedium: [
				{ name: "Organic", users: 39 },
				{ name: "Direct", users: 8 },
				{ name: "Paid", users: 5 },
				{ name: "Referral", users: 4 }
			],
			categories: [
				{ name: "Men Clothing", users: 9 },
				{ name: "Women Clothing", users: 9 },
				{ name: "Gadgets", users: 10 },
				{ name: "Books", users: 5 },
				{ name: "Others", users: 23 }
			],
			states: [
				{ name: "Others", users: 17 },
				{ name: "Pennsylvania", users: 4 },
				{ name: "Florida", users: 5 },
				{ name: "Texas", users: 7 },
				{ name: "New York", users: 11 },
				{ name: "California", users: 12 },
			]
		},
		{
			activeUsers: 57,
			pageViewsPerSecondLowerLimit: 0,
			pageViewsPerSecondUpperLimit: 3,

			device: [
				{ name: "PC", users: 41 },
				{ name: "Comercio", users: 5 },
				{ name: "Mobile", users: 11 }
			],
			trafficMedium: [
				{ name: "Organic", users: 39 },
				{ name: "Direct", users: 9 },
				{ name: "Paid", users: 5 },
				{ name: "Referral", users: 4 }
			],
			categories: [
				{ name: "Men Clothing", users: 8 },
				{ name: "Women Clothing", users: 9 },
				{ name: "Gadgets", users: 11 },
				{ name: "Books", users: 5 },
				{ name: "Others", users: 24 }
			],
			states: [
				{ name: "Others", users: 17 },
				{ name: "Pennsylvania", users: 4 },
				{ name: "Florida", users: 5 },
				{ name: "Texas", users: 7 },
				{ name: "New York", users: 12 },
				{ name: "California", users: 12 },
			]
		},
		{
			activeUsers: 58,
			pageViewsPerSecondLowerLimit: 0,
			pageViewsPerSecondUpperLimit: 3,

			device: [
				{ name: "PC", users: 42 },
				{ name: "Comercio", users: 5 },
				{ name: "Mobile", users: 11 }
			],
			trafficMedium: [
				{ name: "Organic", users: 40 },
				{ name: "Direct", users: 8 },
				{ name: "Paid", users: 6 },
				{ name: "Referral", users: 4 }
			],
			categories: [
				{ name: "Men Clothing", users: 9 },
				{ name: "Women Clothing", users: 10 },
				{ name: "Gadgets", users: 11 },
				{ name: "Books", users: 4 },
				{ name: "Others", users: 24 }
			],
			states: [
				{ name: "Others", users: 17 },
				{ name: "Pennsylvania", users: 4 },
				{ name: "Florida", users: 6 },
				{ name: "Texas", users: 7 },
				{ name: "New York", users: 12 },
				{ name: "California", users: 12 },
			]
		},
		{
			activeUsers: 59,
			pageViewsPerSecondLowerLimit: 0,
			pageViewsPerSecondUpperLimit: 4,

			device: [
				{ name: "PC", users: 43 },
				{ name: "Comercio", users: 4 },
				{ name: "Mobile", users: 12 }
			],
			trafficMedium: [
				{ name: "Organic", users: 41 },
				{ name: "Direct", users: 8 },
				{ name: "Paid", users: 6 },
				{ name: "Referral", users: 4 }
			],
			categories: [
				{ name: "Men Clothing", users: 9 },
				{ name: "Women Clothing", users: 10 },
				{ name: "Gadgets", users: 11 },
				{ name: "Books", users: 5 },
				{ name: "Others", users: 24 }
			],
			states: [
				{ name: "Others", users: 17 },
				{ name: "Pennsylvania", users: 4 },
				{ name: "Florida", users: 6 },
				{ name: "Texas", users: 7 },
				{ name: "New York", users: 12 },
				{ name: "California", users: 13 },
			]
		},
		{
			activeUsers: 60,
			pageViewsPerSecondLowerLimit: 0,
			pageViewsPerSecondUpperLimit: 4,

			device: [
				{ name: "PC", users: 43 },
				{ name: "Comercio", users: 5 },
				{ name: "Mobile", users: 12 }
			],
			trafficMedium: [
				{ name: "Organic", users: 40 },
				{ name: "Direct", users: 9 },
				{ name: "Paid", users: 6 },
				{ name: "Referral", users: 5 }
			],
			categories: [
				{ name: "Men Clothing", users: 9 },
				{ name: "Women Clothing", users: 10 },
				{ name: "Gadgets", users: 12 },
				{ name: "Books", users: 5 },
				{ name: "Others", users: 24 }
			],
			states: [
				{ name: "Others", users: 18 },
				{ name: "Pennsylvania", users: 3 },
				{ name: "Florida", users: 6 },
				{ name: "Texas", users: 8 },
				{ name: "New York", users: 12 },
				{ name: "California", users: 13 },
			]
		},
		{
			activeUsers: 61,
			pageViewsPerSecondLowerLimit: 0,
			pageViewsPerSecondUpperLimit: 5,

			device: [
				{ name: "PC", users: 44 },
				{ name: "Comercio", users: 5 },
				{ name: "Mobile", users: 12 }
			],
			trafficMedium: [
				{ name: "Organic", users: 42 },
				{ name: "Direct", users: 9 },
				{ name: "Paid", users: 5 },
				{ name: "Referral", users: 5 }
			],
			categories: [
				{ name: "Men Clothing", users: 9 },
				{ name: "Women Clothing", users: 10 },
				{ name: "Gadgets", users: 12 },
				{ name: "Books", users: 5 },
				{ name: "Others", users: 25 }
			],
			states: [
				{ name: "Others", users: 18 },
				{ name: "Pennsylvania", users: 3 },
				{ name: "Florida", users: 6 },
				{ name: "Texas", users: 8 },
				{ name: "New York", users: 13 },
				{ name: "California", users: 13 },
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
	
	// CanvasJS doughnut chart to show device type of active users
	var usersDeviceDoughnutChart = new CanvasJS.Chart("users-device-doughnut-chart", {
		animationDuration: 800,
		animationEnabled: true,
		backgroundColor: "transparent",
		colorSet: "customColorSet",
		theme: "theme2",
		legend: {
			fontFamily: "calibri",
			fontSize: 14,
			horizontalAlign: "left",
			verticalAlign: "center",
			itemTextFormatter: function (e) {
				return e.dataPoint.name + ": " + Math.round(e.dataPoint.y / activeUsers * 100) + "%";  
			} 
		},
		title: {
			dockInsidePlotArea: true,
			fontSize: 99,
			fontWeight: "normal",
			horizontalAlign: "center",
			verticalAlign: "center",
			text: "55"
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
				innerRadius: "80%",
				radius: "90%",
				legendMarkerType: "square",
				showInLegend: true,
				startAngle: 90,
				type: "doughnut",
				dataPoints: [
					{  y: 39, name: "PC" },
					{  y: 5, name: "Comercio" },
					{  y: 11, name: "Mobile" }
				]
			}
		]
	});
	
	// CanvasJS pie chart to traffic medium of active users
	var usersMediumPieChart = new CanvasJS.Chart("users-medium-pie-chart", {
		animationDuration: 800,
		animationEnabled: true,
		backgroundColor: "transparent",
		colorSet: "customColorSet",
		legend: {
			fontFamily: "calibri",
			fontSize: 14,
			horizontalAlign: "left",
			verticalAlign: "center",
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
					{  y: 38, name:"Organic" },
					{  y: 8, name:"Direct" },
					{  y: 5, name:"Paid" },
					{  y: 4, name:"Referral" }
				]
			}
		]
	});

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

	// CanvasJS column chart to show live page views per second
	var pageViewsPerSecondColumnChart = new CanvasJS.Chart("page-views-per-second-column-chart", {
		animationDuration: 800,
		animationEnabled: true,
		backgroundColor: "transparent",
		axisX: {						
			interval: 1,
			intervalType: "second",
			labelAutoFit: false,
			labelFontColor: "#717171",
			lineColor: "#a2a2a2",
			tickColor: "transparent",
			tickLength: 2,
			labelFormatter: function(e) {
				var diff, currentTime = (new Date()).getTime();
				diff = Math.floor((e.value.getTime() - currentTime) / 1000);
				return diff % 15 < 0 ? "" : diff + " sec";
			}
		},
		axisY: {
			gridThickness: 0,
			labelFontColor: "#717171",
			lineColor: "#a2a2a2",
			tickColor: "#a2a2a2"
		},
		toolTip: {
			cornerRadius: 0,
			fontStyle: "normal",
		},
		data: [
			{
				color: "#CD5740",
				xValueFormatString: "hh:mm:ss TT",
				type: "column",
				dataPoints : pageViewsPerSecondDataPoints
			}
		]
	});

	
	//----------------------------------------------------------------------------------//
	var allCharts = [
		usersDeviceDoughnutChart,
		usersMediumPieChart,
		usersCategoryPieChart,
		pageViewsPerSecondColumnChart,
		pageViewsPerMinuteColumnChart
		
	];
	
	// generate random number between given range
	function generateRandomNumber (minimum, maximum) {
		return Math.floor(Math.random() * (maximum - minimum + 1)) + minimum;
	}
	
	
	function updateUsersCategoryChart(dataIndex) {
		for (var i = 0; i < data[dataIndex].categories.length; i++)
			usersCategoryPieChart.options.data[0].dataPoints[i].y = data[dataIndex].categories[i].users;
		
		usersCategoryPieChart.render();
	}
	

	
	// update all charts with revelant demo data, except "Page Views Per Second" and "Page Views Per Minute" charts
	function updateCharts(dataIndex) {
		activeUsers = data[dataIndex].activeUsers;
		pageViewsPerSecondLowerLimit = data[dataIndex].pageViewsPerSecondLowerLimit;
		pageViewsPerSecondUpperLimit = data[dataIndex].pageViewsPerSecondUpperLimit;

//dash02 no tiene		updateUsersDeviceChart(dataIndex);
//dash02 no tiene		updateUsersMediumPieChart(dataIndex);
		updateUsersCategoryChart(dataIndex);
		
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
	
	// populate "Chart 0202 " charts with initial data
	function populateChart0202() {

		var time1, time2;
		
        console.log("json banwire populateChart0202 "); 
        
		for (var i = 0; i < 60; i++) {
			time1 = new Date((new Date).getTime() - ((59 - i) * 1000)); // for pageViewsPerSecond chart 
			time1.setMilliseconds(0);
			
			time2 = new Date((new Date).getTime() - ((59 - i) * 60 * 1000)); // for pageViewsPerMinute chart
			time2.setSeconds(0);
			
			pageViewsPerSecondDataPoints.push({ x: time1, y: initialDataPageViewsPerSecond[i] });
//			pageViewsPerMinuteDataPoints.push({ x: time2, y: initialDataPageViewsPerMinute[i] });
		}
	
		pageViewsPerSecondColumnChart.render();

     
	}


	// populate "Chart 0201 " charts with initial data
	function populateChart0201() {
		
        console.log("populateChart0201 json banwire populateChart0201 "); 
                
//        $.getJSON("https://canvasjs.com/services/data/datapoints.php?xstart=1&ystart=10&length=100&type=json", function(data) {  
//        $.getJSON("https://8070-dot-3809294-dot-devshell.appspot.com/v1/getjsondatabanwire", function(data) {  
        $.getJSON("https://8070-dot-3809294-dot-devshell.appspot.com/v1/getDash02Grafica01?reference=10&reference2=nono&dato01=99&dato02=88&dato03=77", function(data) {  
            
            $.each(data, function(key, value){
            
                pageViewsPerMinuteDataPoints.push({ x: parseInt(value[0]),y: parseInt(value[1]) });
            
                 console.log("populateChart0201 json banwire"+value[1]); 
                 
            });	
            pageViewsPerMinuteColumnChart.render();      
        });
     pageViewsPerMinuteColumnChart.render();   
	}


	// update "updateChart0202" chart every second 
	function updateChart0202() {
		var time1, time2;
		time1 = new Date();
		time1.setMilliseconds(0);
		
		yValuePageViewsPerSecond = generateRandomNumber(pageViewsPerSecondLowerLimit, pageViewsPerSecondUpperLimit); 
		
		pageViewsPerSecondDataPoints.push({ x: time1, y: yValuePageViewsPerSecond });
    
    if (pageViewsPerSecondDataPoints.length > 60)
    	pageViewsPerSecondDataPoints.shift();

		pageViewsPerSecondColumnChart.render();		

}

	// update "updateChart0201" chart every set of seconds 
	function updateChart0201() {
			console.log("updateChart0201: json banwire recent GET");
            
//        $.getJSON("https://8070-dot-3809294-dot-devshell.appspot.com/v1/getjsondatabanwire", function(data) {  
        $.getJSON("https://8070-dot-3809294-dot-devshell.appspot.com/v1/getDash02Grafica01?reference=10&reference2=nono&dato01=99&dato02=88&dato03=77", function(data) {              
            var linea=0;

            $.each(data, function(key, value){
                if(linea ==0){//primera linea, viene estatus Success or Error

             
                }
                if(linea ==1){//second linea, viene estatus value 1 or Error

             
                }
                if(linea ==2){//third linea, viene an array with the result data
                    var arrayResultados= value;//here the array of data structres passed
                    arrayResultados.forEach(myFunction0201) //set the values in the  graph points ,below

                    //pageViewsPerMinuteDataPoints.push({ x: parseInt(value[0]),y: parseInt(value[1]) });
             
                }
                linea=linea+1;
                 console.log("json banwire recent"+value[1]); 
                 
            });	
           if (pageViewsPerMinuteDataPoints.length > 60) 
				pageViewsPerMinuteDataPoints.shift();			
           pageViewsPerMinuteColumnChart.render();
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
			
		}else if ($(window).outerWidth() < 1920 && $(window).outerWidth() >= 1200) {
			
			chart.options.legend.fontSize = 14;
			chart.options.legend.horizontalAlign = "left";
			chart.options.legend.verticalAlign = "center";
			chart.options.legend.maxWidth = 140;
			
		} else if ($(window).outerWidth() < 1200 && $(window).outerWidth() >= 992) {
			
			chart.options.legend.fontSize = 12;
			chart.options.legend.horizontalAlign = "center";
			chart.options.legend.verticalAlign = "top";
			chart.options.legend.maxWidth = null;
			
		} else if ($(window).outerWidth() < 992) {
			
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
		for (var i = 0; i < allCharts.length; i++)
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
        window.location.href = "/v1/dash03";  

    }


	(function init() {
		customizeCharts();
		$(window).resize(customizeCharts);
		populateChart0202();
        populateChart0201();
		setInterval(updateChart0202, 1000); //every 1 second
        setInterval(updateChart0201, valueRotacionOriginal);//every 20 seconds
		setTimeout(updateChartsAtRandomIntervals, 4000); //every 4 seconds

        setInterval(checkRotationValues, valueRotacionOriginal);//every 20 seconds
        
        

        setInterval(go_next2, valuePasarNextOriginal);//every 150 seconds
        

		sidebarToggleOnClick();
	})();
	
    
});