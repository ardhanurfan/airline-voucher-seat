"use client";

export default function Button({
  children,
  variant = "primary",
  ...props
}: React.ButtonHTMLAttributes<HTMLButtonElement> & {
  variant?: "primary" | "secondary";
}) {
  const base = "px-4 py-2 rounded-lg font-semibold transition w-full";
  const variants = {
    primary: "bg-primary text-white hover:bg-[#003377]",
    secondary: "bg-secondary text-white hover:bg-[#cc5500]",
  } as const;
  return (
    <button className={`${base} ${variants[variant]}`} {...props}>
      {children}
    </button>
  );
}
