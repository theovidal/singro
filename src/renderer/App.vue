<template>
  <v-app>
    <v-app-bar app>
      <div class="d-flex align-center">
        <v-img
          alt="Highest Logo"
          class="shrink mr-2"
          contain
          src="/img/logo-large.png"
          width="150"/>
      </div>

      <v-spacer/>

      <v-btn
        href="https://github.com/highest-app/highest/releases/latest"
        target="_blank"
        text>
        <span class="mr-2">Latest Release</span>
        <v-icon>mdi-open-in-new</v-icon>
      </v-btn>
    </v-app-bar>

    <v-content>
      <router-view/>
    </v-content>
  </v-app>
</template>

<script>

export default {
  name: 'App',
  data () {
    return {

    }
  }
};
</script>
