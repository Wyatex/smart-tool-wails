<template>
  <n-modal
    v-model:show="store.itemModal.isShow"
    class="max-w-600px"
    :segmented="{
      content: 'soft',
      footer: 'soft',
    }"
    preset="card"
    :title="store.itemModal.title"
    size="huge"
    :bordered="false"
  >
    <n-form
      ref="formRef"
      :model="store.itemModal.formData"
      :rules="rules"
      label-placement="left"
      label-width="auto"
      require-mark-placement="right-hanging"
      size="large"
    >
      <n-form-item label="显示名称" path="label">
        <n-input
          v-model:value="store.itemModal.formData.label"
          placeholder="输入显示名称"
          @keyup.enter.native="handleConfirm"
        />
      </n-form-item>
      <n-form-item label="路径" path="value">
        <n-input
          v-model:value="store.itemModal.formData.value"
          placeholder="输入显示名称"
          @keyup.enter.native="handleConfirm"
        />
      </n-form-item>
    </n-form>
    <template #footer>
      <div class="w-full flex justify-end">
        <n-button class="mr04" @click="store.itemModal.isShow = false"
          >取消</n-button
        >
        <n-button
          type="info"
          @click="handleConfirm"
          :disabled="disabledComputed"
          >确定</n-button
        >
      </div>
    </template>
  </n-modal>
</template>

<script lang="ts" setup>
import { useStore } from '../store'
import { NForm } from 'naive-ui'
const store = useStore()
const formRef = ref<InstanceType<typeof NForm>>()
const rules = ref({
  label: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入 快捷名称',
  },
  value: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入 路径',
  },
})
const handleConfirm = () => {
  formRef.value?.validate(async (errors) => {
    if (errors) {
      window.$message.error('请检查快捷方式名称和路径')
      return
    }
    if (store.itemModal.title === '添加快捷方式') {
      const index = store.data.findIndex(
        (v) => v.label === store.itemModal.cate
      )
      const isExist = store.data[index].list.some(
        (v) => v.label === store.itemModal.formData.label
      )
      if (isExist) {
        window.$message.error('该快捷方式已存在')
        return
      }
      store.data[index].list.push(store.itemModal.formData)
      window.$notification.success({
        title: '添加成功',
        duration: 3000,
      })
      store.itemModal.isShow = false
    } else {
      const cateIndex = store.data.findIndex(
        (v) => v.label === store.itemModal.cate
      )
      if (store.itemModal.prevLabel !== store.itemModal.formData.label) {
        // 修改名称，需要判断是否已存在
        const isExist = store.data[cateIndex].list.some(
          (v) => v.label === store.itemModal.formData.label
        )
        if (isExist) {
          window.$message.error('该快捷方式已存在')
          return
        }
      }
      const itemIndex = store.data[cateIndex].list.findIndex(
        (v) => v.label === store.itemModal.prevLabel
      )
      store.data[cateIndex].list[itemIndex].label =
        store.itemModal.formData.label
      store.data[cateIndex].list[itemIndex].value =
        store.itemModal.formData.value
      window.$notification.success({
        title: '修改成功',
        duration: 3000,
      })
      store.itemModal.isShow = false
    }
  })
}
const disabledComputed = computed(() => {
  if (store.itemModal.title === '添加快捷方式') {
    if (store.itemModal.formData.label === store.itemModal.prevLabel) {
      return true
    }
    return false
  } else {
    if (
      store.itemModal.formData.label === store.itemModal.prevLabel &&
      store.itemModal.formData.value === store.itemModal.prevValue
    ) {
      return true
    }
    return false
  }
})
</script>
