<template>
  <n-modal
    v-model:show="store.cateModal.isShow"
    class="max-w-600px"
    :segmented="{
      content: 'soft',
      footer: 'soft',
    }"
    preset="card"
    :title="store.cateModal.title"
    size="huge"
    :bordered="false"
  >
    <n-form
      label-placement="left"
      label-width="auto"
      require-mark-placement="right-hanging"
      size="large"
    >
      <n-form-item label="分类名称">
        <n-input
          v-model:value="store.cateModal.label"
          placeholder="输入显示名称"
          @keyup.enter.native="handleConfirm"
        />
      </n-form-item>
    </n-form>
    <template #footer>
      <div class="w-full flex justify-end">
        <n-button class="mr04" @click="store.cateModal.isShow = false"
          >取消</n-button
        >
        <n-button
          type="info"
          @click="handleConfirm"
          :disabled="store.cateModal.label === store.cateModal.prevLabel"
          >确定</n-button
        >
      </div>
    </template>
  </n-modal>
</template>

<script lang="ts" setup>
import { useStore } from '../store'
const store = useStore()

const showAdd = () => {
  store.cateModal.title = '添加分类'
  store.cateModal.label = ''
  store.cateModal.prevLabel = ''
  store.cateModal.isShow = true
}
const handleConfirm = () => {
  if (!store.cateModal.label) {
    window.$message.error('请输入分类名称')
    return
  }
  const isExist = store.data.some((v) => v.label === store.cateModal.label)
  if (isExist) {
    window.$message.error('该分类已存在')
    return
  }
  if (store.cateModal.title === '添加分类') {
    store.data.push({
      label: store.cateModal.label,
      list: [],
    })
    window.$notification.success({
      title: '分类：' + store.cateModal.label + '创建成功',
      duration: 3000,
    })
  } else {
    const index = store.data.findIndex(
      (v) => v.label === store.cateModal.prevLabel
    )
    store.data[index].label = store.cateModal.label
    window.$notification.success({
      title:
        '分类：' +
        store.cateModal.prevLabel +
        '修改成：' +
        store.cateModal.label,
      duration: 3000,
    })
  }
  store.cateModal.isShow = false
}
defineExpose({
  showAdd,
})
</script>
