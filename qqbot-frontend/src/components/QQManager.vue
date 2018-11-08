<template>
    <v-card>
        <v-card-title primary-title class="text-md-center">
            <!--<v-layout>-->
            <!--<v-flex xs5>-->
            <div class="headline">
                Service QQ
            </div>
            <!--</v-flex>-->
            <!--</v-layout>-->

        </v-card-title>
        <v-divider></v-divider>
        <v-list>
            <v-list-tile
                    v-for="qq in qq_numbers"
                    :key="qq.number"
                    avatar
            >

                <v-list-tile-content class="text-md-center">
                    <v-list-tile-title  class="text-md-center">
                        机器人账号 {{ qq.number }}
                    </v-list-tile-title>
                </v-list-tile-content>

                <template v-if="qq.status">
                    <v-chip color="primary" text-color="white">
                        Running
                    </v-chip>
                </template>
                <template v-else>
                    <v-chip color="red" text-color="white">
                        Stopped
                    </v-chip>
                </template>

                <v-list-tile-avatar>
                    <v-icon>
                        fab fa-qq
                    </v-icon>
                </v-list-tile-avatar>
                <!--<i class="fab fa-qq"></i>-->

            </v-list-tile>
        </v-list>
    </v-card>
</template>

<script>
export default {
  name: "QQManager",
  data: function() {
    return {
      qq_numbers: [
        {
          number: 312392187,
          status: 0
        }
      ]
    };
  }
};
</script>

<style scoped>
</style>
