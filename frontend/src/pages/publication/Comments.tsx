import React from "react";
import CommentDropdown from "./CommentDropdown";

interface CommentsProps {
  children: React.ReactNode;
}

export const Comments: React.FC<CommentsProps> = ({ children }) => {
  return (
    <div className="publication-comments-block">
      <h3>Comments</h3>
      <form className="publication-comment-form" id="publication-comment-form">
        <input type="hidden" name="publication_id" value="PUBLICATION.ID" />
        <input type="hidden" name="reply_id" id="reply_id_input" />
        <span
          id="reply-to-user"
          className="reply-to-user reply-to-user-closed"
        ></span>
        <textarea id="publication_comment_text" name="comment_text"></textarea>
        <button type="submit" className="send-publication-comment">
          Send
        </button>
      </form>
      <div className="publication-comment-list" id="publication-comment-list">
        {children}
      </div>
    </div>
  );
};

interface CommentProps {
  isHide: boolean;
}

export const Comment: React.FC<CommentProps> = ({ isHide }) => {
  return (
    <div
      className={`publication-comment-item ${
        isHide === true ? "publication-comment-item-hidden" : ""
      }`}
      id="commentNode"
    >
      <button
        className={`publication-comment-item-hidden-panel ${
          isHide === true ? "PCH-enable" : ""
        }`}
        data-comm-id="COMMENT.ID"
        data-publication-id="PUBLICATION.ID"
        onClick={showComment}
      >
        <a>Hidden comment. Click to show comment.</a>
      </button>
      <div className="comment-user-info">
        <button className="profile-icon profile-icon-comment">
          <a>
            <img src="images/default/default.jpg" alt="" />
          </a>
        </button>
        <div className="comment-user-username">
          AUTHOR.NAME
        </div>
        <CommentDropdown />
      </div>
      <div className="comment-content">
        COMMENT.TEXT
      </div>
      <form className="comment-footer">
        <button
          data-reply-name="AUTHOR.NAME"
          data-reply-id="COMMENT.ID"
          type="button"
          className="reply-button"
        >
          Reply
        </button>
        <button
          type="button"
          data-comment-id="COMMENT.ID"
          className="comment-open-close-reply comment_answers_button"
        >
          <a>
            Answers 11111
            <img src="images/icons/expand_right.svg" alt="" />
          </a>
        </button>
      </form>
    </div>
  );
};

export function showComment(ev: React.MouseEvent<HTMLButtonElement>) {
  const target = ev.currentTarget;
  const comment = target.parentElement;
  if (comment){
    target.classList.remove("PCH-enable");
    comment.classList.remove("publication-comment-item-hidden");
  }
}

export function hideComments(ev: React.MouseEvent<HTMLButtonElement>){
  const target = ev.currentTarget
  const comment = target.closest(".publication-comment-item");
  let PCH = null;
  if(comment){
    PCH = comment.querySelector(".publication-comment-item-hidden-panel") as HTMLButtonElement;
  }
  if (comment && PCH){
    comment.classList.add("publication-comment-item-hidden");
    PCH.classList.add("PCH-enable");
  }
}