import clsx from "clsx";

interface DropdownProps {
  name: string;
  label?: string;
  value: string;
  onChange: (e: React.ChangeEvent<HTMLSelectElement>) => void;
  options: string[];
  placeholder: string;
  errorMessage?: string;
}

const Dropdown: React.FC<DropdownProps> = ({
  name,
  label,
  value,
  onChange,
  options,
  placeholder,
  errorMessage,
}) => {
  return (
    <div className="flex flex-col space-y-1 w-full">
      {label && (
        <label htmlFor={name} className="text-sm font-bold text-gray-700 block">
          {label}
        </label>
      )}
      <select
        disabled={options.length === 0}
        id={name}
        name={name}
        value={value}
        onChange={onChange}
        className={clsx(
          "border rounded-lg px-4 py-2 outline-none",
          errorMessage
            ? "border-red-500 text-red-700 focus:ring-red-500"
            : "border-gray-300 text-gray-900 focus:border-primary focus:ring-primary"
        )}
      >
        <option value="" disabled>
          {placeholder}
        </option>

        {options.map((opt) => (
          <option key={opt} value={opt}>
            {opt}
          </option>
        ))}
      </select>
      {errorMessage && <p className="text-sm text-red-500">{errorMessage}</p>}
    </div>
  );
};

export default Dropdown;
