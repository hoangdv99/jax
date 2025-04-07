<template>
  <div ref="list" class="product-list" @scroll="onScroll">
    <ProductCard
      v-for="product in props.products"
      :key="product.id"
      :product="product"
      :platform="props.platform"
    />
  </div>
</template>
<script lang="ts" setup>
import type { PropType } from 'vue'
import type { Product } from '@/services/store/types'
import ProductCard from './ProductCard.vue'

defineOptions({
  name: 'ProductList',
})

const props = defineProps({
  products: {
    type: Array as PropType<Product[]>,
    default: () => [],
  },
  platform: {
    type: String,
    default: '',
  },
})

const emit = defineEmits(['loadMoreProducts'])

const onScroll = (event: Event) => {
  const scrollContainer = event.target as HTMLDivElement
  const isBottom =
    scrollContainer.scrollTop + scrollContainer.clientHeight >=
    scrollContainer.scrollHeight
  if (isBottom) {
    emit('loadMoreProducts')
  }
}
</script>
<style lang="scss" scoped>
.product-list {
  display: flex;
  gap: 16px;
  flex-wrap: wrap;
  max-height: calc(100vh - 140px);
  overflow: auto;
}
</style>
