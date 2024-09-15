import React, { useState } from "react";
import { LayoutProvider, Layout } from "../LayoutContext";
import { useNavigate } from "react-router-dom";
import { CSRFTokenInfer, parseCSRFResponce, SingleErrorInfer, useCsrfToken } from "../../scripts/csrf_token";
import ARequest from "../../scripts/request";
import { singleErrorResponseSchema } from "../../messages/schemas/error.schemas";

function getContent(onSubmit: any, error?: string) {
  return (
    <form
      className="auth-panel"
      method="post"
      action="register-post"
      encType="application/x-www-form-urlencoded"
      onSubmit={onSubmit}
    >
      <div className="error">{error}</div>
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
}

const Register: React.FC = () => {
  const navigate = useNavigate();
  const [error, setError] = useState<string | undefined>(undefined);
  const csrfToken = useCsrfToken();
  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    await parseCSRFResponce(csrfToken, async function (csrfTokenResponce: CSRFTokenInfer) {
      const formData = new FormData(event.currentTarget);
      formData.append("CSRF_TOKEN", csrfTokenResponce.Token);
      const req = new ARequest("POST", "http://localhost:8000/api/register", singleErrorResponseSchema, singleErrorResponseSchema, {
        "Content-Type": "application/x-www-form-urlencoded"
      }, formData);
      const res = await req.send();
      const response = singleErrorResponseSchema.parse(res);
      if (response?.Error == ""){
        navigate("/login");
      } else {
        setError(response?.Error);
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

export default Register;
