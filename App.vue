```vue
<template>
  <div>
    <!-- Sign Up Form -->
    <h2>Sign Up</h2>
    <!-- <h1>csvdatas{{ csvdata }}</h1> -->
    <h1>CSV Data</h1>

<pre>{{ csvdata }}</pre>

<div v-for="(record, index) in csvdata" :key="index">
  <p>{{ record.join(" | ") }}</p>
</div>


    <form @submit.prevent="signUp">
      <label>Username:</label>
      <input
        type="text"
        v-model="username"
        placeholder="Enter username"
        required
      />

      <br /><br />
      <!-- <button @click="fetchCSV">Fetch CSV Data</button> -->
      <!-- <button type="button" @click="fetchCSV">Fetch CSV Data</button> -->
      <button type="button" @click="fetchCSV">fetchdetail</button>

      <label>Phone Number:</label>
      <input
        type="tel"
        v-model="phone"
        placeholder="Enter phone number"
        required
      />

      <br /><br />

      <label>Email:</label>
      <input
        type="email"
        v-model="email"
        placeholder="Enter email"
        required
      />

      <br /><br />

      <label>Password:</label>
      <input
        type="password"
        v-model="password"
        placeholder="Enter password"
        required
      />

      <br /><br />

      <button type="submit">Sign Up</button>
    </form>

    <hr />

    <!-- Bus Search -->
    <h2>Bus Search</h2>

    <label>From:</label>
    <input
      type="text"
      v-model="from"
      placeholder="Enter starting place"
    />

    <br /><br />

    <label>To:</label>
    <input
      type="text"
      v-model="to"
      placeholder="Enter destination"
    />

    <br /><br />

    <button @click="searchBuses">Send Request</button>

    <h3>Available Buses</h3>

    <div v-for="bus in returnarr" :key="bus.busId">
      <p>Bus Name: {{ bus.busName }}</p>
      <p>From: {{ bus.from }}</p>
      <p>To: {{ bus.to }}</p>
      <p>Time: {{ bus.time }}</p>
      <p>Price: {{ bus.rupees }}</p>
      <hr />
    </div>
    <button @click="fetch">fetchdetail</button>
  </div>
</template>

<script>
import eventservices from "./assets/Services/eventservices";

export default {
  name: "BusSearch",

  data() {
    return {
      // Signup form values
      username: "",
      phone: "",
      email: "",
      password: "",

      // Bus search values
      from: "",
      to: "",
      returnarr: [],
      csvdata:[],
    };
  },

  methods: {
    // Signup API method
    signUp() {
      const userData = {
        username: this.username,
        phone: this.phone,
        email: this.email,
        password: this.password
      };

      eventservices
        .signUp(userData)
        .then((response) => {
          console.log("Signup API Response:", response.data);

          if (response.data.status === "S") {
            alert("Signup successful!");

            // Clear signup form
            this.username = "";
            this.phone = "";
            this.email = "";
            this.password = "";
          } else {
            alert(
              "Signup failed: " +
              (response.data.errorMsg || "Unable to create user")
            );
          }
        })
        .catch((error) => {
          console.error("Signup API Error:", error);
          alert("Signup failed. Please check the backend.");
        });
    },
    // fetch(){
    //   eventservices.fetchs().then((response)=>{
    //     // console.log(response.data.data)
    //     if(response.data.status === "s"){
    //       console.log(response.data.data)
    //     } else {
    //       alert("not fetched")
    //     }
    //   })

    // },

    // Bus search API method
    searchBuses() {
      eventservices
        .searchBuses(this.from, this.to)
        .then((response) => {
          console.log("Bus Search API Response:", response.data);

          if (response.data.status === "S") {
            this.returnarr = response.data.data || [];
          } else {
            this.returnarr = [];
            alert(
              "Bus search failed: " +
              (response.data.errorMsg || "Unable to search buses")
            );
          }
        })
        .catch((error) => {
          console.error("Bus Search API Error:", error);
          alert("Bus search failed. Please check the backend.");
        });
    },
    fetchCSV() {
  eventservices
    .fetchCSV()
    .then((response) => {
      console.log("CSV Data:", response.data);

      this.csvdata = response.data;

      console.log("Vue csvdata:", this.csvdata);
    })
    .catch((error) => {
      console.error("CSV Fetch Error:", error);
    });
}

  }
};
</script>
```
