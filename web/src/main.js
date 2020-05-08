import 'babel-polyfill'
import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import VueRouter from 'vue-router';
import store from './plugins/store';

// Components
import Navigation from './components/Navigation';
import List from './components/List';
import Home from './components/Home';
import Item from './components/Item';

// Vue router
const router = new VueRouter({
  mode: 'history',
  routes: [
    // dynamic segments start with a colon
    { path: '/', component: Home },
    { path: '/list', component: List },
    { path: '/navigation', component: Navigation },
    { path: '/item', component: Item,
      children:[
        {
          path:'edit',
          component: Item
        }
      ]
     },
  ]
})

// Initialize router module
Vue.use(VueRouter);


Vue.config.productionTip = false


new Vue({
  vuetify,
  router,
  store,
  render: h => h(App)
}).$mount('#app')
