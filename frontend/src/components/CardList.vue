<template>
  <n-space>
    <n-card
      v-for="item in store.items"
      :key="item.label"
      :title="item.label"
      hoverable
      class="min-w-200px max-w-600px"
      @contextmenu="(e: MouseEvent) => handleContextMenu(e, item)"
    >
      <template #header-extra>
        <n-button quaternary circle class="ml-8" @click="open(item.label)">
          <template #icon>
            <n-icon>
              <arrow />
            </n-icon>
          </template>
        </n-button>
      </template>
    </n-card>
    <n-dropdown
      placement="bottom-start"
      trigger="manual"
      :x="x"
      :y="y"
      :options="options"
      :show="showDropdown"
      :on-clickoutside="onClickoutside"
      @select="handleSelect"
    />
  </n-space>
</template>

<script lang="ts" setup>
import { useStore } from '../store'
import arrow from '~icons/material-symbols/arrow-forward-rounded'
import { Open, Del } from '../../wailsjs/go/main/App'
import { main } from '../../wailsjs/go/models'
const store = useStore()
const x = ref(0)
const y = ref(0)
const contextValue = ref<main.Item>({
  label: '',
  value: '',
})
const showDropdown = ref(false)
const open = (name: string) => {
  Open(name).then((result) => {
    if (result) {
      window.$message.error(result)
    } else {
      window.$message.success('打开成功')
    }
  })
}
const options = [
  {
    label: '编辑',
    key: 'edit',
  },
  {
    label: '删除',
    key: 'del',
  },
]
const handleSelect = (key: string) => {
  if (key === 'edit') {
    store.formData = {...contextValue.value}
    store.modalTitle = '编辑快捷方式'
    store.showModal = true
  } else {
    Del(contextValue.value.label).then((result) => {
      if (result) {
        window.$message.error(result)
      } else {
        window.$message.success('删除成功')
        store.update()
      }
    })
  }
  showDropdown.value = false
}
const handleContextMenu = (e: MouseEvent, label: main.Item) => {
  contextValue.value = label
  e.preventDefault()
  showDropdown.value = false
  nextTick().then(() => {
    showDropdown.value = true
    x.value = e.clientX
    y.value = e.clientY
  })
}
const onClickoutside = () => {
  showDropdown.value = false
}
</script>
