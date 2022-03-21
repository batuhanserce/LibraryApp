<template>
<div class="bg">
  <section class="Form">
    <div class="container scrn ">
      <div class="row no-gutters">
        <div class="col-lg-5 p-0">
          <img src="@/assets/login/book.png" class="img-fluid" alt="" />
        </div>
        <div class="col-lg-7 px-5 pt-5 pl-5" >
          <h1 class="font-weight-bold py-3">Log-in</h1>
          <h4 >Sign in to your account</h4>
          <div>
            <div class="form-row d-flex justify-content-start">
              <div class=" col-lg-7">
                <input
                    name="Mail"
                    type="email"
                    v-model="model.eposta"
                    placeholder="Email Adress"
                    class="form-control my-3 p-4"
                />
              </div>
            </div>
            <div class="form-row d-flex justify-content-start">
              <div class="col-lg-7 ">
                <input
                    name="password"
                    type="password"
                    v-model="model.parola"
                    placeholder="********"
                    class="form-control my-3 p-4"
                />
              </div>
            </div>
            <div class="form-row d-flex justify-content-start">
              <div class="col-lg-7">
                <button @click="login" class="btn1 mt-3 mb-5">Login</button>
              </div>
            </div>
            <p class="mt-5">No account? <a href="/signup">Create one!</a></p>
          </div>
          </div>
        </div>
    </div>
  </section>
</div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      model:{
      eposta:"",
      parola:"",
      },
    }
  },
  mounted() {
    if (localStorage.name) {
      this.$router.push("/")
    }
  },
  methods: {
    async login(){
      let res = await axios.post('http://127.0.0.1:3000/login',this.model)
      if (res.data){
        console.log(res.data)
        localStorage.setItem("name",res.data.isim)
        localStorage.setItem("id",res.data.id)
        localStorage.setItem("surname",res.data.soyisim)
        localStorage.setItem("email",res.data.eposta)
       await this.$router.push("/")
      }else{
        this.$bvToast.toast("e-mail or password not true", {
          title:'Error',
          variant:'danger'})
      }
    }
  },
}
</script>

<style >
.bg {
  height: 100vh;
  background: url('../../assets/login/login.jpg') no-repeat center center fixed;
  -webkit-background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  background-size:100% 100%;
  background-size: cover;
}
* {
  justify-content: center;
  padding: 0;
  margin: 0;
  box-sizing: border-box;
}
.scrn {
  padding-top: 7%;
}
.row {
  background-color: white;
  border-radius: 30px;
  box-shadow: 12px 12px 22px grey;
}
img {
  border-top-left-radius: 30px;
  border-bottom-left-radius: 30px;
}
.btn1 {
  border: none;
  outline: none;
  height: 50px;
  width: 100%;
  background-color: black;
  color: white;
  border-radius: 4px;
  font-weight: bold;
}
.btn1:hover {
  background-color: white;
  border: 1px solid;
  color: black;
}

</style>