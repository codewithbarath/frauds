<template>
  <div>
    <h2>Child Component</h2>

    <h3>Array received from Parent: {{ childName }}</h3>

    <button @click="sendName">Calculate Total</button>
    <div>
      <input type="text" v-model="firstNumber" />
      <input type="text" v-model="secondNumber" />
      <input type="text" v-model="sign" />
      <h1>total{{ total }}</h1>
      <button @click="calculate">calculate</button>
    </div>
    <div>
      <h1>toupper:{{ values }}</h1>
      <button @click="upper">change</button>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    childName: Array,
  },
  data() {
    return {
      firstNumber: "",
      secondNumber: "",
      sign: "",
      total: 0,
      values: ["barath", "kumar"],
    };
  },

  methods: {
    sendName() {
      let total = 0;

      for (let i = 0; i < this.childName.length; i++) {
        total = total + this.childName[i];
      }

      this.$emit("updateName", total);
    },
    calculate() {
      if (this.sign == "+") {
        this.total = this.firstNumber + this.secondNumber;
      } else if (this.sign == "-") {
        this.total = this.firstNumber - this.secondNumber;
      } else if (this.sign == "*") {
        this.total = this.firstNumber * this.secondNumber;
      }
    },
    upper() {
      this.values = this.values.map((vlaues) => {
        return vlaues.toUpperCase();
      });
    },
  },
};
</script>
