import React, { useEffect, useRef } from "react";
import { Dropdown, closeDropdown } from "../../scripts/dropdown";

interface CDropdownProps {
  className?: string;
  targetButtonClass: string;
  children: React.ReactNode;
}

export const CDropdown: React.FC<CDropdownProps> = ({ className, targetButtonClass, children }) => {
  const dropdownRef = useRef<Dropdown | null>(null);

  useEffect(() => {
    dropdownRef.current = new Dropdown(targetButtonClass);
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
    <div className={`dropdown-wrapper ${className}`}>{children}</div>
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
  className?: string;
  children: React.ReactNode;
}

export const Items: React.FC<ItemsProps> = ({ className, children }) => {
  return (
    <div className={`dropdown-for-btn dropdown-for-btn-hide ${className}`}>{children}</div>
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
  dataAttributes?: Record<string, string>;
  onClick?: (event: React.MouseEvent<HTMLButtonElement>) => void;
  children?: React.ReactNode;
}

export const ButtonItem: React.FC<ButtonItemProps> = ({
  id,
  className,
  dataAttributes,
  onClick,
  children,
}) => {
  // const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
  //   if (onClick) {
  //     onClick(event.currentTarget);
  //   }
  // };
  return (
    <button
      id={id}
      className={className}
      onClick={onClick}
      {...dataAttributes}
    >
      {children}
    </button>
  );
};
