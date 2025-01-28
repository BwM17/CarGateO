<script setup>
import ApexCharts from 'vue3-apexcharts'
import { ref } from 'vue'
const series = [66, 77, 88]

const chartOptions = ref({
  chart: {
    type: 'pie',
  },
  labels: ['parked', 'gone', 'leased'],
  responsive: [
    {
      breakpoint: 480,
      options: {
        chart: {
          width: 200,
        },
        legend: {
          position: 'bottom',
        },
      },
    },
  ],
})
</script>

<template>
  <div id="chart">
    <ApexCharts type="pie" width="350" :options="chartOptions" :series="series"></ApexCharts>
  </div>
</template>
