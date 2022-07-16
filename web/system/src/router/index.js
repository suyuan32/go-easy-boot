import { createRouter, createWebHashHistory } from 'vue-router'
import { LogIn } from '@/views/login/index'
export const constantRoutes = [
    {
        path: '/login',
        component: LogIn,
        hidden: false,
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes: constantRoutes,
})

export default router