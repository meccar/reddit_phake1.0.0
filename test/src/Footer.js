import React from "react";

function Footer() {
  return (
    <footer className="footer">
      <div className="footer__content">
        <p>&copy; 2024 Your Company. All rights reserved.</p>
        <ul className="footer__links">
          <li>
            <a href="/about">About Us</a>
          </li>
          <li>
            <a href="/contact">Contact Us</a>
          </li>
          <li>
            <a href="/privacy">Privacy Policy</a>
          </li>
        </ul>
      </div>
    </footer>
  );
}

export default Footer;
