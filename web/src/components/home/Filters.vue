<template>
  <div class="filters-container">
    <div class="item">
      <label class="label">Store Url</label>
      <Select
        :model-value="props.selectedStore"
        :options="props.stores"
        filter
        input-id="store"
        optionLabel="url"
        placeholder="Select a store"
        class="select"
        @update:model-value="emit('updateSelectedStore', $event)"
      />
    </div>
    <div class="item">
      <label class="label">Collection</label>
      <Select
        :model-value="props.selectedCollection"
        :options="props.collections"
        input-id="collection"
        option-label="title"
        class="select"
        @show="onShowDropdown"
        @update:model-value="emit('updateSelectedCollection', $event)"
      />
    </div>
    <div class="item">
      <label class="label">Tags</label>
      <MultiSelect
        v-model="props.selectedTags"
        :options="props.tags"
        :show-toggle-all="false"
        display="chip"
        input-id="tag"
        :maxSelectedLabels="2"
        optionLabel="name"
        placeholder="Select tags"
        class="select"
      />
    </div>
  </div>
</template>
<script setup lang="ts">
import { nextTick, type PropType } from 'vue'
import { Select, MultiSelect } from 'primevue'
import type { Collection, Store, Tag } from '@/services/store/types'

defineOptions({
  name: 'Filters',
})

const emit = defineEmits([
  'getListCollection',
  'updateCollectionPage',
  'updateSelectedStore',
  'updateSelectedCollection',
])

const props = defineProps({
  stores: {
    type: Array,
    default: () => [],
  },
  collections: {
    type: Array,
    default: () => [],
  },
  tags: {
    type: Array,
    default: () => [],
  },
  selectedStore: {
    type: Object as PropType<Store>,
  },
  selectedTags: {
    type: Array as PropType<Tag[]>,
    default: () => [],
  },
  selectedCollection: {
    type: Object as PropType<Collection>,
  },
  collectionPage: {
    type: Number,
    default: 1,
  },
  allCollectionsFetched: {
    type: Boolean,
    default: false,
  },
})

function onShowDropdown() {
  nextTick(() => {
    const panel = window.document.querySelector('.p-select-list-container')
    if (panel) {
      panel.addEventListener('scroll', onScrollToEnd)
    }
  })
}

function onScrollToEnd(event: Event) {
  const { scrollTop, scrollHeight, clientHeight } = event.target as HTMLElement
  if (
    scrollTop + clientHeight >= scrollHeight &&
    !props.allCollectionsFetched &&
    props.selectedStore?.id
  ) {
    emit('getListCollection', props.selectedStore.id, {
      page: props.collectionPage + 1,
    })
    emit('updateCollectionPage', props.collectionPage + 1)
  }
}
</script>
<style lang="scss" scoped>
.filters-container {
  display: flex;
  gap: 16px;
  background-color: #fff;
  padding: 16px;
  border-radius: 4px;
  > .item {
    display: flex;
    flex-direction: column;
  }
  > .item .select {
    width: 250px;
  }
  > .item > .label {
    font-weight: 700;
    margin-bottom: 8px;
  }
}
</style>
