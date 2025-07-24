"use client";

import clsx from "clsx";
import Link from "next/link";
import { usePathname } from "next/navigation";

const links = [
  {
    href: "/",
    label: "Generate",
  },
  {
    href: "/vouchers",
    label: "Vouchers",
  },
];

export default function NavBar() {
  const pathname = usePathname();

  return (
    <nav className="fixed z-50 top-0 w-full shadow-sm bg-primary text-white flex items-center justify-between h-14 px-6 md:px-12 xl:px-16">
      <Link href="/" className="md:text-xl font-medium">
        VOUCHER SEAT APP
      </Link>
      <div className="space-x-4">
        {links.map((link) => (
          <Link
            href={link.href}
            key={link.href}
            className={clsx(
              "hover:underline",
              pathname === link.href ? "text-secondary" : "text:white"
            )}
          >
            {link.label}
          </Link>
        ))}
      </div>
    </nav>
  );
}
