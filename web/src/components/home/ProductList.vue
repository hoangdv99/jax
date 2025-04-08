<template>
  <div ref="list" class="product-list" @scroll="onScroll">
    <ProductCard
      v-if="props.platform"
      v-for="product in props.products"
      :key="product.id"
      :product="product"
      :platform="props.platform"
    />
    <div v-if="!props.products.length || !props.platform" class="notice">
      {{ noticeMessage }}
    </div>
  </div>
</template>
<script lang="ts" setup>
import { computed } from 'vue'
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

const noticeMessage = computed(() => {
  if (!props.platform) {
    return 'Please select a store to view products.'
  }
  if (!props.products.length) {
    return 'No products found.'
  }
})

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
  > .notice {
    width: 100%;
    text-align: center;
    font-size: 24px;
    font-weight: 600;
  }
}
</style>
