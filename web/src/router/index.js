import { createRouter, createWebHashHistory} from 'vue-router'

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/',
    redirect: 'Map',
    component: () => import('@/view/geography/geography.vue'),
  },
  {
    path: '/init',
    name: 'Init',
    component: () => import('@/view/init/index.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/view/login/index.vue')
  },
  {
    path: '/scanUpload',
    name: 'ScanUpload',
    meta: {
      title: '扫码上传',
      client: true
    },
    component: () => import('@/view/example/upload/scanUpload.vue')
  },
  {
    path: '/equManage',
    name: 'EquManage',
    meta: {
      title: '设备列表',
      client: true
    },
    component: () => import('@/view/good/equManage/equManage.vue')
  },
  {
    path: '/:catchAll(.*)',
    meta: {
      closeTab: true
    },
    component: () => import('@/view/error/index.vue')
  },

]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
