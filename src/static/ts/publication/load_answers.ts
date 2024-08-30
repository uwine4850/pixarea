import { AsyncForm } from "../form/asyncForm";

export function sendLoadAswersForm(loadButtonsClass: string){
    let loadButtons = document.getElementsByClassName(loadButtonsClass) as HTMLCollectionOf<HTMLButtonElement>;
    if (!loadButtons){
        return
    }
    for (let i = 0; i < loadButtons.length; i++) {
        const btn = loadButtons[i];
        btn.onclick = function(){
            let comm_id = btn.getAttribute("data-comment-id");
            
            let url = `/publication-load-answers?comm_id=${encodeURIComponent(comm_id)}`;
            let frm = new AsyncForm(null, "GET", url);
            frm.onResponse(function(response: Map<string, string>){
              if (response["success"] == "false"){
                  console.log(response["error"]);
              } else if (response["success"] == "true"){
                console.log(response["comments"]);
              }
            });
            frm.onError(function(error: string){
                console.log(error);
            });
            frm.send();
        }
    }
}