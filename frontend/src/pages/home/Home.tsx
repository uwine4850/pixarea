import React, {useEffect, useState} from "react";
import { LayoutProvider, Layout } from "../LayoutContext";
import Header from "../components/Header";
import {useCsrfToken} from "../../scripts/csrf_token";

const content = (
  <div className="content">
    <div className="explore-content">
      {Array.from({ length: 20 }, (_, index) => (
        <div className="publication-item" key={index}>
          <img src={`/images/temp/${(index % 2) + 1}.jpg`} alt="" />
        </div>
      ))}
    </div>
  </div>
);

const Home: React.FC = () => {
  const csrfToken = useCsrfToken();
  console.log(csrfToken);
  return (
    <LayoutProvider
      value={{
        header: <Header />,
        content: content,
      }}
    >
      <Layout />
    </LayoutProvider>
  );
};

export default Home;
