// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
import Home from './views/Home.vue';
import Collect from './views/Collect.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/collect',
    name: 'Collect',
    component: Collect
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
