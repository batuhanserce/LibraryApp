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
            path: '/books',
            name: 'books',
            component: () => import('@/views/book/books')
        },
        {
            path: '/users',
            name: 'users',
            component: () => import('@/views/user/users')
        },
        {
            path: '/user/:id',
            name: 'user',
            component: () => import('@/views/user/user')
        },

        {
            path: '*',
            name: 'notFound',
            component: () => import('@/views/NotFound')
        }
    ]
})
