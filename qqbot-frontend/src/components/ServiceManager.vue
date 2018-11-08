<template>
    <v-card>
        <v-card-title class="text-md-center">
            <div class="headline">
                Servicer Manager
            </div>
        </v-card-title>
        <v-divider></v-divider>

        <v-list>
            <v-list-tile
                    v-for="servicer in servicers"
                    :key="servicer.addName"
                    avatar
            >
                <v-list-tile-content>
                    <v-list-tile-title v-html="servicer.addName"></v-list-tile-title>
                    <v-list-tile-sub-title v-html="servicer.serviceType"></v-list-tile-sub-title>
                </v-list-tile-content>

                <v-btn color="error" @click="delete_service(servicer.addName)">
                    DELETE
                </v-btn>
            </v-list-tile>
        </v-list>

        <!--service adder-->
        <v-container fluid>
            <v-layout row wrap align-center>
                <v-flex xs12 sm5>
                    <v-text-field
                            v-model="cur_add_name"
                            label="Service Name"
                    ></v-text-field>
                </v-flex>
                <v-flex xs12 sm5>
                    <v-select
                            :items="all_servicers"
                            v-model="cur_service"
                            :menu-props="{ maxHeight: '400' }"
                            label="Select"
                            item-value="text"
                            hint="Pick your New Service"
                            persistent-hint
                            id="selectService"
                    ></v-select>
                </v-flex>
                <v-flex xs12 sm2>
                    <v-btn color="info" @click="add_service">
                        ADD
                    </v-btn>
                </v-flex>
            </v-layout>
        </v-container>
    </v-card>
</template>

<script>
const axios = require("axios");
import constexpr from "../constexpr";

export default {
  name: "ServiceManager",
  data: function() {
    return {
      servicers: [],
      servicerID: 0,
      cur_add_name: "",
      all_servicers: [
        {
          id: 0,
          text: "trace.moe image search service"
        },
        {
          id: 1,
          text: "hitokoto provider"
        },
        {
          id: 2,
          text: "RSS Searcher"
        }
      ],
      cur_service: null,
      cur_name: ""
    };
  },
  created() {
    setTimeout(() => {
      this.cur_name = "trace.moe image search service";
      this.cur_service = this.cur_name;
    }, 1000);

    this.flush_service();
  },
  methods: {
    add_service: function() {
      if (this.cur_add_name === "" || this.cur_service == null) {
        return;
      }
      console.log(this.cur_service);
      console.log(this.cur_name);
      axios
        .post(
          "http://" +
            constexpr.HttpServer +
            "/manager/service/" +
            this.cur_service,
          {
            addName: this.cur_add_name
          }
        )
        .then(resp => {
          this.flush_service();
        });
      this.cur_service = null;
      this.cur_add_name = "";
    },
    delete_service: function(servicer) {
      axios
        .delete(
          "http://" + constexpr.HttpServer + "/manager/service/" + servicer
        )
        .then(() => {
          this.flush_service();
        });
    },
    flush_service: function() {
      this.servicers = [];
      axios
        .get("http://" + constexpr.HttpServer + "/manager/service")
        .then(resp => {
          console.log(resp);
          let cur_map = resp.data["groups"];
          for (let key in cur_map) {
            console.log(key);
            this.servicers.push({
              addName: key,
              serviceType: cur_map[key]
            });
          }
        })
        .catch(() => {
          console.log("http error");
        });
    }
  }
};
</script>

<style scoped>
</style>
