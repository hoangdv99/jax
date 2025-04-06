<template>
  <div class="filters-container">
    <div class="item">
      <label class="label">Store Url</label>
      <Select
        v-model="selectedStore"
        :options="props.stores"
        filter
        input-id="store"
        optionLabel="url"
        placeholder="Select a store"
        class="select"
      />
    </div>
    <div class="item">
      <label class="label">Collection</label>
      <Select
        v-model="selectedCollection"
        :options="props.collections"
        input-id="collection"
        option-label="title"
        class="select"
        @show="onShowDropdown"
      />
    </div>
    <div class="item">
      <label class="label">Tags</label>
      <MultiSelect
        v-model="selectedTags"
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
import { nextTick, ref, watch } from 'vue'
import { Select, MultiSelect } from 'primevue'
import type { Collection, Store } from '@/services/store/types'
import { useStoreStore } from '@/stores/store'
import { PRODUCT_PAGE_LIMIT } from '@/constants/store'

defineOptions({
  name: 'Filters',
})

const emit = defineEmits([
  'getListCollection',
  'getCollectionProducts',
  'getProducts',
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
})

const storeStore = useStoreStore()

const selectedStore = ref<Store | null>(props.stores[0] as Store | null)
const selectedTags = ref([])
const selectedCollection = ref<Collection | null>(
  props.collections[0] as Collection | null
)
const collectionPage = ref(1)

watch(
  () => props.stores,
  newStores => {
    if (newStores.length) {
      selectedStore.value = newStores[0] as Store
    } else {
      selectedStore.value = null
    }
  },
  { immediate: true }
)

watch(
  () => props.collections,
  newCollections => {
    if (newCollections.length) {
      selectedCollection.value = newCollections[0] as Collection
    } else {
      selectedCollection.value = null
    }
  },
  { immediate: true }
)

watch(selectedStore, (val: Store | null) => {
  if (val) {
    collectionPage.value = 1
    storeStore.resetAllCollectionsFetchedFlag()
    emit('getListCollection', val.id)
  }
})

watch(selectedCollection, (val: Collection | null) => {
  if (val?.handle && selectedStore.value) {
    emit('getCollectionProducts', {
      storeId: selectedStore.value.id,
      collectionId: val.id,
      handle: val.handle,
      page: 1,
      limit: PRODUCT_PAGE_LIMIT,
    })
  } else if (val?.title === 'All' && selectedStore.value) {
    emit('getProducts', {
      storeIds: selectedStore.value.id,
      page: 1,
      limit: PRODUCT_PAGE_LIMIT,
    })
  }
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
    scrollTop + clientHeight >= scrollHeight - 3 &&
    !storeStore.allCollectionsFetched
  ) {
    emit('getListCollection', selectedStore.value?.id, {
      page: collectionPage.value + 1,
    })
    collectionPage.value += 1
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
