<template>
    <apexchart type="radialBar"  height="450" :options="options" :series="series"></apexchart>
 </template>

<script>
import VueApexCharts from "vue-apexcharts";
import * as Wails from '@wailsapp/runtime';
export default {    
  name: 'guage',
  data() {
    return {
      series: [0],
      options: {
        chart: {
          height: 450,
          type: 'radialBar',
          toolbar: {
            show: true
          }
        },
        plotOptions: {
          radialBar: {
            startAngle: -135,
            endAngle: 225,
              hollow: {
              margin: 0,
              size: '70%',
              background: '#fff',
              image: undefined,
              imageOffsetX: 0,
              imageOffsetY: 0,
              position: 'front',
              dropShadow: {
                enabled: true,
                top: 3,
                left: 0,
                blur: 4,
                opacity: 0.24
              }
          },
          track: {
              background: '#fff',
              strokeWidth: '67%',
              margin: 0, // margin is in pixels
              dropShadow: {
                enabled: true,
                top: -3,
                left: 0,
                blur: 4,
                opacity: 0.35
              }
          },
        
          dataLabels: {
              show: true,
              name: {
                offsetY: -10,
                show: true,
                color: '#888',
                fontSize: '17px'
              },
              value: {
                formatter: function(val) {
                  return parseInt(val);
                },
                color: '#111',
                fontSize: '36px',
                show: true,
              }
            }
          }
        },
        /////
        fill: {
        type: 'gradient',
          gradient: {
            shade: 'dark',
            type: 'horizontal',
            shadeIntensity: 0.5,
            gradientToColors: ['#ABE5A1'],
            inverseColors: true,
            opacityFrom: 1,
            opacityTo: 1,
            stops: [0, 100]
          }
      },
      stroke: {
        lineCap: 'round'
      },
      labels: ['Fuel Usage'],
      }
    };
  },
  components: {
    apexchart: VueApexCharts,
  },
  mounted: function() {
    Wails.Events.On("sensor_reading", sensor_reading => {
      if (sensor_reading) {
        this.series = [ sensor_reading.value ];
        console.log(sensor_reading.value);
      }
    });
  }
};
</script>