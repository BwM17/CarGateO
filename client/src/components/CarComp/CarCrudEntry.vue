<script setup>
const props = defineProps({
  data: Object,
})

const emit = defineEmits(['delete', 'edit'])
</script>

<template>
  <tr
    v-for="data in props.data"
    :key="data.id"
    class="border-gray-400 border-b border-t py-4 my-4 text-md w-full"
  >
    <th class="font-semibold py-4 text-gray-400">{{ data.id }}</th>
    <th class="font-medium py-4">{{ data.model }}</th>
    <th class="font-medium">{{ data.numberplate }}</th>
    <th class="font-medium flex items-center">{{ data.owner }}</th>
    <th>
      <div v-if="data.type === 'company'" class="flex text-blue-400 space-x-4 justify-center">
        <i class="bi bi-building"></i>
        <span>{{ data.type }}</span>
      </div>
      <div
        v-else-if="data.type === 'private'"
        class="flex text-orange-300 space-x-4 justify-center"
      >
        <i class="bi bi-person-fill"></i>
        <span>{{ data.type }}</span>
      </div>
      <div v-else-if="data.type === 'guest'" class="flex text-green-400 space-x-4 justify-center">
        <i class="bi bi-clock-fill"></i>
        <span>{{ data.type }}</span>
      </div>
      <div v-else>
        <p>Wrong Type</p>
      </div>
    </th>
    <th>
      <div v-if="data.status === 'parked'" class="space-x-4">
        <i class="bi bi-circle-fill text-[10px] text-green-400"></i>
        <span>{{ data.status }}</span>
      </div>
      <div v-else-if="data.status === 'leased'" class="space-x-4">
        <i class="bi bi-circle-fill text-[10px] text-orange-300"></i>
        <span>{{ data.status }}</span>
      </div>
      <div v-else-if="data.status === 'gone'" class="space-x-4">
        <i class="bi bi-circle-fill text-[10px] text-red-400"></i>
        <span>{{ data.status }}</span>
      </div>
    </th>
    <th>
      <div v-if="data.lease === ''" class="text-gray-400 font-bold">-</div>
      <div v-else>
        {{ data.lease }}
      </div>
    </th>
    <th>
      <div v-if="data.return === ''" class="text-gray-400 font-bold">-</div>
      <div v-else>
        {{ data.return }}
      </div>
    </th>
    <th>
      <div class="flex space-x-5 justify-center">
        <span class="text-red-400 space-x-2 cursor-pointer" @click="emit('delete')">
          <i class="bi bi-trash"></i><span>Delete</span></span
        >
        <span class="text-blue-400 space-x-2 cursor-pointer" @click="emit('edit')"
          ><i class="bi bi-pen"></i><span>Edit</span></span
        >
      </div>
    </th>
  </tr>
</template>
