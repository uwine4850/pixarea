class ARequest<T> {
  method: string;
  contentType?: string;
  url: string;
  data?: FormData;
  constructor(
    method: string,
    url: string,
    contentType?: string,
    data?: any,
  ) {
    this.method = method;
    this.contentType = contentType;
    this.url = url;
    this.data = data;
  }

  public async send(): Promise<T> {
    let body = null
    let headers: HeadersInit = {}
    if (this.data){
      body = new URLSearchParams(this.data as any).toString()
    }
    if (this.contentType){
      headers["Content-Type"] = this.contentType
    }
    try { 
      const response = await fetch(this.url, {
        method: this.method,
        headers: headers,
        body: body,
        credentials: "include",
      });

      if (!response.ok) {
        throw new Error("Request failed");
      }

      const responseData: T = await response.json();
      return responseData;

    } catch (error) {
      throw new Error(String(error));
    }
  }
}

export default ARequest;
