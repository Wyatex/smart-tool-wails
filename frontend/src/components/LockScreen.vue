<template>
  <n-layout>
    <div class="w-100vw h-100vh relative">
      <div class="bg absolute w-100vw h-100vh"></div>
      <div
        class="absolute flex justify-between h-40px w-90px top-10vh left-[50%] translate-x-[-50%] cursor-pointer"
      >
        <n-tooltip placement="bottom" trigger="hover">
          <template #trigger>
            <div class="w-40px h-40px" @click="store.screenLocked = false">
              <n-icon class="w-40px h-40px">
                <i-material-symbols-lock-open-outline-rounded
                  class="text-40px"
                />
              </n-icon>
            </div>
          </template>
          解锁
        </n-tooltip>
        <n-tooltip placement="bottom" trigger="hover">
          <template #trigger>
            <div
              class="w-40px h-40px"
              @click="store.fullscreen = !store.fullscreen"
            >
              <n-icon class="w-40px h-40px">
                <i-material-symbols-fullscreen-rounded
                  class="text-40px"
                  v-if="!store.fullscreen"
                />
                <i-material-symbols-fullscreen-exit-rounded
                  class="text-40px"
                  v-else
                />
              </n-icon>
            </div>
          </template>
          {{ store.fullscreen ? '取消全屏' : '全屏' }}
        </n-tooltip>
      </div>

      <div
        class="hitokoto absolute left-50vw top-50vh translate-x-[-50%] translate-y-[-50%]"
      >
        <div class="relative cursor-pointer" @click="handleCopy">
          <div class="absolute left-0 top-0 text-30px">『</div>
          <div class="text-30px py-15px px-50px">{{ word }}</div>
          <div class="absolute right-0 bottom-0 text-30px">』</div>
        </div>
        <div class="float-right text-20px mt-20px">
          —— {{ fromWho }}「{{ from }}」
        </div>
      </div>
      <div class="absolute bottom-60px left-60px">
        <div class="text-70px text-left">{{ hour }}:{{ minute }}</div>
        <div class="text-40px text-left">{{ month }}月{{ day }}号，星期{{ week }}</div>
      </div>
    </div>
  </n-layout>
</template>
<script lang="ts" setup>
import { useTime } from '../hooks/useTime'
import { useStore } from '../store'
import { useHitokoto } from '../hooks/useHitokoto'

const store = useStore()
const { month, day, hour, minute, week } = useTime()
const { word, from, fromWho } = useHitokoto()
const handleCopy = () => {
  navigator.clipboard.writeText(`${word} —— ${fromWho}「${from}」`).then(() => {
    window.$message.success('已复制到剪贴板')
  })
}
const bgc = computed(() => {
  return store.darkTheme ? '#18181c' : '#e493d0'
})
const bgi = computed(() => {
  return store.darkTheme
    ? `radial-gradient(
      closest-side,
      rgba(51, 8, 103, 0.8),
      rgba(235, 105, 78, 0)
    ),
    radial-gradient(closest-side, rgba(255, 209, 253, 0.2), rgba(243, 11, 164, 0)),
    radial-gradient(
      closest-side,
      rgba(244, 246, 138, 0.1),
      rgba(254, 234, 131, 0)
    ),
    radial-gradient(
      closest-side,
      rgba(235, 188, 170, 0.2),
      rgba(170, 142, 245, 0)
    ),
    radial-gradient(
      closest-side,
      rgba(187, 106, 192, 0.5),
      rgba(248, 192, 147, 0)
    )`
    : `radial-gradient(
      closest-side,
      rgba(235, 105, 78, 1),
      rgba(235, 105, 78, 0)
    ),
    radial-gradient(closest-side, rgba(243, 11, 164, 1), rgba(243, 11, 164, 0)),
    radial-gradient(
      closest-side,
      rgba(254, 234, 131, 1),
      rgba(254, 234, 131, 0)
    ),
    radial-gradient(
      closest-side,
      rgba(170, 142, 245, 1),
      rgba(170, 142, 245, 0)
    ),
    radial-gradient(
      closest-side,
      rgba(248, 192, 147, 1),
      rgba(248, 192, 147, 0)
    )`
})
</script>
<style scoped>
.bg {
  background-color: v-bind(bgc);
  background-image: v-bind(bgi);
  background-size: 130vmax 130vmax, 80vmax 80vmax, 90vmax 90vmax,
    110vmax 110vmax, 90vmax 90vmax;
  background-position: -80vmax -80vmax, 60vmax -30vmax, 10vmax 10vmax,
    -30vmax -10vmax, 50vmax 50vmax;
  background-repeat: no-repeat;
  animation: 15s movement ease-in-out infinite;
}
.bg::after {
  content: '';
  display: block;
  position: fixed;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  backdrop-filter: blur(15px);
  -webkit-backdrop-filter: blur(10px);
}
@keyframes movement {
  0%,
  100% {
    background-size: 130vmax 130vmax, 80vmax 80vmax, 90vmax 90vmax,
      110vmax 110vmax, 90vmax 90vmax;
    background-position: -80vmax -80vmax, 60vmax -30vmax, 10vmax 10vmax,
      -30vmax -10vmax, 50vmax 50vmax;
  }
  25% {
    background-size: 100vmax 100vmax, 90vmax 90vmax, 100vmax 100vmax,
      90vmax 90vmax, 60vmax 60vmax;
    background-position: -60vmax -90vmax, 50vmax -40vmax, 0vmax -20vmax,
      -40vmax -20vmax, 40vmax 60vmax;
  }
  50% {
    background-size: 80vmax 80vmax, 110vmax 110vmax, 80vmax 80vmax,
      60vmax 60vmax, 80vmax 80vmax;
    background-position: -50vmax -70vmax, 40vmax -30vmax, 10vmax 0vmax,
      20vmax 10vmax, 30vmax 70vmax;
  }
  75% {
    background-size: 90vmax 90vmax, 90vmax 90vmax, 100vmax 100vmax,
      90vmax 90vmax, 70vmax 70vmax;
    background-position: -50vmax -40vmax, 50vmax -30vmax, 20vmax 0vmax,
      -10vmax 10vmax, 40vmax 60vmax;
  }
}

@media screen and (max-width: 600px) {
  .hitokoto {
    width: 100%;
  }
}

@media screen and (min-width: 600px) {
  .hitokoto {
    width: 50%;
  }
}
</style>
