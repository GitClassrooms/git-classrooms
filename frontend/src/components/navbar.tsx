import GitlabLogo from "@/assets/gitlab_logo.svg";
import { ModeToggle } from "@/components/modeToggle";
import { Link } from "@tanstack/react-router";
import { useCsrf } from "@/provider/csrfProvider";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { LogOut, Menu, Moon, Sun, User as UserIcon } from "lucide-react";
import { Avatar } from "./avatar";
import { GetMeResponse } from "@/swagger-client";
import { Button } from "@/components/ui/button";
import { useTheme } from "@/provider/themeProvider";
import { useEffect, useState } from "react";
import { Switch } from "@/components/ui/switch";
import { Sheet, SheetContent, SheetHeader, SheetTitle, SheetTrigger } from "./ui/sheet";

export function Navbar(props: { auth: GetMeResponse | null }) {
  const { csrfToken } = useCsrf();
  const { theme, setTheme } = useTheme();
  const [isDarkMode, setIsDarkMode] = useState(theme === "dark");

  useEffect(() => {
    setTheme(isDarkMode ? "dark" : "light");
  }, [isDarkMode, setTheme]);

  return (
    <div className="flex justify-end sticky top-0 bg-white dark:bg-slate-900 mb-8">
      <Sheet>
        <SheetTrigger asChild>
          <Button size="icon" className="md:hidden m-2">
            <Menu />
          </Button>
        </SheetTrigger>
        <SheetContent>
          <SheetHeader>
            <SheetTitle>
              <a href="/" className="">
                <img className="h-14" src={GitlabLogo} alt="Gitlab Logo" />
              </a>
            </SheetTitle>
          </SheetHeader>
          <nav></nav>
        </SheetContent>
      </Sheet>
      <nav className="hidden md:flex justify-between px-8 py-2.5 border-b w-full">
        <div className="flex items-center">
          <a href="/" className="">
            <img className="h-14" src={GitlabLogo} alt="Gitlab Logo" />
          </a>
          {props.auth ? (
            <ul className="flex">
              <li className="content-center">
                <Link
                  to="/"
                  className="font-medium text-sm px-4 py-2 hover:underline"
                  activeProps={{ className: "!font-bold" }}
                >
                  Dashboard
                </Link>
              </li>
            </ul>
          ) : (
            <div />
          )}
        </div>
        <div className="flex items-center">
          <div className="px-4 py-2"></div>
          {props.auth ? (
            <DropdownMenu>
              <DropdownMenuTrigger>
                <Avatar
                  avatarUrl={props.auth.gitlabAvatar?.avatarURL}
                  fallbackUrl={props.auth.gitlabAvatar?.fallbackAvatarURL}
                  name={props.auth.name!}
                />
              </DropdownMenuTrigger>
              <DropdownMenuContent>
                <DropdownMenuLabel>
                  <div className="font-medium">{props.auth.name}</div>
                  <div className="text-sm text-muted-foreground md:inline">@{props.auth.gitlabUsername}</div>
                </DropdownMenuLabel>
                <DropdownMenuSeparator />
                <a
                  href={props.auth.gitlabUrl}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="flex items-center w-full"
                >
                  <DropdownMenuItem className="w-full">
                    <UserIcon className="mr-2 h-4 w-4" />
                    <span>Profile</span>
                  </DropdownMenuItem>
                </a>

                <DropdownMenuItem
                  onClick={() => setIsDarkMode((prev) => !prev)}
                  onSelect={(event) => event.preventDefault()}
                >
                  <Sun className="h-[1.2rem] w-[1.2rem] mr-2 dark:text-white text-black" />
                  <Switch checked={isDarkMode} />
                  <Moon className="h-[1.2rem] w-[1.2rem] ml-2 dark:text-white text-black" />
                  <span className="sr-only">Toggle theme</span>
                </DropdownMenuItem>
                <DropdownMenuSeparator />

                <form id="logOutForm" method="POST" action="/api/v1/auth/sign-out">
                  <button type="submit" className="w-full">
                    <DropdownMenuItem className="w-full">
                      <LogOut className="mr-2 h-4 w-4" />
                      <input type="hidden" name="csrf_token" value={csrfToken} />
                      <span className="font-bold">Log out</span>
                    </DropdownMenuItem>
                  </button>
                </form>
              </DropdownMenuContent>
            </DropdownMenu>
          ) : (
            <div className="flex">
              <ModeToggle />
              <Button className="ml-3">
                <Link to="/login" search={{ redirect: location.href }}>
                  Login
                </Link>
              </Button>
            </div>
          )}
        </div>
      </nav>
    </div>
  );
}
