import { clickIfExistById } from "./utils";

clickIfExistById("mini-header-search-btn", (el)=>{
    document.getElementById("mhs1").classList.add("mini-header-side-hide");
    document.getElementById("mhs2").classList.remove("mini-header-side-hide");
});

clickIfExistById("mini-header-close-search", (el)=>{
    document.getElementById("mhs1").classList.remove("mini-header-side-hide");
    document.getElementById("mhs2").classList.add("mini-header-side-hide");
});
