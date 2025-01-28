<script setup>
const props = defineProps({
  entries: Array,
  tableStyle: String,
  headStyle: String,
  bodyStyle: String,
})
</script>

<template>
  <table :class="props.tableStyle" class="align-start">
    <thead>
      <tr>
        <th :class="props.headStyle" v-for="entry in props.entries" scope="col">
          {{ entry }}
        </th>
      </tr>
    </thead>
    <tbody :class="props.bodyStyle">
      <slot />
    </tbody>
  </table>
</template>
