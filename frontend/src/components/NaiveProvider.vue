<template>
  <n-config-provider
    :locale="zhCN"
    :date-locale="dateZhCN"
    :theme="store.darkTheme ? darkTheme : null"
  >
    <n-loading-bar-provider>
      <n-dialog-provider>
        <n-notification-provider>
          <n-message-provider>
            <n-layout class="h-full w-100vw flex">
              <n-layout-header
                class="fixed top-0 left-0 flex justify-between items-center h-80px shadow-lg z-10 pr-8"
                :class="{ 'shadow-gray-600': store.darkTheme }"
              >
                <n-h1 prefix="bar" class="m-0 ml-4">
                  <n-text type="primary"> 一个简单的快捷工具 </n-text>
                </n-h1>
                <n-popover placement="left" trigger="hover">
                  <template #trigger>
                    <n-icon>
                      <question class="text-20px" />
                    </n-icon>
                  </template>
                  <n-p>点击右下角加号按钮添加，右键卡片编辑和删除</n-p>
                </n-popover>
              </n-layout-header>
              <n-layout-content class="flex-1 flex justify-center p-8 pt-95px">
                <slot></slot>
                <n-button
                  class="fixed bottom-25 right-10"
                  circle
                  size="large"
                  @click="modelRef?.showAdd"
                >
                  <template #icon>
                    <n-icon>
                      <add />
                    </n-icon>
                  </template>
                </n-button>
                <n-button
                  class="fixed bottom-10 right-10"
                  circle
                  size="large"
                  @click="themeSwitch"
                >
                  <template #icon>
                    <n-icon>
                      <sun v-if="!store.darkTheme" />
                      <moon v-else />
                    </n-icon>
                  </template>
                </n-button>
                <n-layout-footer
                  class="flex items-center justify-center fixed bottom-0 w-full"
                >
                  <n-p>Copyright ©2022 Wyatex</n-p>
                </n-layout-footer>
              </n-layout-content>
            </n-layout>
            <naive-provider-content />
          </n-message-provider>
        </n-notification-provider>
      </n-dialog-provider>
    </n-loading-bar-provider>
    <item-model ref="modelRef" />
  </n-config-provider>
</template>

<script lang="ts" setup>
import { zhCN, dateZhCN, darkTheme } from 'naive-ui'
import { useStore } from '../store'
import sun from '~icons/line-md/sun-rising-loop'
import moon from '~icons/line-md/moon'
import add from '~icons/icon-park-outline/list-add'
import question from '~icons/ri/question-line'
import NaiveProviderContent from './NaiveProviderContent.vue'
import ItemModel from './ItemModel.vue'
const modelRef = ref<InstanceType<typeof ItemModel> | null>(null)
const store = useStore()
const themeSwitch = () => {
  store.darkTheme = !store.darkTheme
  window.$loading.start()
  setTimeout(() => {
    window.$loading.finish()
  }, 100)
}
</script>
