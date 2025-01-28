<script setup>
const props = defineProps({
  modalStyle: String,
})

const emit = defineEmits(['close'])
</script>

<template>
  <!--BackGround-->
  <div
    class="fixed z-[90] bg-black/65 top-0 bottom-0 left-0 right-0 flex justify-center items-center"
  >
    <!--PopUp-->
    <div :class="props.modalStyle">
      <slot name="header"></slot>
      <slot name="body"></slot>
      <slot name="footer"></slot>
    </div>
  </div>
</template>
