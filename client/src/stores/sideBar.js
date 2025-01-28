import { defineStore } from 'pinia'

export const sideBarStore = defineStore({
  id: 'sideBarStore',
  state: () => ({
    collapsed: true,
    SIDEBAR_WIDTH: 15,
    SIDEBAR_WIDTH_COLLAPSED: 4,
  }),
  getters: {
    sideBarWidth: (state) => {
      return `${state.collapsed ? state.SIDEBAR_WIDTH_COLLAPSED : state.SIDEBAR_WIDTH}rem`
    },
  },
  actions: {
    toggle() {
      this.collapsed = !this.collapsed
    },
  },
})
