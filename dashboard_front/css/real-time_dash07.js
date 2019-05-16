"use strict";
var valueRotacion=10000;
var valueRotacionOriginal=10000;

var valuePasarNext=150000;
var valuePasarNextOriginal=150000;

var previousGroupNumber=0;//used to keep track of what group was displayed previously

var previousGroupNumber0101=0;//used to keep track of what group was displayed previously
var previousGroupNumber0102=0;//used to keep track of what group was displayed previously
var previousGroupNumber0103=0;//used to keep track of what group was displayed previously


function show_selected2() {
    alert("aj");
    var selector = document.getElementById('id_of_select');
    valueRotacion= selector[selector.selectedIndex].value;

    document.getElementById('display').innerHTML = valueRotacion;
}


$(function () {


    var activities = document.getElementById("id_of_selectRotacion");

    activities.addEventListener("change", function() {
        alert("aj id_of_selectRotacion");

        var selector = document.getElementById('id_of_selectRotacion');
        valueRotacion= selector[selector.selectedIndex].value;
        alert("aj id_of_selectRotacion:"+valueRotacion);

    });

    var activities02 = document.getElementById("id_of_select");

    activities02.addEventListener("change", function() {
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
	
	
	// data for demo only
	var data = [
		{
			activeUsers: 99,

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
					{  y: 39, name: "PCs" },
					{  y: 5, name: "ComercioS" },
					{  y: 11, name: "MobileS" },
                    {  y: 6, name: "OtrosS" }
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
			

	
	//----------------------------------------------------------------------------------//
	var allCharts = [
		usersDeviceDoughnutChart,
		usersMediumPieChart,
		usersCategoryPieChart
	];
	
	// generate random number between given range
	function generateRandomNumber (minimum, maximum) {
		return Math.floor(Math.random() * (maximum - minimum + 1)) + minimum;
	}

/*	
	function updateUsersDeviceChart(dataIndex) {
		usersDeviceDoughnutChart.options.title.text = activeUsers.toString();
		
		for (var i = 0; i < data[dataIndex].device.length; i++)
			usersDeviceDoughnutChart.options.data[0].dataPoints[i].y = data[dataIndex].device[i].users;
		
		usersDeviceDoughnutChart.render();
	}
	
	function updateUsersMediumPieChart(dataIndex) {
		for (var i = 0; i < data[dataIndex].trafficMedium.length; i++)
			usersMediumPieChart.options.data[0].dataPoints[i].y = data[dataIndex].trafficMedium[i].users;
		
		usersMediumPieChart.render();
	}
	
	function updateUsersCategoryChart(dataIndex) {
		for (var i = 0; i < data[dataIndex].categories.length; i++)
			usersCategoryPieChart.options.data[0].dataPoints[i].y = data[dataIndex].categories[i].users;
		
		usersCategoryPieChart.render();
	}
*/

	/********************************************
    
    functions that access the Webservices 
    obtain the data
    
    ***************************************+ */
//Device


		

    function updateUsersDeviceChartWithDB()
   {
                 //use a parameter PageTitle {{.PageTitle}} set in net_v1.go with the server url

        $.getJSON("{{.PageTitle}}/v1/getDash02Grafica02?reference="+previousGroupNumber0101+"&reference2=11&dato01=99&dato02=88&dato03=77", function(data) {                  
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
                   previousGroupNumber0101 =grupoNumberResultado;

                }                
                if(linea ==3 && exito=="1"){//foruth linea, viene an array with the result data
                    var arrayResultados= value;//here the array of data structres passed
                    arrayResultados.forEach(myFunction0101) //set the values in the  graph points ,below
             
                }
                linea=linea+1;
                 console.log("json banwire recent"); 
                 
            });	
            if(exito=="1"){ //only update data if new data group was found, more recent than the previuous group number

                 usersDeviceDoughnutChart.render();

            }
        });
      
        usersDeviceDoughnutChart.render();

}


function myFunction0101(item, index) {
   //  demoP.innerHTML = demoP.innerHTML + "index[" + index + "]: " + item + "<br>"; 

   usersDeviceDoughnutChart.options.data[0].dataPoints[index].y = parseInt(item["data_valuea"]);
   usersDeviceDoughnutChart.options.data[0].dataPoints[index].name = item["data_name"];

/*
    Object.keys(item).forEach(function(key) {
    console.log('Key : ' + key + ', Value : ' + item[key])
    })
*/

}

//Medium


    function updateUsersMediumPieChartWithDB()
   {
//use a parameter PageTitle {{.PageTitle}} set in net_v1.go with the server url
        $.getJSON("{{.PageTitle}}/v1/getDash02Grafica02?reference="+previousGroupNumber0102+"&reference2=12&dato01=99&dato02=88&dato03=77", function(data) {                  
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
                   previousGroupNumber0102 =grupoNumberResultado;

                }                
                if(linea ==3 && exito=="1"){//foruth linea, viene an array with the result data
                    var arrayResultados= value;//here the array of data structres passed
                    arrayResultados.forEach(myFunction0102) //set the values in the  graph points ,below
             
                }
                linea=linea+1;
                 console.log("json banwire recent"); 
                 
            });	
            if(exito=="1"){ //only update data if new data group was found, more recent than the previuous group number

                 usersMediumPieChart.render();

            }
        });
      
        usersMediumPieChart.render();

}


function myFunction0102(item, index) {
   //  demoP.innerHTML = demoP.innerHTML + "index[" + index + "]: " + item + "<br>"; 

   usersMediumPieChart.options.data[0].dataPoints[index].y = parseInt(item["data_valuea"]);
   usersMediumPieChart.options.data[0].dataPoints[index].name = item["data_name"];

/*
    Object.keys(item).forEach(function(key) {
    console.log('Key : ' + key + ', Value : ' + item[key])
    })
*/

}


// Category
    function updateUsersCategoryChartWithDB()
   {
//use a parameter PageTitle {{.PageTitle}} set in net_v1.go with the server url
        $.getJSON("{{.PageTitle}}/v1/getDash02Grafica02?reference="+previousGroupNumber0103+"&reference2=13&dato01=99&dato02=88&dato03=77", function(data) {                  
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
                   previousGroupNumber0103 =grupoNumberResultado;

                }                
                if(linea ==3 && exito=="1"){//foruth linea, viene an array with the result data
                    var arrayResultados= value;//here the array of data structres passed
                    arrayResultados.forEach(myFunction0103) //set the values in the  graph points ,below

             
                }
                linea=linea+1;
                 console.log("json banwire recent"); 
                 
            });	
            if(exito=="1"){ //only update data if new data group was found, more recent than the previuous group number

                 usersCategoryPieChart.render();

            }
        });
      
        usersCategoryPieChart.render();

}


function myFunction0103(item, index) {
   //  demoP.innerHTML = demoP.innerHTML + "index[" + index + "]: " + item + "<br>"; 

   usersCategoryPieChart.options.data[0].dataPoints[index].y = parseInt(item["data_valuea"]);
   usersCategoryPieChart.options.data[0].dataPoints[index].name = item["data_name"];

/*
    Object.keys(item).forEach(function(key) {
    console.log('Key : ' + key + ', Value : ' + item[key])
    })
*/

}


	
	// update all charts with revelant data
	function updateCharts(dataIndex) {
		activeUsers = data[dataIndex].activeUsers;

		updateUsersDeviceChartWithDB();
		updateUsersMediumPieChartWithDB();
		// usar datos demo     updateUsersCategoryChart(dataIndex);
        //usar datos de la db
        updateUsersCategoryChartWithDB();
		
	}


	function updateChartsAtRandomIntervals() {
		var dataIndex = generateRandomNumber(0, data.length - 1);
				
		updateCharts(dataIndex);

    // set the next time the functions will be executed, by the value in the valueRotacion		
		if (timeoutIdUpdateCharts)
			clearTimeout(timeoutIdUpdateCharts);
		
		timeoutIdUpdateCharts = setTimeout(function () {
			updateChartsAtRandomIntervals();
		}, valueRotacion);
	}


	
/*
functions for graphs with XY 


	// populate "Chart 0201 " charts with initial data
	function populateChart0201() {
		
        console.log("populateChart0201 json banwire populateChart0201 "); 
                
//        $.getJSON("https://canvasjs.com/services/data/datapoints.php?xstart=1&ystart=10&length=100&type=json", function(data) {  
//        $.getJSON("https://8070-dot-3809294-dot-devshell.appspot.com/v1/getjsondatabanwire", function(data) {  
        $.getJSON("https://8070-dot-3809294-dot-devshell.appspot.com/v1/getDash02Grafica01?reference=0&reference2=21&dato01=99&dato02=88&dato03=77", function(data) {  
            
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

//        $.getJSON("https://8070-dot-3809294-dot-devshell.appspot.com/v1/getDash02Grafica01?reference=10&reference2=21&dato01=99&dato02=88&dato03=77", function(data) {              
        $.getJSON("https://8070-dot-3809294-dot-devshell.appspot.com/v1/getDash02Grafica01?reference="+previousGroupNumber+"&reference2=21&dato01=99&dato02=88&dato03=77", function(data) {                  
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

*/
	
	// chart properties cutomized further based on screen width	
	function chartPropertiesCustomization(chart) {
		if ($(window).outerWidth() >= 1920) {			
			
			chart.options.legend.fontSize = 14;
			chart.options.legend.horizontalAlign = "left";
			chart.options.legend.verticalAlign = "center";
			chart.options.legend.maxWidth = null;
			
		}else if (1920 >= $(window).outerWidth()  && $(window).outerWidth() >= 1200) {
			
			chart.options.legend.fontSize = 14;
			chart.options.legend.horizontalAlign = "left";
			chart.options.legend.verticalAlign = "center";
			chart.options.legend.maxWidth = 140;
			
		} else if ( 1200 >= $(window).outerWidth()   && $(window).outerWidth() >= 992) {
			
			chart.options.legend.fontSize = 12;
			chart.options.legend.horizontalAlign = "center";
			chart.options.legend.verticalAlign = "top";
			chart.options.legend.maxWidth = null;
			
		} else if (992 >= $(window).outerWidth() ) {
			
			chart.options.legend.fontSize = 14;
			chart.options.legend.horizontalAlign = "center";
			chart.options.legend.verticalAlign = "bottom";
			chart.options.legend.maxWidth = null;
			
		}
		
		chart.render();
	}
	
	function customizeCharts() {
		chartPropertiesCustomization(usersDeviceDoughnutChart);
		chartPropertiesCustomization(usersMediumPieChart);
		chartPropertiesCustomization(usersCategoryPieChart);
	}
	
	function renderAllCharts() {
		for (var i = 0;  allCharts.length>= i; i++)
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
//            setInterval(updateChart0201, valueRotacion);//every 20 seconds
            setInterval(updateChartsAtRandomIntervals, valueRotacion);//every number of  seconds decided by the user in the screen
            
            valueRotacionOriginal=valueRotacion;
        }
         
        if(valuePasarNextOriginal==valuePasarNext){
             //continue with the same value
        }else{
            setTimeout(go_next2, valuePasarNext);//every 20 seconds
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

		setTimeout(updateChartsAtRandomIntervals, valueRotacionOriginal); //every 4 seconds

        setInterval(checkRotationValues, valueRotacionOriginal);//every 20 seconds

        setInterval(go_next2, valuePasarNextOriginal);//every 150 seconds


//		setTimeout(updateChartsAtRandomIntervals, 4000);
		sidebarToggleOnClick();
	})();


});

