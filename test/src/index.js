import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import Header from "./Header";
import Footer from "./Footer";
import reportWebVitals from "./reportWebVitals";

ReactDOM.render(
  <React.StrictMode>
    <Header />
    <Footer />
  </React.StrictMode>,
  document.getElementById("index"),
);

reportWebVitals();
