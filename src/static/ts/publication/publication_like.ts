import { showErrorOnPage } from "../errors/displayError";
import { AsyncForm } from "../form/asyncForm";


function changeLike(response){
  let like_count = document.getElementById("publication-like-count");
  if (like_count){
      let likeElem = document.getElementById("publication-like-count");
      if (likeElem){
          document.getElementById("pub-author-like")!.classList.toggle("publcation-liked");
          let likes = parseInt(like_count.innerText);
          if(response["addLike"] == "true"){
              likes++;
          } else if (response["addLike"] == "false"){
              likes--;
          }
          likeElem.innerHTML = String(likes);
      }
  }
}

export function sendLikeForm(formId: string){
    let form = document.getElementById(formId) as HTMLFormElement;
    if (form){
      form.addEventListener('submit', (event: Event) => {
        event.preventDefault();
  
        const formData = new FormData(form);

        let frm = new AsyncForm(formData, "POST", "/publication-like");
        frm.onResponse(function(response: Map<string, string>){
          if (response["success"] == "false"){
            showErrorOnPage(response["error"]);
          } else if (response["success"] == "true"){
            changeLike(response);
          }
        });
        frm.onError(function(error: string){
          showErrorOnPage(error);
        });
        frm.send();
      });
    }
}