<template>
  <Dialog
    :visible="props.visible"
    modal
    header="Add new store"
    :draggable="false"
    :style="{ width: '640px' }"
    class="store-dialog"
    @update:visible="$emit('hide')"
  >
    <div class="field">
      <label for="url" class="label">URL</label>
      <InputText id="url" autocomplete="off" class="input" />
    </div>
    <div class="field">
      <label for="tags" class="label">Tags</label>
      <MultiSelect
        id="tags"
        v-model="store.tags"
        display="chip"
        :options="storeStore.listTag"
        filter
        filterPlaceholder="Search"
        :show-toggle-all="false"
        option-label="name"
        class="select"
      >
        <template #header>
          <div class="header-container">
            <Button
              icon="pi pi-plus"
              label="Create new tag"
              @click="startCreateNewTag"
            />
          </div>
        </template>
        <template #option="slotProps">
          <div class="option-wrapper">
            <div>
              {{ slotProps.option.name }}
            </div>
            <div class="actions">
              <Button
                icon="pi pi-pencil"
                severity="secondary"
                variant="text"
                @click="e => startEdit(e, slotProps.option)"
              />
              <Button
                icon="pi pi-trash"
                severity="danger"
                variant="text"
                @click="e => onDeleteTag(e, slotProps.option.id)"
              />
            </div>
          </div>
        </template>
      </MultiSelect>
    </div>
  </Dialog>
  <TagDialog
    :visible="showTagDialog"
    :originalTag="selectedTag"
    @hide="showTagDialog = false"
  />
</template>
<script setup lang="ts">
import { Dialog, InputText, MultiSelect, Button, useToast } from 'primevue'
import { reactive, ref } from 'vue'
import { useConfirm } from 'primevue/useconfirm'
import TagDialog from './TagDialog.vue'
import type { Tag } from '@/services/store/types'
import { useStoreStore } from '@/stores/store'
import { services } from '@/services'

defineOptions({
  name: 'StoreDialog',
})
defineEmits(['hide'])

const confirm = useConfirm()
const toast = useToast()
const storeStore = useStoreStore()

const props = defineProps({
  visible: { type: Boolean, default: false },
})

const store = reactive({
  url: '',
  tags: [],
})
const showTagDialog = ref(false)

const selectedTag = ref<Tag>()

function onDeleteTag(e: Event, id: number) {
  e.stopPropagation()
  confirm.require({
    header: 'Confirmation',
    message: 'Do you want to delete this tag?',
    rejectLabel: 'Cancel',
    icon: 'pi pi-info-circle',
    rejectProps: {
      label: 'Cancel',
      severity: 'secondary',
      outlined: true,
    },
    acceptProps: {
      label: 'Delete',
      severity: 'danger',
    },
    accept: async () => {
      await deleteTag(id)
    },
  })
}

async function deleteTag(id: number) {
  const res = await services.store.deleteTag(id)
  if (res.success) {
    storeStore.getListTag()
    toast.add({
      severity: 'success',
      summary: 'Success',
      detail: 'Deleted tag',
    })
  } else {
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: res.message?.name || 'Internal server error',
    })
  }
}

function startCreateNewTag() {
  selectedTag.value = undefined
  showTagDialog.value = true
}
function startEdit(e: Event, tag: Tag) {
  e.stopPropagation()
  showTagDialog.value = true
  selectedTag.value = tag
}
</script>
<style lang="scss" scoped>
.store-dialog {
  .field {
    display: flex;
    gap: 24px;
    align-items: center;
    margin-bottom: 16px;
  }
  .field > .label {
    font-weight: 600;
  }
  .field > .input,
  .field > .select {
    width: 100%;
  }
}

.option-wrapper {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.header-container {
  display: flex;
  padding: 4px 14px 0 14px;
}
</style>
