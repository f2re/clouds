<template lang="pug">
  v-app
    
    v-navigation-drawer( class="pa-4" fixed :color="$store.state.themecolor" permanent floating)
      v-list-item
        v-list-item-content
          v-list-item-title( :class="$store.state.themecolor+'--text'" class="text--lighten-5" ) Справочник облаков
          v-list-item-subtitle( :class="$store.state.themecolor+'--text'" class="text--lighten-2" ) v{{version}}
      v-divider

    v-main
      transition(name="fade")
        router-view(:key="$route.fullPath")
    
        
</template>

<script>
import Home from './components/Home';

export default {
  name: 'App',
  components: {
    Home,
  },

  data: () => ({
    
    version: '0.1',
    
    // menu
    menu:[
      { "title": 'Главная',
        "href" : '/',
        "icon" : "mdi-home" },
    ]

  }),
};
</script>

<style >
  .v-content{
    padding-left:256px!important;
  }
  .fade-enter-active, .fade-leave-active {
    transition: opacity .3s;
  }
  .fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
    opacity: 0;
  }
</style>
