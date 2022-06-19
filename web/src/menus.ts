import {Coin, User} from '@element-plus/icons-vue'

const menus = [
  {
    path: '/',
    name: 'query',
    icon: Coin,
    label: '日志查询',
    color: 'rgb(105, 192, 255)',
    component: () => import('./views/query.vue'),
  },
  {
    path: '/user',
    name: 'user',
    label: '用户信息',
    icon: User,
    color: 'rgb(255, 156, 110)',
    component: () => import('./views/user.vue'),
  },
]

export default menus
