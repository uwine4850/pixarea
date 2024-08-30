const comm = `
    <div class="publication-comment-item">
        <!-- ENABLE CLASS PCH-enable -->
        <button class="publication-comment-item-hidden-panel">
            <a>
                Hidden comment. Click to show comment.
            </a>
        </button>
        <div class="comment-user-info">
            <button class="profile-icon profile-icon-comment">
                <a>
                    <img src="/media/avatars/88a5d4f503b26504ed25f8cf38a0b97ddfd234c4f024896a96a3fa363de0660b.jpg" alt="">
                </a>
            </button>
            <div class="comment-user-username">
                qqq1
            </div>
            <div class="dropdown-wrapper dropdown-wrapper-comm-menu">
                <button class="comment-menu comment-menu-dropdown">
                    <a>
                        <img src="/static/img/icons/dot_menu.svg" alt="">
                    </a>
                </button>
                <div class="dropdown-for-btn dropdown-for-btn-comm-menu dropdown-for-btn-hide">
                    <button class="dropdown-for-btn-linkBtn-item comment-hide-btn">
                        <a>
                            Hide
                        </a>
                    </button>
                </div>
            </div>
        </div>
        <div class="comment-content">
            wfwf
        </div>
        <form class="comment-footer">
            <button data-reply-name="qqq1" data-reply-id="8" type="button" class="reply-button">
                Reply
            </button>
            <button type="button" data-comment-id="8" class="comment-open-close-reply comment_answers_button">
                <a>
                    Answers 2
                    <img src="/static/img/icons/expand_right.svg" alt="">
                </a>
            </button>
        </form>
    </div>
`

export class BuildComment{
    private comment_id: string;
    private publication_id: string;
    private commentNode: Node;
    constructor(cloneNodeId: string, comment_id: string, publication_id: string){
        this.comment_id = comment_id;
        this.publication_id = publication_id;
        let nn = document.getElementById(cloneNodeId);
        this.commentNode = nn.cloneNode(true);
    }
    public enablePCH(enable: boolean){
        let enbleClass = "";
        if(enable){
            enbleClass = "PCH-enable";
        }
    }
    public insert(){
        console.log(this.commentNode);
    }
}

class PCH {
    private enable: boolean
    constructor(enable: boolean){
        this.enable = enable
    }

    public get(){
        let enbleClass = "";
        if(this.enable){
            enbleClass = "PCH-enable";
        }
        return `
        <button class="publication-comment-item-hidden-panel ${enbleClass}">
            <a>
                Hidden comment. Click to show comment.
            </a>
        </button>
        `
    }
}