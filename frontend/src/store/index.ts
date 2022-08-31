import { createPinia, defineStore } from 'pinia'
import { ref } from 'vue'
import { main } from '../../wailsjs/go/models'
import { GetList } from '../../wailsjs/go/main/App'

export const store = createPinia()

export const useStore = defineStore('global', () => {
  const count = ref(0)
  const darkTheme = ref(false)
  const items = ref<main.Item[]>([])
  const showModal = ref(false)
  const modalTitle = ref('')
  const formData = ref<main.Item>({
    label: '',
    value: '',
  })
  // 加载数据
  const update = async () => {
    const list = await GetList()
    items.value = list
    console.log('获取到的列表：', items.value)
  }
  update()
  return {
    count,
    darkTheme,
    items,
    update,
    showModal,
    modalTitle,
    formData,
  }
})
