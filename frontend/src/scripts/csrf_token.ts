import { useEffect, useState } from "react";
import ARequest from "./request";
import { CSRFTokenResponse } from "../messages/csrf";

export async function getCsrfToken(): Promise<CSRFTokenResponse> {
  let req = new ARequest<CSRFTokenResponse>(
    "GET",
    "http://localhost:8000/api/csrf"
  );
  const res = await req.send();
  return res;
}

export function useCsrfToken() {
  const [csrfToken, setCsrfToken] = useState<CSRFTokenResponse | null>(null);

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
