var microwaveChart;
var highFrequencyChart;

document.addEventListener("DOMContentLoaded", function() {
  try {

    var date = new Date(Date.now());
    var defaultDate = date.getUTCFullYear() + "-" + (date.getUTCMonth() + 1) + "-" + date.getUTCDate();
    document.getElementById("datetimepicker-dashboard").flatpickr({
      inline: true,
      prevArrow: "<span title=\"Previous month\">&laquo;</span>",
      nextArrow: "<span title=\"Next month\">&raquo;</span>",
      defaultDate: defaultDate,
      onChange: function(){
        var cookie = $.cookie("jwt");
        GetKvMicrowave(cookie);
        GetHighFrequency(cookie);
      }
    });

    var cookie = $.cookie("jwt");
    GetKvMicrowave(cookie);
    GetHighFrequency(cookie);

  } catch (e) {
    alert(e);
  }
});
function GetKvMicrowave(cookie) {
  var strDate = $('#datetimepicker-dashboard').val();
  var req = JSON.parse(JSON.stringify({token: cookie,selDate: strDate}));
  $.ajax({
    url:'./api/getkvMicrowave',
    type:'GET',
    contentType: 'application/json',
    daraType: 'json',
    data: req,
  })
  .done( (data) => {
    SetMChart(data);
    return;
  })
  .fail( (jqXHR, textStatus, errorThrown) => {
    alert('Ajax failed');
    console.log("jqXHR          : " + jqXHR.status);
    console.log("textStatus     : " + textStatus);
    console.log("errorThrown    : " + errorThrown.message);
    return;
  });
}
function GetHighFrequency(cookie) {
  var strDate = $('#datetimepicker-dashboard').val();
  var req = JSON.parse(JSON.stringify({token: cookie,selDate: strDate}));
  $.ajax({
    url:'./api/getHighFrequency',
    type:'GET',
    contentType: 'application/json',
    daraType: 'json',
    data: req,
  })
  .done( (data) => {
    SetFChart(data);
    return;
  })
  .fail( (jqXHR, textStatus, errorThrown) => {
    alert('Ajax failed');
    console.log("jqXHR          : " + jqXHR.status);
    console.log("textStatus     : " + textStatus);
    console.log("errorThrown    : " + errorThrown.message);
    return;
  });
}
function SetMChart(strJson) {

  try {
    var values = JSON.parse(strJson);
    const labels = [];
    const datapoints =[];
    const datapoints2 =[];
    const datapoints3 =[];
    const datapoints4 =[];
    const datapoints5 =[];
    const datapoints6 =[];
    for (var item in values) {
      labels.push(values[item].DataDate);
      datapoints.push(values[item].StopCV);
      datapoints2.push(values[item].WaterLevel);
      datapoints3.push(values[item].FarInfraredHeater);
      datapoints4.push(values[item].MwOutput);
      datapoints5.push(values[item].MwTotalOutput);
      datapoints6.push(values[item].InRoomPressure);
    }

    const data = {
      labels: labels,
      datasets: [
        {
          label: 'StopCV',
          data: datapoints,
          tension: 0.4,
          fill: true,
          backgroundColor: settings.CHARTBG_COLORS.green,
          borderColor: settings.CHART_COLORS.green,
          borderWidth:2,
          pointStyle: 'circle',
          pointRadius: 10,
          pointHoverRadius: 15
        }, {
          label: 'WaterLevel',
          data: datapoints2,
          tension: 0.4,
          fill: true,
          backgroundColor: settings.CHARTBG_COLORS.red,
          borderColor: settings.CHART_COLORS.red,
          borderWidth:2,
          pointStyle: 'circle',
          pointRadius: 10,
          pointHoverRadius: 15
        }, {
          label: 'FarInfraredHeater',
          data: datapoints3,
          tension: 0.4,
          fill: false,
          backgroundColor: settings.CHARTBG_COLORS.blue,
          borderColor: settings.CHART_COLORS.blue,
          borderWidth:2,
          pointStyle: 'circle',
          pointRadius: 10,
          pointHoverRadius: 15
        }, {
          label: 'MwOutput',
          data: datapoints4,
          tension: 0.4,
          fill: false,
          backgroundColor: settings.CHARTBG_COLORS.purple,
          borderColor: settings.CHART_COLORS.purple,
          borderWidth:2,
          pointStyle: 'circle',
          pointRadius: 10,
          pointHoverRadius: 15
        }
      ]
    };
    var ctx = document.getElementById('microwaveChart');
    if (microwaveChart) {
      microwaveChart.destroy();
    }
    microwaveChart = new Chart(ctx, {
      type: 'line',
      data: data,
      options: {
        responsive: true,
        title: {
          display: true,
          fontSize: 18,
          text: "MicroWave"
        },
        legend:{
          display: false,
        }
      },
      });

  } catch (e) {
    alert(e);
  }
}

function SetFChart(strJson) {
  try {
    var values = JSON.parse(strJson);
    const labels = [];
    const datapoints =[];
    const datapoints2 =[];
    const datapoints3 =[];
    const datapoints4 =[];
    const datapoints5 =[];
    const datapoints6 =[];
    const datapoints7 =[];
    for (var item in values) {
      labels.push(values[item].DataDate);
      datapoints.push(values[item].CurrentNow);
      datapoints2.push(values[item].TuningNow);
      datapoints3.push(values[item].WeldingTime);
      datapoints4.push(values[item].CoolingTime);
      datapoints5.push(values[item].WeldingTime1);
      datapoints6.push(values[item].CoolingTime1);
      datapoints7.push(values[item].WeldingTime2);
    }

    const data = {
      labels: labels,
      datasets: [
        {
          label: 'CurrentNow',
          data: datapoints,
          yAxisID: 'y-axis-1',
          tension: 0.4,
          fill: false,
          backgroundColor: settings.CHARTBG_COLORS.green,
          borderColor: settings.CHART_COLORS.green,
          borderWidth:2,
          pointStyle: 'circle',
          pointRadius: 10,
          pointHoverRadius: 15
        }, {
          label: 'TuningNow',
          data: datapoints2,
          yAxisID: 'y-axis-2',
          tension: 0.4,
          fill: false,
          backgroundColor: settings.CHARTBG_COLORS.red,
          borderColor: settings.CHART_COLORS.red,
          borderWidth:2,
          pointStyle: 'circle',
          pointRadius: 10,
          pointHoverRadius: 15
        }, {
          label: 'WeldingTime',
          data: datapoints3,
          yAxisID: 'y-axis-1',
          tension: 0.4,
          fill: false,
          backgroundColor: settings.CHARTBG_COLORS.blue,
          borderColor: settings.CHART_COLORS.blue,
          borderWidth:2,
          pointStyle: 'circle',
          pointRadius: 10,
          pointHoverRadius: 15
        }, {
          label: 'CoolingTime',
          data: datapoints4,
          yAxisID: 'y-axis-1',
          tension: 0.4,
          fill: false,
          backgroundColor: settings.CHARTBG_COLORS.purple,
          borderColor: settings.CHART_COLORS.purple,
          borderWidth:2,
          pointStyle: 'circle',
          pointRadius: 10,
          pointHoverRadius: 15
        }, {
          label: 'WeldingTime1',
          data: datapoints5,
          yAxisID: 'y-axis-1',
          tension: 0.4,
          fill: false,
          backgroundColor: settings.CHARTBG_COLORS.grey,
          borderColor: settings.CHART_COLORS.grey,
          borderWidth:2,
          pointStyle: 'circle',
          pointRadius: 10,
          pointHoverRadius: 15
        }, {
          label: 'CoolingTime1',
          data: datapoints6,
          yAxisID: 'y-axis-1',
          tension: 0.4,
          fill: false,
          backgroundColor: settings.CHARTBG_COLORS.blue,
          borderColor: settings.CHART_COLORS.blue,
          borderWidth:2,
          pointStyle: 'circle',
          pointRadius: 10,
          pointHoverRadius: 15
        }, {
          label: 'WeldingTime2',
          data: datapoints7,
          yAxisID: 'y-axis-2',
          tension: 0.4,
          fill: false,
          backgroundColor: settings.CHARTBG_COLORS.yellow,
          borderColor: settings.CHART_COLORS.yellow,
          borderWidth:2,
          pointStyle: 'circle',
          pointRadius: 10,
          pointHoverRadius: 15
        }
      ]
    };
    var ctx = document.getElementById('highFrequencyChart');
    if (highFrequencyChart){
      highFrequencyChart.destroy();
    }
    highFrequencyChart = new Chart(ctx, {
      type: 'line',
      data: data,
      options: {
        responsive: true,
        title: {
          display: true,
          fontSize: 18,
          text: "HighFrequency"
        },
        legend:{
          display: false,
        }
      },
      scales: {
          yAxes: [{
              id: 'y-axis-1',
              type: 'linear',
              position: 'left',
              ticks: {
                  min: 0,
                  stepSize: 1,
                  callback: function(val){
                      return val.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',');
                  }
              },
              gridLines: {
                  display: false,
              }
          }, {
              id: 'y-axis-2',
              type: 'linear',
              position: 'right',
              ticks: {
                  min: 0,
                  stepSize: 10,
                  callback: function(val){
                      return val.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',');
                  }
              },
              gridLines: {
                  zeroLineColor: '#000',
                  drawBorder: false,
              }
          }]
        }
      });

  } catch (e) {
    alert(e);
  }
}

const settings = {
  MONTHS: [
    'January',
    'February',
    'March',
    'April',
    'May',
    'June',
    'July',
    'August',
    'September',
    'October',
    'November',
    'December'
  ],
  COLORS: [
    '#4dc9f6',
    '#f67019',
    '#f53794',
    '#537bc4',
    '#acc236',
    '#166a8f',
    '#00a950',
    '#58595b',
    '#8549ba'
  ],
  CHART_COLORS: {
    red: 'rgb(255, 99, 132)',
    orange: 'rgb(255, 159, 64)',
    yellow: 'rgb(255, 205, 86)',
    green: 'rgb(75, 192, 192)',
    blue: 'rgb(54, 162, 235)',
    purple: 'rgb(153, 102, 255)',
    grey: 'rgb(201, 203, 207)'
  },
  CHARTBG_COLORS: {
    red: 'rgba(255, 99, 132,0.3)',
    orange: 'rgba(255, 159, 64,0.3)',
    yellow: 'rgba(255, 205, 86,0.3)',
    green: 'rgba(75, 192, 192,0.3)',
    blue: 'rgba(54, 162, 235,0.3)',
    purple: 'rgba(153, 102, 255,0.3)',
    grey: 'rgba(201, 203, 207,0.3)'
  }
};    

