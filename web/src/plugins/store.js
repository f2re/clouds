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
        // Категории техники
        categories:[
          { val:1,
            text: "Измерительная техника"},
          { val:2,
            text: "Радиозондировочная техника"},
          { val:3,
            text: "Телекоммутационная техника"},
        ]
    },
    
})
  