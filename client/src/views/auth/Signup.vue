<template>
  <b-container>
    <b-row class="vh-100" align-v="center" align-h="center">
      <b-col sm="12" md="8" lg="8" xl="6">
        <b-card no-body>
          <div class="p-4 p-sm-5">
            <!-- Logo -->
            <div class="d-flex justify-content-center align-items-center pb-2 mb-4">
              <div class="text-center d-flex">
                <div class="text-center">
                  <img src="../../../public/img/icon.png" width="250" height="250"/>
                </div>
              </div>
            </div>
            <!-- / Logo -->
            <h5 class="text-center text-muted font-weight-normal mb-4">Libbrary App Sign-Up</h5>

            <!-- Form -->
            <form>
              <b-form-group label="Name">
                <b-input v-model="model.isim" placeholder="Name" type="text"/>
              </b-form-group>
              <b-form-group label="Surname">
                <b-input v-model="model.soyisim" placeholder="Surname" type="text"/>
              </b-form-group>
              <b-form-group label="E-Mail">
                <b-input v-model="model.eposta" placeholder="E-Mail" type="E-Mail"/>
              </b-form-group>
              <b-form-group label="Password">
                <b-input v-model="model.parola" placeholder="Password" type="password"/>
              </b-form-group>

              <div class="d-flex justify-content-center align-items-center m-0">
                <b-button type="button" @click="Signup" style="width: 25%;" class="mb-4" data-style="slide-right"
                          variant="primary">
                  Sign-Up
                </b-button>
              </div>
            </form>
            <!-- / Form -->

            <b-card-footer class="py-3 px-4 px-sm-5">
              <div class="text-center text-muted">
              </div>
            </b-card-footer>
          </div>
        </b-card>
      </b-col>
    </b-row>
  </b-container>
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
        axios.post('http://127.0.0.1:3000/sign_up', this.model)
            .then((res) => {
              console.log(res)
            })
            .catch((error) => {
              console.error(error)
            }).finally(() => {
          console.log("bura medem bar")
        });
      }catch (e) {
        console.error(e)
      }
    }
  }
}
</script>

<style scoped>

</style>