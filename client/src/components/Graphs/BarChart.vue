<script setup>
import ApexCharts from 'vue3-apexcharts'
const chartOptions = {
  chart: {
    type: 'bar',
  },
  plotOptions: {
    bar: {
      horizontal: false,
      columnWidth: '30%',
      endingShape: 'rounded',
    },
  },
  dataLabels: {
    enabled: false,
  },
  legend: {
    show: false,
  },
  xaxis: {
    categories: [
      '06:00',
      '07:00',
      '08:00',
      '09:00',
      '10:00',
      '11:00',
      '12:00',
      '13:00',
      '14:00',
      '15:00',
      '16:00',
      '17:00',
      '18:00',
      '19:00',
      '20:00',
      '21:00',
      '22:00',
    ],
    labels: {
      show: true,
    },
  },
}

const series = [
  {
    name: 'Came',
    data: [6, 29, 11, 13, 20, 24, 7, 30, 25, 17, 16, 10, 15, 22, 28, 5, 27],
  },
  {
    name: 'Left',
    data: [19, 8, 14, 27, 23, 12, 26, 9, 18, 21, 6, 29, 11, 13, 20, 24, 7],
  },
]
</script>

<template>
  <ApexCharts type="bar" height="300" :options="chartOptions" :series="series"></ApexCharts>
</template>
