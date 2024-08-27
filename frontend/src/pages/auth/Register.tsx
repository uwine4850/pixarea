import React from "react";
import { LayoutProvider, Layout } from "../LayoutContext";

const content = (
  <form
    className="auth-panel"
    method="post"
    action="register-post"
    encType="application/x-www-form-urlencoded"
  >
    <div className="error"></div>
    <h2 className="auth-panel-title">Register</h2>
    <hr />
    <div className="form-block form-block-auth">
      <div className="form-block-item">
        <label htmlFor="name">Name</label>
        <input id="name" name="name" type="text" />
      </div>
      <div className="form-block-item">
        <label htmlFor="username">Username</label>
        <input id="username" name="username" type="text" />
      </div>
      <div className="form-block-item">
        <label htmlFor="password">Password</label>
        <input id="password" name="password" type="text" />
      </div>
      <div className="form-block-item">
        <label htmlFor="repeat_password">Repeat password</label>
        <input id="repeat_password" name="repeat_password" type="text" />
      </div>
    </div>
    <a href="/login" className="auth-support-link">
      Login
    </a>
    <button type="submit" className="auth-button">
      <a>Register</a>
    </button>
  </form>
);

const Register: React.FC = () => {
  return (
    <LayoutProvider value={{ content: content }}>
      <Layout />
    </LayoutProvider>
  );
};

export default Register;
