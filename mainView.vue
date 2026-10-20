<template>
  <div style="height: 300px; width: 400px; background-color: lightgreen">
    <h2>Parent Component</h2>

    <h3>Array in Parent: {{ name }}</h3>

    <h3>Total received from Child: {{ total }}</h3>

    <MainView :childName="name" @updateName="changeName" />
  </div>
</template>

<script>
import MainView from "@/components/home/component/childView.vue";

export default {
  components: {
    MainView,
  },

  data() {
    return {
      name: [10, 20, 30],
      total: 0,
    };
  },

  methods: {
    changeName(newTotal) {
      this.total = newTotal;
    },
  },
};
</script>
