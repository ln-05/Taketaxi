import { createRouter, createWebHistory } from 'vue-router';
import Login from './Login.vue';
import Captcha from './Captcha.vue';
import Home from './Home.vue';
import Search from './Search.vue';
import Taxi from './Taxi.vue';

const routes = [
  { path: '/', component: Login },
  { path: '/captcha', component: Captcha },
  { path: '/home', component: Home },
  { path: '/search', component: Search },
  { path: '/taxi', component: Taxi },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router; 