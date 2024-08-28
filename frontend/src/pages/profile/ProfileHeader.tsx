import React from "react";
import Profile from "../components/Profile";
import Search from "../components/Search";
import HeaderDropdownButtons from "../components/HeaderDropdownButtons";

const ProfileHeader: React.FC = () => {
  return (
    <>
      <header className="default-heder default-heder-media">
        <div className="header-center">
          <Profile userAvatar="/images/default/default.jpg" userId="1" />
          <button className="header-btn">
            <a href="#">Publications</a>
          </button>
          <button className="header-btn btn-active">
            <a href="#">Collections</a>
          </button>
          <Search />
          <HeaderDropdownButtons />
        </div>
      </header>
      {/* MINI HEADER */}
      <header className="default-heder mini-header">
        <div className="header-center mini-header-center">
          <div className="mini-header-side" id="mhs1">
            <Profile userAvatar="/images/default/default.jpg" userId="1" />
            <button className="header-btn">
              <a href="#">Publications</a>
            </button>
            <button className="header-btn btn-active">
              <a href="#">Collections</a>
            </button>
            <button
              className="mini-header-search-btn"
              id="mini-header-search-btn"
            >
              <a href="#">
                <img src="/static/img/icons/search.svg" alt="" />
              </a>
            </button>
          </div>
          <div className="mini-header-side mini-header-side-hide" id="mhs2">
            <button
              className="mini-header-close-search"
              id="mini-header-close-search"
            >
              <a href="#">
                <img src="/static/img/icons/close.svg" alt="" />
              </a>
            </button>
            <Search />
            <HeaderDropdownButtons />
          </div>
        </div>
        <button className="profile-header-back" id="profile-header-back">
          <a>
            <img src="/static/img/icons/back.svg" alt="" />
          </a>
        </button>
      </header>
    </>
  );
};

export default ProfileHeader;