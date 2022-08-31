<script lang="ts" setup>
import NaiveProvider from './components/NaiveProvider.vue'
import CardList from './components/CardList.vue';
</script>

<template>
  <naive-provider class="h-full">
    <CardList />
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
