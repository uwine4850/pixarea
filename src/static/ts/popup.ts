export class PopUp{
    private triggerButton: HTMLButtonElement;
    private popupWrapper: HTMLDivElement;
    constructor(triggerButtonId: string, popupWrapperId: string){
        this.triggerButton = document.getElementById(triggerButtonId) as HTMLButtonElement;
        this.popupWrapper = document.getElementById(popupWrapperId) as HTMLDivElement;
    }

    public init(){
        this.triggerButton.onclick = () => {
            this.show();
        }
        let popup_bg = this.popupWrapper.getElementsByClassName("popup-bg")[0] as HTMLElement;
        popup_bg.onclick = () => {
            this.hide();
        }
    }

    public show(){
        this.popupWrapper.classList.remove("popup-hide");
    }

    public hide(){
        this.popupWrapper.classList.add("popup-hide");
    }
}