<template lang="pug">
  v-container( height="100%" class="bg" fluid )
    //- Заголовок
    vue-headful( title="Справочник технических средств метеорологической службы" )

    v-row( align="center" justify="center" class="vh-100" )
      v-col( cols=6 align="center" justify="center" class="mt-5")
        h1( :class="$store.state.themecolor+'--text'" class="text--lighten-5 mb-5" ) Справочник технических средств метеорологической службы
        
        v-autocomplete( 
                        :search-input.sync="search"
                        v-model="model"
                        :items="items"
                        :loading="isLoading"
                        color="white" 
                        hide-no-data 
                        hide-selected 
                        item-text="Описание"
                        item-value="API"
                        label="Технические средства"
                        placeholder="Начните ввод, чтобы искать"
                        prepend-icon="mdi-magnify"
                        class="mb-5"
                        return-object
                        dark )

        v-divider(dark)

</template>

<script>
  export default {
    name: 'Home',

    data: () => ({
      // what we search
      entries: [],
      // loading status
      isLoading: false,
      // model
      model: null,
      // search request
      search: null,
      // array of items
      items : []
    }),
  }
</script>

<style scoped>
  .bg{
    background-image: url('../assets/img/bg.jpg') ;
    height: 100%;
  }
  .vh-100{
    height: 100%;
  }
</style>