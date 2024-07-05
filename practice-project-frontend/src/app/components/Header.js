'use client'


import React from "react";
import {Navbar, NavbarBrand, NavbarContent, NavbarItem, Link} from "@nextui-org/react";
import { ThemeSwitcher } from "./ThemeSwitcher";
import { usePathname } from "next/navigation";

export const Header = () => {
  var currentPath = usePathname()
  return (
    <Navbar>
      <NavbarBrand justify="start">
        <p className="font-bold text-inherit">HH-Parser</p>
      </NavbarBrand>
      <NavbarContent className="hidden sm:flex gap-4" justify="center">
        <NavbarItem isActive={(currentPath == "/parse")}>
          <Link color={(currentPath == "/parse") ? "primary": "foreground"} href="/parse">
            Parse to database
          </Link>
        </NavbarItem>
        <NavbarItem isActive={(currentPath == "/get")}>
          <Link href="/get" color={(currentPath == "/get") ? "primary": "foreground"}>
            Get parsed data
          </Link>
        </NavbarItem>
        <NavbarItem isActive={(currentPath == "/analytics")}>
          <Link color={(currentPath == "/analytics") ? "primary": "foreground"} href="/analytics">
            Analytics
          </Link>
        </NavbarItem>
      </NavbarContent>
      <NavbarContent justify="end">
        <NavbarItem>
          <ThemeSwitcher/>
        </NavbarItem>
      </NavbarContent>
    </Navbar>
  );
}
