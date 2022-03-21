<template>
    <b-navbar toggleable="lg" type="dark" variant="success" class="navbar-horizontal">
      <b-navbar-brand to="/">Liberary App</b-navbar-brand>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>

        <!-- Right aligned nav items -->
        <b-navbar-nav class="ml-auto">
          <b-nav-form>
            <li class="nav-item">
              <router-link class="nav-link nav-link-icon" to="/books">
                <span class="nav-link nav-link-inner--text">Books</span>
              </router-link>
            </li>
            <li class="nav-item">
              <router-link class="nav-link nav-link-icon" to="/users">
                <span class="nav-link nav-link-inner--text">Users</span>
              </router-link>
            </li>
            <li class="nav-item">
              <router-link class="nav-link nav-link-icon" to="/favori">
                <span class="nav-link nav-link-inner--text">Favories</span>
              </router-link>
            </li>
            <b-nav-item-dropdown right>
              <template v-slot:button-content>
                <span class="nav-link-inner--text">Other</span>
              </template>
              <b-dropdown-item href="#">bla bla</b-dropdown-item>
              <b-dropdown-item href="#">Another action</b-dropdown-item>
              <b-dropdown-divider></b-dropdown-divider>
              <b-dropdown-item variant="danger" @click="cikis">Log out</b-dropdown-item>
            </b-nav-item-dropdown>
          </b-nav-form>

        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
</template>

<script>
export default {
  name: "navbar",
  methods:{
    cikis(){
      localStorage.removeItem("name")
      localStorage.removeItem("id")
      localStorage.removeItem("surname")
      localStorage.removeItem("email")
      this.$router.push("/login")
    },
  },
}
</script>

<style scoped>
span{
  color: white !important;
}
</style>