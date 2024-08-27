import React from "react";

interface ProfileProps {
  userAvatar: string;
  userId: string;
}

const Profile: React.FC<ProfileProps> = ({ userAvatar, userId }) => {
  return (
    <div className="dropdown-wrapper header-dropdown-wrapper-profile">
      <button className="profile-icon header-btn-dropdown">
        <a>
          <img src={userAvatar} alt="User Avatar" />
        </a>
      </button>
      <div className="dropdown-for-btn dropdown-for-btn-profile dropdown-for-btn-hide">
        <button className="dropdown-for-btn-linkBtn-item">
          <a href={`/profile/${userId}`}>Profile</a>
        </button>
        <button className="dropdown-for-btn-linkBtn-item">
          <a href="/new-publication">New publication</a>
        </button>
      </div>
    </div>
  );
};

export default Profile;
