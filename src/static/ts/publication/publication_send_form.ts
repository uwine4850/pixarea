import { showErrorOnPage } from "../errors/displayError";
import { getSortedImageNames } from "../image/showInputImage";
import { redirect } from "../utils/url/urlUtils";
import { AsyncForm } from "../form/asyncForm";

export function sendNewPublicationForm(formId: string){
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