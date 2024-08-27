import React, { createContext, useContext, ReactNode } from 'react';

interface LayoutContextType {
  header?: ReactNode;
  leftSide?: ReactNode;
  rightSide?: ReactNode;
  content: ReactNode;
}

const LayoutContext = createContext<LayoutContextType | undefined>(undefined);

export const useLayoutContext = () => {
  const context = useContext(LayoutContext);
  if (!context) {
    throw new Error('useLayoutContext must be used within a LayoutProvider');
  }
  return context;
};

interface LayoutProviderProps {
  value: LayoutContextType;
  children: ReactNode;
}

export const LayoutProvider: React.FC<LayoutProviderProps> = ({ value, children }) => {
  return <LayoutContext.Provider value={value}>{children}</LayoutContext.Provider>;
};

export const Layout: React.FC = () => {
  const { header, leftSide, rightSide, content } = useLayoutContext();

  return (
    <>
      <div className="separate-content">
        {leftSide}
        {rightSide}
      </div>
      {header}
      {content}
    </>
  );
};