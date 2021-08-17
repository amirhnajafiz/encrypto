<template>
  <span>
    <v-toolbar app color="black" dark> <!-- Pre defined in vuetify -->
      <v-toolbar-title>
        <v-img src="../../public/tesla.png" width="200">
        </v-img>
      </v-toolbar-title>
      <v-spacer></v-spacer> <!-- Spacer will seperate the items in navigation bar -->
      <v-btn flat class="hidden-sm-and-down white font-weight-bold black--text">
        Specifications
      </v-btn>
      <v-btn
        flat
        class="hidden-sm-and-down white font-weight-bold black--text ml-3"
      >
        Order now
      </v-btn>
    </v-toolbar>
  </span>
</template>

<script>
export default {};
</script>

<style>
</style>