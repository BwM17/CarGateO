import js from '@eslint/js'
import { defineStore } from 'pinia'

export const LogStore = defineStore({
  id: 'logs',
  state: () => ({
    url: null,
    data: null,
    page: 1,
    pageSize: 15,
    totalPages: null,
  }),
  actions: {
    async get() {
      try {
        console.log(this.url + `?page=${this.page}&page_size=${this.pageSize}`)
        const response = await fetch(this.url + `?page=${this.page}&page_size=${this.pageSize}`)
        if (!response.ok) {
          throw new Error(`Error: ${response.status}`)
        }

        const json = await response.json()
        console.log(json)
        this.totalPages = json.totalpages
        this.data = json.data
      } catch (error) {
        console.log(error)
      }
    },
    async post(obj) {},
    async delete(url) {},

    iPage() {
      if (this.page == this.totalPages) {
        return
      }
      this.page++
      this.get()
    },

    dPage() {
      if (this.page == 1) {
        return
      }
      this.page--
      this.get()
    },
  },
})
