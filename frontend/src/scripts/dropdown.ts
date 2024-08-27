interface DropdownItem{
    targetButton: HTMLElement;
    dropdownWrapper: HTMLElement;
    dropdown: HTMLElement;
}

export class Dropdown{
    private targetButtonClassName: string;
    private dropdownWrapperClassName: string = "dropdown-wrapper";
    private dropdownForBtnClassName: string = "dropdown-for-btn";
    private dropdownForBtnHideClassName: string = "dropdown-for-btn-hide";

    public dropdownItems: DropdownItem[] = [];

    constructor(targetButtonClassName: string){
        this.targetButtonClassName = targetButtonClassName;
        this.parse();
    }

    private parse(){
        let targetButtonsItems = Array.from(document.getElementsByClassName(this.targetButtonClassName) as HTMLCollectionOf<HTMLElement>);
        for (let i = 0; i < targetButtonsItems.length; i++) {
            const targetButton: HTMLElement = targetButtonsItems[i];

            let dropdownWrapper = targetButton.parentElement;
            if(!dropdownWrapper!.classList.contains(this.dropdownWrapperClassName)){
                throw new ExpectedClassNotFoundError("Dropdown.parse", this.dropdownWrapperClassName);
            }
            
            let dropdownForBtn = dropdownWrapper!.querySelector("." + this.dropdownForBtnClassName) as HTMLElement;
            if(!dropdownForBtn){
                throw new ExpectedClassNotFoundError("Dropdown.parse", this.dropdownForBtnClassName);
            }
            
            this.dropdownItems.push({
                targetButton: targetButton, 
                dropdownWrapper: dropdownWrapper!, 
                dropdown: dropdownForBtn}
            )
        }
    }

    public run(){
        for (let i = 0; i < this.dropdownItems.length; i++) {
            const item = this.dropdownItems[i];
            item.targetButton.onclick = () => {
                item.dropdown.classList.toggle(this.dropdownForBtnHideClassName);
            }
        }
    }

    public hide(index: number){
        this.dropdownItems[index].dropdown.classList.add(this.dropdownForBtnHideClassName);
    }
}

class ExpectedClassNotFoundError extends Error{
    constructor(caller: string, className: string) {
        super(`${caller}: Class ${className} not found.`);
    }
}

export function closeDropdown(drpd: Dropdown, clickTarget: Node){
    for (let i = 0; i < drpd.dropdownItems.length; i++) {
        const drp = drpd.dropdownItems[i].dropdownWrapper;
        if (!drp.contains(clickTarget)){
            drpd.hide(i);
        }
    }
}
