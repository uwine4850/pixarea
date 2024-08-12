import "../scss/style.scss";
import { Dropdown, closeDropdown } from "./dropdown";
import "./base";
import "./profile";
import { publicationRunDropdown, newPublicationRunDropdown } from "./publication/publication";
import { RemoveUrlParameters } from "./utils/url/urlUtils";

let drpd = new Dropdown("header-btn-dropdown");
drpd.run();

let pubDrpd = publicationRunDropdown()
let newPubDrpd = newPublicationRunDropdown()

document.addEventListener('click', (event) => {
    const clickTarget = event.target as Node;
    closeDropdown(drpd, clickTarget);
    closeDropdown(pubDrpd, clickTarget);
    closeDropdown(newPubDrpd, clickTarget);
});

RemoveUrlParameters(["REDIRECT_ERROR"]);