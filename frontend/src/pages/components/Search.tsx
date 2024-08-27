import React from "react";

const Search: React.FC = () => {
  return (
    <form className="search">
      <input placeholder="Search..." type="text" />
      <button>
        <img src="/static/img/icons/search.svg" alt="" />
      </button>
    </form>
  );
};

export default Search;
    