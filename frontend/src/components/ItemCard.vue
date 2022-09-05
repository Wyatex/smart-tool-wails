<template>
  <div>
    <n-card
      hoverable
      embedded
      class="w-full cursor-pointer"
      @contextmenu="handleContextMenu"
      @click="open(item.value)"
    >
      <div class="w-full h-full flex-1 flex items-center justify-center">
        <n-h4 class="m-0">{{ item.label }}</n-h4>
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
import { NIcon } from 'naive-ui'
import editIcon from '~icons/line-md/edit'
import delIcon from '~icons/line-md/close-circle'
import leftIcon from '~icons/line-md/chevron-small-left'
import leftDoubleIcon from '~icons/line-md/chevron-small-double-left'
import { getCateItemIndex } from '../utils'
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
const isTop = () => {
  const indexs = getCateItemIndex(props.cate, props.item.label)
  return indexs.itemIndex === 0
}
const isBottom = () => {
  const indexs = getCateItemIndex(props.cate, props.item.label)
  return indexs.itemIndex === store.data[indexs.cateIndex].list.length - 1
}
const options = [
  {
    label: '编辑',
    key: 'edit',
    icon: () => h(NIcon, null, { default: () => h(editIcon) }),
  },
  {
    label: '删除',
    key: 'del',
    icon: () => h(NIcon, null, { default: () => h(delIcon) }),
  },
  {
    label: '上移',
    key: 'up',
    disabled: isTop(),
    icon: () => h(NIcon, {class: 'rotate-90'}, { default: () => h(leftIcon) }),
  },
  {
    label: '置顶',
    key: 'top',
    disabled: isTop(),
    icon: () => h(NIcon, {class: 'rotate-90'}, { default: () => h(leftDoubleIcon) }),
  },
  {
    label: '下移',
    key: 'down',
    disabled: isBottom(),
    icon: () => h(NIcon, {class: 'rotate-270'}, { default: () => h(leftIcon) }),
  },
  {
    label: '置底',
    key: 'bottom',
    disabled: isBottom(),
    icon: () => h(NIcon, {class: 'rotate-270'}, { default: () => h(leftDoubleIcon) }),
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
        const indexs = getCateItemIndex(props.cate, props.item.label)
        store.data[indexs.cateIndex].list.splice(indexs.itemIndex, 1)
        window.$notification.success({
          title: '删除成功',
          duration: 3000,
        })
      },
    })
  } else if (key === 'up') {
    const indexs = getCateItemIndex(props.cate, props.item.label)
    const temp = store.data[indexs.cateIndex].list[indexs.itemIndex]
    store.data[indexs.cateIndex].list[indexs.itemIndex] =
      store.data[indexs.cateIndex].list[indexs.itemIndex - 1]
    store.data[indexs.cateIndex].list[indexs.itemIndex - 1] = temp
  } else if (key === 'top') {
    const indexs = getCateItemIndex(props.cate, props.item.label)
    const temp = store.data[indexs.cateIndex].list[indexs.itemIndex]
    store.data[indexs.cateIndex].list[indexs.itemIndex] =
      store.data[indexs.cateIndex].list[0]
    store.data[indexs.cateIndex].list[0] = temp
  } else if (key === 'down') {
    const indexs = getCateItemIndex(props.cate, props.item.label)
    const temp = store.data[indexs.cateIndex].list[indexs.itemIndex]
    store.data[indexs.cateIndex].list[indexs.itemIndex] =
      store.data[indexs.cateIndex].list[indexs.itemIndex + 1]
    store.data[indexs.cateIndex].list[indexs.itemIndex + 1] = temp
  } else if (key === 'bottom') {
    const indexs = getCateItemIndex(props.cate, props.item.label)
    const temp = store.data[indexs.cateIndex].list[indexs.itemIndex]
    store.data[indexs.cateIndex].list[indexs.itemIndex] =
      store.data[indexs.cateIndex].list[
        store.data[indexs.cateIndex].list.length - 1
      ]
    store.data[indexs.cateIndex].list[
      store.data[indexs.cateIndex].list.length - 1
    ] = temp
  }
  showDropdown.value = false
}
const onClickoutside = () => {
  showDropdown.value = false
}
</script>
