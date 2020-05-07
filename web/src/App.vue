<template lang="pug">
  v-app
    
    v-navigation-drawer( fixed :color="themecolor" permanent floating)
      v-list-item
        v-list-item-content
          v-list-item-title( :class="themecolor+'--text'" class="text--lighten-5" ) Справочник ТСМ
          v-list-item-subtitle( :class="themecolor+'--text'" class="text--lighten-2" ) v0.1
      v-divider

      v-list( dense nav )
        v-list-item( v-for='(_m,_i) in menu' :key="_i" :to="_m.href" link )
          v-list-item-icon
            v-icon( :class="themecolor+'--text'" class="text--lighten-5" ) {{_m.icon}}
          v-list-item-content
            v-list-item-title( :class="themecolor+'--text'" class="text--lighten-5" ) {{_m.title}}
          
        v-footer( dark padless absolute  )
          v-card( flat tile width="100%" :color="themecolor" class="text-center"   )
            
            v-divider
            
            v-card-text( class="white--text" )        
              div &copy; {{ new Date().getFullYear() }}

    v-content
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
    
    // variable of theme color
    themecolor: "indigo",
    // menu
    menu:[
      { "title": 'Главная',
        "href" : '/',
        "icon" : "mdi-home" },
      { "title": 'Разделы',
        "href" : '/navigation',
        "icon" : "mdi-bookmarks" },
      { "title": 'Список ТСМ',
        "href" : '/list',
        "icon" : "mdi-info" },
    ]

  }),
};
</script>

<style >
  .v-content{
    padding-left:256px!important;
  }
</style>