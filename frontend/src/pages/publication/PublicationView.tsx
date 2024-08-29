import React from "react";
import { LayoutProvider, Layout } from "../LayoutContext";
import Header from "../components/Header";
import { Comments, Comment } from "./Comments";

const PublicationView: React.FC = () => {
  return (
    <LayoutProvider value={{ leftSide: leftSide, rightSide: rightSide }}>
      <Layout />
    </LayoutProvider>
  );
};

export default PublicationView;

const leftSide = (
  <div className="pub-left-content">
    <span className="error" id="error"></span>
    <div className="pub-author-block">
      <button className="profile-icon">
        <a>
          <img src="images/default/default.jpg" alt="" />
        </a>
      </button>
      <div className="pub-author-right">
        <div className="pub-author-username">AUTHOR.USERNAME</div>
        <div className="pub-author-buttons">
          <button className="pub-author-subscribe">
            <a>Subscribe</a>
          </button>
          <form id="publication-like" className="publication-like-form">
            <input type="hidden" name="publication-id" value="PUBLICATION.ID" />
            <button id="pub-author-like" className="pub-author-like">
              <a>
                <img src="images/icons/like.svg" alt="" />
              </a>
            </button>
          </form>
        </div>
      </div>
    </div>
    <hr />
    <div className="pub-description-block">
      <h2 className="pub-name">PUBLICATION.NAME</h2>
      <div className="pub-description">PUBLICATION.DESCRPTION</div>
    </div>
    <hr />
    <div className="pub-info">
      <div className="pub-info-categories">
        <div className="pub-info-category">CATEGORY.NAME</div>
      </div>
      <div className="pub-info-with-icons">
        <div className="pub-info-with-icons-item">
          <img src="images/icons/date.svg" alt="" />
          <div className="pub-info-with-icons-text">PUBLICATION.NAME</div>
        </div>
        <div className="pub-info-with-icons-item">
          <img src="images/icons/star.svg" alt="" />
          <div className="pub-info-with-icons-text">Best...</div>
        </div>
        <div className="pub-info-with-icons-item">
          <img src="images/icons/like.svg" alt="" />
          <div className="pub-info-with-icons-text" id="publication-like-count">
            2222
          </div>
        </div>
      </div>
      <hr />
      <div className="publication-images mini-publication-images">
        <div className="publication-images-item">
          <img src="images/temp/1.jpg" alt="" />
        </div>
        <div className="publication-images-item">
          <img src="images/temp/2.jpg" alt="" />
        </div>
      </div>
      <div className="previuos-publication-block">
        <h3>Previuos publication</h3>
        <div className="previuos-publication-list">
          <button className="previuos-publication-item">
            <a href="">
              <img src="images/temp/1.jpg" alt="" />
            </a>
          </button>
        </div>
      </div>
      <hr />
      <Comments>
        <Comment isHide={true} />
        <Comment isHide={false} />
      </Comments>
    </div>
  </div>
);

const rightSide = (
  <div className="separate-content-right separate-content-right-publication">
    <Header />
    <div className="publication-images">
      <div className="publication-images-item">
        <img src="images/temp/1.jpg" alt="" />
      </div>
    </div>
  </div>
);
