<template>
  <n-space>
    <category-card v-for="item in store.data" :key="item.label" :cate="item" />
  </n-space>
</template>

<script lang="ts" setup>
import { useStore } from '../store'
import CategoryCard from './CategoryCard.vue';
const store = useStore()
</script>
