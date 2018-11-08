<template>
  <v-app>
    <v-toolbar app>
      <v-toolbar-title class="headline text-uppercase">
        <span>Mwish-QQBot</span>

      </v-toolbar-title>
      <v-spacer></v-spacer>

      <v-btn>
        <flat>
          <span class="font-weight-light">风空之枫的QQ群机器人管理器</span>
        </flat>
      </v-btn>
      <!--<v-btn-->
        <!--flat-->
        <!--href="https://github.com/vuetifyjs/vuetify/releases/latest"-->
        <!--target="_blank"-->
      <!--&gt;-->
        <!--<span class="mr-2">Latest Release</span>-->
      <!--</v-btn>-->
    </v-toolbar>

    <v-content>
      <ManagerPage/>
    </v-content>
  </v-app>
</template>

<script>
import ManagerPage from "./components/ManagerPage";

export default {
  name: "App",
  components: {
    ManagerPage
  },
  data() {
    return {
      //
    };
  }
};
</script>
