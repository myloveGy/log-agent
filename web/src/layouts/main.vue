<template>
  <el-container>
    <el-aside class="menubar">
      <Menu :isCollapsed="isCollapsed"></Menu>
    </el-aside>
    <el-container>

      <el-header
          style="display:flex;justify-content: space-between;border-bottom: 1px solid var(--el-card-border-color);"
      >
        <div style="display:flex;align-items: center">
          <el-link @click="isCollapsed=!isCollapsed">
            <el-icon :size="26">
              <fold v-if="!isCollapsed"/>
              <expand v-else/>
            </el-icon>

          </el-link>
          <span style="margin-left:10px;" v-text="name"></span>
        </div>
        <div style="display:flex;align-items:center">
          <el-icon @click="toggle()" style="margin-right: 10px">
            <FullScreen/>
          </el-icon>
          主题切换：
          <el-switch v-model="isDark">切换主题</el-switch>
          <span style="margin-right:10px;"></span>
          <el-dropdown style="margin-right: 10px">
            <span class="el-dropdown-link" style="display:flex;align-items:center">
              <el-icon style="margin-right: 8px; color: rgb(105, 192, 255)"><User/></el-icon>
              {{ user.username }}
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="logout">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>
      <el-main class="main">
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import Menu from '@/components/Menu.vue'
import {Expand, Fold} from '@element-plus/icons-vue'
import {useDark, useFullscreen, useToggle} from '@vueuse/core'
import router from '@/router'
import {computed, ref} from 'vue'
import {useRoute, useRouter} from 'vue-router'
import {LoginUser} from '@/api'
import {userStore} from '@/utils'

const {isFullscreen, toggle} = useFullscreen()
const route = useRoute()
const isDark = useDark()
const toggleDark = useToggle(isDark)
const isCollapsed = ref(false)

const user = computed<LoginUser>(() => {
  return userStore.fetch()
})

const name = computed(() => {
  return router.options.routes.filter(item => item.path === '/' && item.children?.length)
      .find(item => item.path === route.path)?.name || ''
})

const vueRouter = useRouter()
const logout = () => {
  userStore.flush()
  vueRouter.push('/')
}
</script>


