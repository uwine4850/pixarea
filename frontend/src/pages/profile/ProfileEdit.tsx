import React from "react";

const ProfileEdit: React.FC = () => {
    return (
        <div className="profile-edit-form">
    <h2 className="profile-edit-form-title">
        Profile edit
    </h2>
    <hr/>
    <form className="form-block" method="post" action="/profile-edit-post" encType="multipart/form-data">
        <div className="form-block-item">
            <label htmlFor="avatar">Avatar</label>
            <input id="avatar" name="avatar" type="file"/>
            <div className="form-block-item form-block-item-checkbox form-block-item-inner-checkbox">
                <input id="delete_avatar" name="delete_avatar" type="checkbox"/>
                <label htmlFor="delete_avatar">delete avatar</label>
            </div>
        </div>
        <div className="form-block-item">
            <label htmlFor="background">Background</label>
            <input id="background" name="background" type="file"/>
            <div className="form-block-item form-block-item-checkbox form-block-item-inner-checkbox">
                <input id="delete_background" name="delete_background" type="checkbox"/>
                <label htmlFor="delete_background">delete background</label>
            </div>
        </div>
        <div className="form-block-item">
            <label htmlFor="name">Name</label>
            <input id="name" name="name" type="text" value="profile.Name"/>
        </div>
        <div className="form-block-item">
            <label htmlFor="description">Description</label>
            <textarea name="description" id="description">PROFILE.DESCRIPTION</textarea>
        </div>
        <button className="save-profile-edit"><a>Save</a></button>
    </form>
</div>
    );
}

export default ProfileEdit;
