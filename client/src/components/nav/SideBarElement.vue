<script setup>
import { sideBarStore } from '@/stores/sideBar'
import { Transition, computed } from 'vue'
import { useRoute } from 'vue-router'
const props = defineProps({
  to: {
    type: String,
    required: true,
  },
  icon: {
    type: String,
    required: true,
  },
})
const route = useRoute()
const isActive = computed(() => route.path === props.to)
const sideBarState = sideBarStore()
</script>

<template>
  <router-link
    :to="to"
    class="flex items-center space-x-4 text-white px-2 py-1 mx-2 rounded-xl"
    :class="{ active: isActive, 'bg-blue-500': isActive }"
  >
    <i class="text-2xl p-1" :class="props.icon"></i>
    <transition name="fade">
      <span v-if="!sideBarState.collapsed" class="text-xl">
        <slot />
      </span>
    </transition>
  </router-link>
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

.active {
  background-color: lightblue;
}
</style>
