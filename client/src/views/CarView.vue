<script setup>
import BaseCrud from '@/components/base/BaseCrud.vue'
import BaseInput from '@/components/base/BaseInput.vue'
import BaseModal from '@/components/base/BaseModal.vue'
import CarAddForm from '@/components/CarComp/CarAddForm.vue'
import CarCrudEntry from '@/components/CarComp/CarCrudEntry.vue'
import CarEditForm from '@/components/CarComp/CarEditForm.vue'

import { CarStore } from '@/stores/CarStore'
import { ref } from 'vue'

const carState = CarStore()
carState.url = 'http://127.0.0.1:3030/api/car'

const showAddModal = ref(false)
const showDeletModal = ref(false)
const showEditModal = ref(false)
</script>

<template>
  <main class="mt-5 px-28">
    <BaseCrud
      :store="carState"
      :headerStyle="'p-4 flex'"
      :mainStyle="'rounded-xl border-gray-400 border min-w-[30rem] flex'"
      :headStyle="'bg-gray-100 p-4 bg-gray-100 text-lg font-semibold text-gray-500 border-t border-gray-500'"
      :entries="[
        'ID',
        'Model',
        'Numberplate',
        'Owner',
        'Type',
        'Status',
        'Lease',
        'Return',
        'Action',
      ]"
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
        <CarCrudEntry
          :data="data"
          @edit="showEditModal = true"
          @delete="showDeletModal = true"
        ></CarCrudEntry>
      </template>

      <!--Crud Footer-->
      <template v-slot:footer>
        <div class="flex justify-between p-4 items-center">
          <div class="font-semibold text-gray-600">
            Page {{ carState.page }} from {{ carState.totalPages }}
          </div>
          <div class="flex space-x-2">
            <button class="btn" @click="carState.dPage">
              <i class="bi bi-caret-left-fill"></i>
            </button>
            <button class="btn" @click="carState.iPage">
              <i class="bi bi-caret-right-fill"></i>
            </button>
          </div>
        </div>
      </template>
    </BaseCrud>
  </main>

  <!--Add Modal-->
  <BaseModal v-show="showAddModal" :modalStyle="'bg-white rounded-2xl p-4 w-[40rem]'">
    <!--Add Modal Header-->
    <template v-slot:header>
      <div class="flex justify-between border-b-gray-300 border-b pb-3 items-end">
        <h3 class="text-2xl font-semibold">Add A Car</h3>
        <button class="btn btn-danger" @click="showAddModal = false">Close</button>
      </div>
    </template>

    <!--Add Modal Body-->
    <template v-slot:body>
      <CarAddForm></CarAddForm>
    </template>

    <!--Add Modal Footer-->
    <template v-slot:footer>
      <div class="mt-10 border-gray-500 mx-10">
        <button class="btn-normal w-full py-3 rounded-xl text-xl font-semibold">Submit</button>
      </div>
    </template>
  </BaseModal>

  <!--Edit Modal-->
  <BaseModal v-show="showEditModal" :modalStyle="'bg-white rounded-2xl p-4 w-[40rem]'">
    <!--Edit Modal Header-->
    <template v-slot:header>
      <div class="flex justify-between border-b-gray-300 border-b pb-3 items-end">
        <h3 class="text-2xl font-semibold">Add A Car</h3>
        <button class="btn btn-danger" @click="showEditModal = false">Close</button>
      </div>
    </template>

    <!--Edit Modal Body-->
    <template v-slot:body>
      <CarEditForm></CarEditForm>
    </template>

    <!--Edit Modal Footer-->
    <template v-slot:footer>
      <div class="mt-10 border-gray-500">
        <button class="btn-normal w-full px-3 py-1 rounded-xl">Submit</button>
      </div>
    </template>
  </BaseModal>

  <!--Delete Modal-->
  <BaseModal v-show="showDeletModal" :modalStyle="'bg-white px-10 py-4 rounded-2xl'">
    <template v-slot:header>
      <div class="text-xl font-semibold flex justify-center mb-5">
        <span>Delete this Entry ?</span>
      </div>
    </template>
    <template v-slot:body>
      <div class="space-x-5 flex justify-center">
        <button @click="showDeletModal = false" class="btn btn-normal">Cancel</button>
        <button @click="showDeletModal = false" class="btn btn-danger">Delete</button>
      </div>
    </template>
  </BaseModal>
</template>
