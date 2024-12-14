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
      <Button icon="pi pi-send" />
    </InputGroup>
  </Dialog>
</template>
<script setup lang="ts">
import { Dialog, InputGroup, InputText, Button } from 'primevue'
import { computed, ref, watch, type PropType } from 'vue'
import type { Tag } from '@/services/store/types'

defineOptions({
  name: 'TagDialog',
})
defineEmits(['hide'])

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
</script>
