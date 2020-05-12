<template lang="pug">
  v-container( :class="$store.state.themecolor+' lighten-5'" dark fluid height )
    //- Заголовок
    vue-headful( title="Добавление изделия в справочник" )

    v-row( align="center" justify="center" class="vh-100" )
      v-col( cols=8  justify="center" class="mt-5") 
        v-card( :loading="loading" :disabled="loading" class="mb-7 "   )
            div( class="d-flex align-center justify-center grey theme--light lighten-3 item-placeholder-img relative " )
                //- image uploaded
                v-img( v-if="item.image!=''" :src="$store.state.addressprefix+item.image" )
                //- image placeholder
                div(v-if="item.image==''")
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
                            v-text-field( v-model="item.name" :counter="200" :rules="reqRules" label="Наименование изделия" 
                                        prepend-icon="mdi-sign-text"
                                        required clearable )
                    v-row
                        v-col(cols=12)
                            v-select( :items="$store.state.categories" v-model="item.category"
                                      item-text="text" item-value="val"
                                      prepend-icon="mdi-format-list-bulleted-type"
                                      label="Категория техники" )
                                
                    v-row                        
                        //- принят на снабжение
                        v-col(cols=12)
                            v-text-field( v-model="item.snabjenie" :counter="50"  label="Принят на снабжение" 
                                        prepend-icon="mdi-calendar-check"
                                        required clearable )
                    v-row
                        //- индекс иделия
                        v-col(cols=3)
                            v-text-field( v-model="item.index" :counter="20"  label="Индекс изделия" 
                                        prepend-icon="mdi-numeric"
                                        required clearable )
                        //- код КВТ
                        v-col(cols=3)
                            v-text-field( v-model="item.kvt" :counter="20"  label="Код КВТ" 
                                        prepend-icon="mdi-numeric"
                                        required clearable )
                        //- номенклатурный номер
                        v-col(cols=6)
                            v-text-field( v-model="item.nomenklature" :counter="20"  label="Номенклатурный номер" 
                                        prepend-icon="mdi-numeric"
                                        required clearable )
                    
                    

                    v-row
                        //- довольствующий орган
                        v-col(cols=6)
                            v-textarea( v-model="item.dovorgan" filled label="Довольствующий орган" rows=2 )
                        
                        //- заказывающий орган
                        v-col(cols=6)
                            v-textarea( v-model="item.reqorgan" filled label="Заказывающий орган" rows=2 )
                    
                    v-row
                        //- эксплуатирующий орган
                        v-col(cols=6)
                            v-textarea( v-model="item.explorgan" filled label="Эксплуатирующий орган" rows=2 )
                        
                        //- изготовитель орган
                        v-col(cols=6)
                            v-textarea( v-model="item.creator" filled label="Предприятие изготовитель" rows=2 )
                    
                    v-row
                        //- Состав
                        v-col(cols=12)
                            v-textarea( filled v-model="item.composition" label="Состав изделия" rows=3 )
                    v-row
                        //- назначение
                        v-col(cols=12)
                            v-textarea( filled v-model="item.destination" label="Назначение изделия" rows=3 )

                    v-divider
                        
                    //- ТТХ
                    v-row(v-for="(_t,i) in item.tth" :key="i" class="mb-5")
                        v-col(cols=6 )
                            v-text-field( v-model="_t.label" :counter="40"  label="Название параметра ТТХ" 
                                        prepend-icon="mdi-tag"
                                        required clearable )
                        v-col(cols=5 )
                            v-text-field( v-model="_t.value" :counter="20"  label="Значение параметра ТТХ" 
                                        required clearable )
                        v-col(cols=1 class="d-flex align-center justify-center" )
                            //- удаление парамтера ТТХ
                            v-btn( @click="$delete(item.tth,i)" color="red"  text)
                                v-icon mdi-tag-remove
                    v-row
                        v-col( cols=3 offset=9 class="d-flex align-end justify-end" )
                            //- удаление парамтера ТТХ
                            v-btn( @click="item.tth.push({label:'',value:''})" color="teal"  text) Добавить
                                v-icon mdi-tag-plus
                    v-divider

                    //- нормы табелизации
                    h2( class="mt-9 mb-5" ) Нормы табелизации
                    v-row(v-if="typeof item.tabel !=='undefined'")
                        v-col
                            v-text-field( v-model="item.tabel.Escadrile" type="number" min=0 step=1 label="Отдельная эскадрилься" )
                        v-col
                            v-text-field( v-model="item.tabel.Polk" type="number" min=0 step=1 label="Авиационный полк" )
                        v-col
                            v-text-field( v-model="item.tabel.Division" type="number" min=0 step=1 label="Ав. дивизия" )
                        v-col
                            v-text-field( v-model="item.tabel.Army" type="number" min=0 step=1 label="А ВВС и ПВО" )
                        v-col
                            v-text-field( v-model="item.tabel.Center" type="number" min=0 step=1 label="Учебно авиационный центр" )
                        v-col
                            v-text-field( v-model="item.tabel.CY" type="number" min=0 step=1 label="ЦУ вида/рода" )
                        v-col
                            v-text-field( v-model="item.tabel.Polygon" type="number" min=0 step=1 label="Межвидовой полигон" )
                        v-col
                            v-text-field( v-model="item.tabel.Komend" type="number" min=0 step=1 label="Aв. комендатура" )
                        v-col
                            v-text-field( v-model="item.tabel.School" type="number" min=0 step=1 label="Учебное заведение" )
            v-card-actions(  )
                v-spacer
                v-btn( @click="reset" ) Очистить
                v-btn( @click="submit" :dark="valid" :color="$store.state.themecolor" :disabled="!valid" ) Сохранить
                    v-icon(class="ml-5") mdi-content-save
            
    v-bottom-sheet(v-model="popup")
        v-sheet(class="text-center" height="200px")
            v-btn( @click="popup= !popup" text class="mt-6" :color="popup_style" ) Закрыть
            div(class="my-3") {{popup_text}}
            
</template>

<script>
  import axios from 'axios'; 

  export default {
    name: 'AddItem',

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
        item:{},
        // info popup
        popup: false,
        // style
        popup_style:'error',
        // text
        popup_text:''
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
        // отправляем форму
        submit(){
            // сначала валидируем
            this.validate();
            // if all valid
            
            if ( this.valid ){
                let _data = new FormData();
                if( this.$refs.img.files[0] ){
                    _data.append( 'Image', this.$refs.img.files[0] );
                }
                _data.append( 'name', this.item.name );
                _data.append( 'category', this.item.category );
                _data.append( 'index', this.item.index );
                _data.append( 'snabjenie', this.item.snabjenie );
                _data.append( 'kvt', this.item.kvt );
                _data.append( 'nomenklature', this.item.nomenklature );
                _data.append( 'dovorgan', this.item.dovorgan );
                _data.append( 'reqorgan', this.item.reqorgan );
                _data.append( 'explorgan', this.item.explorgan );
                _data.append( 'creator', this.item.creator );
                _data.append( 'description', this.item.description );
                _data.append( 'destination', this.item.destination );
                _data.append( 'composition', this.item.composition );
                _data.append( 'tth', JSON.stringify(this.item.tth) );
                _data.append( 'tabel', JSON.stringify(this.item.tabel) );

                axios.post( this.$store.state.addressprefix+'/api/items', _data )
                     .then( ()=>{
                        this.popup = true;
                        this.popup_style = 'success';
                        this.popup_text = 'Изделие добавлено. Можно перейти к просмотру';
                        // console.log(res);
                        this.resetData();
                     })
                     .catch(error => {
                        console.log(error.response)
                        this.popup = true;
                        this.popup_style = 'error';
                        this.popup_text = 'Ошибка добавления "'+error.response.data+'"';
                     });
            }
        },
        resetData(){
            this.item = {
                image       : '',
                name        : '',
                category    : 1,
                index       : '',
                snabjenie   : '',
                kvt         : '',
                nomenklature: '',
                dovorgan    : '',
                reqorgan    : '',
                explorgan   : '',
                creator     : '',
                description : '',
                destination : '',
                composition : '',
                tth         : [
                    {label:'Объем ГМИ поступающей в сутки, МБайт',
                    value:'200'},
                ],
                tabel       : {
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
                images      : '',
            }
        },
        // upload image
        uploadImg(){
            // console.log(ev);
            // console.log(this.$refs.img.files[0]);
            this.item.image = URL.createObjectURL(this.$refs.img.files[0]);
        }
    },
    // загружаем данные после того как перешли по ссылке
    mounted() {
        // обнуляем данные
        this.resetData();
    }
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