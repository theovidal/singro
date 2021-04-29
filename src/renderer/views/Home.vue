<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-img
          src="/img/logo.png"
          class="my-3"
          transition="scale-transition"
          height="200"
          contain/>
      </v-col>

      <v-col class="mb-4">
        <h1 class="display-2 font-weight-bold mb-3">
          Highest Template
        </h1>

        <p class="subheading font-weight-regular">
          This is a front-end template including : Vue.js, Vuetify, Router, PWA
        </p>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'Home',
  data () {
    return {

    }
  }
}
</script>
