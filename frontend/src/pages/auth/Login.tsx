import React, {useState, useEffect} from "react";
import { useNavigate } from 'react-router-dom';
import { LayoutProvider, Layout } from "../LayoutContext";
import ARequest from "../../scripts/request";
import { useCsrfToken, parseCSRFResponce, CSRFTokenInfer, SingleErrorInfer } from "../../scripts/csrf_token";
import { cSRFTokenResponseSchema } from "../../messages/schemas/csrf.schemas";
import { singleErrorResponseSchema } from "../../messages/schemas/error.schemas";

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
    await parseCSRFResponce(csrfTokenResult, async function (csrfTokenResponce: CSRFTokenInfer) {
      const parseResult = cSRFTokenResponseSchema.safeParse(csrfTokenResponce);
      if(parseResult.success){
        const formData = new FormData(event.currentTarget);
        const data = {
          username: formData.get('username'),
          password: formData.get('password'),
          CSRF_TOKEN: csrfTokenResponce?.Token,
        };
        const req = new ARequest("POST", "http://localhost:8000/api/login", singleErrorResponseSchema, singleErrorResponseSchema, {
          "Content-Type": "application/x-www-form-urlencoded"
        }, data);
        const res = await req.send();
        const response = singleErrorResponseSchema.parse(res);
        if (response?.Error == ""){
          navigate("/");
        } else {
          setError(response?.Error);
        }
      }
    },
      async function (singleErrorResponse: SingleErrorInfer) {
        navigate(singleErrorResponse.Redirect);
      }
    )
  }
  return (
    <LayoutProvider value={{ content: getContent(handleSubmit, error) }}>
      <Layout />
    </LayoutProvider>
  );
};

export default Login;
