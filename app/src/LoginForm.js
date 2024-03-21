import React, { useState } from "react";

function LoginForm() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [errors, setErrors] = useState({});

  const handleSubmit = (event) => {
    event.preventDefault();
    // Validate form data here
    // If there are errors, setErrors(errors)
    // If form data is valid, submit it
    fetch("/api/v1/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: new URLSearchParams({ Username: username, Password: password }),
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Failed to login");
        }
        // Handle successful login response
      })
      .catch((error) => {
        console.error("Login error:", error);
        // Handle login error
      });
  };

  return (
    <div className="container">
      <h1>Đăng Nhập</h1>
      <form onSubmit={handleSubmit} noValidate>
        <div>
          <p>
            <label>Tên Đăng nhập:</label>
          </p>
          <p>
            <input
              type="text"
              name="Username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
            />
          </p>
          {errors.UsernameLogin && (
            <p className="error">{errors.UsernameLogin}</p>
          )}
        </div>
        <div>
          <p>
            <label>Mật khẩu:</label>
          </p>
          <p>
            <input
              type="password"
              name="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </p>
          {errors.PasswordLogin && (
            <p className="error">{errors.PasswordLogin}</p>
          )}
        </div>
        <div>
          <input type="submit" value="Login" />
        </div>
      </form>
      <p>
        Bạn chưa có tài khoản? <a href="/register">Đăng ký ngay</a>
      </p>
    </div>
  );
}

export default LoginForm;
