<script setup>
import { computed } from 'vue'
const props = defineProps({
  data: Object,
})

const formatedDate = computed(() => {
  const date = new Date(props.data.timewhen)
  const options = { day: '2-digit', month: 'short' }
  return `${date.toLocaleDateString('en-GB', options)} ${date.getHours()}:${date.getMinutes()}`
})
</script>

<template>
  <th>
    <div v-if="data.allowed === true" class="bg-gray-300 py-1 pl-3 px-7 text-gray-500">
      {{ formatedDate }}
    </div>
    <div v-else-if="data.allowed === false" class="bg-red-300 py-1 pl-3 px-7 text-gray-500">
      {{ formatedDate }}
    </div>
  </th>
  <th class="y-1 px-10 text-gray-500">
    {{ data.numberplate }}
  </th>
  <th class="px-5 flex justify-start">
    <div v-if="data.allowed === true" class="text-green-500 space-x-4">
      <i class="bi bi-check-circle-fill"></i>
      {{ data.allowed }}
    </div>
    <div v-else-if="data.allowed === false" class="text-red-500 space-x-4">
      <i class="bi bi-x-circle-fill"></i>
      {{ data.allowed }}
    </div>
  </th>
  <th>
    <div class="flex justify-start font-normal text-gray-600">
      {{ data.status }}
    </div>
  </th>
</template>
