<template>
  <div>
    <navbar></navbar>
    <div class="m-5">
      <vue-good-table
          :line-numbers="true"

          :pagination-options="{
            enabled: true,
            perPage: 30,
            perPageDropdown: [100, 500, 1000],
            allLabel: 'Tüm Kayıtlar',
            rowsPerPageLabel: 'Sayfada Kaç Kayıt Gösterilecek',
            nextLabel: 'Sonraki',
            prevLabel: 'Önceki',
          }"
          :columns="columns"
          :rows="rows">
        <div slot="table-actions" class="d-flex m-2">
          <b-btn variant="primary" @click="yeniClick"> <b-icon  class="mx-2" icon="book" font-scale="1"></b-icon>Add Book</b-btn>
        </div>
        <div slot="emptystate">
          <b>Arama Kriterine Uygun Kayıt Bulunamadı</b>
        </div>
        <template slot="table-row" slot-scope="props">
          <span v-if="props.column.field === 'islem'"> </span>
          <span v-if="props.column.field === 'islem'">
            <b-col class="d-flex justify-content-around">
            <b-btn pill variant="success" @click="duzenleClick(props.row)">Update</b-btn>
            <b-btn pill variant="danger" @click="deleteClick(props.row)">Delete</b-btn>
            </b-col>
            </span>
          <span v-if="props.column.field === 'katagori_id'">
        {{ getKatagoriString(props.row.katagori_id) }}
        </span>
          <span v-else> {{ props.formattedRow[props.column.field] }} </span>
        </template>
      </vue-good-table>
      <b-modal id="editModal" size="lg" hide-footer>
        <div slot="modal-title">
          {{ model.id === 0 ? "Yeni" : "" }} {{ isim }} {{ model.id === 0 ? "" : "Güncelle" }}
          <br />
        </div>
        <b-card class="mb-3">
          <b-form-row>
            <b-form-group label="İsim" class="col">
              <b-input :disabled="isBusy"  v-model.trim="model.adi" />
            </b-form-group>
          </b-form-row>
          <b-form-row>
            <b-form-group label="Yazar" class="col">
              <b-input :disabled="isBusy"  v-model.trim="model.yazar" />
            </b-form-group>
          </b-form-row>
          <b-form-row>
            <b-form-group label="Katagori" class="col">
              <b-form-select :disabled="isBusy" :options="katagoriler" v-model="model.katagori_id">
                <template v-slot:first>
                  <b-form-select-option :value="0" disabled>Seçiniz</b-form-select-option>
                </template>
              </b-form-select>
            </b-form-group>
          </b-form-row>
          <hr />
          <b-form-row>
            <b-form-group class="col text-right">
              <b-button @click="kaydetClick" :disabled="isBusy" variant="success">{{ "Kaydet" }}</b-button>
            </b-form-group>
          </b-form-row>
        </b-card>
      </b-modal>
    </div>
  </div>
</template>

<script>
import navbar from "@/component/navbar";
import axios from "axios";

export default {
  name: "book",
  components: {
    navbar,
  },
  data() {
    return {
      isim:"Book",
      isBusy:false,
      katagoriler:[],
      columns: [
        {
          label: 'Name',
          field: 'adi',
          sortable: true,
          filterOptions: {
            enabled: true,
            placeholder: "Name",
            filterValue: null,
            trigger: "enter",
          },
        },
        {
          label: 'Author',
          field: 'yazar',
          sortable: true,
          filterOptions: {
            enabled: true,
            placeholder: "Author",
            filterValue: null,
            trigger: "enter",
          },
        },
        {
          label: 'Katagori',
          field: 'katagori_id',
          sortable: true,
          filterOptions: {
            enabled: true,
            placeholder: "Katagori",
            filterValue: null,
            filterDropdownItems: [],
            trigger: "enter",
          },
        },
        {
          label:'İşlemler',
          field:'islem',
        },
      ],
      model:{
        id:"",
        adi:"",
        yazar:"",
        katagori_id:0,
      },
      rows: [],
    }
  },
  async created() {
   await this.getKatagori();
   await this.getBook();
  },
  methods: {
    async getBook() {
      try {
        let res = await axios.get("http://127.0.0.1:3000/books")
        if (res.data) {
          this.rows = res.data
        }

      } catch (e) {
        this.$bvToast.toast("Books not come", {
          title:'Error',
          variant:'danger'})
      }
    },
    async getKatagori(){
      try {
        let res = await axios.get("http://127.0.0.1:3000/katagori")
        if (res.data) {
          this.katagoriler = res.data.map(katagori => {
            return {value : katagori.ID, text: katagori.cinsi}
          })
          this.columns.find((x) => x.field === "katagori_id").filterOptions.filterDropdownItems = this.katagoriler;
        }
      } catch (e) {
        this.$bvToast.toast("Katagori not comde", {
          title:'Error',
          variant:'danger'})
      }
    },
    async yeniClick() {
      this.$bvModal.show("editModal");
    },
    async duzenleClick(model) {
      this.model = model;
      this.$bvModal.show("editModal");
    },
    async kaydetClick() {
      this.isBusy = true;
      try {
        await axios.post("http://localhost:3000/book", this.model);
        this.$bvModal.hide("editModal");
        this.$bvToast.toast("Book Created", {
          title:'Success',
          variant:'success'})
        await this.getBook();
      } catch (e) {
        console.error(e);
      }
      this.isBusy = false;
    },
    async deleteClick(model){
      this.isBusy = true
      try {
        await axios.delete("http://localhost:3000/book", {data: model});
        this.$bvToast.toast("Book Deleted", {
          title:'Success',
          variant:'success'})
        await this.getBook();
      }catch (e) {
        this.$bvToast.toast(e, {
          title:'Error',
          variant:'danger'})
      }
    },
    getKatagoriString(id){
      return this.katagoriler.find((x) => x.value === id).text
    }
  },


}
</script>

<style scoped>

</style>