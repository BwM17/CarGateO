<script setup>
import BaseCrud from '@/components/base/BaseCrud.vue'
import LogCrudEntry from '@/components/LogComp/LogCrudEntry.vue'
import PieChart from '@/components/Graphs/PieChart.vue'
import { LogStore } from '@/stores/LogStore'

const logState = LogStore()
logState.url = 'http://127.0.0.1:3030/api/logs'
</script>

<template>
  <main class="mx-[15rem]">
    <div className="grid grid-cols-4 grid-rows-4 gap-4 ">
      <div className="col-span-2 row-span-4 rounded-xl border-gray-400 border shadow-lg">
        <BaseCrud
          :store="logState"
          :entries="['Timestamp', 'Numberplate', 'Allowed', 'Status']"
          :mainStyle="'p-4'"
          :headStyle="'py-4'"
          :tableStyle="''"
        >
          <template #data="{ data }">
            <tr v-for="data in data" :key="data.id">
              <LogCrudEntry :data="data" />
            </tr>
          </template>
          <template v-slot:footer>
            <div class="flex justify-between p-4 items-center">
              <div class="font-semibold text-gray-600">
                Page {{ logState.page }} from {{ logState.totalPages }}
              </div>
              <div class="flex space-x-2">
                <button class="btn" @click="logState.dPage">
                  <i class="bi bi-caret-left-fill"></i>
                </button>
                <button class="btn" @click="logState.iPage">
                  <i class="bi bi-caret-right-fill"></i>
                </button>
              </div>
            </div>
          </template>
        </BaseCrud>
      </div>
    </div>
  </main>
</template>
