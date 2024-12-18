<template>
  <Dialog
    :visible="props.visible"
    modal
    :header="dialogHeader"
    :draggable="false"
    :style="{ width: '480px' }"
    @update:visible="$emit('hide')"
  >
    <InputGroup>
      <InputText v-model="tag.name" />
      <Button icon="pi pi-send" @click="createNewTag" />
    </InputGroup>
  </Dialog>
</template>
<script setup lang="ts">
import { Dialog, InputGroup, InputText, Button, useToast } from 'primevue'
import { computed, ref, watch, type PropType } from 'vue'
import type { Tag } from '@/services/store/types'
import { services } from '@/services'
import { useStoreStore } from '@/stores/store'

defineOptions({
  name: 'TagDialog',
})
const emit = defineEmits(['hide'])

const storeStore = useStoreStore()
const toast = useToast()

const props = defineProps({
  visible: { type: Boolean, default: false },
  originalTag: { type: Object as PropType<Tag> },
})

const tag = ref<Tag>({
  id: 0,
  name: '',
})

watch(
  () => props.originalTag ?? { id: 0, name: '' },
  (newVal: Tag) => {
    tag.value = newVal
  }
)

const dialogHeader = computed(() => {
  return props.originalTag ? 'Edit tag' : 'Add new tag'
})

async function createNewTag() {
  const res = await services.store.createNewTag({ name: tag.value.name.trim() })
  if (res.success) {
    toast.add({
      severity: 'success',
      summary: 'Success',
      detail: 'Created new tag',
      life: 3000,
    })
    storeStore.getListTag()
    emit('hide')
  } else {
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: res.message?.name || 'Internal server error',
    })
  }
}
</script>
