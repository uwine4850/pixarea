import { clickIfExistById } from "./utils";

clickIfExistById("profile-button-open-publication", (el)=>{
    document.getElementById("separate-content-left-profile").classList.add("separate-content-left-profile-hide");
    document.getElementById("separate-content-right-profile").classList.remove("separate-content-right-profile-hide");
});

clickIfExistById("profile-header-back", (el)=>{
    document.getElementById("separate-content-left-profile").classList.remove("separate-content-left-profile-hide");
    document.getElementById("separate-content-right-profile").classList.add("separate-content-right-profile-hide");
});