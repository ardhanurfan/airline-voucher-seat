import clsx from "clsx";
import React from "react";

interface FieldInputProps {
  label?: string;
  value: string;
  onChange: React.ChangeEventHandler<HTMLInputElement> | undefined;
  type?: "text" | "date";
  placeholder: string;
  errorMessage?: string;
  name: string;
}

const FieldInput: React.FC<FieldInputProps> = ({
  label,
  value,
  onChange,
  type = "text",
  placeholder,
  errorMessage,
  name,
}) => {
  return (
    <div className="flex flex-col space-y-1 w-full">
      {label && (
        <label htmlFor={name} className="text-sm font-bold text-gray-700">
          {label}
        </label>
      )}
      <div
        className={clsx(
          "flex items-center border rounded-md overflow-hidden",
          errorMessage
            ? "border-red-500"
            : "border-gray-300 focus-within:border-primary"
        )}
      >
        <input
          id={name}
          name={name}
          type={type}
          value={value}
          onChange={onChange}
          className={clsx(
            "flex-1 px-4 py-2 outline-none",
            errorMessage ? "text-red-500" : "text-gray-700"
          )}
          placeholder={placeholder}
        />
      </div>
      {errorMessage && <p className="text-sm text-red-500">{errorMessage}</p>}
    </div>
  );
};

export default FieldInput;
