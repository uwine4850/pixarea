export function clickIfExistById(_id: string, fn: (el: HTMLElement) => void){
    let elementById = document.getElementById(_id) as HTMLElement;
    if(elementById){
        elementById.onclick = () => {
            fn(elementById);
        }
    }
}