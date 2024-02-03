import { createFileRoute, redirect } from "@tanstack/react-router";
import GitlabLogo from './../assets/gitlab_logo.svg'
import { Button } from "@/components/ui/button";
import { isAuthenticated } from '@/lib/utils'

export const Route = createFileRoute("/login")({
  beforeLoad: async () => {
    if (await isAuthenticated()) {
      throw redirect({
        to: '/dashboard',
      })

    }
  },
  component: Login,
});

function Login() {
  return (
    <div className="h-screen w-screen ">
      <div className="flex flex-col max-w-md mx-auto items-center">
        <img src={GitlabLogo} className="h-96 w-96" />
        <Button asChild>
          <a href="/auth">
            Login with Gitlab
          </a>
        </Button>
      </div>
    </div>
  );
}
