<template>
  <div class="product-card">
    <Galleria
      v-if="product.images.length"
      :value="product.images"
      :num-visible="numVisible"
      :show-thumbnails="showThumbnails"
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
    <Skeleton
      v-else
      class="image"
      width="250px"
      height="250px"
    />
    <div class="content">
      <div class="wrapper">
        <div v-if="product.price" class="price">${{ product.price }}</div>
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
        @click="openSourcePost"
      />
    </div>
  </div>
</template>
<script lang="ts" setup>
import axios from 'axios'
import { computed, watch, type PropType } from 'vue'
import { Galleria, Button, Skeleton } from 'primevue'
import { format } from 'date-fns'
import type { Product } from '@/services/store/types'
import { PLATFORM } from '@/constants/store'

defineOptions({
  name: 'ProductCard',
})

const props = defineProps({
  product: {
    type: Object as PropType<Product>,
    default: () => ({}),
  },
  platform: {
    type: String,
    default: '',
  },
})

const numVisible = 4

const showThumbnailNavigators = computed(
  () => props.product.images.length > numVisible
)
const showThumbnails = computed(() => props.product.images.length > 1)
const openSourcePost = () => {
  window.open(props.product.sourcePostUrl, '_blank')
}
const getWoocommerceProductImage = async () => {
  if (!props.product.featureMedia) {
    return
  }
  const res = await axios.get(
    `${props.product.storeUrl}/wp-json/wp/v2/media/${props.product.featureMedia}`
  )
  if (res.data && res.data.source_url) {
    props.product.images = [res.data.source_url]
  }
}

watch(
  () => props.product.featureMedia,
  val => {
    if (val && props.platform === PLATFORM.WOOCOMMERCE.key) {
      getWoocommerceProductImage()
    }
  },
  { immediate: true }
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
