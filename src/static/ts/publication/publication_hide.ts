import { showErrorOnPage } from "../errors/displayError";
import { AsyncForm } from "../form/asyncForm";

export function sendHideCommentForm(comm_id: string, comm_publication_id: string, onSucces: ()=>void){
    const formData = new FormData();
    formData.append("comm_id", comm_id);
    formData.append("comm_publication_id", comm_publication_id);

    let frm = new AsyncForm(formData, "POST", "/publication-comment-hide");
    frm.onResponse(function(response: Map<string, string>){
      if (response["success"] == "false"){
        showErrorOnPage(response["error"]);
      } else if (response["success"] == "true"){
        onSucces();
      }
    });
    frm.onError(function(error: string){
      showErrorOnPage(error);
    });
    frm.send();
}
