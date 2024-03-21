import "./Header.css";
function Header() {
  return (
    <header className="header">
      <a href="/" className="logo">
        <img src="" alt="Logo" />
      </a>
      <nav className="navigation">
        <ul className="navigation__items">
          <li className="navigation__item">
            <a href="/gioithieu">Giới Thiệu</a>
          </li>
          <li className="navigation__item">
            <a href="/dichvu">Dịch Vụ</a>
          </li>
          <li className="navigation__item">
            <a href="/r/community">Tin Tức</a>
          </li>
          <li className="navigation__item">
            <a href="/lienhe">Liên Hệ</a>
          </li>
          <li className="navigation__item">
            <a href="/login">Đăng Nhập</a>
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
