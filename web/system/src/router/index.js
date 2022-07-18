import { createRouter, createWebHashHistory } from 'vue-router'
export const constantRoutes = [
    {
        path: '/user/login',
        name: 'login',
        component: ()=> import('@/views/login'),
        hidden: false,
    },
    {
        path: '/user/signup',
        name: 'signup',
        component: ()=> import('@/views/signup'),
    },
    {
        path: '/dashboard',
        name: 'dashboard',
        component: ()=> import('@/views/dashboard'),
    },
    {
        path: '/',
        redirect: "/user/login"
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes: constantRoutes,
})

export default router