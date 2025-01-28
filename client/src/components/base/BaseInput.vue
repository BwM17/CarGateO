<script setup>
const props = defineProps({
  label: 'text',
  inputStyle: String,
})

const modelValue = defineModel()
</script>

<template>
  <input
    :placeholder="props.label"
    :class="props.inputStyle"
    :vlaue="modelValue"
    @input="$emit('update:modelValue', $event.target.value)"
  />
</template>
