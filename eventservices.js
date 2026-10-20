import axios from "axios";

const baseclientupi = axios.create({
  withCredentials: false,

  headers: {
    accept: "application/json",
    "Content-Type": "application/json"
  },

  baseURL: "http://localhost:8080"
});

export default {

  // Signup API
  signUp(userData) {
    return baseclientupi.post("/signup", userData);
  },

  // Fetch CSV data
  fetchCSV() {
    return baseclientupi.get("/fetch-csv");
  }

};