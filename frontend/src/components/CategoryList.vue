<template>
  <n-layout class="h-full w-100vw flex">
    <n-layout-header
      class="fixed top-0 left-0 flex justify-between items-center h-64px shadow-lg z-10 pr-8"
      :class="{ 'shadow-gray-600': store.darkTheme }"
    >
      <n-h2 prefix="bar" class="m-0 ml-4">
        <n-text type="primary"> 一个简单的快捷工具 </n-text>
      </n-h2>
      <n-popover placement="left" trigger="hover">
        <template #trigger>
          <n-icon>
            <i-ri-question-line class="text-20px" />
          </n-icon>
        </template>
        <n-p>点击右下角加号按钮添加分类，右键卡片编辑和删除</n-p>
      </n-popover>
    </n-layout-header>
    <n-scrollbar>
      <n-layout-content class="flex-1 flex justify-center mt-64px mb-2 p-8">
        <n-space>
          <category-card
            v-for="item in store.data"
            :key="item.label"
            :cate="item"
          />
        </n-space>
      </n-layout-content>
    </n-scrollbar>
    <n-layout-footer
      class="flex items-center justify-center fixed bottom-0 w-full"
    >
      <div class="flex items-center justify-center p-1">
        <n-p>Copyright ©2022 Wyatex</n-p>
      </div>
    </n-layout-footer>

    <cate-modal />
    <item-modal />
    <float-btn />
  </n-layout>
</template>

<script lang="ts" setup>
import { useStore } from '../store'
import CategoryCard from './CategoryCard.vue'
import CateModal from './CateModal.vue'
import ItemModal from './ItemModal.vue'
import FloatBtn from './FloatBtn.vue'
const store = useStore()
</script>
