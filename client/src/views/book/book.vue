<template>
  <div>
    <navbar></navbar>
    <div class="container-fluid pr-5 pl-5 pt-3">
      <div>
        <b-card no-body class="overflow-hidden w-100" style=" box-shadow: 15px 20px #e6e6e6; border-radius: 10px">
          <b-row no-gutters>
            <b-col md="6">
              <b-card-img
                  src="https://images.pexels.com/photos/448835/pexels-photo-448835.jpeg?cs=srgb&dl=pexels-victor-448835.jpg&fm=jpg"
                  alt="Image" class="rounded-0"></b-card-img>
            </b-col>
            <b-col md="6">
              <b-card-body :title=book.adi>
                <b-card-text>
                  <b-row class="m-2">
                    {{ book.ozet }}
                  </b-row>
                  <b-row class="m-3 d-flex justify-content-between">
                    <b-col class="d-flex justify-content-center"><b>Yazar : </b> {{ book.yazar }}</b-col>
                    <b-col><b>Katagori : </b>{{ book.katagori.cinsi }}</b-col>
                  </b-row>
                  <b-row class="m-3">
                    <b-button v-if="favoriMi" variant="danger" @click="favorilerdenCikar">Favorilerden çıkar</b-button>
                    <b-button v-else variant="success" @click="favorilererEkle">Favorilere ekle</b-button>
                  </b-row>
                </b-card-text>
              </b-card-body>
            </b-col>
          </b-row>
        </b-card>
      </div>
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
      book: {
        ID: 0,
        adi: "",
        katagori: {
          cinsi: "",
        },
      },
      id: 0,
      favoriMi: false,
    }
  },
  async created() {
    this.id = this.$route.params
    await this.getBook()
    await this.getFavori()
  },
  methods: {
    async getFavori() {
      this.isBusy = true
      let userID = localStorage.getItem("id")
      try {
        let res = await axios.get("http://127.0.0.1:3000/favori/" + this.book.ID + "/" + userID)
        if (res.data) {
          this.favoriMi = true
        }
      } catch (e) {
        console.log(e)
      }
      this.isBusy = false
    },
    async getBook() {
      this.isBusy = true
      try {
        let res = await axios.get("http://127.0.0.1:3000/book/" + this.id.id)
        if (res.data) {
          this.book = res.data;
        }
        console.log(res.data)

      } catch (e) {
        console.log(e)
        this.$bvToast.toast("Book not come", {
          title: 'Error',
          variant: 'danger'
        })
      }
      this.isBusy = false
    },
    async favorilererEkle() {
      this.isBusy = true
      try {
        let favoriEkle = {
          kullanici_id: +localStorage.getItem("id"),
          kitap_id: +this.book.ID,
        }
        let res = await axios.post("http://127.0.0.1:3000/favori/", favoriEkle)
        if (res.data) {
          this.book = res.data;
        }
        location.reload();
      } catch (e) {
        this.$bvToast.toast("Book not come", {
          title: 'Error',
          variant: 'danger'
        })
      }
      this.isBusy = false
    },
    async favorilerdenCikar() {
      this.isBusy = true
      let favoriCikar = {
        kullanici_id: +localStorage.getItem("id"),
        kitap_id: +this.book.ID,
      }
      try {
        let res = await axios.delete("http://127.0.0.1:3000/favori/", {data: favoriCikar})
        if (res.data) {
          this.book = res.data;
        }
        location.reload();
      } catch (e) {
        this.$bvToast.toast("Book not come", {
          title: 'Error',
          variant: 'danger'
        })
      }
      this.isBusy = false
    },

  },
}
</script>

<style scoped>

</style>