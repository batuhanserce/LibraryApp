<template>
  <div>
    <b-container fluid class="bg mt-0">
      <b-row align-v="stretch" class="h-100">
        <b-col lg="8" md="8" sm="0">
        </b-col>
        <b-col>
          <b-card no-body class="d-flex justify-content-center align-content-center w-100 h-100">
            <div class="p-4 p-sm-5">
              <!-- Logo -->
              <div class="d-flex justify-content-center align-items-center pb-2 mb-4">
                <div class="text-center d-flex">
                  <div class="text-center">
                    <img src="/img/hayrat.png" width="150" height="150"/>
                  </div>
                </div>
              </div>
              <!-- / Logo -->
              <!-- Title -->
              <h5 class="text-center text-muted font-weight-normal mb-4">Hayrat Vakfı Online Zikir Giriş</h5>
              <!--/Title -->
              <div>
                <b-form-group
                    id="fieldset-1"
                    label="Kod"
                    label-for="input-1"
                >
                  <b-form-input id="input-1" v-model="form.kod" :disabled="kodDisable" trim></b-form-input>
                </b-form-group>

                <b-form-group
                    id="fieldset-1"
                    label="Ad"
                    label-for="input-2"
                >
                  <b-form-input id="input-2" v-model="form.isim" :disabled="disable" trim></b-form-input>
                </b-form-group>

                <b-form-group
                    id="fieldset-1"
                    label="Soyad"
                    label-for="input-3"
                >
                  <b-form-input id="input-3" v-model="form.soyisim" :disabled="disable" trim></b-form-input>
                </b-form-group>
              </div>
              <div class="d-flex justify-content-center align-content-center">
                <b-button style="background-color: #28B4FF; border:0px; width: 100%;" size="lg" type="button"
                          @click="katil">Giriş
                </b-button>
              </div>
              <div class="d-flex justify-content-center align-content-center mt-1">
                <b-button style="background-color: darkred; border:0px; width: 100%;" size="lg" type="button"
                          v-show="disable" @click="cikisYap">Alanları Temizle
                </b-button>
              </div>
            </div>
          </b-card>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script>

export default {

  name: "home",
  data() {
    return {
      form: {
        isim: "",
        soyisim: "",
        kod: ""
      },
      isBusy: false,
      disable: false,
      kodDisable: false,
    }
  },
  methods: {

  },
  created() {
  }
}
</script>

<style scoped>
</style>
