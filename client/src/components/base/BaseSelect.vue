<script setup>
import { defineModel } from 'vue'

const props = defineProps({
  options: Array,
  selectStyle: String,
})

const modelValue = defineModel()
</script>

<template>
  <select :class="props.selectStyle" @input="$emit('update:modelValue', $event.target.value)">
    <option v-for="option in props.options" :value="option">
      {{ option }}
    </option>
  </select>
  <p>{{ modelValue }}</p>
</template>
