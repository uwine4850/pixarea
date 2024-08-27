import React from "react";
import { LayoutProvider, Layout } from "../LayoutContext";

const content = (
  <form
    className="auth-panel"
    method="post"
    action="/login-post"
    encType="application/x-www-form-urlencoded"
  >
    <div className="error"></div>
    <h2 className="auth-panel-title">Login</h2>
    <hr />
    <div className="form-block form-block-auth">
      <div className="form-block-item">
        <label htmlFor="username">Username</label>
        <input id="username" name="username" type="text" />
      </div>
      <div className="form-block-item">
        <label htmlFor="password">Password</label>
        <input id="password" name="password" type="text" />
      </div>
    </div>
    <a href="/register" className="auth-support-link">
      Register
    </a>
    <button type="submit" className="auth-button">
      <a>Login</a>
    </button>
  </form>
);

const Login: React.FC = () => {
  return (
    <LayoutProvider value={{ content: content }}>
      <Layout />
    </LayoutProvider>
  );
};

export default Login;
