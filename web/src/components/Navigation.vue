<template lang="pug">
  v-container( :class="$store.state.themecolor+' lighten-5 vh-100 '" dark fluid height )
    //- Заголовок
    vue-headful( title="Список технических средств по категориям" )

    h1(class="mt-2 mb-2" :class="$store.state.themecolor+'--text'") {{getCategoryName($route.params.category)}}
      v-btn( to="/list" :color="$store.state.themecolor" icon large class="float-left mt-1 mr-5" )
        v-icon mdi-chevron-left
    v-divider
      
    div(class="d-flex flex-wrap mt-5")
      //- preloader
      div(class="col-4 pa-3" v-if="loading")
        v-card(  :loading="loading" :disabled="loading"    )
          v-img( lazy-src="@/assets/img/no-image.jpg" aspect-ratio="2" class="white--text align-end")
      div(class="col-4 pa-3" v-if="loading")
        v-card(  :loading="loading" :disabled="loading"    )
          v-img( lazy-src="@/assets/img/no-image.jpg" aspect-ratio="2" class="white--text align-end")
      div(class="col-4 pa-3" v-if="loading" )
        v-card(  :loading="loading" :disabled="loading"    )
          v-img( lazy-src="@/assets/img/no-image.jpg" aspect-ratio="2" class="white--text align-end")
      //- real loader
      div(v-for="(_item,_k) in items" :key="_k" class="col-4 pa-3")
        v-card(  :loading="loading" :disabled="loading" :to="'/item/'+_item.Slug"   )
          v-img( lazy-src="@/assets/img/no-image.jpg" :src="_item.Image&&_item.Image.Path!=''? ($store.state.addressprefix+_item.Image.Path) : ''" aspect-ratio="2" class="white--text align-end"
              gradient="to bottom, rgba(0,0,0,.0),rgba(0,0,0,.1),rgba(0,0,0,.2), rgba(0,0,0,.5)")
            v-card-title() {{_item.Name}}
          v-card-text() Индекс изделия 
            b {{_item.Index}}
</template>

<script>
  import axios from 'axios'; 

  export default {
    name: 'Navigation',

    data: () => ({
      // loader
      loading: true,
      // список изделий
      items:[],
      // routename
      // name: this.$route.params.name
    }),
    methods:{
      // получаем индекс категории из сторейджа
      getCategoryIndex(indx){
        let r = false;
        this.$store.state.categories.forEach( obj => {
          if ( obj.val==indx ){
            r=obj.val
          }
        } );
        return r;
      },
      // получаем имя категории из сторейджа
      getCategoryName(indx){
        let r = false;
        this.$store.state.categories.forEach( obj => {
          if ( obj.val==indx ){
            r= obj.text
          }
        } );
        return r;
      },
    },
    // загружаем данные после того как перешли по ссылке
    mounted() {
      var _vue = this;
      // set loader
      _vue.loading = true;
      // load data
      axios.get(this.$store.state.addressprefix+'/api/category/'+this.$route.params.category)
      .then( res => { 
        _vue.items = res.data.Value;
        // console.log(_vue.items[0].Image.Path);
        _vue.loading=false;
      } );
    }
  }
</script>

<style scoped>
  .vh-100{
    min-height: 100vh;
  }
</style>