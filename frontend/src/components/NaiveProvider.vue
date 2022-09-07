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
            <slot></slot>
            <naive-provider-content />
          </n-message-provider>
        </n-notification-provider>
      </n-dialog-provider>
    </n-loading-bar-provider>
  </n-config-provider>
</template>

<script lang="ts" setup>
import { zhCN, dateZhCN, darkTheme } from 'naive-ui'
import { useStore } from '../store'
import NaiveProviderContent from './NaiveProviderContent.vue'
const store = useStore()
</script>
