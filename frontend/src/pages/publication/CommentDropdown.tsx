import React from "react";
import {
  CDropdown,
  TargetButton,
  Items,
  ButtonItem,
} from "../components/Dropdown";
import { hideComments } from "./Comments";

const CommentDropdown: React.FC = () => {
  return (
    <CDropdown className="dropdown-wrapper-comm-menu" targetButtonClass="comment-menu-dropdown">
      <TargetButton className="comment-menu comment-menu-dropdown">
        <a>
          <img src="images/icons/dot_menu.svg" alt="" />
        </a>
      </TargetButton>
      <Items className="dropdown-for-btn-comm-menu">
        <ButtonItem
          className="dropdown-for-btn-linkBtn-item comment-hide-btn"
          dataAttributes={{
            "data-comm-id": "COMMENT.ID",
            "data-publication-id": "PUB.ID",
          }}
          onClick={hideComments}
        >
          <a>Hide</a>
        </ButtonItem>
      </Items>
    </CDropdown>
  );
};

export default CommentDropdown;
