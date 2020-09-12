<template lang="pug">
  v-container( height="100%" class="bg" fluid )
    //- Заголовок
    vue-headful( title="Справочник облаков" )

    v-row( align="center" justify="center" class="vh-100" )
      v-col( cols=6 align="center" justify="center" class="mt-5")
        h1( :class="$store.state.themecolor+'--text'" class="text--lighten-5 mb-5" ) Справочник облаков
        

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