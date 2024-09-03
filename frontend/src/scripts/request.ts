class ARequest<T, U> {
  method: string;
  contentType?: string;
  url: string;
  header?: HeadersInit;
  data?: FormData;
  constructor(
    method: string,
    url: string,
    header?: HeadersInit,
    data?: any,
  ) {
    this.method = method;
    this.url = url;
    this.header = header;
    this.data = data;
  }

  public async send(): Promise<T | U> {
    let body = null
    if (this.data){
      body = new URLSearchParams(this.data as any).toString()
    }
    try { 
      const response = await fetch(this.url, {
        method: this.method,
        headers: this.header,
        body: body,
        credentials: "include",
      });

      if (!response.ok) {
        throw new Error("Request failed");
      }

      // const responseData: T = await response.json();
      // return responseData;
      try {
        const responseData: T = await response.json();
        return responseData as unknown as U; // Пробуем вернуть как T
      } catch (jsonError) {
        // Если возникла ошибка преобразования, используем второй тип
        const fallbackData: U = await response.json();
        return fallbackData;
      }

    } catch (error) {
      throw new Error(String(error));
    }
  }
}

export default ARequest;
