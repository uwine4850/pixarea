import React from "react";
import { CDropdown, TargetButton, Items, CheckboxItem } from "./Dropdown";

const HeaderDropdownButtons: React.FC = () => {
  return (
    <CDropdown className="header-dropdown-wrapper" targetButtonClass="header-btn-dropdown">
    <TargetButton className="header-btn header-btn-dropdown-style header-btn-dropdown">
      <a href="#">Categories</a>
    </TargetButton>
    <Items>
      <CheckboxItem
        id="ctgr_PIXELART"
        name="category"
        label="Pixelart"
      />
      <CheckboxItem
        id="ctgr_3D_CHARACTER"
        name="category"
        label="3D Character"
      />
      <CheckboxItem id="ctgr_GAMEDEV" name="category" label="Gamedev" />
    </Items>
  </CDropdown>
  );
};

export default HeaderDropdownButtons;
