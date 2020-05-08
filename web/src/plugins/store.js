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
    },
    mutations: {
      increment (state) {
        state.count++
      }
    }
})
  