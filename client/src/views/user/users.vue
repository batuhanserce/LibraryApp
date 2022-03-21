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
            Add User
          </b-btn>
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
            <b-btn pill variant="info" :to="'user/'+props.row.ID"> kullaniciya git</b-btn>
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
          <br/>
        </div>
        <b-card class="mb-3">
          <b-form-row>
            <b-form-group label="Name" class="col">
              <b-input :disabled="isBusy" v-model.trim="model.isim"/>
            </b-form-group>
          </b-form-row>
          <b-form-row>
            <b-form-group label="Surname" class="col">
              <b-input :disabled="isBusy" v-model.trim="model.soyisim"/>
            </b-form-group>
          </b-form-row>
          <b-form-row>
            <b-form-group label="E-Mail" class="col">
              <b-input :disabled="isBusy" v-model.trim="model.eposta"/>
            </b-form-group>
          </b-form-row>
          <b-form-row>
            <b-form-group label="Password" class="col">
              <b-input :disabled="isBusy" v-model.trim="model.parola"/>
            </b-form-group>
          </b-form-row>
          <hr/>
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
  name: "user",
  components: {
    navbar,
  },
  data() {
    return {
      isim: "user",
      isBusy: false,
      katagoriler: [],
      columns: [
        {
          label: 'Name',
          field: 'isim',
          sortable: true,
          filterOptions: {
            enabled: true,
            placeholder: "Name",
            filterValue: null,
            trigger: "enter",
          },
        },
        {
          label: 'Lastname',
          field: 'soyisim',
          sortable: true,
          filterOptions: {
            enabled: true,
            placeholder: "Lastname",
            filterValue: null,
            trigger: "enter",
          },
        },
        {
          label: 'E-Mail',
          field: 'eposta',
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
          label: 'İşlemler',
          field: 'islem',
        },
      ],
      model: {
        id: "",
        isim: "",
        soyisim: "",
        eposta: "",
        parola: "",
        kullanici_kitap: [],
      },
      rows: [],
    }
  },
  async created() {
    await this.getUsers();
  },
  methods: {
    async getUsers() {
      try {
        let res = await axios.get("http://127.0.0.1:3000/user")
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
      this.model.isim = "";
      this.model.soyisim = "";
      this.model.eposta = "";
      this.model.parola = "";
      this.model.kullanici_kitap = [];
      this.$bvModal.show("editModal");
    },
    async duzenleClick(model) {
      this.model = model;
      this.$bvModal.show("editModal");
    },
    async kaydetClick() {
      this.isBusy = true;
      try {
        await axios.post("http://localhost:3000/user", this.model);
        this.$bvModal.hide("editModal");
        this.$bvToast.toast("User Created", {
          title: 'Success',
          variant: 'success'
        })
        await this.getUser();
      } catch (e) {
        console.error(e);
      }
      this.isBusy = false;
    },
    async deleteClick(model) {
      this.isBusy = true
      try {
        await axios.delete("http://localhost:3000/user", {data: model});
        this.$bvToast.toast("user Deleted", {
          title: 'Success',
          variant: 'success'
        })
        await this.getUser();
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