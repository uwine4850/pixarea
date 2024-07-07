import { DragImage } from "./image/dragImages";
import { Dropdown} from "./dropdown";
import { ShowInputImages } from "./image/showInputImage";

export function publicationRunDropdown(): Dropdown {
  let drpd = new Dropdown("comment-menu-dropdown");
  drpd.run();
  return drpd;
}

export function newPublicationRunDropdown(): Dropdown {
  let drpd = new Dropdown("new-publication-categories-btn");
  drpd.run();
  return drpd;
}

let inputImages = new ShowInputImages("new-pub-images", "load-images", ["load-image", "draggable"]);
inputImages.onInputChange(function(){
  let dragImage = new DragImage("load-images", "draggable");
  dragImage.run();
});
inputImages.run();

// HIDE COMMENT
let commentHideBtns = document.getElementsByClassName("comment-hide-btn") as HTMLCollectionOf<HTMLButtonElement>;
for (let i = 0; i < commentHideBtns.length; i++) {
  const btn = commentHideBtns[i];
  btn.onclick = function() {
    let comment = btn.closest(".publication-comment-item");
    let PCH = null;
    if(comment){
      PCH = comment.getElementsByClassName("publication-comment-item-hidden-panel")[0] as HTMLButtonElement;
    }
    if(comment && PCH){
      comment.classList.add("publication-comment-item-hidden");
      PCH.classList.add("PCH-enable");
    }
  }
}

// SHOW HIDDEN COMMENT
let PCH = document.getElementsByClassName("publication-comment-item-hidden-panel") as HTMLCollectionOf<HTMLButtonElement>;
for (let i = 0; i < PCH.length; i++) {
  PCH[i].onclick = function() {
    let comment = PCH[i].parentElement;
    comment.classList.remove("publication-comment-item-hidden");
    PCH[i].classList.remove("PCH-enable");
  }
}