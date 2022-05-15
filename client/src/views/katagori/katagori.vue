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
          <b-btn variant="primary" @click="yeniClick">
            <b-icon class="mx-2" icon="people" font-scale="1"></b-icon>
            Add Katagori
          </b-btn>
        </div>
        <div slot="emptystate">
          <b>Arama Kriterine Uygun Kayıt Bulunamadı</b>
        </div>
        <template slot="table-row" slot-scope="props">
          <span v-if="props.column.field === 'islem'"> </span>
          <span v-else> {{ props.formattedRow[props.column.field] }} </span>
        </template>
      </vue-good-table>
      <b-modal id="editModal" size="lg" hide-footer>
        <div slot="modal-title">
          {{ model.id === 0 ? "Yeni" : "" }} {{ isim }} {{ model.id === 0 ? "" : "Güncelle" }}
          <br/>
        </div>
        <b-card class="mb-3">
          <b-form-row>
            <b-form-group label="Name" class="col">
              <b-input :disabled="isBusy" v-model.trim="model.cinsi"/>
            </b-form-group>
          </b-form-row>
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
import 'vue-select/dist/vue-select.css';
import axios from "axios";

export default {
  name: "Katagori",
  components: {
    navbar,
  },
  data() {
    return {
      isim: "Katagori",
      isBusy: false,
      katagoriler: [],
      columns: [
        {
          label: 'Cinsi',
          field: 'cinsi',
          sortable: true,
          filterOptions: {
            enabled: true,
            placeholder: "Cinsi",
            filterValue: null,
            trigger: "enter",
          },
        },

      ],
      model: {
        id: 0,
        cinsi: "",
      },
      rows: [],
    }
  },
  async created() {
    await this.getKatagori();
  },
  methods: {
    async getKatagori() {
      try {
        let res = await axios.get("http://127.0.0.1:3000/katagori")
        if (res.data) {
          this.rows = res.data
        }
      } catch (e) {
        this.$bvToast.toast("User not come", {
          title: 'Error',
          variant: 'danger'
        })
      }
    },
    async yeniClick() {
      this.model.cinsi = "";
      this.$bvModal.show("editModal");
    },
    async duzenleClick(model) {
      this.model = model;
      this.$bvModal.show("editModal");
    },
    async kaydetClick() {
      this.isBusy = true;
      try {
        await axios.post("http://localhost:3000/katagori", this.model);
        this.$bvModal.hide("editModal");
        this.$bvToast.toast("Katagori Created", {
          title: 'Success',
          variant: 'success'
        })
        await this.getKatagori()
      } catch (e) {
        console.error(e);
      }
      this.isBusy = false;
    },
    async deleteClick(model) {
      this.isBusy = true
      try {
        await axios.delete("http://localhost:3000/katagori", {data: model});
        await this.getUsers()
        this.$bvToast.toast("user Deleted", {
          title: 'Success',
          variant: 'success'
        })
      } catch (e) {
        this.$bvToast.toast(e, {
          title: 'Error',
          variant: 'danger'
        })
      }
    },
  },


}
</script>

<style scoped>

</style>