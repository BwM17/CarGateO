<script setup>
import PieChart from '@/components/Graphs/PieChart.vue'
import BarChart from '@/components/Graphs/BarChart.vue'
import DueTo from '@/components/Graphs/DueTo.vue'
</script>

<template>
  <div className="grid grid-cols-3 grid-rows-6 gap-4 mx-60 h-[85vh] ">
    <div className="row-span-3 border-gray-200 border rounded-xl shadow-xl p-4">
      <span class="flex text-3xl font-semibold space-x-5 text-gray-400">
        <i class="bi bi-camera-fill"></i>
        <p>Latest Car</p></span
      >
      <img src="../assets/car.svg" class="rounded-xl px-10" />
    </div>
    <div className="bg-green-500 text-white rounded-xl shadow-lg border-gray-200 border py-4 px-7">
      <DueTo :title="'Cars Currently Leased'" :stats="30" />
    </div>
    <div className="bg-red-500 text-white rounded-xl shadow-lg border-gray-200 border py-4 px-7">
      <DueTo :title="'Cars needed to Return today'" :stats="10" />
    </div>
    <div
      className="row-span-2 col-start-2 border-gray-200 border rounded-xl shadow-xl py-4 px-7 bg-blue-500 text-white"
    >
      <p class="text-3xl font-semibold mb-7">Latest Car</p>
      <div class="flex space-x-5">
        <div class="space-y-1">
          <p class="font-semibold">Model:</p>
          <p class="font-semibold">Owner:</p>
          <p class="font-semibold">Sign:</p>
          <p class="font-semibold">Type:</p>
        </div>
        <div class="space-y-1 text-gray-200">
          <p>Opel</p>
          <p>MB</p>
          <p>RODMB1717</p>
          <p>Company</p>
        </div>
      </div>
    </div>
    <div
      className="row-span-2 col-start-3 border border-gray-200 rounded-xl shadow-lg justify-center"
    >
      <PieChart />
    </div>
    <div
      className="col-span-3 row-span-3 row-start-4 border border-gray-200 rounded-xl shadow-lg pt-10 px-10"
    >
      <BarChart />
    </div>
  </div>
</template>
