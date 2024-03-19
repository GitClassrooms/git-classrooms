import { createFileRoute, redirect } from "@tanstack/react-router";
import GitlabLogo from "./../assets/gitlab_logo.svg";
import { Button } from "@/components/ui/button";
import { isAuthenticated } from "@/lib/utils";

export const Route = createFileRoute("/login")({
  validateSearch: (search: Record<string, unknown>) => {
    return {
      redirect: (search.redirect as string) || "",
    };
  },
  beforeLoad: async () => {
    if (await isAuthenticated()) {
      throw redirect({
        to: "/classrooms",
      });
    }
  },
  component: Login,
});

function Login() {
  return (
    <div className="flex flex-col items-center">
      <img src={GitlabLogo} className="h-96 w-96" />
      <form method="POST" action="/api/v1/auth/sign-in">
        <Button type="submit">Login with Gitlab</Button>
      </form>
    </div>
  );
}
