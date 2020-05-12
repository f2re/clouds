<template lang="pug">
  v-container( height="100%" class="bg" fluid )
    //- Заголовок
    vue-headful( title="Справочник технических средств метеорологической службы" )

    v-row( align="center" justify="center" class="vh-100" )
      v-col( cols=6 align="center" justify="center" class="mt-5")
        h1( :class="$store.state.themecolor+'--text'" class="text--lighten-5 mb-5" ) Справочник технических средств метеорологической службы
        //- :debounce-search="0"
        v-autocomplete( 
                        :search-input.sync="search"
                        v-model="model"
                        :items="items"
                        :loading="isLoading"
                        item-text="Name"
                        item-value="Slug"
                        color="white" 
                        hide-no-data 
                        hide-selected 
                        label="Технические средства"
                        placeholder="Начните ввод, чтобы искать"
                        prepend-icon="mdi-magnify"
                        class="mb-5"
                        return-object
                        dark )

        v-divider(dark)

        v-row( v-if="model" align="center" justify="center" class="vh-100" )
          v-col( cols=12  justify="center" class="mt-5") 
            v-card( class="mb-7 "   )
                div( class="d-flex align-center justify-center grey theme--light lighten-3 item-placeholder-img relative " )
                  //- image uploaded
                  v-img( v-if="model.Image.Path!=''" :src="$store.state.addressprefix+model.Image.Path" lazy-src="@/assets/img/no-image.jpg" aspect-ratio="2" )
                  //- upload image button
                  v-btn(  dark top right absolute fab
                      :to="'/item/'+model.Slug"
                      :color="$store.state.themeaccentcolor" )
                      v-icon mdi-link
                v-card-title(class="pl-8 display-1 text--primary") {{model.Name}}
                v-card-text 
                  //- номернклатура изделия и коды с номерами
                  v-row
                    v-col(cols=4)
                      v-card(flat outlined)
                        v-card-text
                          div(class="d-flex flex-no-wrap justify-space-between") Индекс изделия
                            v-btn( class="ml-8" icon color="grey lighten-1" @click="copyToClipboard(model.Index)")
                              v-icon mdi-content-copy
                          p( class="display-1 text--primary mb-0" ) {{model.Index}}
                    v-col(cols=4)
                      v-card(flat outlined)
                        v-card-text
                          div(class="d-flex flex-no-wrap justify-space-between") Код КВТ
                            v-btn( class="ml-8" icon color="grey lighten-1" @click="copyToClipboard(model.KVT)")
                              v-icon mdi-content-copy
                          p( class="display-1 text--primary mb-0" ) {{model.KVT}}
                    v-col(cols=4)
                      v-card(flat outlined)
                        v-card-text
                          div(class="d-flex flex-no-wrap justify-space-between") Номенклатурный номер
                            v-btn( class="ml-8" icon color="grey lighten-1" @click="copyToClipboard(model.Nomenklature)")
                              v-icon mdi-content-copy
                          p( class="display-1 text--primary mb-0" ) {{model.Nomenklature}}

                  v-row
                    v-col(cols=4 class="pl-8")
                      div(class="body-1 d-flex w-100 justify-space-between align-center" ) Довольствующий орган
                        v-btn( class="ml-8" icon color="grey lighten-1" @click="copyToClipboard(model.Dovorgan)")
                          v-icon mdi-content-copy
                    v-col(cols=8)
                      div( class="body-1  grey--text text--darken-4" ) {{model.Dovorgan}}
                  v-row
                    v-col(cols=4 class="pl-8")
                      div( class="body-1 d-flex w-100 justify-space-between align-center"  ) Заказывающий орган
                        v-btn( class="ml-8" icon color="grey lighten-1" @click="copyToClipboard(model.Reqorgan)")
                            v-icon mdi-content-copy
                    v-col(cols=8)
                      div( class="body-1 grey--text text--darken-4" ) {{model.Reqorgan}}
                  v-row
                    v-col(cols=4 class="pl-8")
                      div( class="body-1 d-flex w-100 justify-space-between align-center"  ) Эксплуатирующий орган
                        v-btn( class="ml-8" icon color="grey lighten-1" @click="copyToClipboard(model.Explorgan)")
                            v-icon mdi-content-copy
                    v-col(cols=8)
                      div( class="body-1 grey--text text--darken-4" ) {{model.Explorgan}}
                  v-row
                    v-col(cols=4 class="pl-8")
                      div( class="body-1 d-flex w-100 justify-space-between align-center"  ) Предприятие изготовитель
                        v-btn( class="ml-8" icon color="grey lighten-1" @click="copyToClipboard(model.Creator)")
                            v-icon mdi-content-copy
                    v-col(cols=8)
                      div( class="body-1 grey--text text--darken-4" ) {{model.Creator}}
                  
                  v-divider

                  //- Назначение
                  v-row
                    v-col(cols=4 class="pl-8")
                      div( class="body-1 d-flex w-100 justify-space-between align-center"  ) Назначение
                        v-btn( class="ml-8" icon color="grey lighten-1" @click="copyToClipboard(model.Destination)")
                            v-icon mdi-content-copy
                    v-col(cols=8)
                      div( class="body-1 grey--text text--darken-4" ) {{model.Destination}}
                  //- Состав
                  v-row
                    v-col(cols=4 class="pl-8")
                      div( class="body-1 d-flex w-100 justify-space-between align-center"  ) Состав
                        v-btn( class="ml-8" icon color="grey lighten-1" @click="copyToClipboard(model.Composition)")
                            v-icon mdi-content-copy
                    v-col(cols=8)
                      div( class="body-1 grey--text text--darken-4" ) {{model.Composition}}

</template>

<script>
  
  import axios from 'axios'; 

  export default {
    name: 'Home',

    data: () => ({
      descriptionLimit: 60,
      // what we search
      entries: [],
      // loading status
      isLoading: false,
      // model
      model: null,
      // search request
      search: null,
      items:[],
    }),

    computed: {
      fields () {
        if (!this.model) return []

        return Object.keys(this.model).map(key => {
          return {
            key,
            value: this.model[key] || 'n/a',
          }
        })
      },
      
    },

    watch: {
      search (val) {
        // Items have already been loaded
        if (this.items.length > 0) return

        // Items have already been requested
        if (this.isLoading) return

        this.isLoading = true

        // Lazily load input items
        axios.get(this.$store.state.addressprefix+'/api/search/'+val)
          .then(res => {
              if ( typeof res.data.Value!=='undefined' ){
                this.count   = res.data.Value.length
                this.entries = res.data.Value
                this.items = res.data.Value.map(_e => {
                  return _e
                } )
              }
            // console.log(this.entries)
          })
          .catch(err => {
            console.log(err)
          })
          .finally(() => (this.isLoading = false))
      },
    },
  }
</script>

<style scoped>
  .bg{
    background-image: url('../assets/img/bg.jpg') ;
    background-attachment: fixed;
    height: 100%;
  }
  .vh-100{
    height: 100%;
  }
</style>