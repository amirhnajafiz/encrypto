<template>
  <div>
    <promo />
    <interior />
    <specifications></specifications>
    <core-features />
    <order-now />
  </div>
</template>

<script>
import CoreFeatures from '../components/CoreFeatures.vue';
import Interior from '../components/Interior.vue';
import OrderNow from '../components/OrderNow.vue';
import Promo from "../components/Promo.vue";
import Specifications from "../components/Specifications.vue";

export default {
  components: { Promo, Specifications, Interior, CoreFeatures, OrderNow },
  name: "Home",
};
</script>
