import { Table, TableBody, TableCell, TableRow } from "@/components/ui/table.tsx";
import { Button } from "@/components/ui/button.tsx";
import { Edit, Gitlab } from "lucide-react";
import { Link } from "@tanstack/react-router";
import { formatDate } from "@/lib/utils.ts";
import { Avatar } from "@/components/avatar.tsx";
import { Separator } from "@/components/ui/separator.tsx";
import { ProjectResponse, TeamResponse, UserClassroomResponse } from "@/swagger-client";

export function ClassroomTeamModal({
  classroomId,
  team,
  projects,
}: {
  classroomId: string;
  team: TeamResponse;
  projects: ProjectResponse[];
}) {
  return (
    <div>
      {(team.members.length == 0 || team.members[0].user.gitlabUsername != team.name) && (
        <>
          <h1 className="text-2xl">{team.name}</h1>
          <Separator className="my-1" />
          <h2 className="text-xl mt-4">Members</h2>
        </>
      )}
      <ClassroomTeamMemberTable members={team.members} />
      <Separator className="my-1" />
      <h2 className="text-xl mt-4">Assignments</h2>
      <ClassroomTeamAssignmentTable classroomId={classroomId} projects={projects} />
    </div>
  );
}

function ClassroomTeamMemberTable({ members }: { members: UserClassroomResponse[] }) {
  return (
    <Table>
      <TableBody>
        {members.length > 0
          ? members.map((m) => (
              <TableRow key={m.user.id}>
                <TableCell className="p-2">
                  <ClassroomTeamMemberListElement member={m} />
                </TableCell>
              </TableRow>
            ))
          : "No member in this team"}
      </TableBody>
    </Table>
  );
}

function ClassroomTeamMemberListElement({ member }: { member: UserClassroomResponse }) {
  return (
    <div className="flex">
      <div className="pr-2">
        <Avatar
          avatarUrl={member.user.gitlabAvatar?.avatarURL}
          fallbackUrl={member.user.gitlabAvatar?.fallbackAvatarURL}
          name={member.user.name}
        />
      </div>
      <div>
        <div className="font-medium">{member.user.name}</div>
        <div className="text-sm text-muted-foreground mt-[-0.3rem]">@{member.user.gitlabUsername}</div>
      </div>
    </div>
  );
}

export function ClassroomTeamAssignmentTable({
  classroomId,
  projects,
}: {
  classroomId: string;
  projects: ProjectResponse[];
}) {
  return (
    <Table>
      <TableBody>
        {projects.map((p) => (
          <TableRow key={p.id}>
            <TableCell className="p-2">
              <div className="cursor-default flex justify-between">
                <div>
                  <div className="font-medium">{p.assignment.name}</div>
                  <div className="text-sm text-muted-foreground md:inline">
                    {p.assignmentAccepted ? "Accepted" : "Pending"}
                  </div>
                </div>
                <div className="flex items-end">
                  <div className="ml-auto">
                    <div className="font-medium text-right">Due date</div>
                    <div className="text-sm text-muted-foreground md:inline">
                      {p.assignment.dueDate ? formatDate(p.assignment.dueDate) : "No Due Date"}
                    </div>
                  </div>
                  <Button variant="ghost" size="icon" asChild>
                    <Link
                      to="/classrooms/$classroomId/assignments/$assignmentId"
                      params={{ classroomId: classroomId, assignmentId: p.assignment.id }}
                    >
                      <Edit className="h-6 w-6 text-gray-600" />
                    </Link>
                  </Button>
                  <Button variant="ghost" size="icon" asChild>
                    {p.assignmentAccepted ? (
                      <a href={p.webUrl} target="_blank" rel="noreferrer">
                        <Gitlab className="h-6 w-6 text-gray-600" />
                      </a>
                    ) : (
                      <div>
                        <Gitlab className="h-6 w-6 text-gray-400" />
                      </div>
                    )}
                  </Button>
                </div>
              </div>
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
}
