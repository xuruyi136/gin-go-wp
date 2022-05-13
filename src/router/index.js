import Vue from 'vue'
import VueRouter from 'vue-router'
import store from '../store'
import HelloWorld from '@/components/HelloWorld'
import userRoutes from './module/user.js'

// import Regiter from '../components/Register'
Vue.use(VueRouter)
const routes = [
  {
    path: '/',
    redirect: '/home',
  },

  {
    path: '/home',
    name: 'home',
    component: HelloWorld
  },
  ...userRoutes,
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});
//https://router.vuejs.org/guide/advanced/navigation-guards.html#optional-third-argument-next
router.beforeEach((to, from, next) => {
  if (to.meta.auth) {//判断是否需要登陆
    //判断用户是否登陆
    if (store.state.userModule.token) {
      //这里判断token的有效性，有没有过期.需要后台发放token的时候带上token的有效期
      next()

    } else {
      //跳转登陆
      router.push({ name: 'login' })
    }
  } else {
    next()
  }
})
export default router;

/* const originalPush = Router.prototype.push

Router.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
} */