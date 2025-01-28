<script setup>
import { useRoute } from 'vue-router'

const route = useRoute()
</script>

<template>
  <div class="flex justify-end space-x-4 px-5 py-4">
    <span class="text-xl font-semibold text-gray-500">
      {{ route.name }}
    </span>
    <span class="flex">
      <img src="../../assets/doge.jpg" class="h-10 w-10 rounded-full" alt="pfp" />
    </span>
  </div>
</template>
