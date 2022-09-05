<template>
  <n-card
    :title="cate.label"
    class="max-w-400px min-w-200px"
    :segmented="{
      content: 'soft',
    }"
    hoverable
  >
    <template #header-extra>
      <n-dropdown :options="cateOptions" @select="handleSelect">
        <n-button quaternary circle style="margin-left: 16px">
          <template #icon>
            <n-icon>
              <ArrowDown />
            </n-icon>
          </template>
        </n-button>
      </n-dropdown>
    </template>
    <n-space class="min-h-20px" vertical>
      <item-card
        class="w-full"
        v-for="item in cate.list"
        :item="item"
        :cate="cate.label"
      ></item-card>
    </n-space>
  </n-card>
</template>
<script lang="ts" setup>
import ItemCard from './ItemCard.vue'
import ArrowDown from '~icons/tabler/chevron-down'
import AddItem from '~icons/line-md/document-add'
import EditCate from '~icons/line-md/edit'
import DeleteCate from '~icons/line-md/close-circle'
import ClearCate from '~icons/line-md/document-remove'
import { NIcon } from 'naive-ui'
import { useStore } from '../store'
const store = useStore()
const cateOptions = [
  {
    label: '添加快捷方式',
    key: 'add',
    icon: () => h(NIcon, null, { default: () => h(AddItem) }),
  },
  {
    label: '编辑分类名',
    key: 'edit',
    icon: () => h(NIcon, null, { default: () => h(EditCate) }),
  },
  {
    label: '清空分类',
    key: 'clear',
    icon: () => h(NIcon, null, { default: () => h(ClearCate) }),
  },
  {
    label: '删除分类',
    key: 'delete',
    icon: () => h(NIcon, null, { default: () => h(DeleteCate) }),
  },
]
const handleSelect = (key: string) => {
  if (key === 'add') {
    store.itemModal.title = '添加快捷方式'
    store.itemModal.formData = {
      label: '',
      value: '',
    }
    store.itemModal.cate = props.cate.label
    store.itemModal.isShow = true
  } else if (key === 'edit') {
    store.cateModal.label = props.cate.label
    store.cateModal.prevLabel = props.cate.label
    store.cateModal.title = '修改分类'
    store.cateModal.isShow = true
  } else if (key === 'clear') {
    window.$dialog.warning({
      title: '警告',
      content: '确定清空分类列表？',
      positiveText: '确定',
      negativeText: '不确定',
      onPositiveClick: () => {
        const index = store.data.findIndex((v) => v.label === props.cate.label)
        store.data[index].list = []
        window.$notification.success({
          title: '清空成功',
          duration: 3000,
        })
      },
    })
  } else if (key === 'delete') {
    window.$dialog.warning({
      title: '警告',
      content: '确定删除分类列表？',
      positiveText: '确定',
      negativeText: '不确定',
      onPositiveClick: () => {
        const index = store.data.findIndex((v) => v.label === props.cate.label)
        store.data.splice(index, 1)
        window.$notification.success({
          title: '分类 ' + props.cate.label + ' 删除成功',
          duration: 3000,
        })
      },
    })
  }
}

const props = defineProps<{
  cate: Category
}>()
</script>
