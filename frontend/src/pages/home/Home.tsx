import React from "react";
import { LayoutProvider, Layout } from "../LayoutContext";
import Header from "../components/Header";

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
