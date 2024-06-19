import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/home', name: 'Home', component: () => import('../views/Home.vue') },
    { path: '/edit', name: 'Edit', component: () => import('../views/Edit.vue') },
    { path: '/run', name: 'Run', component: () => import('../views/Run.vue') },
    { path: '/manager', name: 'Manager', component: () => import('../views/Manager.vue') },
    { path: '/register', name: 'Register', component: () => import('../views/Register.vue') },
    { path: '/forget', name: 'Forget', component: () => import('../views/Forget.vue') },
    { path: '/', name: 'Login', component: () => import('../views/Login.vue') },
  ]
})

export default router
