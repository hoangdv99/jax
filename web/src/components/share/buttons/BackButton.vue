<template>
  <Button
    icon="pi pi-arrow-left"
    variant="text"
    severity="secondary"
    @click="back"
  />
</template>
<script lang="ts" setup>
import router from '@/router'
import { Button } from 'primevue'

defineOptions({
  name: 'BackButton',
})

const props = defineProps({
  name: { type: String, default: '' },
})

function back() {
  if (props.name) {
    router.push({ name: props.name })
  } else {
    router.go(-1)
  }
}
</script>
