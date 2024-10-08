import { showErrorOnPage } from "../errors/displayError";
import { AsyncForm } from "../form/asyncForm";
import { removePrefix } from "../utils/string/string";
import inscludes_pub_comm_menu from "/templates/publication/pub_comm_menu.html";

function commentText(name: string, avatarPath: string, text: string, comm_id: string){
    let avatar = ""
    if (avatarPath != "") {
        avatar = avatarPath;
    } else {
        avatar = "/static/img/default/default.jpg";
    }
    return `
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
            ${name}
        </div>
        ${inscludes_pub_comm_menu}
    </div>
    <div class="comment-content">
        ${text}
    </div>
    <form class="comment-footer">
        <input type="hidden" name="comm_id" value="${comm_id}">
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
}

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
                newElement.innerHTML = commentText(response["name"], response["avatar"], response["text"], response["comm_id"]);
                publication_comment_list.prepend(newElement);
                let publication_comment_text = document.getElementById("publication_comment_text") as HTMLTextAreaElement;
                if (publication_comment_text){
                    publication_comment_text.value = "";
                }
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

function replyComment(){
    let reply_buttons = document.getElementsByClassName("reply-button") as HTMLCollectionOf<HTMLButtonElement>;
    for (let index = 0; index < reply_buttons.length; index++) {
        const reply_button = reply_buttons[index];
        reply_button.onclick = function(){
            let reply_id = reply_button.getAttribute("data-reply-id");
            let replyName = reply_button.getAttribute("data-reply-name");
            let reply_id_input = document.getElementById("reply_id_input") as HTMLInputElement;
            reply_id_input.value = reply_id;
            let reply_to_user = document.getElementById("reply-to-user");
            reply_to_user.classList.remove("reply-to-user-closed");
            reply_to_user.innerHTML = replyName;
        }
    }
}

if (document.getElementById("reply-to-user")){
    document.getElementById("reply-to-user").onclick = function(){
        document.getElementById("reply-to-user").innerHTML = "";
        document.getElementById("reply-to-user").classList.add("reply-to-user-closed");
        let reply_id_input = document.getElementById("reply_id_input") as HTMLInputElement;
        reply_id_input.value = null;
    }
}

replyComment();
