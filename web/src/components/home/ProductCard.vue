<template>
  <div class="product-card">
    <Galleria
      :value="product.images"
      :num-visible="numVisible"
      :show-thumbnail-navigators="showThumbnailNavigators"
      class="galleria"
    >
      <template #item="slotProps">
        <img :src="slotProps.item" class="image" />
      </template>
      <template #thumbnail="slotProps">
        <img :src="slotProps.item" class="thumb" />
      </template>
    </Galleria>
    <div class="content">
      <div class="wrapper">
        <div class="price">${{ product.price }}</div>
        <div class="date">
          {{ format(new Date(product.createdDate), 'yyyy-MM-dd') }}
        </div>
      </div>
      <Button
        label="View"
        icon="pi pi-arrow-right"
        iconPos="right"
        size="small"
        variant="text"
      />
    </div>
  </div>
</template>
<script lang="ts" setup>
import { computed, type PropType } from 'vue'
import { Galleria, Button } from 'primevue'
import { format } from 'date-fns'
import type { Product } from '@/services/store/types'

defineOptions({
  name: 'ProductCard',
})

const props = defineProps({
  product: {
    type: Object as PropType<Product>,
    default: () => ({}),
  },
})

const numVisible = 4

const showThumbnailNavigators = computed(
  () => props.product.images.length > numVisible
)
</script>
<style lang="scss" scoped>
.product-card {
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
  width: 250px;
  :deep(.p-galleria-thumbnails-viewport) {
    width: fit-content;
  }
  :deep(.p-galleria-thumbnail-items) {
    gap: 8px;
  }
  .image {
    width: 250px;
    height: 250px;
  }
  .thumb {
    width: 40px;
    height: 40px;
  }
  :deep(.p-galleria) {
    border: none;
    border-radius: 0;
  }
  > .content {
    padding: 8px;
    background-color: #fff;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  > .content > .wrapper > .price {
    margin-bottom: 4px;
    font-weight: 600;
  }
}
</style>
