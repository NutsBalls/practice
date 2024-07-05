import {NextUIProvider} from "@nextui-org/system";
import { ThemeProvider as NextThemeProvider} from "next-themes";

export default function Provider({ children }) {
  return (
      <NextUIProvider>
        <NextThemeProvider defaultTheme="dark" attribute="class">
          {children}
        </NextThemeProvider>
      </NextUIProvider>
  );
}