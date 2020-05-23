<template lang="pug">
  v-container( :class="$store.state.themecolor+' lighten-5'" dark fluid height )
    //- Заголовок
    vue-headful( :title="'Изделие подразделения '+item.Name" )
    
    v-row( align="center" justify="center" class="vh-100" )
      v-col( cols=8  justify="center" class="mt-5") 
        
            
        v-card( :loading="loading" :disabled="loading" class="mb-7 "   )
            div( class="d-flex align-center justify-center grey theme--light lighten-3 item-placeholder-img relative " )
              //- image uploaded
              v-img( v-if="item.Image&&item.Image.Path!=''" :src="$store.state.addressprefix+item.Image.Path" lazy-src="@/assets/img/no-image.jpg" )

            v-card-title(class="pl-8 display-1 text--primary") {{item.Name}} {{'#'+userItem.ID}}
            v-card-text 
              v-row
                v-col(cols=4)
                  v-card(flat outlined)
                    v-card-text
                      div(class="d-flex flex-no-wrap justify-space-between") Дата ввода в эксплуатацию
                        v-spacer
                        v-btn(  icon color="lighten-1"  :color="edit.DateStart?$store.state.themecolor:'grey'" @click="edit.DateStart=true;" )
                          v-icon mdi-pencil
                        v-btn( class="ml-2" icon color="grey lighten-1"  @click="copyToClipboard(userItem.DateStart)" )
                          v-icon mdi-content-copy
                      p( v-if="!edit.DateStart" class="display-1 text--primary mb-0" ) {{userItem.DateStart!=''?userItem.DateStart:'не задано'}}
                      v-menu( v-if="edit.DateStart"
                              ref="menu" 
                              v-model="menu" 
                              :close-on-content-click="false"
                              :return-value.sync="userItem.DateStart"
                              transition="scale-transition"
                              offset-y )
                        template(v-slot:activator="{ on }")
                          v-text-field( v-model="userItem.DateStart" 
                                            append-outer-icon="mdi-content-save"
                                            @click:append-outer="saveParam('DateStart')"
                                            v-on:keyup.enter="saveParam('DateStart')"
                                            v-on="on"
                                            clearable )
                        v-date-picker( v-model="userItem.DateStart" 
                                        locale="ru-ru"
                                        :max="new Date().toISOString().substr(0, 10)"
                                        min="1950-01-01"
                                        color="green lighten-1"
                                        scrollable  )
                          v-spacer
                          v-btn( @click="menu=false" ) Отмена
                          v-btn( @click="$refs.menu.save(userItem.DateStart)" ) Ок
                              
                        
                v-col(cols=4)
                  v-card(flat outlined)
                    v-card-text
                      div(class="d-flex flex-no-wrap justify-space-between") Время работы, часов в день
                        v-spacer
                        v-btn(  icon color="lighten-1" :color="edit.DayHours?$store.state.themecolor:'grey'" @click="edit.DayHours=true" )
                          v-icon mdi-pencil
                        v-btn( class="ml-2" icon color="grey lighten-1" @click="copyToClipboard(userItem.DayHours)" )
                          v-icon mdi-content-copy
                      p( v-if="!edit.DayHours" class="display-1 text--primary mb-0" ) {{userItem.DayHours>0?userItem.DayHours:'не задано'}}
                      v-text-field( v-model="userItem.DayHours" v-if="edit.DayHours" :counter="10" 
                                        append-outer-icon="mdi-content-save"
                                        @click:append-outer="saveParam('DayHours')"
                                        v-on:keyup.enter="saveParam('DayHours')"
                                        clearable )
                v-col(cols=4)
                  v-card(flat outlined)
                    v-card-text
                      div(class="d-flex flex-no-wrap justify-space-between") Срок наработки на отказ, лет
                        v-spacer
                        v-btn(  icon color="lighten-1" :color="edit.Srok?$store.state.themecolor:'grey'" @click="edit.Srok=true" )
                          v-icon mdi-pencil
                        v-btn( class="ml-2" icon color="grey lighten-1" @click="copyToClipboard(userItem.Srok)" )
                          v-icon mdi-content-copy
                      p( v-if="!edit.Srok" class="display-1 text--primary mb-0" ) {{userItem.Srok>0?userItem.Srok:'не задано'}}
                      v-text-field( v-model="userItem.Srok" v-if="edit.Srok" :counter="10"  
                                        append-outer-icon="mdi-content-save"
                                        @click:append-outer="saveParam('Srok')"
                                        v-on:keyup.enter="saveParam('Srok')"
                                        clearable )
              v-row
                v-col(cols=8)
                  v-card(flat )
                    v-card-text
                      div(class="d-flex flex-no-wrap justify-space-between") Примечание
                        v-spacer
                        v-btn(  icon color=" lighten-1" :color="edit.Primechanie?$store.state.themecolor:'grey'" @click="edit.Primechanie=true" )
                          v-icon mdi-pencil
                        v-btn( class="ml-2" icon color="grey lighten-1" @click="copyToClipboard(userItem.Primechanie)" )
                          v-icon mdi-content-copy
                      
                      v-text-field( v-model="userItem.Primechanie" :disabled="!edit.Primechanie" :counter="500"  label="Примечание" 
                                        prepend-icon="mdi-information-outline"
                                        append-outer-icon="mdi-content-save"
                                        @click:append-outer="saveParam('Primechanie')"
                                        v-on:keyup.enter="saveParam('Primechanie')"
                                        required clearable )
                v-col(cols=4)
                  v-card(flat )
                    v-card-text
                      div(class="d-flex flex-no-wrap justify-space-between") Инвентарный номер
                        v-spacer
                        v-btn(  icon color="lighten-1" :color="edit.Inventory?$store.state.themecolor:'grey'" @click="edit.Inventory=true" )
                          v-icon mdi-pencil
                        v-btn( class="ml-2" icon color="grey lighten-1" @click="copyToClipboard(userItem.Inventory)")
                          v-icon mdi-content-copy
                      p( v-if="!edit.Inventory" class="display-1 text--primary mb-0" ) {{userItem.Inventory>0?userItem.Inventory:'не задано'}}
                      v-text-field( v-model="userItem.Inventory" v-if="edit.Inventory" :counter="100" 
                                        append-outer-icon="mdi-content-save"
                                        @click:append-outer="saveParam('Inventory')"
                                        v-on:keyup.enter="saveParam('Inventory')"
                                        clearable )
              
                
              v-expansion-panels( flat hover )
                v-expansion-panel
                  v-expansion-panel-header Больше информации об изделии
                  v-expansion-panel-content
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
                        div(class="body-1 d-flex w-100 justify-space-between align-center" ) Категория техники
                          v-btn( class="ml-8" icon color="grey lighten-1" :to="'/navigation/'+getCategoryIndex(item.Category)")
                            v-icon mdi-open-in-new
                      v-col(cols=8)
                        div( class="body-1  grey--text text--darken-4" ) {{getCategoryName(item.Category)}}
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
    name: 'UserItem',
    
    data: () => ({
      // loader
      loading:true,
      // menu activator
      menu:false,
      // editflags
      edit:{
        DateStart : false,
        DayHours : false,
        Srok : false,
        Inventory : false,
        Primechanie : false,
      },
      
      // useritem data
      userItem:{},
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
      axios.get(this.$store.state.addressprefix+'/api/useritem/'+this.$route.params.id)
      .then( res => { 
        _vue.item     = res.data.Value.Item;
        _vue.userItem = res.data.Value;
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
      },
      // save only one param to DB
      saveParam(param){
        let _data = new FormData();
        _data.append( param, this.userItem[param] );
        _data.append( 'paramName', param );

        axios.put( this.$store.state.addressprefix+'/api/useritem/'+this.$route.params.id, _data )
                     .then( (res)=>{
                        this.popup       = true;
                        this.popup_style = 'success';
                        this.popup_text  = 'Изменения сохранены!';
                        this.newslug     = '/item/'+res.data.URL;
                        
                        // /api/useritem/:id
                        this.edit[param]=false;
                     })
                     .catch(error => {
                        // console.log(error.response)
                        if ( typeof error.response !=='undefined' ){
                            this.popup       = true;
                            this.popup_style = 'error';
                            this.popup_text  = 'Ошибка сохранения изменений "'+error.response.data+'"';
                        }
                     });

        
      },
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
