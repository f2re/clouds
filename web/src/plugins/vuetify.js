import Vue from 'vue';
import Vuetify from 'vuetify/lib';
import '@mdi/font/css/materialdesignicons.css'

Vue.use(Vuetify);

export default new Vuetify({
  defaultAssets: {
    font: true,
    icons: 'mdi'
  },
  icons: {
    iconfont: 'mdi',
  }
});
