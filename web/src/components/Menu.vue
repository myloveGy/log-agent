<template>
  <div style="text-align:center;width: 200px;color:white" v-if="!isCollapsed">
    <h1 style="margin-top:20px;">日志收集系统</h1>
  </div>
  <el-menu
      :default-active="active"
      :collapse="isCollapsed"
      class="el-menu-vertical-demo"
      background-color="#191a22"
      text-color="#fff"
      active-text-color="#ffd04b"
  >
    <router-link v-for="item in items" :to="{name:item.name}" :key="item.name">
      <el-menu-item :index="item.path">
        <Icon :name="item.icon" :color="item.color"/>
        <template #title>{{ item.name }}</template>
      </el-menu-item>
    </router-link>
  </el-menu>
</template>

<script setup lang="ts">
import menus from '@/menus'
import Icon from './Icon.vue'
import {computed, defineProps} from 'vue'
import {useRoute} from 'vue-router'

const props = defineProps<{
  isCollapsed: boolean
}>()

const items = menus.find(item => item.path === '/' && item.children?.length)?.children || []
const route = useRoute()
const active = computed<string>(() => {
  return items.find(item => `/${item.path}` === route.path)?.path || ''
})

</script>

<style scoped>
.el-menu a {
  text-decoration: none;
}
</style>
