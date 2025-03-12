<template>
  <Dialog :visible="props.visible" modal header="Add new store" :draggable="false" :style="{ width: '640px' }"
    class="store-dialog" @update:visible="$emit('hide')">
    <div class="field">
      <label for="url" class="label">URL</label>
      <InputText id="url" v-model="store.url" :disabled="originalStore.id" autocomplete="off"
        placeholder="Ex: https://example.com" class="input" />
    </div>
    <div class="field">
      <label for="tags" class="label">Tags</label>
      <MultiSelect id="tags" v-model="store.tags" display="chip" :options="storeStore.listTag" filter
        filterPlaceholder="Search" :show-toggle-all="false" option-label="name" placeholder="Select tags"
        class="select">
        <template #header>
          <div class="header-container">
            <Button icon="pi pi-plus" label="Create new tag" @click="startCreateNewTag" />
          </div>
        </template>
        <template #option="slotProps">
          <div class="option-wrapper">
            <div>
              {{ slotProps.option.name }}
            </div>
            <div class="actions">
              <Button icon="pi pi-pencil" severity="secondary" variant="text"
                @click="e => startEdit(e, slotProps.option)" />
              <Button icon="pi pi-trash" severity="danger" variant="text" tabindex="-1"
                @click="e => onDeleteTag(e, slotProps.option.id)" />
            </div>
          </div>
        </template>
      </MultiSelect>
    </div>
    <div class="actions">
      <Button type="button" label="Cancel" severity="secondary" variant="outlined" @click="$emit('hide')"></Button>
      <Button type="button" label="Save" @click="saveStore"></Button>
    </div>
  </Dialog>
  <TagDialog :visible="showTagDialog" :originalTag="selectedTag" @hide="showTagDialog = false" />
</template>
<script setup lang="ts">
import { Dialog, InputText, MultiSelect, Button, useToast } from 'primevue'
import { reactive, ref, watch } from 'vue'
import { useConfirm } from 'primevue/useconfirm'
import TagDialog from './TagDialog.vue'
import type { Store, Tag } from '@/services/store/types'
import { useStoreStore } from '@/stores/store'
import { services } from '@/services'

defineOptions({
  name: 'StoreDialog',
})
const emit = defineEmits(['hide'])

const confirm = useConfirm()
const toast = useToast()
const storeStore = useStoreStore()

const props = defineProps({
  visible: { type: Boolean, default: false },
  originalStore: { type: Object, default: () => ({}) },
})

watch(() => props.originalStore, (value) => {
  store.url = value.url
  store.tags = value.tags
})

const store = reactive<Store>({
  id: null,
  url: '',
  tags: [],
})
const showTagDialog = ref(false)
const selectedTag = ref<Tag>()

function onDeleteTag(e: Event, id: number) {
  e.preventDefault()
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
async function saveStore() {
  if (props.originalStore.id) {
    await updateStore()
  } else {
    await addStore()
  }
}

async function addStore() {
  const res = await services.store.addStore({
    url: store.url,
    tagIds: store.tags.map((tag) => tag.id),
  })
  if (res.success) {
    storeStore.getListStore()
    toast.add({
      severity: 'success',
      summary: 'Success',
      detail: 'Added store',
    })
    store.id = null
    store.url = ''
    store.tags = []
    emit('hide')
  } else {
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: res.message.name || res.message || 'Internal server error',
    })
  }
}

async function updateStore() {
  const res = await services.store.updateStore(
    props.originalStore.id,
    {
      tagIds: store.tags.map((tag) => tag.id),
    }
  )
  if (res.success) {
    storeStore.getListStore()
    toast.add({
      severity: 'success',
      summary: 'Success',
      detail: 'Edited store',
    })
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
<style lang="scss" scoped>
.store-dialog {
  .field {
    display: flex;
    gap: 24px;
    align-items: center;
    margin-bottom: 16px;
  }

  .field>.label {
    font-weight: 600;
  }

  .field>.input,
  .field>.select {
    width: 100%;
  }

  .actions {
    display: flex;
    justify-content: flex-end;
    gap: 16px;
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
