<template lang="pug">
  v-container( :class="$store.state.themecolor+' lighten-5'" dark fluid height )
    //- Заголовок
    vue-headful( title="Список технических средств" )

    div(class="d-flex flex-wrap ")
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
          v-img( lazy-src="@/assets/img/no-image.jpg" :src="$store.state.addressprefix+_item.Image.Path" aspect-ratio="2" class="white--text align-end"
              gradient="to bottom, rgba(0,0,0,.0),rgba(0,0,0,.1),rgba(0,0,0,.2), rgba(0,0,0,.5)")
            v-card-title() {{_item.Name}}
          v-card-text() Индекс изделия 
            b {{_item.Index}}
</template>

<script>
  import axios from 'axios'; 

  export default {
    name: 'List',

    data: () => ({
      // loader
      loading: true,
      // список изделий
      items:[],
      // routename
      // name: this.$route.params.name
    }),
    // загружаем данные после того как перешли по ссылке
    mounted() {
      var _vue = this;
      // set loader
      _vue.loading = true;
      // load data
      axios.get(this.$store.state.addressprefix+'/api/items')
      .then( res => { 
        _vue.items = res.data.Value;
        // console.log(_vue.items[0].Image.Path);
        _vue.loading=false;
      } );
      console.log('sdf');

    }
  }
</script>

<style scoped>
  .vh-100{
    height: 100%;
  }
</style>