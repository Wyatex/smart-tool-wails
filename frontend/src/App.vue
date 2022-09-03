<script lang="ts" setup>
import NaiveProvider from './components/NaiveProvider.vue'
import { useStore } from './store'
import CategoryList from './components/CategoryList.vue'
const store = useStore()

// import CardList from './components/CardList.vue';
// setTimeout(() => {
//   store.data = [
//     {
//       label: '123',
//       list: [],
//     },
//   ]
// }, 1000)
</script>

<template>
  <naive-provider class="h-full">
    <!-- <CardList /> -->
    <category-list />
  </naive-provider>
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
