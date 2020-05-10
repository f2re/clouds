import Vue from 'vue';
import Vuex from 'vuex'

// Initialize Vuex store
Vue.use(Vuex);

// Vuex store component
export default new Vuex.Store({
    state: {
        // variable of theme color
        themecolor: "teal",
        // accent color
        themeaccentcolor: "orange",
        // local address prefix
        addressprefix:'http://localhost:8000',
        // addressprefix:'',
    },
    mutations: {
      increment (state) {
        state.count++
      }
    }
})
  