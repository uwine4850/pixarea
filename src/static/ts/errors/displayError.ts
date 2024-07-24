export function showErrorOnPage(error: string){
    let err = document.getElementById("error");
    if (err){
        err.innerText = error;
    }
}