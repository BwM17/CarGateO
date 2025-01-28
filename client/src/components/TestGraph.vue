<template>
  <div id="chart">
    <apexchart type="radialBar" :options="chartOptions" :series="series"></apexchart>
  </div>
</template>

<script>
import VueApexCharts from 'vue3-apexcharts'

export default {
  name: 'MyChart',
  components: {
    apexchart: VueApexCharts,
  },
  data() {
    return {
      series: [76, 57, 78],
      chartOptions: {
        chart: {
          type: 'radialBar',
          offsetY: -20,
          sparkline: {
            enabled: true,
          },
        },
        plotOptions: {
          radialBar: {
            startAngle: -90,
            endAngle: 90,
            track: {
              background: '#e7e7e7',
              strokeWidth: '97%',
              margin: 3, // margin is in pixels
            },
            dataLabels: {
              name: {
                show: false,
              },
              value: {
                offsetY: -2,
                fontSize: '22px',
              },
            },
          },
        },
        grid: {
          padding: {
            top: -10,
          },
        },

        labels: ['Average Results'],
      },
    }
  },
}
</script>
