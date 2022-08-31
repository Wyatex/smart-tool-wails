<script lang="ts" setup>
import { reactive } from 'vue'
import { Greet, Add } from '../../wailsjs/go/main/App'
import { useStore } from '../store'
const store = useStore()

const data = reactive({
  name: '',
  resultText: 'Please enter your name below 👇',
})

function greet() {
  store.count++
  Greet(data.name).then((result) => {
    data.resultText = result
  })
}
</script>

<template>
  <main>
    <div id="result" class="result">
      <n-p>{{ data.resultText }} - count: {{ store.count }}</n-p>
    </div>
    <div id="input" class="input-box">
      <n-input
        v-model:value="data.name"
        placeholder="请输入"
        autocomplete="off"
        class="max-w-200px"
      />
      <n-button type="info" class="ml-2" @click="greet">Greet</n-button>
    </div>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}
</style>
