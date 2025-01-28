<script setup>
import BaseCrud from '@/components/base/BaseCrud.vue'
import BaseInput from '@/components/base/BaseInput.vue'
import UserCrudEntry from '@/components/UserComp/UserCrudEntry.vue'

import { UserStore } from '@/stores/UserStore'

const userState = UserStore()
userState.url = 'http://127.0.0.1:3030/api/user'
</script>

<template>
  <main class="mt-5 px-28">
    <BaseCrud
      :store="userState"
      :headerStyle="'p-4 flex'"
      :mainStyle="'rounded-xl border-gray-400 border flex'"
      :headStyle="'bg-gray-100 p-4 bg-gray-100 text-lg font-semibold text-gray-500 border-t border-gray-500'"
      :entries="['ID', 'Sirname', 'Lastname', 'Email', 'Role', 'Userage', 'Action']"
    >
      <template v-slot:header>
        <div class="flex justify-between w-full">
          <button
            @click="showAddModal = true"
            class="btn font-medium btn-normal space-x-2 items-center"
          >
            <span>+</span>
            <span>Add Entry</span>
          </button>
          <div class="min-w-80">
            <BaseInput :inputStyle="'input'" :label="'search'"></BaseInput>
          </div>
        </div>
      </template>

      <template #data="{ data }">
        <UserCrudEntry :data="data"></UserCrudEntry>
      </template>

      <!--Crud Footer-->
      <template v-slot:footer>
        <div class="flex justify-between p-4 items-center">
          <div class="font-semibold text-gray-600">
            Page {{ userState.page }} from {{ userState.totalPages }}
          </div>
          <div class="flex space-x-2">
            <button class="btn" @click="userState.dPage">
              <i class="bi bi-caret-left-fill"></i>
            </button>
            <button class="btn" @click="userState.iPage">
              <i class="bi bi-caret-right-fill"></i>
            </button>
          </div>
        </div>
      </template>
    </BaseCrud>
  </main>
</template>
