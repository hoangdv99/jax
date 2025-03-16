<template>
  <div class="filters-container">
    <div class="item">
      <label for="store" class="label">Store Url</label>
      <Select
        v-model="selectedStore"
        :options="storeSelections"
        filter
        input-id="store"
        optionLabel="url"
        placeholder="Select a store"
        class="select"
      />
    </div>
    <div class="item">
      <label for="tag" class="label">Tags</label>
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
import { computed, ref } from 'vue'
import { Select, MultiSelect } from 'primevue'

defineOptions({
  name: 'Filters',
})

const props = defineProps({
  stores: {
    type: Array,
    default: () => [],
  },
  tags: {
    type: Array,
    default: () => [],
  },
})

const storeSelections = computed(() => {
  return [
    {
      id: null,
      url: 'All',
      platform: null,
      tags: [],
    },
    ...props.stores,
  ]
})

const selectedStore = ref(storeSelections.value[0])
const selectedTags = ref([])
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
