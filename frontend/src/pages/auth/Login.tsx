import React, {useState, useEffect} from "react";
import { useNavigate } from 'react-router-dom';
import { LayoutProvider, Layout } from "../LayoutContext";
import ARequest from "../../scripts/request";
import { useCsrfToken } from "../../scripts/csrf_token";
import { CSRFTokenResponse } from "../../messages/csrf";
import { SingleErrorResponse } from "../../messages/messages";
import { checkType } from "../../scripts/typecheck";

function getContent(onSubmit?: any, error?: string) {
  return (
    <form
      className="auth-panel"
      onSubmit={onSubmit}
    >
      <div className="error">{error}</div>
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
  const navigate = useNavigate();
  const [error, setError] = useState<string | undefined>(undefined);

  const csrfTokenResult = useCsrfToken();
  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (checkType<CSRFTokenResponse>(csrfTokenResult as CSRFTokenResponse)) {
      const csrf = csrfTokenResult as CSRFTokenResponse;
      const formData = new FormData(event.currentTarget);
      const data = {
        username: formData.get('username'),
        password: formData.get('password'),
        CSRF_TOKEN: csrf?.Token,
      };
      const req = new ARequest<SingleErrorResponse, undefined>("POST", "http://localhost:8000/api/login", {
        "Content-Type": "application/x-www-form-urlencoded"
      }, data);
      const res = await req.send();
      if (res?.Error == ""){
        navigate("/");
      } else {
        setError(res?.Error);
      }
    } else {
      const error = csrfTokenResult as SingleErrorResponse;
      navigate(error.Redirect);
    }
  }
  return (
    <LayoutProvider value={{ content: getContent(handleSubmit, error) }}>
      <Layout />
    </LayoutProvider>
  );
};

export default Login;
