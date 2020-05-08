<template lang="pug">
  v-app
    
    v-navigation-drawer( class="pa-4" fixed :color="$store.state.themecolor" permanent floating)
      v-list-item
        v-list-item-content
          v-list-item-title( :class="$store.state.themecolor+'--text'" class="text--lighten-5" ) Справочник ТСМ
          v-list-item-subtitle( :class="$store.state.themecolor+'--text'" class="text--lighten-2" ) v0.1
      v-divider

      v-list( dense nav )
        v-list-item( v-for='(_m,_i) in menu' :key="_i" :to="_m.href" link )
          v-list-item-icon
            v-icon( :class="$store.state.themecolor+'--text'" class="text--lighten-5" ) {{_m.icon}}
          v-list-item-content
            v-list-item-title( :class="$store.state.themecolor+'--text'" class="text--lighten-5" ) {{_m.title}}
          
        v-footer( dark padless absolute  )
          v-card( flat tile width="100%" :color="$store.state.themecolor" class="text-center"   )
            
            v-btn(  to="/item/add"
                    dark
                    :color="$store.state.themeaccentcolor"
                    class="accent-4"
                    absolute
                    top
                    right
                    fab)
              v-icon mdi-plus
                
            v-divider
            
            v-card-text( class="white--text" )        
              div &copy; {{ new Date().getFullYear() }}

    v-content
      transition(name="fade")
        router-view
    
        
</template>

<script>
import Navigation from './components/Navigation';
import List from './components/List';
import Home from './components/Home';
import Item from './components/Item';



export default {
  name: 'App',
  components: {
    Home,
    List,
    Navigation,
    Item,
  },

  data: () => ({
    
    
    
    // menu
    menu:[
      { "title": 'Главная',
        "href" : '/',
        "icon" : "mdi-home" },
      { "title": 'Разделы',
        "href" : '/navigation',
        "icon" : "mdi-book-multiple" },
      { "title": 'Список ТСМ',
        "href" : '/list',
        "icon" : "mdi-view-list" },
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