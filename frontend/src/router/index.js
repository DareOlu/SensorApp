import VueRouter from 'vue-router'
import History from '../views/History.vue'
import Fuel from '../views/Fuel.vue'
import Optimise from '../views/Optimise.vue'


const routes = [
  { path: '/', component: Fuel},
  { path: '/History', component: History},
  { path: '/Optimise', component: Optimise},
]

const router = new VueRouter({
  routes,
  mode: 'history'
})

export default router
