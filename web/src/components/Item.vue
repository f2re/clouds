<template lang="pug">
  v-container( :class="$store.state.themecolor+' lighten-5'" dark fluid height )
    //- Заголовок
    vue-headful( :title="'Изделие '+item.Name" )
    
    v-row( align="center" justify="center" class="vh-100" )
      v-col( cols=8  justify="center" class="mt-5") 
        v-card( :loading="loading" :disabled="loading" class="mb-7 "   )
            div( class="d-flex align-center justify-center grey theme--light lighten-3 item-placeholder-img relative " )
              //- image uploaded
              v-img( v-if="item.Image.Path!=''" :src="$store.state.addressprefix+item.Image.Path" lazy-src="@/assets/img/no-image.jpg" )
              //- upload image button
              v-btn(  dark bottom right absolute fab
                  :to="'/item/edit/'+$route.params.slug"
                  :color="$store.state.themeaccentcolor" )
                  v-icon mdi-pencil
            v-card-title(class="pl-8 display-1 text--primary") {{item.Name}}
            v-card-text 
              //- номернклатура изделия и коды с номерами
              v-row
                v-col(cols=4)
                  v-card(flat outlined)
                    v-card-text
                      div(class="d-flex flex-no-wrap justify-space-between") Индекс изделия
                        v-btn( class="ml-8" icon color="grey lighten-1" @click="copyToClipboard(item.Index)")
                          v-icon mdi-content-copy
                      p( class="display-1 text--primary mb-0" ) {{item.Index}}
                v-col(cols=4)
                  v-card(flat outlined)
                    v-card-text
                      div(class="d-flex flex-no-wrap justify-space-between") Код КВТ
                        v-btn( class="ml-8" icon color="grey lighten-1" @click="copyToClipboard(item.KVT)")
                          v-icon mdi-content-copy
                      p( class="display-1 text--primary mb-0" ) {{item.KVT}}
                v-col(cols=4)
                  v-card(flat outlined)
                    v-card-text
                      div(class="d-flex flex-no-wrap justify-space-between") Номенклатурный номер
                        v-btn( class="ml-8" icon color="grey lighten-1" @click="copyToClipboard(item.Nomenklature)")
                          v-icon mdi-content-copy
                      p( class="display-1 text--primary mb-0" ) {{item.Nomenklature}}

              v-row
                v-col(cols=4 class="pl-8")
                  div(class="body-1 d-flex w-100 justify-space-between align-center" ) Довольствующий орган
                    v-btn( class="ml-8" icon color="grey lighten-1" @click="copyToClipboard(item.Dovorgan)")
                      v-icon mdi-content-copy
                v-col(cols=8)
                  div( class="body-1  grey--text text--darken-4" ) {{item.Dovorgan}}
              v-row
                v-col(cols=4 class="pl-8")
                  div( class="body-1 d-flex w-100 justify-space-between align-center"  ) Заказывающий орган
                    v-btn( class="ml-8" icon color="grey lighten-1" @click="copyToClipboard(item.Reqorgan)")
                        v-icon mdi-content-copy
                v-col(cols=8)
                  div( class="body-1 grey--text text--darken-4" ) {{item.Reqorgan}}
              v-row
                v-col(cols=4 class="pl-8")
                  div( class="body-1 d-flex w-100 justify-space-between align-center"  ) Эксплуатирующий орган
                    v-btn( class="ml-8" icon color="grey lighten-1" @click="copyToClipboard(item.Explorgan)")
                        v-icon mdi-content-copy
                v-col(cols=8)
                  div( class="body-1 grey--text text--darken-4" ) {{item.Explorgan}}
              v-row
                v-col(cols=4 class="pl-8")
                  div( class="body-1 d-flex w-100 justify-space-between align-center"  ) Предприятие изготовитель
                    v-btn( class="ml-8" icon color="grey lighten-1" @click="copyToClipboard(item.Creator)")
                        v-icon mdi-content-copy
                v-col(cols=8)
                  div( class="body-1 grey--text text--darken-4" ) {{item.Creator}}
              
              v-divider

              //- Назначение
              v-row
                v-col(cols=4 class="pl-8")
                  div( class="body-1 d-flex w-100 justify-space-between align-center"  ) Назначение
                    v-btn( class="ml-8" icon color="grey lighten-1" @click="copyToClipboard(item.Destination)")
                        v-icon mdi-content-copy
                v-col(cols=8)
                  div( class="body-1 grey--text text--darken-4" ) {{item.Destination}}
              //- Состав
              v-row
                v-col(cols=4 class="pl-8")
                  div( class="body-1 d-flex w-100 justify-space-between align-center"  ) Состав
                    v-btn( class="ml-8" icon color="grey lighten-1" @click="copyToClipboard(item.Composition)")
                        v-icon mdi-content-copy
                v-col(cols=8)
                  div( class="body-1 grey--text text--darken-4" ) {{item.Composition}}
              
              v-divider
                
              //- ТТХ
              v-card-title Тактико-технические характеристики
              v-row( v-for="(_t,_i) in item.TTH" :key="_i" )
                v-col(cols=4 class="pl-8")
                  div( class="body-1 d-flex w-100 justify-space-between align-center"  ) {{_t.label}}
                v-col(cols=8)
                  div( class="body-1 grey--text text--darken-4" ) {{_t.value}}
              
              v-divider

              v-card-title Нормы табелизации
              v-simple-table( v-if="typeof item.Tabel !=='undefined'" )
                template( v-slot:default )
                  thead(class="caption")
                    tr
                      th Отд эскадрилья
                      th Ав. полк
                      th Ав. дивизия
                      th А ВВС и ПВО
                      th Уч. ав. центр
                      th ЦУ вида/рода
                      th Межвидовой полигон
                      th Ав. комендатура
                      th Уч. заведение
                  tbody
                    tr
                      td {{item.Tabel.Escadrile}}
                      td {{item.Tabel.Polk}}
                      td {{item.Tabel.Division}}
                      td {{item.Tabel.Army}}
                      td {{item.Tabel.Center}}
                      td {{item.Tabel.CY}}
                      td {{item.Tabel.Polygon}}
                      td {{item.Tabel.Komend}}
                      td {{item.Tabel.School}}

</template>

<script>
  import axios from 'axios'; 

  export default {
    name: 'Item',
    
    data: () => ({
      // loader
      loading:true,
      // item data
      item:{
        Image:{
          Path:''
        }
      },
    }),
    // загружаем данные после того как перешли по ссылке
    mounted() {
      var _vue = this;
      // set loader
      _vue.loading = true;
      // load data
      axios.get(this.$store.state.addressprefix+'/api/item/'+this.$route.params.slug)
      .then( res => { 
        _vue.item = res.data.Value;
        if ( typeof _vue.item.TTH !=='undefined' ){
          _vue.item.TTH  = JSON.parse( _vue.item.TTH );
        }
        _vue.loading=false;
      } );
    },
    // methods
    methods:{
      // copy text to clipboard
      copyToClipboard(text){
        navigator.clipboard.writeText(text);
      }
    }
  }
</script>
<style scoped>
.relative{
    position: relative;
}
.w-100{
  width: 100%;
}
.d-flex{
  display: flex;
}
</style>
