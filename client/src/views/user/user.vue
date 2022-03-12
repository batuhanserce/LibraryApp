<template>
  <div>
    <navbar></navbar>

    <b-card
        class="m-5"
        header-tag="header"
        footer="BatuhanSerce"
        border-variant="primary"
        header="User "
        header-bg-variant="primary"
        header-text-variant="white"
        align="center"
        footer-tag="footer"
    >
      <b-card no-body class="d-flex justify-content-around">
        <b-row >
          <b-col>
            <b-form-group label="Name">
              <b-form-input v-model="row.isim" placeholder="Name"></b-form-input>
            </b-form-group>
          </b-col>
          <b-col>
            <b-form-group label="Surname">
              <b-form-input v-model="row.soyisim" placeholder="Surname"></b-form-input>
            </b-form-group>
          </b-col>
        </b-row>
        <b-row >
          <b-col>
            <b-form-group label="EMail">
              <b-form-input v-model="row.eposta" placeholder="E-Mail"></b-form-input>
            </b-form-group>
          </b-col>
          <b-col>
            <b-form-group label="Password">
              <b-form-input v-model="row.parola" placeholder="Password"></b-form-input>
            </b-form-group>
          </b-col>
        </b-row>
      </b-card>
    </b-card>
  </div>
</template>

<script>

import axios from "axios";
import navbar from "@/component/navbar";

export default {
  name: "user",
  components:{navbar},
  data() {
    return {
      user_name: "",
      user_id: 0,
      user:{
        id:"",
        isim:"",
        soyisim:"",
        parola:"",
        eposta:"",
        kitaplar:[]
      },
      row: {},
    }
  },

  async created() {
    this.user_name = sessionStorage.getItem("name")
    this.user_id = this.$route.params.id;
    await this.getUser()
  },
  methods: {
    async getUser() {
      try {
        let res = await axios.get("http://127.0.0.1:3000/user/" + this.user_id)
        if (res.data) {
          this.row = res.data
        }

      } catch (e) {
        this.$bvToast.toast("User not come", {
          title: 'Error',
          variant: 'danger'
        })
      }
    },

  }
}
</script>

<style scoped>

</style>