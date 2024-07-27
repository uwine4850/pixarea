import { DragImage } from "./image/dragImages";
import { Dropdown, closeDropdown} from "./dropdown";
import { ShowInputImages } from "./image/showInputImage";
import { showErrorOnPage } from "./errors/displayError";
import { getSortedImageNames } from "./image/showInputImage";
import { redirect } from "./routing/urlUtils";
import { AsyncForm } from "./form/asyncForm";
import { sendLikeForm } from "./publication_like";
import { sendCommentForm } from "./publication_comment";


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

export function hideShowComments(){
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
}

hideShowComments();

function sendForm(formId: string){
  let form = document.getElementById(formId) as HTMLFormElement;
  if (form){
    form.addEventListener('submit', (event: Event) => {
      event.preventDefault();

      const formData = new FormData(form);
      const fileInput = document.getElementById('new-pub-images') as HTMLInputElement;
      const images = fileInput.files;
        
      if (fileInput.files.length > 0){
        formData.delete("images");
      }

      let names = getSortedImageNames("load-images", "load-image");
      const filesArray = Array.from(images);
      const sortedFiles = filesArray.sort((a, b) => names.indexOf(a.name) - names.indexOf(b.name));
        
      for (let i = 0; i < sortedFiles.length; i++) {
        const image = sortedFiles[i];
        formData.append('images', image, image.name);
      }
      let frm = new AsyncForm(formData, "POST", "/new-publication-post");
      frm.onResponse(function(response: Map<string, string>){
        if (response["success"] == "false"){
          showErrorOnPage(response["error"]);
        } else if (response["success"] == "true"){
          redirect("/explore");
        }
      });
      frm.onError(function(error: string){
        showErrorOnPage(error);
      });
      frm.send();
    });
  }
}
sendForm("new-publication-form");

sendLikeForm("publication-like");

sendCommentForm("publication-comment-form", function () {
  let drpd = publicationRunDropdown();
  hideShowComments();
  document.addEventListener('click', (event) => {
    const clickTarget = event.target as Node;
    closeDropdown(drpd, clickTarget);
  });
});