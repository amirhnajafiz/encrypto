<template>
  <div style="background-color: black; margin-top:-10px">
      <div fluid class="background" style="display:table">
          <div style="display: table-cell; vertical-align:bottom; padding-bottom: 100px;"> 
              <v-row>
                  <v-col>
                      <h1 class="white--text font-weight-light text-right"> 
                          1020hp
                      </h1>
                      <p class="white--text text-right">
                          Peak Power
                      </p>
                  </v-col>
                  <v-col>
                      <h1 class="white--text font-weight-light text-center"> 
                        9.9s
                      </h1>
                      <p class="white--text text-center">
                          1/4 mile
                      </p>
                  </v-col>
                  <v-col>
                      <h1 class="white--text font-weight-light"> 
                          2.5s
                      </h1>
                      <p class="white--text">
                          0-60 mph+
                      </p>
                  </v-col>
              </v-row>
          </div>
      </div>
  </div>
</template>

<script>
export default {

}
</script>

<style>
    .background {
        background: url("https://tesla-cdn.thron.com/delivery/public/image/tesla/e13368c8-5657-468e-aa16-a28d9dc763c6/bvlatuR/std/2880x1800/MX-Performance-Hero-Desktop");
        background-size: cover;
        width: 100%;
        height: 100vh;
    }
</style>