'use client';

import React, { useEffect, useState } from "react";
import {Switch} from "@nextui-org/react";
import {MoonIcon} from "./MoonIcon";
import {SunIcon} from "./SunIcon";
import { useTheme } from "next-themes";

export const ThemeSwitcher = () => {

  const [mounted, setMounted] = useState(false)
  const { theme, setTheme } = useTheme()

  // useEffect only runs on the client, so now we can safely show the UI
  useEffect(() => {
    setMounted(true)
  }, [])

  if (!mounted) {
    return null
  }


  return (
    <Switch
      defaultSelected={global?.localStorage.getItem("theme") == 'light' ? true : false}
      size="lg"
      color="primary"
      startContent={<SunIcon />}
      endContent={<MoonIcon />}
      onValueChange={(isSelected) => {
        if (isSelected){
          global?.localStorage.setItem("theme", "light")
          setTheme("light")
          return
        }
        global?.localStorage.setItem("theme", "dark")
        setTheme("dark")
      }}>
    </Switch>
  );
}