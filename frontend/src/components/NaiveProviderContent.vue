<template>
  <div></div>
</template>

<script lang="ts" setup>
import { useLoadingBar, useDialog, useMessage, useNotification } from 'naive-ui'

// 挂载naive组件的方法至window
function registerNaiveTools() {
  window.$loading = useLoadingBar()
  window.$dialog = useDialog()
  window.$message = useMessage()
  window.$notification = useNotification()
}

registerNaiveTools()
</script>