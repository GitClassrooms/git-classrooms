/* prettier-ignore-start */

/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file is auto-generated by TanStack Router

import { createFileRoute } from "@tanstack/react-router"

// Import Routes

import { Route as rootRoute } from "./routes/__root"
import { Route as LoginImport } from "./routes/login"
import { Route as AuthImport } from "./routes/_auth"
import { Route as IndexImport } from "./routes/index"
import { Route as AuthClassroomsCreateImport } from "./routes/_auth/classrooms/create"
import { Route as AuthClassroomsIndexImport } from "./routes/_auth/classrooms/_index"
import { Route as AuthClassroomsOwnedIndexImport } from "./routes/_auth/classrooms/owned/index"
import { Route as AuthClassroomsJoinedIndexImport } from "./routes/_auth/classrooms/joined/index"
import { Route as AuthClassroomsIndexIndexImport } from "./routes/_auth/classrooms/_index/index"
import { Route as AuthClassroomsJoinedClassroomIdIndexImport } from "./routes/_auth/classrooms/joined/$classroomId/index"
import { Route as AuthClassroomsOwnedClassroomIdInviteImport } from "./routes/_auth/classrooms/owned/$classroomId/invite"
import { Route as AuthClassroomsOwnedClassroomIdIndexImport } from "./routes/_auth/classrooms/owned/$classroomId/_index"
import { Route as AuthClassroomsIndexCreateModalImport } from "./routes/_auth/classrooms/_index/create.modal"
import { Route as AuthClassroomsJoinedClassroomIdTeamsRouteImport } from "./routes/_auth/classrooms/joined/$classroomId/teams/route"
import { Route as AuthClassroomsOwnedClassroomIdIndexIndexImport } from "./routes/_auth/classrooms/owned/$classroomId/_index/index"
import { Route as AuthClassroomsOwnedClassroomIdTeamsCreateImport } from "./routes/_auth/classrooms/owned/$classroomId/teams/create"
import { Route as AuthClassroomsOwnedClassroomIdTeamsIndexImport } from "./routes/_auth/classrooms/owned/$classroomId/teams/_index"
import { Route as AuthClassroomsOwnedClassroomIdAssignmentsCreateImport } from "./routes/_auth/classrooms/owned/$classroomId/assignments/create"
import { Route as AuthClassroomsJoinedClassroomIdTeamsCreateImport } from "./routes/_auth/classrooms/joined/$classroomId/teams/create"
import { Route as AuthClassroomsJoinedClassroomIdTeamsIndexImport } from "./routes/_auth/classrooms/joined/$classroomId/teams/_index"
import { Route as AuthClassroomsJoinedClassroomIdInvitationsInvitationIdImport } from "./routes/_auth/classrooms/joined/$classroomId/invitations/$invitationId"
import { Route as AuthClassroomsJoinedClassroomIdTeamsTeamIdRouteImport } from "./routes/_auth/classrooms/joined/$classroomId/teams/$teamId/route"
import { Route as AuthClassroomsOwnedClassroomIdTeamsIndexIndexImport } from "./routes/_auth/classrooms/owned/$classroomId/teams/_index/index"
import { Route as AuthClassroomsOwnedClassroomIdAssignmentsAssignmentIdIndexImport } from "./routes/_auth/classrooms/owned/$classroomId/assignments/$assignmentId/index"
import { Route as AuthClassroomsJoinedClassroomIdTeamsJoinIndexImport } from "./routes/_auth/classrooms/joined/$classroomId/teams/join.index"
import { Route as AuthClassroomsJoinedClassroomIdTeamsIndexIndexImport } from "./routes/_auth/classrooms/joined/$classroomId/teams/_index/index"
import { Route as AuthClassroomsJoinedClassroomIdAssignmentsAssignmentIdAcceptImport } from "./routes/_auth/classrooms/joined/$classroomId/assignments/$assignmentId/accept"
import { Route as AuthClassroomsOwnedClassroomIdTeamsIndexCreateModalImport } from "./routes/_auth/classrooms/owned/$classroomId/teams/_index/create.modal"
import { Route as AuthClassroomsOwnedClassroomIdIndexTeamsTeamIdModalImport } from "./routes/_auth/classrooms/owned/$classroomId/_index/teams.$teamId.modal"
import { Route as AuthClassroomsOwnedClassroomIdIndexTeamCreateModalImport } from "./routes/_auth/classrooms/owned/$classroomId/_index/team.create.modal"
import { Route as AuthClassroomsJoinedClassroomIdTeamsIndexCreateModalImport } from "./routes/_auth/classrooms/joined/$classroomId/teams/_index/create.modal"

// Create Virtual Routes

const AuthClassroomsImport = createFileRoute("/_auth/classrooms")()
const AuthClassroomsOwnedClassroomIdImport = createFileRoute(
  "/_auth/classrooms/owned/$classroomId",
)()
const AuthClassroomsOwnedClassroomIdTeamsImport = createFileRoute(
  "/_auth/classrooms/owned/$classroomId/teams",
)()

// Create/Update Routes

const LoginRoute = LoginImport.update({
  path: "/login",
  getParentRoute: () => rootRoute,
} as any)

const AuthRoute = AuthImport.update({
  id: "/_auth",
  getParentRoute: () => rootRoute,
} as any)

const IndexRoute = IndexImport.update({
  path: "/",
  getParentRoute: () => rootRoute,
} as any)

const AuthClassroomsRoute = AuthClassroomsImport.update({
  path: "/classrooms",
  getParentRoute: () => AuthRoute,
} as any)

const AuthClassroomsCreateRoute = AuthClassroomsCreateImport.update({
  path: "/create",
  getParentRoute: () => AuthClassroomsRoute,
} as any)

const AuthClassroomsIndexRoute = AuthClassroomsIndexImport.update({
  id: "/_index",
  getParentRoute: () => AuthClassroomsRoute,
} as any)

const AuthClassroomsOwnedClassroomIdRoute =
  AuthClassroomsOwnedClassroomIdImport.update({
    path: "/owned/$classroomId",
    getParentRoute: () => AuthClassroomsRoute,
  } as any)

const AuthClassroomsOwnedIndexRoute = AuthClassroomsOwnedIndexImport.update({
  path: "/owned/",
  getParentRoute: () => AuthClassroomsRoute,
} as any)

const AuthClassroomsJoinedIndexRoute = AuthClassroomsJoinedIndexImport.update({
  path: "/joined/",
  getParentRoute: () => AuthClassroomsRoute,
} as any)

const AuthClassroomsIndexIndexRoute = AuthClassroomsIndexIndexImport.update({
  path: "/",
  getParentRoute: () => AuthClassroomsIndexRoute,
} as any)

const AuthClassroomsOwnedClassroomIdTeamsRoute =
  AuthClassroomsOwnedClassroomIdTeamsImport.update({
    path: "/teams",
    getParentRoute: () => AuthClassroomsOwnedClassroomIdRoute,
  } as any)

const AuthClassroomsJoinedClassroomIdIndexRoute =
  AuthClassroomsJoinedClassroomIdIndexImport.update({
    path: "/joined/$classroomId/",
    getParentRoute: () => AuthClassroomsRoute,
  } as any)

const AuthClassroomsOwnedClassroomIdInviteRoute =
  AuthClassroomsOwnedClassroomIdInviteImport.update({
    path: "/invite",
    getParentRoute: () => AuthClassroomsOwnedClassroomIdRoute,
  } as any)

const AuthClassroomsOwnedClassroomIdIndexRoute =
  AuthClassroomsOwnedClassroomIdIndexImport.update({
    id: "/_index",
    getParentRoute: () => AuthClassroomsOwnedClassroomIdRoute,
  } as any)

const AuthClassroomsIndexCreateModalRoute =
  AuthClassroomsIndexCreateModalImport.update({
    path: "/create/modal",
    getParentRoute: () => AuthClassroomsIndexRoute,
  } as any)

const AuthClassroomsJoinedClassroomIdTeamsRouteRoute =
  AuthClassroomsJoinedClassroomIdTeamsRouteImport.update({
    path: "/joined/$classroomId/teams",
    getParentRoute: () => AuthClassroomsRoute,
  } as any)

const AuthClassroomsOwnedClassroomIdIndexIndexRoute =
  AuthClassroomsOwnedClassroomIdIndexIndexImport.update({
    path: "/",
    getParentRoute: () => AuthClassroomsOwnedClassroomIdIndexRoute,
  } as any)

const AuthClassroomsOwnedClassroomIdTeamsCreateRoute =
  AuthClassroomsOwnedClassroomIdTeamsCreateImport.update({
    path: "/create",
    getParentRoute: () => AuthClassroomsOwnedClassroomIdTeamsRoute,
  } as any)

const AuthClassroomsOwnedClassroomIdTeamsIndexRoute =
  AuthClassroomsOwnedClassroomIdTeamsIndexImport.update({
    id: "/_index",
    getParentRoute: () => AuthClassroomsOwnedClassroomIdTeamsRoute,
  } as any)

const AuthClassroomsOwnedClassroomIdAssignmentsCreateRoute =
  AuthClassroomsOwnedClassroomIdAssignmentsCreateImport.update({
    path: "/assignments/create",
    getParentRoute: () => AuthClassroomsOwnedClassroomIdRoute,
  } as any)

const AuthClassroomsJoinedClassroomIdTeamsCreateRoute =
  AuthClassroomsJoinedClassroomIdTeamsCreateImport.update({
    path: "/create",
    getParentRoute: () => AuthClassroomsJoinedClassroomIdTeamsRouteRoute,
  } as any)

const AuthClassroomsJoinedClassroomIdTeamsIndexRoute =
  AuthClassroomsJoinedClassroomIdTeamsIndexImport.update({
    id: "/_index",
    getParentRoute: () => AuthClassroomsJoinedClassroomIdTeamsRouteRoute,
  } as any)

const AuthClassroomsJoinedClassroomIdInvitationsInvitationIdRoute =
  AuthClassroomsJoinedClassroomIdInvitationsInvitationIdImport.update({
    path: "/joined/$classroomId/invitations/$invitationId",
    getParentRoute: () => AuthClassroomsRoute,
  } as any)

const AuthClassroomsJoinedClassroomIdTeamsTeamIdRouteRoute =
  AuthClassroomsJoinedClassroomIdTeamsTeamIdRouteImport.update({
    path: "/$teamId",
    getParentRoute: () => AuthClassroomsJoinedClassroomIdTeamsRouteRoute,
  } as any)

const AuthClassroomsOwnedClassroomIdTeamsIndexIndexRoute =
  AuthClassroomsOwnedClassroomIdTeamsIndexIndexImport.update({
    path: "/",
    getParentRoute: () => AuthClassroomsOwnedClassroomIdTeamsIndexRoute,
  } as any)

const AuthClassroomsOwnedClassroomIdAssignmentsAssignmentIdIndexRoute =
  AuthClassroomsOwnedClassroomIdAssignmentsAssignmentIdIndexImport.update({
    path: "/assignments/$assignmentId/",
    getParentRoute: () => AuthClassroomsOwnedClassroomIdRoute,
  } as any)

const AuthClassroomsJoinedClassroomIdTeamsJoinIndexRoute =
  AuthClassroomsJoinedClassroomIdTeamsJoinIndexImport.update({
    path: "/join/",
    getParentRoute: () => AuthClassroomsJoinedClassroomIdTeamsRouteRoute,
  } as any)

const AuthClassroomsJoinedClassroomIdTeamsIndexIndexRoute =
  AuthClassroomsJoinedClassroomIdTeamsIndexIndexImport.update({
    path: "/",
    getParentRoute: () => AuthClassroomsJoinedClassroomIdTeamsIndexRoute,
  } as any)

const AuthClassroomsJoinedClassroomIdAssignmentsAssignmentIdAcceptRoute =
  AuthClassroomsJoinedClassroomIdAssignmentsAssignmentIdAcceptImport.update({
    path: "/joined/$classroomId/assignments/$assignmentId/accept",
    getParentRoute: () => AuthClassroomsRoute,
  } as any)

const AuthClassroomsOwnedClassroomIdTeamsIndexCreateModalRoute =
  AuthClassroomsOwnedClassroomIdTeamsIndexCreateModalImport.update({
    path: "/create/modal",
    getParentRoute: () => AuthClassroomsOwnedClassroomIdTeamsIndexRoute,
  } as any)

const AuthClassroomsOwnedClassroomIdIndexTeamsTeamIdModalRoute =
  AuthClassroomsOwnedClassroomIdIndexTeamsTeamIdModalImport.update({
    path: "/teams/$teamId/modal",
    getParentRoute: () => AuthClassroomsOwnedClassroomIdIndexRoute,
  } as any)

const AuthClassroomsOwnedClassroomIdIndexTeamCreateModalRoute =
  AuthClassroomsOwnedClassroomIdIndexTeamCreateModalImport.update({
    path: "/team/create/modal",
    getParentRoute: () => AuthClassroomsOwnedClassroomIdIndexRoute,
  } as any)

const AuthClassroomsJoinedClassroomIdTeamsIndexCreateModalRoute =
  AuthClassroomsJoinedClassroomIdTeamsIndexCreateModalImport.update({
    path: "/create/modal",
    getParentRoute: () => AuthClassroomsJoinedClassroomIdTeamsIndexRoute,
  } as any)

// Populate the FileRoutesByPath interface

declare module "@tanstack/react-router" {
  interface FileRoutesByPath {
    "/": {
      preLoaderRoute: typeof IndexImport
      parentRoute: typeof rootRoute
    }
    "/_auth": {
      preLoaderRoute: typeof AuthImport
      parentRoute: typeof rootRoute
    }
    "/login": {
      preLoaderRoute: typeof LoginImport
      parentRoute: typeof rootRoute
    }
    "/_auth/classrooms": {
      preLoaderRoute: typeof AuthClassroomsImport
      parentRoute: typeof AuthImport
    }
    "/_auth/classrooms/_index": {
      preLoaderRoute: typeof AuthClassroomsIndexImport
      parentRoute: typeof AuthClassroomsRoute
    }
    "/_auth/classrooms/create": {
      preLoaderRoute: typeof AuthClassroomsCreateImport
      parentRoute: typeof AuthClassroomsImport
    }
    "/_auth/classrooms/_index/": {
      preLoaderRoute: typeof AuthClassroomsIndexIndexImport
      parentRoute: typeof AuthClassroomsIndexImport
    }
    "/_auth/classrooms/joined/": {
      preLoaderRoute: typeof AuthClassroomsJoinedIndexImport
      parentRoute: typeof AuthClassroomsImport
    }
    "/_auth/classrooms/owned/": {
      preLoaderRoute: typeof AuthClassroomsOwnedIndexImport
      parentRoute: typeof AuthClassroomsImport
    }
    "/_auth/classrooms/joined/$classroomId/teams": {
      preLoaderRoute: typeof AuthClassroomsJoinedClassroomIdTeamsRouteImport
      parentRoute: typeof AuthClassroomsImport
    }
    "/_auth/classrooms/_index/create/modal": {
      preLoaderRoute: typeof AuthClassroomsIndexCreateModalImport
      parentRoute: typeof AuthClassroomsIndexImport
    }
    "/_auth/classrooms/owned/$classroomId": {
      preLoaderRoute: typeof AuthClassroomsOwnedClassroomIdImport
      parentRoute: typeof AuthClassroomsImport
    }
    "/_auth/classrooms/owned/$classroomId/_index": {
      preLoaderRoute: typeof AuthClassroomsOwnedClassroomIdIndexImport
      parentRoute: typeof AuthClassroomsOwnedClassroomIdRoute
    }
    "/_auth/classrooms/owned/$classroomId/invite": {
      preLoaderRoute: typeof AuthClassroomsOwnedClassroomIdInviteImport
      parentRoute: typeof AuthClassroomsOwnedClassroomIdImport
    }
    "/_auth/classrooms/joined/$classroomId/": {
      preLoaderRoute: typeof AuthClassroomsJoinedClassroomIdIndexImport
      parentRoute: typeof AuthClassroomsImport
    }
    "/_auth/classrooms/joined/$classroomId/teams/$teamId": {
      preLoaderRoute: typeof AuthClassroomsJoinedClassroomIdTeamsTeamIdRouteImport
      parentRoute: typeof AuthClassroomsJoinedClassroomIdTeamsRouteImport
    }
    "/_auth/classrooms/joined/$classroomId/invitations/$invitationId": {
      preLoaderRoute: typeof AuthClassroomsJoinedClassroomIdInvitationsInvitationIdImport
      parentRoute: typeof AuthClassroomsImport
    }
    "/_auth/classrooms/joined/$classroomId/teams/_index": {
      preLoaderRoute: typeof AuthClassroomsJoinedClassroomIdTeamsIndexImport
      parentRoute: typeof AuthClassroomsJoinedClassroomIdTeamsRouteImport
    }
    "/_auth/classrooms/joined/$classroomId/teams/create": {
      preLoaderRoute: typeof AuthClassroomsJoinedClassroomIdTeamsCreateImport
      parentRoute: typeof AuthClassroomsJoinedClassroomIdTeamsRouteImport
    }
    "/_auth/classrooms/owned/$classroomId/assignments/create": {
      preLoaderRoute: typeof AuthClassroomsOwnedClassroomIdAssignmentsCreateImport
      parentRoute: typeof AuthClassroomsOwnedClassroomIdImport
    }
    "/_auth/classrooms/owned/$classroomId/teams": {
      preLoaderRoute: typeof AuthClassroomsOwnedClassroomIdTeamsImport
      parentRoute: typeof AuthClassroomsOwnedClassroomIdImport
    }
    "/_auth/classrooms/owned/$classroomId/teams/_index": {
      preLoaderRoute: typeof AuthClassroomsOwnedClassroomIdTeamsIndexImport
      parentRoute: typeof AuthClassroomsOwnedClassroomIdTeamsRoute
    }
    "/_auth/classrooms/owned/$classroomId/teams/create": {
      preLoaderRoute: typeof AuthClassroomsOwnedClassroomIdTeamsCreateImport
      parentRoute: typeof AuthClassroomsOwnedClassroomIdTeamsImport
    }
    "/_auth/classrooms/owned/$classroomId/_index/": {
      preLoaderRoute: typeof AuthClassroomsOwnedClassroomIdIndexIndexImport
      parentRoute: typeof AuthClassroomsOwnedClassroomIdIndexImport
    }
    "/_auth/classrooms/joined/$classroomId/assignments/$assignmentId/accept": {
      preLoaderRoute: typeof AuthClassroomsJoinedClassroomIdAssignmentsAssignmentIdAcceptImport
      parentRoute: typeof AuthClassroomsImport
    }
    "/_auth/classrooms/joined/$classroomId/teams/_index/": {
      preLoaderRoute: typeof AuthClassroomsJoinedClassroomIdTeamsIndexIndexImport
      parentRoute: typeof AuthClassroomsJoinedClassroomIdTeamsIndexImport
    }
    "/_auth/classrooms/joined/$classroomId/teams/join/": {
      preLoaderRoute: typeof AuthClassroomsJoinedClassroomIdTeamsJoinIndexImport
      parentRoute: typeof AuthClassroomsJoinedClassroomIdTeamsRouteImport
    }
    "/_auth/classrooms/owned/$classroomId/assignments/$assignmentId/": {
      preLoaderRoute: typeof AuthClassroomsOwnedClassroomIdAssignmentsAssignmentIdIndexImport
      parentRoute: typeof AuthClassroomsOwnedClassroomIdImport
    }
    "/_auth/classrooms/owned/$classroomId/teams/_index/": {
      preLoaderRoute: typeof AuthClassroomsOwnedClassroomIdTeamsIndexIndexImport
      parentRoute: typeof AuthClassroomsOwnedClassroomIdTeamsIndexImport
    }
    "/_auth/classrooms/joined/$classroomId/teams/_index/create/modal": {
      preLoaderRoute: typeof AuthClassroomsJoinedClassroomIdTeamsIndexCreateModalImport
      parentRoute: typeof AuthClassroomsJoinedClassroomIdTeamsIndexImport
    }
    "/_auth/classrooms/owned/$classroomId/_index/team/create/modal": {
      preLoaderRoute: typeof AuthClassroomsOwnedClassroomIdIndexTeamCreateModalImport
      parentRoute: typeof AuthClassroomsOwnedClassroomIdIndexImport
    }
    "/_auth/classrooms/owned/$classroomId/_index/teams/$teamId/modal": {
      preLoaderRoute: typeof AuthClassroomsOwnedClassroomIdIndexTeamsTeamIdModalImport
      parentRoute: typeof AuthClassroomsOwnedClassroomIdIndexImport
    }
    "/_auth/classrooms/owned/$classroomId/teams/_index/create/modal": {
      preLoaderRoute: typeof AuthClassroomsOwnedClassroomIdTeamsIndexCreateModalImport
      parentRoute: typeof AuthClassroomsOwnedClassroomIdTeamsIndexImport
    }
  }
}

// Create and export the route tree

export const routeTree = rootRoute.addChildren([
  IndexRoute,
  AuthRoute.addChildren([
    AuthClassroomsRoute.addChildren([
      AuthClassroomsIndexRoute.addChildren([
        AuthClassroomsIndexIndexRoute,
        AuthClassroomsIndexCreateModalRoute,
      ]),
      AuthClassroomsCreateRoute,
      AuthClassroomsJoinedIndexRoute,
      AuthClassroomsOwnedIndexRoute,
      AuthClassroomsJoinedClassroomIdTeamsRouteRoute.addChildren([
        AuthClassroomsJoinedClassroomIdTeamsTeamIdRouteRoute,
        AuthClassroomsJoinedClassroomIdTeamsIndexRoute.addChildren([
          AuthClassroomsJoinedClassroomIdTeamsIndexIndexRoute,
          AuthClassroomsJoinedClassroomIdTeamsIndexCreateModalRoute,
        ]),
        AuthClassroomsJoinedClassroomIdTeamsCreateRoute,
        AuthClassroomsJoinedClassroomIdTeamsJoinIndexRoute,
      ]),
      AuthClassroomsOwnedClassroomIdRoute.addChildren([
        AuthClassroomsOwnedClassroomIdIndexRoute.addChildren([
          AuthClassroomsOwnedClassroomIdIndexIndexRoute,
          AuthClassroomsOwnedClassroomIdIndexTeamCreateModalRoute,
          AuthClassroomsOwnedClassroomIdIndexTeamsTeamIdModalRoute,
        ]),
        AuthClassroomsOwnedClassroomIdInviteRoute,
        AuthClassroomsOwnedClassroomIdAssignmentsCreateRoute,
        AuthClassroomsOwnedClassroomIdTeamsRoute.addChildren([
          AuthClassroomsOwnedClassroomIdTeamsIndexRoute.addChildren([
            AuthClassroomsOwnedClassroomIdTeamsIndexIndexRoute,
            AuthClassroomsOwnedClassroomIdTeamsIndexCreateModalRoute,
          ]),
          AuthClassroomsOwnedClassroomIdTeamsCreateRoute,
        ]),
        AuthClassroomsOwnedClassroomIdAssignmentsAssignmentIdIndexRoute,
      ]),
      AuthClassroomsJoinedClassroomIdIndexRoute,
      AuthClassroomsJoinedClassroomIdInvitationsInvitationIdRoute,
      AuthClassroomsJoinedClassroomIdAssignmentsAssignmentIdAcceptRoute,
    ]),
  ]),
  LoginRoute,
])

/* prettier-ignore-end */
