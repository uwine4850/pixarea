export class DragImage{
    private container_id: string;
    private drag_class_name: string;
    private customEvent: (draggable: HTMLElement) =>void;
    private customTouchEvent: (draggable: HTMLElement) =>void;
    private dragOver: (event)=>void;
    private touchMove: (event)=>void;
    public container: HTMLElement;
    public draggedElement: HTMLElement;
    constructor(container_id: string, drag_class_name: string){
        this.container_id = container_id;
        this.drag_class_name = drag_class_name;
        this.container = document.getElementById(this.container_id);
    }

    public addCustomEvent(fn: (draggable: HTMLElement) =>void){
        this.customEvent = fn;
    }

    public addCustomTouchEvent(fn: (draggable: HTMLElement) =>void){
        this.customTouchEvent = fn;
    }

    public setDragOver(fn: (event)=>void){
        this.dragOver = fn;
    }

    public setTouchMove(fn: (event)=>void){
        this.touchMove = fn;
    }

    private onDragStart(draggable: HTMLElement){
        draggable.addEventListener('dragstart', (event) => {
            let _image = event.target as HTMLElement;
            this.draggedElement = _image.parentNode as HTMLElement;       
        });
    }

    private onTouchStart(draggable: HTMLElement){
        draggable.addEventListener('touchstart', (event) => {
            let _image = event.target as HTMLElement;
            this.draggedElement = _image.parentNode as HTMLElement;
        });
    }

    private replaceElements(target: HTMLElement){
        if (this.draggedElement && target !== this.draggedElement && target.classList.contains(this.drag_class_name)) {
                
            const draggedIndex = Array.from(this.container.children).indexOf(this.draggedElement);
            const targetIndex = Array.from(this.container.children).indexOf(target);
            
            if (draggedIndex < targetIndex) {
                this.container.insertBefore(this.draggedElement, target.nextSibling);
            } else {
                this.container.insertBefore(this.draggedElement, target);
            }
        }
    }

    private onDrop(draggable: HTMLElement){        
        draggable.addEventListener('drop', (event) => {
            event.preventDefault();
            const targetImage = event.target as HTMLElement;
            const target = targetImage.parentNode as HTMLElement;
            this.replaceElements(target);
        });
    }

    private onTouchEnd(draggable: HTMLElement){
        draggable.addEventListener('touchend', (event) => {
            const touchEvent = event as TouchEvent;
            const touch = touchEvent.changedTouches[0];
            const targetImage = document.elementFromPoint(touch.clientX, touch.clientY) as HTMLElement;
            const target = targetImage.parentNode as HTMLElement;
            this.replaceElements(target);
        });
    }

    public run(){
        let draggables = document.getElementsByClassName(this.drag_class_name) as HTMLCollectionOf<HTMLAnchorElement>;
        Array.from(draggables).forEach(draggable => {            
            this.onDragStart(draggable);            
            draggable.addEventListener('dragover', (event) => {            
                event.preventDefault();
                if(this.dragOver)
                    this.dragOver(event);
            });
            if (this.customEvent)
                this.customEvent(draggable);            
            this.onDrop(draggable);

            // TOUCH SCREEN
            this.onTouchStart(draggable)
            draggable.addEventListener('touchmove', (event) => {
                event.preventDefault();
                if(this.touchMove)
                    this.touchMove(event);
            });
            if(this.customTouchEvent)
                this.customTouchEvent(draggable);
            this.onTouchEnd(draggable);
        });
    }
}