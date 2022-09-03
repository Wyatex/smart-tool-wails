<template>
  <div>
    <n-card hoverable embedded class="w-full" @contextmenu="handleContextMenu">
      <div class="w-full flex">
        <div class="flex-1 flex items-center justify-center">
          <n-h4 class="m-0">{{ item.label }}</n-h4>
        </div>
        <n-button class="ml-4" quaternary circle @click="open(item.value)">
          <template #icon>
            <n-icon>
              <arrow />
            </n-icon>
          </template>
        </n-button>
      </div>
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
  </div>
</template>

<script lang="ts" setup>
import { useStore } from '../store'
import { Open } from '../../wailsjs/go/main/App'
import arrow from '~icons/material-symbols/arrow-forward-rounded'
const props = defineProps<{
  item: Item
  cate: string
}>()
const store = useStore()
const x = ref(0)
const y = ref(0)
const showDropdown = ref(false)
const handleContextMenu = (e: MouseEvent) => {
  e.preventDefault()
  x.value = e.clientX
  y.value = e.clientY
  showDropdown.value = true
}
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
    store.itemModal.title = '编辑快捷方式'
    store.itemModal.formData = {
      label: props.item.label,
      value: props.item.value,
    }
    store.itemModal.prevValue = props.item.value
    store.itemModal.prevLabel = props.item.label
    store.itemModal.cate = props.cate
    store.itemModal.isShow = true
  } else if (key === 'del') {
    window.$dialog.warning({
      title: '警告',
      content: '确定删除该快捷方式？',
      positiveText: '确定',
      negativeText: '不确定',
      onPositiveClick: () => {
        const cateIndex = store.data.findIndex((v) => v.label === props.cate)
        const itemIndex = store.data[cateIndex].list.findIndex(
          (v) => v.label === props.item.label
        )
        store.data[cateIndex].list.splice(itemIndex, 1)
        window.$notification.success({
          title: '删除成功',
          duration: 3000,
        })
      },
    })
  }
  showDropdown.value = false
}
const onClickoutside = () => {
  showDropdown.value = false
}
</script>
