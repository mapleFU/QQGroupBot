<template>
    <div class="servicers">
        <v-card>
            <v-card-title primary-title class="text-md-center">
                <!--<v-layout>-->
                <!--<v-flex xs5>-->
                <div class="headline">
                    Currently Available Services
                </div>
                <!--</v-flex>-->
                <!--</v-layout>-->

            </v-card-title>
            <v-divider></v-divider>

            <v-list>
                <v-list-tile
                        v-for="item in items"
                        :key="item.title"
                        avatar

                >
                    <v-list-tile-content class="text-md-center">
                        <v-list-tile-title class="text-md-center" v-text="item.title"></v-list-tile-title>
                    </v-list-tile-content>

                </v-list-tile>
            </v-list>
        </v-card>
    </div>
</template>

<script>
export default {
  name: "Servicers",
  data: function() {
    return {
      items: [
        {
          id: 0,
          title: "trace.moe image search service"
        },
        {
          id: 1,
          title: "hitokoto provider"
        },
        {
          id: 2,
          title: "RSS Searcher"
        }
      ]
    };
  }
};
</script>

<style scoped>
</style>
