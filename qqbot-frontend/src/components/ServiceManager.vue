<template>
    <v-card>
        <v-card-title class="text-md-center">
            <div class="headline">
                Servicer Manager
            </div>
        </v-card-title>

        <v-list>
            <v-list-tile
                    v-for="servicer in servicers"
                    :key="servicer"
                    avatar
            >
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
  },
  methods: {
    add_service: function() {
      if (this.cur_add_name === "" || this.cur_service == null) {
        return;
      }
      this.cur_service = null;
      this.cur_add_name = "";
    }
  }
};
</script>

<style scoped>
</style>
