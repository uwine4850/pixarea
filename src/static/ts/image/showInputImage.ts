import { CompressImage } from './compress';

export class ShowInputImages{
    private imageWrapperClasses: string[];
    private inputChange: () => void;
    public imageContainer: HTMLDivElement;
    public input: HTMLInputElement;
    constructor(inputId: string, imageContainerId: string, imageWrapperClasses: string[]){
        this.imageWrapperClasses = imageWrapperClasses;
        
        this.imageContainer = document.getElementById(imageContainerId) as HTMLDivElement;
        this.input = document.getElementById(inputId) as HTMLInputElement;
    }

    public onInputChange(fn: ()=>void){
        this.inputChange = fn;
    }

    private addImageWrapperClasses(wrapper: HTMLDivElement){
        for (let i = 0; i < this.imageWrapperClasses.length; i++) {
            const className = this.imageWrapperClasses[i];
            wrapper.classList.add(className);
        }
    }

    private displayImages(files: File[]){
        this.imageContainer.innerHTML = "";
        
        for (const file of files) {
            if (!file.type.startsWith('image/'))
                continue;
      
            const img = document.createElement('img');
            
            const wrapper = document.createElement("div");
            wrapper.setAttribute("data-name", file.name);            

            this.addImageWrapperClasses(wrapper);
            wrapper.appendChild(img);

            const reader = new FileReader();
            reader.onload = (e) => {
                img.src = e.target?.result as string;
            };
            reader.readAsDataURL(file);

            this.imageContainer.appendChild(wrapper);
        }
    }

    public run(){
        if (!this.input) {
            return;
        }
        this.input.onchange = async (event) => {
            if (this.input.files) {                
                let compresedFiles = await CompressImage.compressFromInput(this.input, 0.2);
                this.displayImages(Array.from(compresedFiles));
                if (this.inputChange)
                    this.inputChange();
            }
        }
    }
}

export function getSortedImageNames(containerId: string, imageClass: string): string[]{
    let names = []
    let images = document.getElementById(containerId)?.getElementsByClassName(imageClass) as HTMLCollectionOf<HTMLElement>
    if (images){
        for (let i = 0; i < images.length; i++) {
            const image = images[i];
            names.push(image.getAttribute("data-name"));
        }
    }
    return names
}
