import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter as Router } from "react-router-dom";
import "./login.css";
import Header from "./Header";
import Footer from "./Footer";
import LoginForm from "./LoginForm";
import reportWebVitals from "./reportWebVitals";

ReactDOM.render(
  <Router>
    <Header />
    <LoginForm />
    <Footer />
  </Router>,
  document.getElementById("login"),
);

reportWebVitals();
