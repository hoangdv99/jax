<template>
  <div class="login-page">
    <Button label="submit" />
  </div>
</template>
<script lang="ts">
import { defineComponent } from 'vue'
import Button from 'primevue/button'

export default defineComponent({
  name: 'LoginPage',
  components: {
    Button,
  },
})
</script>
