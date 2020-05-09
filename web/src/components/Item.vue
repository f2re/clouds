<template lang="pug">
  v-container( :class="$store.state.themecolor+' lighten-5'" dark fluid height )
    //- Заголовок
    vue-headful( :title="'Изделие '+item.Name" )
    
</template>

<script>
  import axios from 'axios'; 

  export default {
    name: 'Item',
    
    data: () => ({
      // item data
      item:{},
    }),
    // загружаем данные после того как перешли по ссылке
    mounted() {
      var _vue = this;
      // set loader
      _vue.loading = true;
      // load data
      axios.get('http://localhost:8000/api/item/'+this.$route.params.slug)
      .then( res => { 
        _vue.item = res.data.Value;
        _vue.loading=false;
      } );
    }
  }
</script>
