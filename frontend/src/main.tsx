import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import { RouterProvider, createRouteMask, createRouter } from "@tanstack/react-router";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import * as Sentry from "@sentry/react";

// Set sentry
if (import.meta.env.MODE === "production") {
  Sentry.init({
    dsn: import.meta.env.VITE_SENTRY_DSN,
    integrations: [
      // See docs for support of different versions of variation of react router
      // https://docs.sentry.io/platforms/javascript/guides/react/configuration/integrations/react-router/
      Sentry.replayIntegration(),
    ],

    // Set tracesSampleRate to 1.0 to capture 100%
    // of transactions for performance monitoring.
    tracesSampleRate: 1.0,

    // Set `tracePropagationTargets` to control for which URLs distributed tracing should be enabled
    tracePropagationTargets: ["localhost", /^https:\/\/staging\.hs-flensburg\.dev\/api/],

    // Capture Replay for 10% of all sessions,
    // plus for 100% of sessions with an error
    replaysSessionSampleRate: 0.1,
    replaysOnErrorSampleRate: 1.0,
  });
}

const queryClient = new QueryClient();

// Import the generated route tree
import { routeTree } from "./routeTree.gen";

const classroomCreateModalToClassroomCreateMask = createRouteMask({
  routeTree,
  from: "/classrooms/create/modal",
  to: "/classrooms/create",
  params: true,
});

const joinedClassroomTeamCreateMask = createRouteMask({
  routeTree,
  from: "/classrooms/joined/$classroomId/teams/create/modal",
  to: "/classrooms/joined/$classroomId/teams/create",
  params: true,
});

const ownedClassroomTeamCreateMask = createRouteMask({
  routeTree,
  from: "/classrooms/owned/$classroomId/team/create/modal",
  to: "/classrooms/owned/$classroomId/teams/create",
  params: true,
});

const ownedClassroomTeamsCreateMask = createRouteMask({
  routeTree,
  from: "/classrooms/owned/$classroomId/teams/create/modal",
  to: "/classrooms/owned/$classroomId/teams/create",
  params: true,
});

// Create a new router instance
const router = createRouter({
  routeTree,
  context: {
    queryClient,
  },
  routeMasks: [
    classroomCreateModalToClassroomCreateMask,
    ownedClassroomTeamCreateMask,
    ownedClassroomTeamsCreateMask,
    joinedClassroomTeamCreateMask,
  ],
  defaultPreload: "intent",
  unmaskOnReload: true,
  // Since we're using React Query, we don't want loader calls to ever be stale
  // This will ensure that the loader is always called when the route is preloaded or visited
  defaultPreloadStaleTime: 0,
});

// Register the router instance for type safety
declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
    </QueryClientProvider>
  </React.StrictMode>,
);
