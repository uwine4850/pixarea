import { showErrorOnPage } from "./errors/displayError";
import { AsyncForm } from "./form/asyncForm";
import { removePrefix } from "./utils/string/string";
import inscludes_pub_comm_menu from "/templates/publication/pub_comm_menu.html";

export function sendCommentForm(formId: string, onSuccess: () => void){
    let form = document.getElementById(formId) as HTMLFormElement;
    if (form){
      form.addEventListener('submit', (event: Event) => {
        event.preventDefault();
  
        const formData = new FormData(form);

        let frm = new AsyncForm(formData, "POST", "/publication-comment");
        frm.onResponse(function(response: Map<string, string>){
          if (response["success"] == "false"){
            console.log(response["error"]);
            
            showErrorOnPage(response["error"]);
          } else if (response["success"] == "true"){
            let publication_comment_list = document.getElementById("publication-comment-list");
            if (publication_comment_list){
                const newElement = document.createElement('div');
                newElement.classList.add("publication-comment-item");
                let avatar = ""
                if (response["avatar"] != "") {
                    avatar = response["avatar"];
                } else {
                    avatar = "/static/img/default/default.jpg";
                }
                newElement.innerHTML = `
                <button class="publication-comment-item-hidden-panel">
                    <a>
                        Hidden comment. Click to show comment.
                    </a>
                 </button>
                <div class="comment-user-info">
                    <button class="profile-icon profile-icon-comment">
                        <a>
                            <img src="${removePrefix(avatar, "src")}" alt="">
                        </a>
                    </button>
                    <div class="comment-user-username">
                        ${response["name"]}
                    </div>
                    ${inscludes_pub_comm_menu}
                </div>
                <div class="comment-content">
                    ${response["text"]}
                </div>
                <form class="comment-footer">
                    <input type="hidden" name="comm_id" value="${ response["comm_id"] }">
                    <button class="reply-button">
                        Reply
                    </button>
                    <button class="comment-open-close-reply">
                        <a>
                            Answers
                            <img src="/static/img/icons/expand_right.svg" alt="">
                        </a>
                    </button>
                </form>
                `
                publication_comment_list.prepend(newElement);
                onSuccess();
            }
          }
        });
        frm.onError(function(error: string){
          showErrorOnPage(error);
        });
        frm.send();
      });
    }
}