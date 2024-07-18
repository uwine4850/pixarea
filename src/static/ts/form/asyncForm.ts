export class AsyncForm{
    private formData: FormData;
    private method: string;
    private url: string;
    private fnOnResponse: (reponse: Map<string, string>) => void;
    private fnOnError: (error: string) => void;
    constructor(formData: FormData, method: string, url: string){
        this.formData = formData;
        this.method = method;
        this.url = url;
    }
    
    public onResponse(fn: (reponse: Map<string, string>) => void){
        this.fnOnResponse = fn;
    }
    public onError(fn: (error: string) => void){
        this.fnOnError = fn;
    }

    public send(){
        fetch(this.url, {
            method: this.method,
            body: this.formData,
        }).then(response => response.json()).then(
            data => {
                if(this.fnOnResponse)
                    this.fnOnResponse(data);
            }
        )
        .catch(
            error => {
                if(this.fnOnError)
                    this.fnOnError(error);
            }
        );
    }
}
