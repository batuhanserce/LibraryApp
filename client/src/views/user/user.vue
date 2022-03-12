<template>
  <div>
    <b-card
        class="text-center m-5 "
        header-tag="header"
        footer="Card Footer"
        border-variant="primary"
        header="User "
        header-bg-variant="primary"
        header-text-variant="white"
        align="center"
        footer-tag="footer"
    >
      <b-card no-body class="overflow-hidden d-flex justify-content-center" style="max-width: 540px;">
        <b-row no-gutters>
          <b-col md="6">
            <b-card-img src="" alt="Image" class="rounded-0"></b-card-img>
          </b-col>
          <b-col md="6">
            <b-card-body title="Horizontal Card">
              <b-card-text>
                This is a wider card with supporting text as a natural lead-in to additional content.
                This content is a little bit longer.
              </b-card-text>
            </b-card-body>
          </b-col>
        </b-row>
      </b-card>
    </b-card>
  </div>
</template>

<script>

import axios from "axios";

export default {
  name: "user",
  data() {
    return {
      user_name: "",
      user_id: 0,
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
          console.log(this.row)
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