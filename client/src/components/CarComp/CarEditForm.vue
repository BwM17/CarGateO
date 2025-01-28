<script setup>
import BaseInput from '../base/BaseInput.vue'
import BaseSelect from '../base/BaseSelect.vue'

const props = defineProps({
  entry: Object,
})
</script>

<template>
  <form class="mt-6 space-y-5" @submit.prevent>
    <div>
      <label class="text-lg font-semibold">Car Model</label>
      <BaseInput :inputStyle="'input'" :placeholder="'Car Model'"></BaseInput>
    </div>
    <div>
      <label class="text-lg font-semibold">Numberplate</label>
      <BaseInput :inputStyle="'input'" placeholder="Numberplate"></BaseInput>
    </div>
    <div>
      <label class="text-lg font-semibold">Owner</label>
      <BaseInput :inputStyle="'input'" :placeholder="'Owner'"></BaseInput>
    </div>
    <div>
      <label class="text-lg font-semibold">Type</label>
      <BaseSelect :selectStyle="'select'" :options="['private', 'company', 'guest']"></BaseSelect>
    </div>
    <div>
      <label class="text-lg font-semibold mb-20">Status</label>
      <BaseSelect :selectStyle="'select'" :options="['parked', 'gone', 'leased']"></BaseSelect>
    </div>
  </form>
</template>
