<script setup>
const props = defineProps({
  title: String,
  stats: Number,
})
</script>

<template>
  <div class="space-y-3">
    <div class="text-xl font-semibold">{{ props.title }}</div>
    <div class="text-5xl font-bold">{{ props.stats }}</div>
  </div>
</template>
