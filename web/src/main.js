// import 'babel-polyfill'
import 'idempotent-babel-polyfill';
import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import VueRouter from 'vue-router';
import store from './plugins/store';
import vueHeadful from 'vue-headful';
// import axios from 'axios';

// Components
import Navigation from './components/Navigation';
import List from './components/List';
import Home from './components/Home';
import Item from './components/Item';
import AddItem from './components/AddItem';
import ItemEdit from './components/ItemEdit';
import User from './components/User';
import UserItem from './components/UserItem';

// Vue router
const router = new VueRouter({
  mode: 'history',
  routes: [
    // dynamic segments start with a colon
    { path: '/', component: Home },
    { path: '/list', component: List },
    { path: '/navigation/:category', component: Navigation },
    { path: '/item/add', component: AddItem },
    { path: '/item/:slug', component: Item },
    { path: '/item/edit/:slug', component: ItemEdit },
    { path: '/user', component: User },
    { path: '/user/:slug', component: UserItem },
  ]
})

// Initialize router module
Vue.use(VueRouter);
// headful
Vue.component('vue-headful', vueHeadful);

Vue.config.productionTip = false


new Vue({
  vuetify,
  router,
  store,
  render: h => h(App)
}).$mount('#app')
