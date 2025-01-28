<script setup>
import { sideBarStore } from '@/stores/sideBar'
import { Transition } from 'vue'

const sideBarState = sideBarStore()
</script>

<template>
  <!--Wrapper-->
  <div
    class="bg-blue-950 fixed flex flex-col h-full transition-all float-left z-10 top-0 left-0 bottom-0 text-white"
    :style="{ width: sideBarState.sideBarWidth }"
  >
    <!--Header-->
    <div class="px-4 py-3">
      <span v-if="sideBarState.collapsed">
        <img class="w-8 h-8" src="../../assets/car.svg" />
      </span>
      <transition v-else name="fade">
        <span class="flex text-2xl space-x-4">
          <img class="w-8 h-8" src="../../assets/car.svg" />
          <span>CarGate</span>
        </span>
      </transition>
    </div>

    <!--Links-->
    <div class="mt-10 space">
      <slot />
    </div>

    <!--Bottom Button-->
    <div
      class="bottom-0 absolute text-2xl py-4 px-5 trasition-all ease-in-out duration-200 right-0"
      :class="{ 'transform rotate-180 ': sideBarState.collapsed }"
    >
      <span @click="sideBarState.toggle">
        <i class="bi bi-chevron-double-left"></i>
      </span>
    </div>
  </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.1s;
}

.fade-enter,
.fade-leave-to {
  opacity: 0;
}
</style>
