import {Coin, User} from '@element-plus/icons-vue'

const menus = [
  {
    path: '/',
    name: 'login',
    component: () => import('./views/login.vue'),
  },
  {
    path: '/register',
    name: 'register',
    component: () => import('./views/register.vue'),
  },
  {
    path: '/',
    component: () => import('./layouts/main.vue'),
    children: [
      {
        path: 'database',
        icon: Coin,
        name: '日志查询',
        color: 'rgb(105, 192, 255)',
        component: () => import('./views/query.vue'),
      },
      {
        path: 'user',
        name: '登录用户',
        icon: User,
        color: 'rgb(255, 156, 110)',
        component: () => import('./views/user/index.vue'),
      },
    ],
  },
]

export default menus
