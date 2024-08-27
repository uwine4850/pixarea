import React, { useEffect, useRef } from "react";
import { Dropdown, closeDropdown } from "../../scripts/dropdown";

interface CDropdownProps {
  children: React.ReactNode;
}

export const CDropdown: React.FC<CDropdownProps> = ({ children }) => {
  const dropdownRef = useRef<Dropdown | null>(null);

  useEffect(() => {
    dropdownRef.current = new Dropdown("header-btn-dropdown");
    dropdownRef.current.run();

    const handleClickOutside = (event: MouseEvent) => {
      const clickTarget = event.target as Node;
      if (dropdownRef.current) closeDropdown(dropdownRef.current, clickTarget);
    };
    document.addEventListener("click", handleClickOutside);

    return () => {
      dropdownRef.current = null;
      document.removeEventListener("click", handleClickOutside);
    };
  }, []);
  return (
    <div className="dropdown-wrapper header-dropdown-wrapper">{children}</div>
  );
};

interface TargetButtonProps {
  className: string;
  children: React.ReactNode;
}

export const TargetButton: React.FC<TargetButtonProps> = ({
  className,
  children,
}) => {
  return <button className={className}>{children}</button>;
};

interface ItemsProps {
  children: React.ReactNode;
}

export const Items: React.FC<ItemsProps> = ({ children }) => {
  return (
    <div className="dropdown-for-btn dropdown-for-btn-hide">{children}</div>
  );
};

interface CheckboxItemProps {
  id?: string;
  className?: string;
  name: string;
  label: string;
}

export const CheckboxItem: React.FC<CheckboxItemProps> = ({
  id,
  className,
  name,
  label,
}) => {
  return (
    <div className="dropdown-for-btn-checkbox-item">
      <input type="checkbox" id={id} name={name} className={className} />
      <label htmlFor={id}>{label}</label>
    </div>
  );
};

interface ButtonItemProps {
  id?: string;
  className?: string;
  dataAttributes: Record<string, string>;
  onClick?: (target: HTMLElement) => void;
  children: React.ReactNode;
}

export const ButtonItem: React.FC<ButtonItemProps> = ({
  id,
  className,
  dataAttributes,
  onClick,
  children,
}) => {
  const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    if (onClick) {
      onClick(event.currentTarget);
    }
  };
  return (
    <button
      id={id}
      className={className}
      onClick={handleClick}
      {...dataAttributes}
    >
      {children}
    </button>
  );
};
