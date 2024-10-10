import "./globals.css";

export const metadata = {
  title: "ColorFun",
  description: "RSC and WS Demo",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>
        {children}
      </body>
    </html>
  );
}
