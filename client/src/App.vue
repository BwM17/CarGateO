<script setup>
import { RouterView } from 'vue-router'
import SideBar from './components/nav/SideBar.vue'
import StatHeader from './components/GlobComp/StatHeader.vue'
import SideBarElement from './components/nav/SideBarElement.vue'
</script>

<template>
  <div class="flex">
    <SideBar>
      <SideBarElement to="/" icon="bi bi-house">Home</SideBarElement>
      <SideBarElement to="/cars" icon="bi bi-car-front-fill">Cars</SideBarElement>
      <SideBarElement to="/logs" icon="bi bi-journal-bookmark">Logs</SideBarElement>
      <SideBarElement to="/user" icon="bi bi-people-fill">User</SideBarElement>
    </SideBar>
    <main class="ml-16 w-full">
      <div>
        <StatHeader />
        <RouterView />
      </div>
    </main>
  </div>
</template>
