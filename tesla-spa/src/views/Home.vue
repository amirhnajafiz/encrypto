<template>
  <div>
    <promo />
    <specifications></specifications>
    <interior />
  </div>
</template>

<script>
import Interior from '../components/Interior.vue';
import Promo from "../components/Promo.vue";
import Specifications from "../components/Specifications.vue";

export default {
  components: { Promo, Specifications, Interior },
  name: "Home",
};
</script>
