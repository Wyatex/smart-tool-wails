<template>
  <n-modal
    v-model:show="store.showModal"
    class="max-w-600px"
    :segmented="{
      content: 'soft',
      footer: 'soft',
    }"
    preset="card"
    :title="store.modalTitle"
    size="huge"
    :bordered="false"
  >
    <n-form
      ref="formRef"
      :model="store.formData"
      :rules="rules"
      label-placement="left"
      label-width="auto"
      require-mark-placement="right-hanging"
      size="large"
    >
      <n-form-item label="显示名称" path="label">
        <n-input
          v-model:value="store.formData.label"
          placeholder="输入显示名称"
          :disabled="store.modalTitle === '编辑快捷方式'"
          @keyup.enter.native="handleConfirm"
        />
      </n-form-item>
      <n-form-item label="路径" path="value">
        <n-input
          v-model:value="store.formData.value"
          placeholder="输入显示名称"
          @keyup.enter.native="handleConfirm"
        />
      </n-form-item>
    </n-form>
    <template #footer>
      <div class="w-full flex justify-end">
        <n-button class="mr04" @click="store.showModal = false">取消</n-button>
        <n-button type="info" @click="handleConfirm">确定</n-button>
      </div>
    </template>
  </n-modal>
</template>

<script lang="ts" setup>
import { main } from '../../wailsjs/go/models'
import { Add, Edit } from '../../wailsjs/go/main/App'
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
const showAdd = () => {
  store.modalTitle = '添加快捷方式'
  store.formData = {
    label: '',
    value: '',
  }
  store.showModal = true
}
const handleConfirm = () => {
  formRef.value?.validate(async (errors) => {
    if (errors) {
      window.$message.error('请检查快捷方式名称和路径')
      return
    }
    if (store.modalTitle === '添加快捷方式') {
      const res = await Add(store.formData)
      if (res !== '') {
        window.$message.error(res)
        return
      }
      window.$notification.success({
        title: '添加成功',
        duration: 3000,
      })
      await store.update()
      store.showModal = false
    } else {
      const res = await Edit(store.formData)
      if (res !== '') {
        window.$message.error(res)
        return
      }
      window.$notification.success({
        title: '修改成功',
        duration: 3000,
      })
      await store.update()
      store.showModal = false
    }
  })
}
defineExpose({
  showAdd,
})
</script>
