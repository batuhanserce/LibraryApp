<template>
  <div class="bg">
    <section class="Form">
      <div class="container scrn ">
        <div class="row no-gutters">
          <div class="col-lg-5 p-0">
            <img src="@/assets/signup/signup.jpg" class="img-fluid" alt="" />
          </div>
          <div class="col-lg-7 px-5 pt-5 pl-5" >
            <h1 class="font-weight-bold py-3">Sign-up</h1>
            <h4 >Sign up to new account</h4>
            <div>
              <div class="form-row d-flex justify-content-between">
                <div class=" col-lg-6">
                  <input
                      name="Name"
                      type="text"
                      v-model="model.isim"
                      placeholder="Name"
                      class="form-control my-3 p-4"
                  />
                </div>
                <div class=" col-lg-6">
                  <input
                      name="Surname"
                      type="text"
                      v-model="model.soyisim"
                      placeholder="Surname"
                      class="form-control my-3 p-4"
                  />
                </div>
              </div>
              <div class="form-row d-flex justify-content-start">
                <div class="col-lg-6 ">
                  <input
                      name="mail"
                      type="email"
                      v-model="model.eposta"
                      placeholder="temp@mail.com"
                      class="form-control my-3 p-4"
                  />
                </div>
                <div class="col-lg-6 ">
                  <input
                      name="password"
                      type="password"
                      v-model="model.parola"
                      placeholder="********"
                      class="form-control my-3 p-4"
                  />
                </div>
              </div>
              <div class="form-row d-flex justify-content-center">
                <div class="col-lg-7">
                  <button @click="Signup" class="btn1 mt-5 mb-5">Sign-Up</button>
                </div>
              </div>
              <p class="mt-5">Have You an Account <a href="/login">Login</a></p>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: "Signup",
  data (){
    return{
      model:{
        isim:"",
        soyisim:"",
        eposta:"",
        parola:"",
      }
    }
  },
  methods:{
    async Signup(){
      try {
       let res = await axios.post('http://127.0.0.1:3000/sign_up', this.model)
           if (res.data){
             await this.$router.push("/login")
           }
      }catch (e) {
        console.error(e)
      }
    }
  }
}
</script>

<style >
.bg {
  height: 100vh;
  background: url('../../assets/signup/bg1.jpg') no-repeat center center fixed;
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
  box-shadow: -12px -12px 22px grey;
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