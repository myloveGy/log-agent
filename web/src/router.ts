import {createRouter, createWebHashHistory} from 'vue-router'

import menus from '@/menus'
import {userStore} from '@/utils'

const router = createRouter({
  history: createWebHashHistory(),
  routes: menus,
})

router.beforeEach(async (to, form) => {
  if (to.path !== '/' && !userStore.fetch('token')) {
    return {name: 'login'}
  }
})

export default router
