import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter)

export const router = new VueRouter({
    base: '/',
    mode: 'history',
    routes: [
        {
            path: '',
            name: 'home',
            component: () => import('@/views/home')
        },
        {
            path: '/login',
            name: 'login',
            component: () => import('@/views/auth/Login')
        },
        {
            path: '/signup',
            name: 'signup',
            component: () => import('@/views/auth/Signup')
        },

        {
            path: '*',
            name: 'notFound',
            component: () => import('@/views/NotFound')
        }
    ]
})
