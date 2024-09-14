import { useEffect, useState } from "react";
import ARequest from "./request";
import { singleErrorResponseSchema } from "../messages/schemas/error.schemas";
import { cSRFTokenResponseSchema } from "../messages/schemas/csrf.schemas";
import { z } from 'zod';

export async function getCsrfToken(): Promise<typeof cSRFTokenResponseSchema | typeof singleErrorResponseSchema> {
  let req = new ARequest<typeof cSRFTokenResponseSchema, typeof singleErrorResponseSchema>(
    "GET",
    "http://localhost:8000/api/csrf",
    cSRFTokenResponseSchema,
    singleErrorResponseSchema
  );
  const res = await req.send();
  return res;
}

export function useCsrfToken(): typeof cSRFTokenResponseSchema | typeof singleErrorResponseSchema | null{
  const [csrfToken, setCsrfToken] = useState<typeof cSRFTokenResponseSchema | typeof singleErrorResponseSchema | null >(null);

  useEffect(() => {
    const fetchCsrfToken = async () => {
      try {
        const token = await getCsrfToken();
        setCsrfToken(token);
      } catch (error) {
        console.error("Failed to fetch CSRF token:", error);
      }
    };

    fetchCsrfToken();
  }, []);

  return csrfToken;
}

export type CSRFTokenInfer = z.infer<typeof cSRFTokenResponseSchema>;
export type SingleErrorInfer = z.infer<typeof singleErrorResponseSchema>;
type CSRFResponse = typeof cSRFTokenResponseSchema | typeof singleErrorResponseSchema | null;
type okFunc = (csrfTokenResponse: CSRFTokenInfer) => Promise<void>;
type errorFunc = (singleErrorResponse: SingleErrorInfer) => Promise<void>;

export async function parseCSRFResponce(csrfToken: CSRFResponse, okFunc: okFunc, errorFunc: errorFunc){
  if (csrfToken == null){
    console.error("CSRF token is null.");
    return;
  }
  try {
    const validCSRFToken = cSRFTokenResponseSchema.parse(csrfToken);
    await okFunc(validCSRFToken);
  } catch {
    try {
      const validErrorResponse = singleErrorResponseSchema.parse(csrfToken);
      await errorFunc(validErrorResponse);
    } catch {
      console.error("Response does not match any expected type.");
    }
  }
}
