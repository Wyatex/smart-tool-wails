import { ref, onMounted, onUnmounted } from 'vue'
import axios from 'axios'

/**
 * @description 获取本地时间
 */
export function useHitokoto() {
  let timer: NodeJS.Timeout
  const word = ref('')
  const from = ref('')
  const fromWho = ref('')

  // 更新时间
  const updateWord = async () => {
    const res = await axios.get('https://v1.hitokoto.cn/?encode=json')
    if (res.status === 200) {
      word.value = res.data.hitokoto
      from.value = res.data.from
      fromWho.value = res.data.from_who
    }
    console.log(res)
  }

  updateWord()

  onMounted(() => {
    clearInterval(timer)
    timer = setInterval(() => updateWord(), 7000)
  })

  onUnmounted(() => {
    clearInterval(timer)
  })

  return {
    word,
    from,
    fromWho,
  }
}
