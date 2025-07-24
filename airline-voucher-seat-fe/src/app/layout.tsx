import type { Metadata } from "next";
import "./globals.css";
import NavBar from "@/components/Navbar";
import { ToastContainer } from "react-toastify";

export const metadata: Metadata = {
  title: "Voucher Seat Management",
  description: "Voucher seat assignment application",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`antialiased bg-white`}>
        <ToastContainer />
        <NavBar />
        {children}
      </body>
    </html>
  );
}
