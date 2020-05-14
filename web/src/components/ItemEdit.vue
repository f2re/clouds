<template lang="pug">
  v-container( :class="$store.state.themecolor+' lighten-5'" dark fluid height )
    //- Заголовок
    vue-headful( title="Добавление изделия в справочник" )

    v-row( align="center" justify="center" class="vh-100" )
      v-col( cols=8  justify="center" class="mt-5") 
        v-card( :loading="loading" :disabled="loading" class="mb-7 "   )
            div( class="d-flex align-center justify-center grey theme--light lighten-3 item-placeholder-img relative " )
                //- image uploaded
                v-img( v-if="item.Image!=''" :src="(isblob?'':$store.state.addressprefix)+item.Image" )
                //- image placeholder
                div(v-if="item.Image==''")
                    v-icon( size="84px") mdi-image
                //- upload image button
                v-btn(  dark bottom right absolute fab
                    @click="$refs.img.click()"
                    :color="$store.state.themecolor" )
                    v-icon mdi-camera
            
            v-card-title Добавление изделия
            v-card-subtitle заполните все поля, установите изображение и нажмите кнопку "сохранить"
            v-card-text 
                v-form( ref="form" v-model="valid" lazy-validation  )
                    input( class="d-none" type="file" @change="uploadImg" ref="img" id="img" )
                        
                    //- название
                    v-row
                        v-col(cols=12)
                            v-text-field( v-model="item.Name" :counter="200" :rules="reqRules" label="Наименование изделия" 
                                        prepend-icon="mdi-sign-text"
                                        required clearable )
                    v-row
                        v-col(cols=12)
                            v-select( :items="$store.state.categories" v-model="item.Category"
                                      item-text="text" item-value="val"
                                      prepend-icon="mdi-format-list-bulleted-type"
                                      label="Категория техники" )
                    v-row                        
                        //- принят на снабжение
                        v-col(cols=12)
                            v-text-field( v-model="item.Snabjenie" :counter="50"  label="Принят на снабжение" 
                                        prepend-icon="mdi-calendar-check"
                                        required clearable )
                    v-row
                        //- индекс иделия
                        v-col(cols=3)
                            v-text-field( v-model="item.Index" :counter="20"  label="Индекс изделия" 
                                        prepend-icon="mdi-numeric"
                                        required clearable )
                        //- код КВТ
                        v-col(cols=3)
                            v-text-field( v-model="item.KVT" :counter="20"  label="Код КВТ" 
                                        prepend-icon="mdi-numeric"
                                        required clearable )
                        //- номенклатурный номер
                        v-col(cols=6)
                            v-text-field( v-model="item.Nomenklature" :counter="20"  label="Номенклатурный номер" 
                                        prepend-icon="mdi-numeric"
                                        required clearable )
                    
                    v-row
                        //- довольствующий орган
                        v-col(cols=6)
                            v-textarea( v-model="item.Dovorgan" filled label="Довольствующий орган" rows=2 )
                        
                        //- заказывающий орган
                        v-col(cols=6)
                            v-textarea( v-model="item.Reqorgan" filled label="Заказывающий орган" rows=2 )
                    
                    v-row
                        //- эксплуатирующий орган
                        v-col(cols=6)
                            v-textarea( v-model="item.Explorgan" filled label="Эксплуатирующий орган" rows=2 )
                        
                        //- изготовитель орган
                        v-col(cols=6)
                            v-textarea( v-model="item.Creator" filled label="Предприятие изготовитель" rows=2 )
                    
                    v-row
                        //- Состав
                        v-col(cols=12)
                            v-textarea( filled v-model="item.Composition" label="Состав изделия" rows=3 )
                    v-row
                        //- назначение
                        v-col(cols=12)
                            v-textarea( filled v-model="item.Destination" label="Назначение изделия" rows=3 )

                    v-divider
                        
                    //- ТТХ
                    v-row(v-for="(_t,i) in item.TTH" :key="i" class="mb-5")
                        v-col(cols=6 )
                            v-text-field( v-model="_t.label" :counter="40"  label="Название параметра ТТХ" 
                                        prepend-icon="mdi-tag"
                                        required clearable )
                        v-col(cols=5 )
                            v-text-field( v-model="_t.value" :counter="20"  label="Значение параметра ТТХ" 
                                        required clearable )
                        v-col(cols=1 class="d-flex align-center justify-center" )
                            //- удаление парамтера ТТХ
                            v-btn( @click="$delete(item.TTH,i)" color="red"  text)
                                v-icon mdi-tag-remove
                    v-row
                        v-col( cols=3 offset=9 class="d-flex align-end justify-end" )
                            //- удаление парамтера ТТХ
                            v-btn( @click="item.TTH.push({label:'',value:''})" color="teal"  text) Добавить
                                v-icon mdi-tag-plus
                    v-divider

                    //- нормы табелизации
                    h2( class="mt-9 mb-5" ) Нормы табелизации
                    v-row( v-if="typeof item.Tabel !== 'undefined' " )
                        v-col
                            v-text-field( v-model="item.Tabel.Escadrile" type="number" min=0 step=1 label="Отдельная эскадрилься" )
                        v-col
                            v-text-field( v-model="item.Tabel.Polk" type="number" min=0 step=1 label="Авиационный полк" )
                        v-col
                            v-text-field( v-model="item.Tabel.Division" type="number" min=0 step=1 label="Ав. дивизия" )
                        v-col
                            v-text-field( v-model="item.Tabel.Army" type="number" min=0 step=1 label="А ВВС и ПВО" )
                        v-col
                            v-text-field( v-model="item.Tabel.Center" type="number" min=0 step=1 label="Учебно авиационный центр" )
                        v-col
                            v-text-field( v-model="item.Tabel.CY" type="number" min=0 step=1 label="ЦУ вида/рода" )
                        v-col
                            v-text-field( v-model="item.Tabel.Polygon" type="number" min=0 step=1 label="Межвидовой полигон" )
                        v-col
                            v-text-field( v-model="item.Tabel.Komend" type="number" min=0 step=1 label="Aв. комендатура" )
                        v-col
                            v-text-field( v-model="item.Tabel.School" type="number" min=0 step=1 label="Учебное заведение" )
            v-card-actions(  )

                v-btn( @click="deleteItem" text color="red" ) Удалить
                    v-icon(class="ml-5") mdi-delete
                v-spacer
                v-btn( @click="submit" :dark="valid" :color="$store.state.themecolor" :disabled="!valid" ) Сохранить
                    v-icon(class="ml-5") mdi-content-save

    v-bottom-sheet(v-model="popup" persistent)
        v-sheet(class="text-center" height="150px")
            v-btn( v-if="popup_style=='error'" @click="popup= !popup" text class="mt-6" :color="popup_style" ) Закрыть
            v-btn( v-if="popup_style!='error'" @click="popup= !popup" text class="mt-6" :color="popup_style" :to="newslug" ) Закрыть
                v-icon(class="ml-4") mdi-open-in-new
            div(class="my-3") {{popup_text}}

        
</template>

<script>
  import axios from 'axios'; 

  export default {
    name: 'ItemEdit',

    data: () => ({
        // is valid form
        valid: false,
        // loader
        loading: false,
        // required rules
        reqRules:[
            v => !!v || 'Заполните поле',
        ],
        // item
        item:{
            Image       : '',
            Name        : '',
            Category    : 1,
            Index       : '',
            Snabjenie   : '',
            KVT         : '',
            Nomenklature: '',
            Dovorgan    : '',
            Reqorgan    : '',
            Explorgan   : '',
            Creator     : '',
            Description : '',
            Destination : '',
            Composition : '',
            TTH         : [
                {label:'Объем ГМИ поступающей в сутки, МБайт',
                 value:'200'},
            ],
            Tabel       : {
                Escadrile :0,
                Polk :0,
                Division :0,
                Army :0,
                Center :0,
                CY :0,
                Polygon :0,
                Komend :0,
                School :0,
            },
            Images      : '',
        },
        // info popup
        popup: false,
        // style
        popup_style:'error',
        // text
        popup_text:'',
        // newslug
        newslug:'',
        // image blobbing whitout prefix
        isblob: false,
    }),

    methods:{
        // validate form
        validate(){
            this.$refs.form.validate()
        },
        // reset form
        reset () {
            this.$refs.form.reset()
        },
        // удаляем 
        deleteItem(){
            axios.delete( this.$store.state.addressprefix+'/api/item/'+this.$route.params.slug )
                .then( ()=>{
                    this.popup       = true;
                    this.popup_style = 'success';
                    this.popup_text  = 'Карточка изделия удалена!';
                    this.newslug     = '/list/';
                    // console.log(res);
                    this.resetData();
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
        // отправляем форму
        submit(){
            // сначала валидируем
            this.validate();
            // if all valid
            // console.log(this.item.image);
            if ( this.valid ){
                let _data = new FormData();
                if ( this.$refs.img.files[0] ){
                    _data.append( 'Image', this.$refs.img.files[0] );
                }
                _data.append( 'name', this.item.Name );
                _data.append( 'slug', this.item.Slug );
                _data.append( 'category', this.item.Category );
                _data.append( 'index', this.item.Index );
                _data.append( 'snabjenie', this.item.Snabjenie );
                _data.append( 'kvt', this.item.KVT );
                _data.append( 'nomenklature', this.item.Nomenklature );
                _data.append( 'dovorgan', this.item.Dovorgan );
                _data.append( 'reqorgan', this.item.Reqorgan );
                _data.append( 'explorgan', this.item.Explorgan );
                _data.append( 'creator', this.item.Creator );
                _data.append( 'description', this.item.Description );
                _data.append( 'destination', this.item.Destination );
                _data.append( 'composition', this.item.Composition );
                _data.append( 'tth', JSON.stringify(this.item.TTH) );
                _data.append( 'tabel', JSON.stringify(this.item.Tabel) );

                axios.put( this.$store.state.addressprefix+'/api/item/'+this.$route.params.slug, _data )
                     .then( (res)=>{
                        this.popup       = true;
                        this.popup_style = 'success';
                        this.popup_text  = 'Изменения сохранены!';
                        this.newslug     = '/item/'+res.data.URL;
                        // console.log(res);
                        this.resetData();
                     })
                     .catch(error => {
                        // console.log(error.response)
                        if ( typeof error.response !=='undefined' ){
                            this.popup       = true;
                            this.popup_style = 'error';
                            this.popup_text  = 'Ошибка сохранения изменений "'+error.response.data+'"';
                        }
                     });
            }
        },
        
        // upload image
        uploadImg(){
            this.isblob     = true;
            this.item.Image = URL.createObjectURL(this.$refs.img.files[0]);
        }
    },
    // загружаем данные после того как перешли по ссылке
    mounted() {
        var _vue = this;
        // устанавливаем путь по тому поути, по которому зашли
        _vue.newslug = this.$route.params.slug;

        // set loader
        _vue.loading = true;
        // load data
        axios.get(this.$store.state.addressprefix+'/api/item/'+this.$route.params.slug)
        .then( res => { 
            _vue.item = res.data.Value;
            if ( typeof _vue.item.TTH !=='undefined' ){
                _vue.item.TTH  = JSON.parse( _vue.item.TTH );
            }
            if ( typeof _vue.item.Image !=='undefined' ){
                _vue.item.Image  = _vue.item.Image.Path;
            }
            _vue.loading=false;
        } );
    },
  }
</script>
<style scoped>
.height{
    height: 100%;
}
.item-placeholder-img{
    height:300px;
}
.item-placeholder-img .v-image{
    max-height: 300px;
}
.relative{
    position: relative;
}
</style>