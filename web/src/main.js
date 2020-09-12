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
import Home from './components/Home';

// Vue router
const router = new VueRouter({
  mode: 'history',
  routes: [
    // dynamic segments start with a colon
    { path: '/', component: Home },
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
