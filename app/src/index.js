import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter as Router } from "react-router-dom";
import "./index.css";
import Header from "./Header";
import Footer from "./Footer";
import reportWebVitals from "./reportWebVitals";

ReactDOM.render(
  <Router>
    <Header />
    <Footer />
  </Router>,
  document.getElementById("root"),
);

reportWebVitals();
