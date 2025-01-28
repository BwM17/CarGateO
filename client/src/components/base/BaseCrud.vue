<script setup>
import { onMounted } from 'vue'
import BaseTable from './BaseTable.vue'
const props = defineProps({
  store: Object,
  entries: Array,
  mainStyle: String,
  tableStyle: String,
  headStyle: String,
  bodyStyle: String,
  headerStyle: String,
})

onMounted(() => {
  props.store.get()
})

/*ToDos
    -[ ]search 
*/
</script>

<template>
  <div class="flex-col" :class="props.mainStyle">
    <div class="" :class="props.headerStyle">
      <slot name="header"></slot>
    </div>
    <BaseTable
      :entries="props.entries"
      :tableStyle="props.tableStyle"
      :headStyle="props.headStyle"
      :body-style="props.bodyStyle"
    >
      <slot name="data" v-bind:data="props.store.data"></slot>
    </BaseTable>
    <slot name="footer"></slot>
  </div>
</template>
