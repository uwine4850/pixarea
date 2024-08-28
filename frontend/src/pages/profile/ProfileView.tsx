import React from "react";
import { LayoutProvider, Layout } from "../LayoutContext";
import ProfileHeader from "./ProfileHeader";

const ProfileView: React.FC = () => {
  return (
    <LayoutProvider value={{ leftSide: leftSide, rightSide: rightSide }}>
      <Layout />
    </LayoutProvider>
  );
};

export default ProfileView;

const leftSide = (
  <div
    className="separate-content-left separate-content-left-profile"
    id="separate-content-left-profile"
  >
    <div className="profile-wrap">
      <div className="profile-bg-avatar">
        <div className="profile-bg-avatar-blur"></div>
        <img src="images/default/default.jpg" alt="" />
        <button className="profile-edit-button">
          <a href="/profile-edit/USER_ID">
            <img src="images/icons/edit.svg" alt="" />
          </a>
        </button>
        <div className="profile-avatar">
          <img src="images/default/default.jpg" alt="" />
        </div>
      </div>
      <div className="profile-description">
        <div className="profile-id">
          <div className="profile-id-username">profile.Name</div>
          <div className="profile-id-uniq-username">
            profile.Username
          </div>
        </div>
        <div className="profile-description-text">
          DESCRIPTION
        </div>
        <hr className="profile-hr" />
        <div className="profile-info">
          <div className="profile-info-item">
            <div className="profile-info-item-title">Publications</div>
            <div className="profile-info-item-value">3573</div>
          </div>
          <div className="profile-info-item">
            <div className="profile-info-item-title">Subscribers</div>
            <div className="profile-info-item-value">1.2m</div>
          </div>
          <div className="profile-info-item">
            <div className="profile-info-item-title">Likes</div>
            <div className="profile-info-item-value">13.5m</div>
          </div>
        </div>
        <hr className="profile-hr" />
        <div className="profile-buttons">
          <button className="profile-button">
            <a>Subscribe</a>
          </button>
          <button
            className="profile-button profile-button-open-publication"
            id="profile-button-open-publication"
          >
            <a>Publications</a>
          </button>
        </div>
        <hr className="profile-hr" />
        <div className="profile-pin-publication">
          <div className="profile-pin-publication-text">Text for pin item</div>
          <button className="profile-pin-publication-pub">
            <a href="" className="profile-pin-publication-unpin-btn">
              <img src="images/icons/pin.svg" alt="" />
            </a>
            <a href="">
              <img src="images/temp/2.jpg" alt="" />
            </a>
          </button>
        </div>
      </div>
    </div>
  </div>
);

const rightSide = (
  <div
    className="separate-content-right separate-content-right-profile separate-content-right-profile-hide"
    id="separate-content-right-profile"
  >
    <div className="separate-content-right-header">
        <ProfileHeader/>
    </div>
    <div className="separate-content-right-content">
      <div className="collections-list">
        <button className="collection-item">
          <a href="">
            <div className="collection-item-image">
              <img src="images/temp/2.jpg" alt="" />
            </div>
            <div className="collection-item-name">Coll name</div>
          </a>
        </button>
        <button className="collection-item">
          <a href="">
            <div className="collection-item-image">
              <img src="images/temp/2.jpg" alt="" />
            </div>
            <div className="collection-item-name">Coll name</div>
          </a>
        </button>
        <button className="collection-item">
          <a href="">
            <div className="collection-item-image">
              <img src="images/temp/1.jpg" alt="" />
            </div>
            <div className="collection-item-name">Coll name</div>
          </a>
        </button>
        <button className="collection-item">
          <a href="">
            <div className="collection-item-image">
              <img src="images//temp/1.jpg" alt="" />
            </div>
            <div className="collection-item-name">Coll name</div>
          </a>
        </button>
      </div>
    </div>
  </div>
);
