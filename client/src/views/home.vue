<template>
  <div>
    <Navbar></Navbar>
    <div class="container-fluid">
    <b-row class="m-5">
      <b-col lg="3" v-for="(book,index) in books" :key="index"><div
      >
        <b-card
            :title="book.adi"
            img-src="https://images.pexels.com/photos/415071/pexels-photo-415071.jpeg?cs=srgb&dl=pexels-pixabay-415071.jpg&fm=jpg"
            img-alt="Image"
            img-top
            tag="article"
            style="max-width: 20rem;"
            class="mb-2"
        >
          <b-card-text>
            {{ book.ozet }}
          </b-card-text>
          <b-btn variant="success" :to="'book/'+book.ID">Detaylar</b-btn>
            <template #footer >
              <div class="d-flex justify-content-between">
              <em>Yazar : {{book.yazar}} </em>
              <em>Cinsi : {{book.katagori.cinsi}} </em>
              </div>
            </template>
        </b-card>
      </div></b-col>
    </b-row>
    </div>

  </div>
</template>

<script>
import Navbar from "@/component/navbar";
import axios from "axios";
export default {
  components: {
    Navbar
  },
  name: "home",
  data() {
    return {
      form: {
        isim: "",
        soyisim: "",
        kod: ""
      },
      books:[],
      isBusy: false,
      disable: false,
      kodDisable: false,
    }
  },
  async created() {
    if (!localStorage.name) {
      this.$router.push("/login")
    }
    await this.getBooks();
  },
  methods:{
     async getBooks(){
        this.isBusy= true
        try {
          let res = await axios.get("http://127.0.0.1:3000/book")
          if (res.data) {
            this.row = res.data
          }
          this.books= res.data;
          console.log(res.data);
        }catch (e) {
          console.log(e)
          this.$bvToast.toast("Books not come", {
            title: 'Error',
            variant: 'danger'
          })
        }
        this.isBusy= false

      },
  },

}
</script>

<style scoped>
</style>
