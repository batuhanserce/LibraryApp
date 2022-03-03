<template>
  <div>
    <navbar></navbar>
    <div class="m-5">
      <b-table striped hover :items="books" ></b-table>
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
  data(){
    return{
      books:[],
    }
  },
  methods:{
    async getBook(){
    try{
      let res = await axios.get("http://127.0.0.1:3000/books")
      if (res.data ){
        res.data.forEach(deger => {
        let model={}
          console.log(deger)
          model.ID = deger.ID
          model.Name = deger.adi
          model.Author = deger.yazari
          model.Cins = deger.katagori.Cinsi
          this.books.push(model)
        })
      }

    }catch (e) {
      console.log("lan sikerim ha ")
    }
    }
  },
  created() {
  this.getBook()
  }
}
</script>

<style scoped>

</style>