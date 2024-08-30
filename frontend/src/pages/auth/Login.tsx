import React, {useState} from "react";
import { LayoutProvider, Layout } from "../LayoutContext";

function getContent(onSubmit?: any) {
  return (
    <form
      className="auth-panel"
      onSubmit={onSubmit}
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
}

const Login: React.FC = () => {
  return (
    <LayoutProvider value={{ content: getContent() }}>
      <Layout />
    </LayoutProvider>
  );
};

export default Login;
