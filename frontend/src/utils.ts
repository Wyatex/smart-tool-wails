import { useStore } from './store'

const getCateItemIndex = (cateLabel: string, itemLabel: string) => {
  const store = useStore()
  const cateIndex = store.data.findIndex((v) => v.label === cateLabel)
  const itemIndex = store.data[cateIndex].list.findIndex(
    (v) => v.label === itemLabel
  )
  return {
    cateIndex,
    itemIndex,
  }
}

export { getCateItemIndex }
