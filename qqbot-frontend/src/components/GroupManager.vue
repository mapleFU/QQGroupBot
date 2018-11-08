<template>
    <v-card>
        <v-card-title class="text-md-center">
            <div class="headline">
                Group Manager
            </div>
        </v-card-title>
        <v-divider></v-divider>

        <v-list>
            <v-list-tile
                    v-for="current_group in groups"
                    :key="current_group"
                    avatar
            >
                <v-list-tile-content>
                    <v-list-tile-title v-html="current_group"></v-list-tile-title>
                    <!--<v-list-tile-sub-title v-html="servicers.serviceType"></v-list-tile-sub-title>-->
                </v-list-tile-content>

                <v-btn color="error" @click="delete_group(current_group)">
                    DELETE
                </v-btn>
            </v-list-tile>
        </v-list>

        <!--service adder-->
        <v-container fluid>
            <v-layout row wrap align-center>
                <v-flex xs12 sm8>
                    <v-text-field
                            v-model="cur_add_group"
                            label="Group ID"
                    ></v-text-field>
                </v-flex>
                <v-flex xs12 sm2>
                    <v-btn color="info" @click="add_group">
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
  name: "GroupManager",
  data: function() {
    return {
      groups: [],
      cur_add_group: "000000"
    };
  },
  created: function() {
    this.flush_group();
  },
  methods: {
    add_group: function() {
      if (this.cur_add_group === null || this.cur_add_group === "") {
        return;
      }
      axios
        .post(
          "http://" +
            constexpr.HttpServer +
            "/manager/group/" +
            this.cur_add_group
        )
        .then(() => {
          this.flush_group();
        })
        .catch(() => {
          // console.log("add group error");
        });

      this.cur_add_group = "123456";
    },
    delete_group: function(group_name) {
      axios
        .delete(
          "http://" + constexpr.HttpServer + "/manager/group/" + group_name
        )
        .then(resp => {
          this.flush_group();
        })
        .catch(() => {
          // console.log("error");
        });
    },
    // 刷新小组
    flush_group: function() {
      this.groups = [];
      axios
        .get("http://" + constexpr.HttpServer + "/manager/group")
        .then(resp => {
          // console.log(resp);
          for (let key of resp.data["groups"]) {
            this.groups.push(key);
          }
        })
        .catch(() => {
          // console.log("http error in group manager");
        });
    }
  }
};
</script>

<style scoped>
</style>
