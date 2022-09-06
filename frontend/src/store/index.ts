import { createPinia, defineStore } from 'pinia'
import { ref, watchEffect } from 'vue'

export const store = createPinia()

export const useStore = defineStore('global', () => {
  const darkTheme = ref(false)
  const data = ref<Category[]>([])
  const itemModal = ref({
    isShow: false,
    title: '',
    formData: {
      label: '',
      value: '',
    },
    prevLabel: '', // 用来对比是否已修改和是否已存在
    prevValue: '', // 用来对比是否已修改
    cate: '', // 针对哪个分类
  })
  const cateModal = ref({
    isShow: false,
    title: '',
    label: '',
    prevLabel: '',
  })
  const searchModal = ref<{
    input: string
    result: Item[]
  }>({
    input: '',
    result: [],
  })
  const screenLocked = ref(false)
  // 加载数据
  const getData = () => {
    try {
      const storageData = window.localStorage.getItem('data')
      if (!storageData) {
        data.value = []
        return
      }
      const jsonParse = JSON.parse(storageData)
      data.value = jsonParse
    } catch (e) {}
    console.log('获取到的列表：', data.value)
  }
  const getDarkTheme = () => {
    const storageDarkTheme = window.localStorage.getItem('darkTheme')
    if (storageDarkTheme) {
      if (storageDarkTheme === 'true') {
        darkTheme.value = true
      } else {
        darkTheme.value = false
      }
    } else {
      darkTheme.value = false
    }
  }
  getData()
  getDarkTheme()
  watchEffect(() => {
    // 修改data后会立刻保存
    window.localStorage.setItem('data', JSON.stringify(data.value))
    window.localStorage.setItem('dataVersion', 'v1') // 如果以后数据结构发生改变可以用这个进行自动化升级
    window.localStorage.setItem('darkTheme', darkTheme.value ? 'true' : 'false')
  })
  return {
    darkTheme,
    itemModal,
    cateModal,
    searchModal,
    data,
    screenLocked,
    getData,
  }
})
