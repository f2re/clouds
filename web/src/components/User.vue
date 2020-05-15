<template lang="pug">
  v-container( :class="$store.state.themecolor+' lighten-5 vh-100 '" dark fluid height )
    //- Заголовок
    vue-headful( title="Список технических средств подразделения" )

    h1(class="mt-2 mb-2" :class="$store.state.themecolor+'--text'") Технические средства метеоподразделения
      v-btn( to="/list" :color="$store.state.themecolor" icon large class="float-left mt-1 mr-5" )
        v-icon mdi-chevron-left
      v-subheader(class="ml-12 relative") Для добавления ТСМ в штат подразделения выберите технику из общего списка
        v-btn( @click="dialog=true" 
                    dark
                    :color="$store.state.themeaccentcolor"
                    class="accent-4"
                    absolute
                    bottom
                    right
                    fab)
            v-icon mdi-book-plus
    v-divider
          
      
    div(class="d-flex flex-wrap mt-5")
      //- preloader
      div(class="col-4 pa-3" v-if="loadingUser")
        v-card(  :loading="loadingUser" :disabled="loadingUser"    )
          v-img( lazy-src="@/assets/img/no-image.jpg" aspect-ratio="2" class="white--text align-end")
      div(class="col-4 pa-3" v-if="loadingUser")
        v-card(  :loading="loadingUser" :disabled="loadingUser"    )
          v-img( lazy-src="@/assets/img/no-image.jpg" aspect-ratio="2" class="white--text align-end")
      div(class="col-4 pa-3" v-if="loadingUser" )
        v-card(  :loading="loadingUser" :disabled="loadingUser"    )
          v-img( lazy-src="@/assets/img/no-image.jpg" aspect-ratio="2" class="white--text align-end")
      //- real loader
      div(v-for="(_item,_k) in userItems" :key="_k" class="col-4 pa-3")
        v-card(  :loading="loadingUser" :disabled="loadingUser" :to="'/item/'+_item.Item.Slug"   )
          v-img( lazy-src="@/assets/img/no-image.jpg" :src="_item.Item.Image&&_item.Item.Image.Path!=''? ($store.state.addressprefix+_item.Item.Image.Path) : ''" aspect-ratio="2" class="white--text align-end"
              gradient="to bottom, rgba(0,0,0,.0),rgba(0,0,0,.1),rgba(0,0,0,.2), rgba(0,0,0,.5)")
            v-card-title() {{_item.Item.Name}}
          v-card-text() Индекс изделия 
            b {{_item.Item.Index}}
    
    v-dialog( v-model="dialog" 
              fullscreen 
              transition="dialog-bottom-transition"
              scrollable )
      v-card( tile )
        v-toolbar( dense dark :color="$store.state.themecolor" )
          v-btn( icon dark @click="dialog=false" )
            v-icon mdi-close
          v-toolbar-title Ввыберите технику из списка для добавления в штат метеоподразделения
          v-spacer
          v-toolbar-items
              v-btn( @click="dialog=false" dark text ) Сохранить
        v-card-text 
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
                v-card(  :loading="loading" :disabled="loading" @click="addItem(_item.Slug)"   )
                    v-img( lazy-src="@/assets/img/no-image.jpg" :src="_item.Image&&_item.Image.Path!=''? ($store.state.addressprefix+_item.Image.Path) : ''" aspect-ratio="2" class="white--text align-end"
                        gradient="to bottom, rgba(0,0,0,.0),rgba(0,0,0,.1),rgba(0,0,0,.2), rgba(0,0,0,.5)")
                        v-card-title() {{_item.Name}}
                    v-card-text() Индекс изделия 
                        b {{_item.Index}}
                    v-fade-transition
                      v-overlay(v-if="_item.Slug==addedSlug" absolute :color="$store.state.themecolor")
                        v-btn( text fab x-large)
                          v-icon(class="mdi-48px") mdi-check-circle
                        
              
            
          
        
</template>

<script>
  import axios from 'axios'; 

  export default {
    name: 'User',

    data: () => ({
      // loader
      loadingUser: true,
      loading: true,
      // список изделий в штате
      userItems:[],
      // список изделий всех
      items:[],
      // routename
      // name: this.$route.params.name
      //   open dialog to select item
      dialog:false,
      // 
      addedSlug:''
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
      //   add Item to user store
      addItem(slug){
        //   let _vue = this;
          let data = new FormData();
          data.append('slug',slug);
          axios.post(this.$store.state.addressprefix+'/api/adduseritem', data )
               .then( () => { 
                  //  console.log(res);
                   this.addedSlug = slug;
               } );
      }
      
    },
    // загружаем данные после того как перешли по ссылке
    mounted() {
      var _vue = this;
      // set loader
      _vue.loadingUser = true;
      // load data
      axios.get(this.$store.state.addressprefix+'/api/useritems')
      .then( res => { 
        _vue.userItems = res.data.Value;
        // console.log(_vue.userItems[0].Item.Image.Path);
        _vue.loadingUser=false;
      } );

       _vue.loading = true;
      axios.get(this.$store.state.addressprefix+'/api/items')
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
  .relative{
      position: relative;
  }
</style>