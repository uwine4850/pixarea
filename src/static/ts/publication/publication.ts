import { DragImage } from "../image/dragImages";
import { Dropdown, closeDropdown} from "../dropdown";
import { ShowInputImages } from "../image/showInputImage";
import { sendLikeForm } from "./publication_like";
import { sendCommentForm } from "./publication_comment";
import { sendNewPublicationForm } from "./publication_send_form";
import { sendHideCommentForm } from "./publication_hide";
import { sendLoadAswersForm } from "./load_answers";
import { BuildComment } from "./comment";


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
        let comm_id = btn.getAttribute("data-comm-id");
        let publication_id = btn.getAttribute("data-publication-id");
        if(comm_id && publication_id){
          sendHideCommentForm(comm_id, publication_id, function(){
            comment.classList.add("publication-comment-item-hidden");
            PCH.classList.add("PCH-enable");
          });
        }
      }
    }
  }

  // SHOW HIDDEN COMMENT
  let PCH = document.getElementsByClassName("publication-comment-item-hidden-panel") as HTMLCollectionOf<HTMLButtonElement>;
  for (let i = 0; i < PCH.length; i++) {
    PCH[i].onclick = function() {
      let comment = PCH[i].parentElement;
      let comm_id = PCH[i].getAttribute("data-comm-id");
      let publication_id = PCH[i].getAttribute("data-publication-id");
      if(comm_id && publication_id){
        sendHideCommentForm(comm_id, publication_id, function(){
          comment.classList.add("publication-comment-item-hidden");
          comment.classList.remove("publication-comment-item-hidden");
          PCH[i].classList.remove("PCH-enable");
        });
      }
    }
  }
}

hideShowComments();

sendNewPublicationForm("new-publication-form");
sendLikeForm("publication-like");
sendCommentForm("publication-comment-form", function () {
  let drpd = publicationRunDropdown();
  hideShowComments();
  document.addEventListener('click', (event) => {
    const clickTarget = event.target as Node;
    closeDropdown(drpd, clickTarget);
  });
});
sendLoadAswersForm("comment_answers_button");
let bb = new BuildComment("commentNode", "", "");
bb.insert();