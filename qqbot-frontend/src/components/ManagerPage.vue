<template>
    <div class="container">
        <v-layout>
            <v-flex xs12 sm10 offset-sm1>
                <QQManager></QQManager>
            </v-flex>
        </v-layout>

        <v-layout>
            <v-flex xs12 sm10 offset-sm1>
                <Servicers/>
            </v-flex>
        </v-layout>

        <v-layout>
            <v-flex xs12 sm10 offset-sm1>
                <ServiceManager></ServiceManager>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
import QQManager from "./QQManager";
import ServiceManager from "./ServiceManager";
import Servicers from "./Servicers";
export default {
  name: "ManagerPage",
  components: { Servicers, ServiceManager, QQManager },
  data: function() {
    return {
      qq_numbers: [312392187]
    };
  }
};
</script>

<style scoped>
</style>
