<template>
  <div class="default-layout">
    <Sidebar />
    <div class="content-container">
      <RouterView />
    </div>
  </div>
</template>
<script setup lang="ts">
import Sidebar from './Sidebar.vue'

defineOptions({
  name: 'DefaultLayout',
})
</script>
<style lang="scss" scoped>
.default-layout {
  display: flex;
}

.content-container {
  width: 100%;
  padding: 16px;
}
</style>
