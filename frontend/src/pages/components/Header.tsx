import Profile from "./Profile";
import Search from "./Search";
import {
  CDropdown,
  TargetButton,
  Items,
  CheckboxItem,
  ButtonItem,
} from "./Dropdown";

const Header: React.FC = () => {
  return (
    <>
      {/* Base header */}
      <header className="default-heder default-heder-media">
        <div className="header-center">
          <Profile userAvatar="/images/default/default.jpg" userId="1" />
          <button className="header-btn">
            <a href="#">Explore</a>
          </button>
          <button className="header-btn btn-active">
            <a href="#">Subscribed</a>
          </button>
          <Search />
          <CDropdown>
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
        </div>
      </header>

      {/* Mini-header */}
      <header className="default-heder mini-header">
        <div className="header-center mini-header-center">
          <div className="mini-header-side" id="mhs1">
            <Profile userAvatar="/images/default/default.jpg" userId="1" />
            <button className="header-btn">
              <a href="#">Explore</a>
            </button>
            <button className="header-btn btn-active">
              <a href="#">Subscribed</a>
            </button>
            <button
              className="mini-header-search-btn"
              id="mini-header-search-btn"
            >
              <a href="#">
                <img src="/static/img/icons/search.svg" alt="Search" />
              </a>
            </button>
          </div>
          <div className="mini-header-side mini-header-side-hide" id="mhs2">
            <button
              className="mini-header-close-search"
              id="mini-header-close-search"
            >
              <a href="#">
                <img src="/static/img/icons/close.svg" alt="Close" />
              </a>
            </button>
            <Search />
            <CDropdown>
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
                <CheckboxItem
                  id="ctgr_GAMEDEV"
                  name="category"
                  label="Gamedev"
                />
              </Items>
            </CDropdown>
          </div>
        </div>
      </header>
    </>
  );
};

export default Header;
