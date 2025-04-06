<template>
  <div class="product-list">
    <ProductCard
      v-for="product in props.products"
      :key="product.id"
      :product="product"
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
})
</script>
<style lang="scss" scoped>
.product-list {
  display: flex;
  gap: 16px;
  flex-wrap: wrap;
}
</style>
