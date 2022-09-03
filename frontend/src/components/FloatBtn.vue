<template>
  <div class="fixed flex flex-col bottom-8 right-10">
    <n-tooltip placement="top-end" trigger="hover">
      <template #trigger>
        <n-button class="mb-2" circle size="large" @click="handleExport">
          <template #icon>
            <n-icon>
              <export-icon />
            </n-icon>
          </template>
        </n-button>
      </template>
      导出到文件
    </n-tooltip>
    <n-tooltip placement="top-end" trigger="hover">
      <template #trigger>
        <n-button class="mb-2" circle size="large" @click="handleImport">
          <template #icon>
            <n-icon>
              <import-icon />
            </n-icon>
          </template>
        </n-button>
      </template>
      从文件导入
    </n-tooltip>
    <n-tooltip placement="top-end" trigger="hover">
      <template #trigger>
        <n-button class="mb-2" circle size="large" @click="clearAll">
          <template #icon>
            <n-icon>
              <clear-icon />
            </n-icon>
          </template>
        </n-button>
      </template>
      清除所有数据
    </n-tooltip>
    <n-tooltip placement="top-end" trigger="hover">
      <template #trigger>
        <n-button class="mb-2" circle size="large" @click="handleAddCate">
          <template #icon>
            <n-icon>
              <add />
            </n-icon>
          </template>
        </n-button>
      </template>
      添加一个分类
    </n-tooltip>
    <n-tooltip placement="top-end" trigger="hover">
      <template #trigger>
        <n-button class="mb-2" circle size="large" @click="themeSwitch">
          <template #icon>
            <n-icon>
              <sun v-if="!store.darkTheme" />
              <moon v-else />
            </n-icon>
          </template>
        </n-button>
      </template>
      {{ store.darkTheme ? '切换到浅色模式' : '切换到深色模式' }}
    </n-tooltip>
  </div>
</template>
<script lang="ts" setup>
import { useStore } from '../store'
import sun from '~icons/line-md/sun-rising-loop'
import moon from '~icons/line-md/moon'
import add from '~icons/carbon/category-new'
import exportIcon from '~icons/uil/file-import'
import importIcon from '~icons/uil/file-export'
import clearIcon from '~icons/icon-park-outline/clear'
import { Load, Save } from '../../wailsjs/go/main/App'
const store = useStore()
const handleAddCate = () => {
  store.cateModal.title = '添加分类'
  store.cateModal.label = ''
  store.cateModal.prevLabel = ''
  store.cateModal.isShow = true
}
const themeSwitch = () => {
  store.darkTheme = !store.darkTheme
  window.$loading.start()
  setTimeout(() => {
    window.$loading.finish()
  }, 100)
}
const clearAll = () => {
  window.$dialog.warning({
    title: '警告',
    content: '确定删除所有数据？建议先点击上方按钮导出备份',
    positiveText: '确定',
    negativeText: '不确定',
    onPositiveClick: () => {
      store.data = []
      window.$notification.success({
        title: '清空成功',
        duration: 3000,
      })
    },
  })
}
const handleExport = () => {
  Save(JSON.stringify(store.data)).then((err) => {
    if (err === '') {
      window.$notification.success({
        title: '导出成功',
        duration: 3000,
      })
    } else {
      window.$notification.error({
        title: '导出失败',
        content: '原因：' + err,
      })
    }
  })
}
const handleImport = () => {
  Load().then((returnData) => {
    console.log(returnData)
    if (returnData.err !== '') {
      window.$notification.error({
        title: '导入失败',
        content: '原因：' + returnData.err,
      })
      return
    }
    try {
      const data = JSON.parse(returnData.data)
      store.data = data
      window.$notification.success({
        title: '导入成功',
        duration: 3000,
      })
    } catch (e) {
      window.$notification.error({
        title: '数据异常',
        content: '原因：' + e,
      })
    }
  })
}
</script>
