"use strict";
var valueRotacion=20000;
var valueRotacionOriginal=60000;

var valuePasarNext=450000;
var valuePasarNextOriginal=450000;

var previousGroupNumber=0;//used to keep track of what group was displayed previously

var previousGroupNumber0101=0;//used to keep track of what group was displayed previously
var previousGroupNumber0102=0;//used to keep track of what group was displayed previously
var previousGroupNumber0103=0;//used to keep track of what group was displayed previously
var previousGroupNumber0104=0;//used to keep track of what group was displayed previously



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
			numberOfSeconds = 1,
			updateChartsInterval,
			updateChartsIntervalLowerLimit = 4000, // milliseconds
			updateChartsIntervalUpperLimit = 6000, // milliseconds
			timeoutIdUpdateCharts;
	
	
	
	// data for demo only
	var data = [
		{
			activeUsers: 99,

			trafficMedium: [
				{ name: "Organic", users: 38 },
				{ name: "Direct", users: 8 },
				{ name: "Paid", users: 5 },
				{ name: "Referral", users: 4 }
			],
			states: [
				{ name: "Others", users: 16 },
				{ name: "Pennsylvania", users: 4 },
				{ name: "Florida", users: 5 },
				{ name: "Texas", users: 7 },
				{ name: "New York", users: 11 },
                { name: "New York", users: 11 },
                { name: "New York", users: 11 },
                { name: "New York", users: 11 },
                { name: "New York", users: 11 },
				{ name: "California", users: 12 },
			],
			states24: [
				{ name: "Others", users: 16 },
				{ name: "Pennsylvania", users: 4 },
				{ name: "Florida", users: 5 },
				{ name: "Texas", users: 7 },
				{ name: "New York", users: 11 },
                { name: "New York", users: 11 },
                { name: "New York", users: 11 },
                { name: "New York", users: 11 },
                { name: "New York", users: 11 },
				{ name: "California", users: 12 },
			],
			states48: [
				{ name: "Others", users: 16 },
				{ name: "Pennsylvania", users: 4 },
				{ name: "Florida", users: 5 },
				{ name: "Texas", users: 7 },
				{ name: "New York", users: 11 },
                { name: "New York", users: 11 },
                { name: "New York", users: 11 },
                { name: "New York", users: 11 },
                { name: "New York", users: 11 },
				{ name: "California", users: 12 },
			]                        
		},
		{
			activeUsers: 56,

			trafficMedium: [
				{ name: "Organic", users: 39 },
				{ name: "Direct", users: 8 },
				{ name: "Paid", users: 5 },
				{ name: "Referral", users: 4 }
			],
			states: [
				{ name: "Others", users: 17 },
				{ name: "Pennsylvania", users: 4 },
				{ name: "Florida", users: 5 },
				{ name: "Texas", users: 7 },
				{ name: "New York", users: 11 },
                { name: "New York", users: 11 },
                { name: "New York", users: 11 },
                { name: "New York", users: 11 },
				{ name: "California", users: 12 },
			]
		},
		{
			activeUsers: 57,

			trafficMedium: [
				{ name: "Organic", users: 39 },
				{ name: "Direct", users: 9 },
				{ name: "Paid", users: 5 },
				{ name: "Referral", users: 4 }
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

			trafficMedium: [
				{ name: "Organic", users: 40 },
				{ name: "Direct", users: 8 },
				{ name: "Paid", users: 6 },
				{ name: "Referral", users: 4 }
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
			trafficMedium: [
				{ name: "Organic", users: 41 },
				{ name: "Direct", users: 8 },
				{ name: "Paid", users: 6 },
				{ name: "Referral", users: 4 }
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

			trafficMedium: [
				{ name: "Organic", users: 40 },
				{ name: "Direct", users: 9 },
				{ name: "Paid", users: 6 },
				{ name: "Referral", users: 5 }
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

			trafficMedium: [
				{ name: "Organic", users: 42 },
				{ name: "Direct", users: 9 },
				{ name: "Paid", users: 5 },
				{ name: "Referral", users: 5 }
			],
			states: [
				{ name: "Others", users: 18 },
				{ name: "Pennsylvania", users: 3 },
				{ name: "Florida", users: 6 },
				{ name: "Texas", users: 8 },
				{ name: "New York", users: 13 },
				{ name: "California", users: 13 },
			],
           	states24: [
				{ name: "Others", users: 18 },
				{ name: "Pennsylvania", users: 3 },
				{ name: "Florida", users: 6 },
				{ name: "Texas", users: 8 },
				{ name: "New York", users: 13 },
				{ name: "California", users: 13 },
			]
,
			states48: [
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
//				return e.entries[0].dataPoint.name + ": " + Math.round(e.entries[0].dataPoint.y / activeUsers * 100) + "% (" + e.entries[0].dataPoint.y  + ")";  
                return e.entries[0].dataPoint.name + ": " + e.entries[0].dataPoint.y + " (" + e.entries[0].dataPoint.y  + ")";  
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

			

	// CanvasJS bar chart to show active users by state
	var usersStateBarChart = new CanvasJS.Chart("users-state-bar-chart", {
		animationDuration: 800,
		animationEnabled: true,
		backgroundColor: "transparent",
		colorSet: "customColorSet",
		axisX: {
			labelFontColor: "#717171",
			labelFontSize: 10,
			lineThickness: 0,
			tickThickness: 0
		},
		axisY: {
			gridThickness: 0,
			lineThickness: 0,
			tickThickness: 0,
			valueFormatString: " "
		},
        
		toolTip: {
			cornerRadius: 0,
			fontStyle: "normal",
			contentFormatter: function (e) {
				//return e.entries[0].dataPoint.label + ": " + Math.round(e.entries[0].dataPoint.y / activeUsers * 100) + "% (" + e.entries[0].dataPoint.y  + ")";
                return  e.entries[0].dataPoint.label.substring(0, 8) + ": " + e.entries[0].dataPoint.y  + " (" + e.entries[0].dataPoint.y  + ")";
			} 
		},
		data: [
			{
				indexLabelFontColor: "#717171",
				indexLabelFontFamily: "calibri",
				indexLabelFontSize: 18,
				indexLabelPlacement: "outside",
				indexLabelFormatter: function (e) {
					//return Math.round(e.dataPoint.y / activeUsers * 100) + "%";  
                    return "$" + e.dataPoint.y  ;  
				},
				type: "bar",
				dataPoints: [
					{ y: 16,  label: "Others" },
					{ y: 4, label: "Pennsylvania" },
					{ y: 5,  label: "Florida" },
					{ y: 7, label: "Texas" },
					{ y: 11, label: "New York" },
                    { y: 11, label: "New York" },
                    { y: 11, label: "New York" },
					{ y: 12, label: "California" }
				]
			}
		]
	});

	usersStateBarChart.render();
	
    //24 hrs

    	// CanvasJS bar chart to show active users by state
	var u24StateBarChart = new CanvasJS.Chart("u24-state-bar-chart", {
		animationDuration: 800,
		animationEnabled: true,
		backgroundColor: "transparent",
		colorSet: "customColorSet",
		axisX: {
			labelFontColor: "#717171",
			labelFontSize: 10,
			lineThickness: 0,
			tickThickness: 0
		},
		axisY: {
			gridThickness: 0,
			lineThickness: 0,
			tickThickness: 0,
			valueFormatString: " "
		},
        
		toolTip: {
			cornerRadius: 0,
			fontStyle: "normal",
			contentFormatter: function (e) {
				//return e.entries[0].dataPoint.label + ": " + Math.round(e.entries[0].dataPoint.y / activeUsers * 100) + "% (" + e.entries[0].dataPoint.y  + ")";
                return  e.entries[0].dataPoint.label.substring(0, 8) + ": " + e.entries[0].dataPoint.y  + " (" + e.entries[0].dataPoint.y  + ")";
			} 
		},
		data: [
			{
				indexLabelFontColor: "#717171",
				indexLabelFontFamily: "calibri",
				indexLabelFontSize: 18,
				indexLabelPlacement: "outside",
				indexLabelFormatter: function (e) {
					//return Math.round(e.dataPoint.y / activeUsers * 100) + "%";  
                    return "$"+ e.dataPoint.y  ;  
				},
				type: "bar",
				dataPoints: [
					{ y: 16,  label: "Others" },
					{ y: 4, label: "Penn24" },
					{ y: 5,  label: "Flori24" },
					{ y: 7, label: "Texas24" },
					{ y: 11, label: "New York" },
                    { y: 11, label: "New York" },
                    { y: 11, label: "New York" },
					{ y: 12, label: "Cal24" }
				]
			}
		]
	});

	u24StateBarChart.render();
	


    //48hrs
    	// CanvasJS bar chart to show active users by state
	var u48StateBarChart = new CanvasJS.Chart("u48-state-bar-chart", {
		animationDuration: 800,
		animationEnabled: true,
		backgroundColor: "transparent",
		colorSet: "customColorSet",
		axisX: {
			labelFontColor: "#717171",
			labelFontSize: 10,
			lineThickness: 0,
			tickThickness: 0
		},
		axisY: {
			gridThickness: 0,
			lineThickness: 0,
			tickThickness: 0,
			valueFormatString: " "
		},
        
		toolTip: {
			cornerRadius: 0,
			fontStyle: "normal",
			contentFormatter: function (e) {
				//return e.entries[0].dataPoint.label + ": " + Math.round(e.entries[0].dataPoint.y / activeUsers * 100) + "% (" + e.entries[0].dataPoint.y  + ")";
                return  e.entries[0].dataPoint.label.substring(0, 8) + ": " + e.entries[0].dataPoint.y  + " (" + e.entries[0].dataPoint.y  + ")";
			} 
		},
		data: [
			{
				indexLabelFontColor: "#717171",
				indexLabelFontFamily: "calibri",
				indexLabelFontSize: 18,
				indexLabelPlacement: "outside",
				indexLabelFormatter: function (e) {
					//return Math.round(e.dataPoint.y / activeUsers * 100) + "%";  
                    return "$" + e.dataPoint.y  ;  
				},
				type: "bar",
				dataPoints: [
					{ y: 16,  label: "Others" },
					{ y: 4, label: "Penn48" },
					{ y: 5,  label: "Flor48" },
					{ y: 7, label: "Texas48" },
					{ y: 11, label: "New York" },
                    { y: 11, label: "New York" },
                    { y: 11, label: "New York" },
					{ y: 12, label: "Cal48" }
				]
			}
		]
	});

	u48StateBarChart.render();
	
		var chartHoy = new CanvasJS.Chart("chartContainer", {
		animationEnabled: true,
		title:{
			text: "Aceptadas (Hoy)"
		},
		axisY: {
			title: "Monto en $"
		},
		legend: {
			cursor:"pointer"
//			,
//			itemclick : toggleDataSeries
		},
		toolTip: {
			shared: true
//			,
//			content: toolTipFormatter
		},
		data: [{
			type: "bar",
			showInLegend: true,
			name: "Now",
			color: "gold",
			dataPoints: [
				{ y: 243, label: "Italy" },
				{ y: 236, label: "China" },
				{ y: 243, label: "France" },
				{ y: 273, label: "Great Britain" },
				{ y: 269, label: "Germany" },
				{ y: 196, label: "Russia" },
				{ y: 188, label: "Russia" },
				{ y: 1118, label: "USA" }
			]
		},
		{
			type: "bar",
			showInLegend: true,
			name: "-1Hr",
			color: "silver",
			dataPoints: [
				{ y: 212, label: "Italy" },
				{ y: 186, label: "China" },
				{ y: 272, label: "France" },
				{ y: 299, label: "Great Britain" },
				{ y: 270, label: "Germany" },
				{ y: 165, label: "Russia" },
				{ y: 188, label: "Russia" },
				{ y: 896, label: "USA" }
			]
		},
		{
			type: "bar",
			showInLegend: true,
			name: "-2Hrs",
			color: "#A57164",
			dataPoints: [
				{ y: 236, label: "Italy" },
				{ y: 172, label: "China" },
				{ y: 309, label: "France" },
				{ y: 302, label: "Great Britain" },
				{ y: 285, label: "Germany" },
				{ y: 188, label: "Russia" },
				{ y: 188, label: "Russia" },
				{ y: 788, label: "USA" }
			]
		}]
	});
	chartHoy.render();


			var chart7dias = new CanvasJS.Chart("chartContainer7dias", {
		animationEnabled: true,
		title:{
			text: "Hace 7 dias"
		},
		axisY: {
			title: "Monto en $"
		},
		legend: {
			cursor:"pointer"
//			,
//			itemclick : toggleDataSeries
		},
		toolTip: {
			shared: true
//			,
//			content: toolTipFormatter
		},
		data: [{
			type: "bar",
			showInLegend: true,
			name: "Now",
			color: "gold",
			dataPoints: [
				{ y: 243, label: "Italy" },
				{ y: 236, label: "China" },
				{ y: 243, label: "France" },
				{ y: 273, label: "Great Britain" },
				{ y: 269, label: "Germany" },
				{ y: 196, label: "Russia" },
				{ y: 188, label: "Russia" },
				{ y: 1118, label: "USA" }
			]
		},
		{
			type: "bar",
			showInLegend: true,
			name: "-1Hr",
			color: "silver",
			dataPoints: [
				{ y: 212, label: "Italy" },
				{ y: 186, label: "China" },
				{ y: 272, label: "France" },
				{ y: 299, label: "Great Britain" },
				{ y: 270, label: "Germany" },
				{ y: 165, label: "Russia" },
				{ y: 188, label: "Russia" },
				{ y: 896, label: "USA" }
			]
		},
		{
			type: "bar",
			showInLegend: true,
			name: "-2Hrs",
			color: "#A57164",
			dataPoints: [
				{ y: 236, label: "Italy" },
				{ y: 172, label: "China" },
				{ y: 309, label: "France" },
				{ y: 302, label: "Great Britain" },
				{ y: 285, label: "Germany" },
				{ y: 188, label: "Russia" },
				{ y: 188, label: "Russia" },
				{ y: 788, label: "USA" }
			]
		}]
	});
	chart7dias.render();

			var chart30dias = new CanvasJS.Chart("chartContainer30dias", {
		animationEnabled: true,
		title:{
			text: "Hace 30 dias"
		},
		axisY: {
			title: "Monto en $"
		},
		legend: {
			cursor:"pointer"
//			,
//			itemclick : toggleDataSeries
		},
		toolTip: {
			shared: true
//			,
//			content: toolTipFormatter
		},
		data: [{
			type: "bar",
			showInLegend: true,
			name: "Now",
			color: "gold",
			dataPoints: [
				{ y: 243, label: "Italy" },
				{ y: 236, label: "China" },
				{ y: 243, label: "France" },
				{ y: 273, label: "Great Britain" },
				{ y: 269, label: "Germany" },
				{ y: 196, label: "Russia" },
				{ y: 188, label: "Russia" },
				{ y: 1118, label: "USA" }
			]
		},
		{
			type: "bar",
			showInLegend: true,
			name: "-1Hr",
			color: "silver",
			dataPoints: [
				{ y: 212, label: "Italy" },
				{ y: 186, label: "China" },
				{ y: 272, label: "France" },
				{ y: 299, label: "Great Britain" },
				{ y: 270, label: "Germany" },
				{ y: 165, label: "Russia" },
				{ y: 188, label: "Russia" },
				{ y: 896, label: "USA" }
			]
		},
		{
			type: "bar",
			showInLegend: true,
			name: "-2Hrs",
			color: "#A57164",
			dataPoints: [
				{ y: 236, label: "Italy" },
				{ y: 172, label: "China" },
				{ y: 309, label: "France" },
				{ y: 302, label: "Great Britain" },
				{ y: 285, label: "Germany" },
				{ y: 188, label: "Russia" },
				{ y: 188, label: "Russia" },
				{ y: 788, label: "USA" }
			]
		}]
	});
	chart30dias.render();



	//----------------------------------------------------------------------------------//
	var allCharts = [
		usersMediumPieChart,
		usersStateBarChart,
        u24StateBarChart,
        u48StateBarChart,
		chartHoy,
		chart7dias,
		chart30dias
	];
	
	// generate random number between given range
	function generateRandomNumber (minimum, maximum) {
		return Math.floor(Math.random() * (maximum - minimum + 1)) + minimum;
	}
	
	
//Medium


    function updateUsersMediumPieChartWithDB()
   {
             //use a parameter PageTitle {{.PageTitle}} set in net_v1.go with the server url

        $.getJSON("{{.PageTitle}}/v1/getDash02Grafica02?reference="+previousGroupNumber0101+"&reference2=221&dato01=00&dato02=88&dato03=77", function(data) {
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

                 usersMediumPieChart.render();

            }
        });
      
        usersMediumPieChart.render();

}


function myFunction0101(item, index) {
   //  demoP.innerHTML = demoP.innerHTML + "index[" + index + "]: " + item + "<br>"; 

   usersMediumPieChart.options.data[0].dataPoints[index].y = parseInt(item["data_valuea"]);
   usersMediumPieChart.options.data[0].dataPoints[index].label = item["data_valuea"];
   usersMediumPieChart.options.data[0].dataPoints[index].name = item["data_name"];

/*
    Object.keys(item).forEach(function(key) {
    console.log('Key : ' + key + ', Value : ' + item[key])
    })
*/

}


////pagados de la ultima hora
///start
    function updateUsersStateChartWithDB()
   {
             //use a parameter PageTitle {{.PageTitle}} set in net_v1.go with the server url
        $.getJSON("{{.PageTitle}}/v1/getDash02Grafica02?reference="+previousGroupNumber0102+"&reference2=221&dato01=00&dato02=88&dato03=77", function(data) {
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

                 usersStateBarChart.render();

            }
        });
      
        usersStateBarChart.render();

}


function myFunction0102(item, index) {
   //  demoP.innerHTML = demoP.innerHTML + "index[" + index + "]: " + item + "<br>"; 

   usersStateBarChart.options.data[0].dataPoints[index].y = parseInt(item["data_valuea"]);
   usersStateBarChart.options.data[0].dataPoints[index].label = item["data_name"];
   usersStateBarChart.options.data[0].dataPoints[index].name = item["data_name"];

/*
    Object.keys(item).forEach(function(key) {
    console.log('Key : ' + key + ', Value : ' + item[key])
    })
*/

}
///end
////pagados ultima hora



////pagados de la ultimas 24hrs
///start
    function updateU24StateChartWithDB()
   {
             //use a parameter PageTitle {{.PageTitle}} set in net_v1.go with the server url
        $.getJSON("{{.PageTitle}}/v1/getDash02Grafica02?reference="+previousGroupNumber0103+"&reference2=221&dato01=01&dato02=88&dato03=77", function(data) {
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

                 u24StateBarChart.render();

            }
        });
      
        u24StateBarChart.render();

}


function myFunction0103(item, index) {
   //  demoP.innerHTML = demoP.innerHTML + "index[" + index + "]: " + item + "<br>"; 

   u24StateBarChart.options.data[0].dataPoints[index].y = parseInt(item["data_valuea"]);
   u24StateBarChart.options.data[0].dataPoints[index].label = item["data_name"];
   u24StateBarChart.options.data[0].dataPoints[index].name = item["data_name"];

/*
    Object.keys(item).forEach(function(key) {
    console.log('Key : ' + key + ', Value : ' + item[key])
    })
*/

}
///end
////pagados ultimas 24hrs

	////pagados ultimas 48hrs
///start
    function updateU48StateChartWithDB()
   {
             //use a parameter PageTitle {{.PageTitle}} set in net_v1.go with the server url
        $.getJSON("{{.PageTitle}}/v1/getDash02Grafica02?reference="+previousGroupNumber0104+"&reference2=221&dato01=02&dato02=88&dato03=77", function(data) {
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
                   previousGroupNumber0104 =grupoNumberResultado;

                }                
                if(linea ==3 && exito=="1"){//foruth linea, viene an array with the result data
                    var arrayResultados= value;//here the array of data structres passed
                    arrayResultados.forEach(myFunction0104) //set the values in the  graph points ,below
             
                }
                linea=linea+1;
                 console.log("json banwire recent"); 
                 
            });	
            if(exito=="1"){ //only update data if new data group was found, more recent than the previuous group number

                 u48StateBarChart.render();

            }
        });
      
        u48StateBarChart.render();

}


function myFunction0104(item, index) {
   //  demoP.innerHTML = demoP.innerHTML + "index[" + index + "]: " + item + "<br>"; 

   u48StateBarChart.options.data[0].dataPoints[index].y = parseInt(item["data_valuea"]);
   u48StateBarChart.options.data[0].dataPoints[index].label = item["data_name"];
   u48StateBarChart.options.data[0].dataPoints[index].name = item["data_name"];

/*
    Object.keys(item).forEach(function(key) {
    console.log('Key : ' + key + ', Value : ' + item[key])
    })
*/

}
///end
////pagados ultimas 48hrs



////pagados Hoy -1hr -2hr
///start
    function updateHOYChartWith(paramDataType)
   {
             //use a parameter PageTitle {{.PageTitle}} set in net_v1.go with the server url
        $.getJSON("{{.PageTitle}}/v1/getDash02Grafica02?reference="+previousGroupNumber0104+"&reference2=221&dato01="+paramDataType +"&dato02=88&dato03=77", function(data) {
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
                   previousGroupNumber0104 =grupoNumberResultado;

                }                
                if(linea ==3 && exito=="1"){//foruth linea, viene an array with the result data
                    var arrayResultados= value;//here the array of data structres passed
			if(paramDataType=="00" ){ //hoy now
				arrayResultados.forEach(updateHOYChartWithNow) //set the values in the  graph points ,below
			}
			if(paramDataType=="01" ){ //hoy -1hr
				arrayResultados.forEach(updateHOYChartWithMenos1) //set the values in the  graph points ,below
			}                    
			if(paramDataType=="02" ){ //hoy -2hrs
				arrayResultados.forEach(updateHOYChartWithMenos2) //set the values in the  graph points ,below
			}	
       
                }
                linea=linea+1;
                 console.log("json banwire recent"); 
                 
            });	
            if(exito=="1"){ //only update data if new data group was found, more recent than the previuous group number

                 chartHoy.render();

            }
        });
      
        chartHoy.render();

}

//func for now
function updateHOYChartWithNow(item, index) {
  
   chartHoy.options.data[0].dataPoints[index].y = parseInt(item["data_valuea"]);
   chartHoy.options.data[0].dataPoints[index].label = item["data_name"];
   chartHoy.options.data[0].dataPoints[index].name = item["data_name"];



}
function updateHOYChartWithMenos1(item, index) {
   

   chartHoy.options.data[1].dataPoints[index].y = parseInt(item["data_valuea"]);
   chartHoy.options.data[1].dataPoints[index].label = item["data_name"];
   chartHoy.options.data[1].dataPoints[index].name = item["data_name"];



}
function updateHOYChartWithMenos2(item, index) {
   

   chartHoy.options.data[2].dataPoints[index].y = parseInt(item["data_valuea"]);
   chartHoy.options.data[2].dataPoints[index].label = item["data_name"];
   chartHoy.options.data[2].dataPoints[index].name = item["data_name"];



}
///end
////pagados pagados Hace 7 dias    now -1hr -2hr

	
////pagados hace 7 dias -1hr -2hr
///start
    function updateChartWith7dias(paramDataType)
   {
             //use a parameter PageTitle {{.PageTitle}} set in net_v1.go with the server url
        $.getJSON("{{.PageTitle}}/v1/getDash02Grafica02?reference="+previousGroupNumber0104+"&reference2=221&dato01="+paramDataType +"&dato02=88&dato03=77", function(data) {
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
                   previousGroupNumber0104 =grupoNumberResultado;

                }                
                if(linea ==3 && exito=="1"){//foruth linea, viene an array with the result data
                    var arrayResultados= value;//here the array of data structres passed
			if(paramDataType=="70" ){ //hace 7dias now
				arrayResultados.forEach(updateHOYChartWith7dias) //set the values in the  graph points ,below
			}
			if(paramDataType=="71" ){ //hace 7dias -1hr
				arrayResultados.forEach(updateHOYChartWith7diasMenos1) //set the values in the  graph points ,below
			}                    
			if(paramDataType=="72" ){ //hace 7dias -2hrs
				arrayResultados.forEach(updateHOYChartWith7diasMenos2) //set the values in the  graph points ,below
			}					
                }
                linea=linea+1;
                 console.log("json banwire recent"); 
                 
            });	
            if(exito=="1"){ //only update data if new data group was found, more recent than the previuous group number

                 chart7dias.render();

            }
        });
      
        chart7dias.render();

}

//func for now
function updateHOYChartWith7dias(item, index) {
  
   chart7dias.options.data[0].dataPoints[index].y = parseInt(item["data_valuea"]);
   chart7dias.options.data[0].dataPoints[index].label = item["data_name"];
   chart7dias.options.data[0].dataPoints[index].name = item["data_name"];



}
function updateHOYChartWith7diasMenos1(item, index) {
   

   chart7dias.options.data[1].dataPoints[index].y = parseInt(item["data_valuea"]);
   chart7dias.options.data[1].dataPoints[index].label = item["data_name"];
   chart7dias.options.data[1].dataPoints[index].name = item["data_name"];



}
function updateHOYChartWith7diasMenos2(item, index) {
   

   chart7dias.options.data[2].dataPoints[index].y = parseInt(item["data_valuea"]);
   chart7dias.options.data[2].dataPoints[index].label = item["data_name"];
   chart7dias.options.data[2].dataPoints[index].name = item["data_name"];



}
///end
////pagados pagados 7dias -1hr -2hr
	

////pagados hace 30 dias -1hr -2hr
///start
    function updateChartWith30dias(paramDataType)
   {
             //use a parameter PageTitle {{.PageTitle}} set in net_v1.go with the server url
        $.getJSON("{{.PageTitle}}/v1/getDash02Grafica02?reference="+previousGroupNumber0104+"&reference2=221&dato01="+paramDataType +"&dato02=88&dato03=77", function(data) {
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
                   previousGroupNumber0104 =grupoNumberResultado;

                }                
                if(linea ==3 && exito=="1"){//foruth linea, viene an array with the result data
                    var arrayResultados= value;//here the array of data structres passed
			if(paramDataType=="30" ){ //hace 30dias now
				arrayResultados.forEach(updateHOYChartWith30dias) //set the values in the  graph points ,below
			}
			if(paramDataType=="31" ){ //hace 30dias -1hr
				arrayResultados.forEach(updateHOYChartWith30diasMenos1) //set the values in the  graph points ,below
			}                    
			if(paramDataType=="32" ){ //hace 30dias -2hrs
				arrayResultados.forEach(updateHOYChartWith30diasMenos2) //set the values in the  graph points ,below
			}             
                }
                linea=linea+1;
                 console.log("json banwire recent"); 
                 
            });	
            if(exito=="1"){ //only update data if new data group was found, more recent than the previuous group number

                 chart30dias.render();

            }
        });
      
        chart30dias.render();

}

//func for now
function updateHOYChartWith30dias(item, index) {
  
   chart30dias.options.data[0].dataPoints[index].y = parseInt(item["data_valuea"]);
   chart30dias.options.data[0].dataPoints[index].label = item["data_name"];
   chart30dias.options.data[0].dataPoints[index].name = item["data_name"];



}
function updateHOYChartWith30diasMenos1(item, index) {
   

   chart30dias.options.data[1].dataPoints[index].y = parseInt(item["data_valuea"]);
   chart30dias.options.data[1].dataPoints[index].label = item["data_name"];
   chart30dias.options.data[1].dataPoints[index].name = item["data_name"];



}
function updateHOYChartWith30diasMenos2(item, index) {
   

   chart30dias.options.data[2].dataPoints[index].y = parseInt(item["data_valuea"]);
   chart30dias.options.data[2].dataPoints[index].label = item["data_name"];
   chart30dias.options.data[2].dataPoints[index].name = item["data_name"];



}
///end
////pagados pagados 30dias y -1hr -2hr
		

	// update all charts with revelant data
	function updateCharts(dataIndex) {
		activeUsers = data[dataIndex].activeUsers;

		updateUsersMediumPieChartWithDB();
    	updateUsersStateChartWithDB();

       updateU24StateChartWithDB();
       updateU48StateChartWithDB();
		
		updateHOYChartWith("00");
		updateHOYChartWith("01");
		updateHOYChartWith("02");

		updateChartWith7dias("70");
		updateChartWith7dias("71");
		updateChartWith7dias("72");
		
		updateChartWith30dias("30");
		updateChartWith30dias("31");
		updateChartWith30dias("32");
		
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
			
		} else if ( 1200 >= $(window).outerWidth()  && $(window).outerWidth() >= 992) {
			
			chart.options.legend.fontSize = 12;
			chart.options.legend.horizontalAlign = "center";
			chart.options.legend.verticalAlign = "top";
			chart.options.legend.maxWidth = null;
			
		} else if ( 992 >=  $(window).outerWidth()) {
			
			chart.options.legend.fontSize = 14;
			chart.options.legend.horizontalAlign = "center";
			chart.options.legend.verticalAlign = "bottom";
			chart.options.legend.maxWidth = null;
			
		}
		
		chart.render();
	}
	
	function customizeCharts() {
		chartPropertiesCustomization(usersMediumPieChart);
	}
	
	function renderAllCharts() {
		for (var i = 0; allCharts.length >=i  ; i++)
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
//        alert("next");
        window.location.href = "/v1/dash05";  

    }


	(function init() {
		customizeCharts();
		$(window).resize(customizeCharts);

		setTimeout(updateChartsAtRandomIntervals, valueRotacionOriginal); //every 4 seconds
             //initial
                updateChartsAtRandomIntervals();
        setInterval(checkRotationValues, valueRotacionOriginal);//every 20 seconds

        setInterval(go_next2, valuePasarNextOriginal);//every 150 seconds
//		setTimeout(updateChartsAtRandomIntervals, 4000);
		sidebarToggleOnClick();
	})();
	
});
