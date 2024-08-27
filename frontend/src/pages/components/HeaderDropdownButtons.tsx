import React from "react";

const HeaderDropdownButtons: React.FC = () => {
  return (
    <div className="dropdown-wrapper header-dropdown-wrapper">
      <button className="header-btn header-btn-dropdown-style header-btn-dropdown">
        <a href="#">Categories</a>
      </button>
      <div className="dropdown-for-btn dropdown-for-btn-hide">
        <div className="dropdown-for-btn-checkbox-item">
          <input type="checkbox" id="ctgr_PIXELART" />
          <label htmlFor="ctgr_PIXELART">Pixelart</label>
        </div>
        <div className="dropdown-for-btn-checkbox-item">
          <input type="checkbox" id="ctgr_3D_CHARACTER" />
          <label htmlFor="ctgr_3D_CHARACTER">3D Character</label>
        </div>
        <div className="dropdown-for-btn-checkbox-item">
          <input type="checkbox" id="ctgr_GAMEDEV" />
          <label htmlFor="ctgr_GAMEDEV">Gamedev</label>
        </div>
      </div>
    </div>
  );
};

export default HeaderDropdownButtons;
