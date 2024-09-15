import { z, ZodType } from 'zod';

class ARequest<T extends ZodType<any>, U extends ZodType<any>> {
  method: string;
  contentType?: string;
  url: string;
  header?: HeadersInit;
  data?: FormData;
  private successSchema: T;
  private errorSchema: U;

  constructor(
    method: string,
    url: string,
    successSchema: T,
    errorSchema: U,
    header?: HeadersInit,
    data?: any,
  ) {
    this.method = method;
    this.url = url;
    this.successSchema = successSchema;
    this.errorSchema = errorSchema;
    this.header = header;
    this.data = data;
  }

  public async send(): Promise<T | U> {
    let body = null;
    if (this.data) {
      body = new URLSearchParams(this.data as any).toString();
    }

    try {
      const response = await fetch(this.url, {
        method: this.method,
        headers: this.header,
        body: body,
        credentials: 'include',
      });

      if (!response.ok) {
        throw new Error('Request failed');
      }

      const responseData = await response.json();
      try {
        return this.successSchema.parse(responseData);
      } catch (successValidationError) {
        return this.errorSchema.parse(responseData);
      }
    } catch (error) {
      throw new Error(String(error));
    }
  }
}

export default ARequest;
