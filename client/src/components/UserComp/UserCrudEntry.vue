<script setup>
import { computed } from 'vue'
const props = defineProps({
  data: Object,
})

const emit = defineEmits(['delete', 'edit'])

const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('en-GB')
}
</script>

<template>
  <tr
    v-for="data in props.data"
    :key="data.id"
    class="border-gray-400 border-b border-t py-4 my-4 text-md w-full"
  >
    <th class="font-semibold py-4 text-gray-400">{{ data.id }}</th>
    <th class="font-medium">{{ data.sirname }}</th>
    <th class="font-medium">{{ data.lastname }}</th>
    <th class="font-medium text-gray-400">{{ data.email }}</th>
    <th>
      <div v-if="data.role === 'Admin'" class="flex text-blue-400 space-x-4 justify-center">
        <i class="bi bi-lock-fill"></i>
        <span>{{ data.role }}</span>
      </div>
      <div
        v-else-if="data.role === 'Watcher'"
        class="flex text-orange-300 space-x-4 justify-center"
      >
        <i class="bi bi-eye"></i>
        <span>{{ data.role }}</span>
      </div>
      <div v-else-if="data.role === 'Manager'" class="flex text-green-400 space-x-4 justify-center">
        <i class="bi bi-clipboard-data"></i>
        <span>{{ data.role }}</span>
      </div>
      <div v-else>
        <p>Wrong Type</p>
      </div>
    </th>
    <th class="font-normal text-gray-400">
      {{ formatDate(data.userage) }}
    </th>
    <th>
      <div class="flex space-x-5 justify-center">
        <span class="text-red-400 space-x-2 cursor-pointer" @click="emit('delete')">
          <i class="bi bi-trash"></i><span>Delete</span></span
        >
        <span class="text-blue-400 space-x-2 cursor-pointer" @click="log"
          ><i class="bi bi-pen"></i><span>Edit</span></span
        >
      </div>
    </th>
  </tr>
</template>
