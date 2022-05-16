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

        <div slot="emptystate">
          <b>Arama Kriterine Uygun Kayıt Bulunamadı</b>
        </div>
        <template slot="table-row" slot-scope="props">
          <span v-if="props.column.field === 'islem'">
            <b-col class="d-flex justify-content-around">
            <b-btn pill variant="danger" @click="deleteClick(props.row)">Delete</b-btn>
            </b-col>
            </span>
          <span v-if="props.column.field === 'name'">
            {{props.row.kullanici.isim}}
        </span>
          <span v-if="props.column.field === 'surname'">
            {{props.row.kullanici.soyisim}}
        </span>
          <span v-if="props.column.field === 'books'">
            {{props.row.kitap.adi}} - ({{props.row.kitap.yazar}})
        </span>
          <span v-else> {{ props.formattedRow[props.column.field] }} </span>
        </template>
      </vue-good-table>
    </div>
  </div>
</template>

<script>
import navbar from "@/component/navbar";
import axios from "axios";

export default {
  name: "favori",
  components: {
    navbar,
  },
  data() {
    return {
      isim:"favori",
      isBusy:false,
      katagoriler:[],
      columns: [
        {
          label: 'Name',
          field: 'name',
          sortable: true,
          filterOptions: {
            enabled: true,
            placeholder: "Name",
            filterValue: null,
            trigger: "enter",
          },
        },
        {
          label: 'Surname',
          field: 'surname',
          sortable: true,
          filterOptions: {
            enabled: true,
            placeholder: "Surname",
            filterValue: null,
            trigger: "enter",
          },
        },
        {
          label: 'Books',
          field: 'books',
          sortable: true,
          filterOptions: {
            enabled: true,
            placeholder: "E-Mail",
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
        kullanici_id:0,
        kitap_id:0,
      },
      rows: [],
    }
  },
  async created() {
    await this.getFavori();
  },
  methods: {
    async getFavori() {
      try {
        let res = await axios.get("http://127.0.0.1:3000/favori")
        if (res.data) {
          console.log(res)
          this.rows = res.data
        }
      } catch (e) {
        this.$bvToast.toast("User not come", {
          title:'Error',
          variant:'danger'})
      }
    },

    async deleteClick(model){
      this.isBusy = true
      try {
        await axios.delete("http://localhost:3000/favori", {data: model});
        this.$bvToast.toast("favori Deleted", {
          title:'Success',
          variant:'success'})
        await this.getFavori();
      }catch (e) {
        this.$bvToast.toast(e, {
          title:'Error',
          variant:'danger'})
      }
    },
  },


}
</script>

<style scoped>

</style>