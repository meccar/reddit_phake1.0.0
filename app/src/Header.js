import React from "react";
import { Link } from "react-router-dom";

function Header() {
  return (
    <header className="header">
      <Link to="/" className="logo">
        <img src="" alt="Logo" />
      </Link>
      <nav className="navigation">
        <ul className="navigation__items">
          <li className="navigation__item">
            <Link to="/gioithieu">Giới Thiệu</Link>
          </li>
          <li className="navigation__item">
            <Link to="/dichvu">Dịch Vụ</Link>
          </li>
          <li className="navigation__item">
            <Link to="/r/community">Tin Tức</Link>
          </li>
          <li className="navigation__item">
            <Link to="/lienhe">Liên Hệ</Link>
          </li>
          <li className="navigation__item">
            <Link to="/login">Đăng Nhập</Link>
          </li>
        </ul>
      </nav>
      <div className="contact">
        <div className="phone"></div>
        <div className="form"></div>
      </div>
    </header>
  );
}

export default Header;
