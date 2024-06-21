import { classroomQueryOptions } from "@/api/classroom";
import { teamsQueryOptions, useJoinTeam } from "@/api/team";
import { CreateTeamForm } from "@/components/createTeamForm";
import { Header } from "@/components/header";
import { Button } from "@/components/ui/button";
import { Dialog, DialogContent, DialogHeader, DialogTrigger } from "@/components/ui/dialog";
import { TableCaption, TableHeader, TableRow, TableHead, TableBody, TableCell, Table } from "@/components/ui/table";
import { TeamResponse } from "@/swagger-client";
import { useSuspenseQuery } from "@tanstack/react-query";
import { createFileRoute, redirect, useNavigate } from "@tanstack/react-router";
import { Code } from "lucide-react";
import { Role } from "@/types/classroom.ts";

export const Route = createFileRoute("/_auth/classrooms/$classroomId/teams/join/")({
  loader: async ({ context: { queryClient }, params }) => {
    const userClassroom = await queryClient.fetchQuery(classroomQueryOptions(params.classroomId));
    const teams = await queryClient.ensureQueryData(teamsQueryOptions(params.classroomId));

    if (userClassroom.team || userClassroom.role !== Role.Student) {
      throw redirect({
        to: "/classrooms/$classroomId",
        params,
        replace: true,
      });
    }

    return { teams, userClassroom };
  },
  component: JoinTeam,
});

function JoinTeam() {
  const navigate = useNavigate();
  const { classroomId } = Route.useParams();
  const { data: joinedClassroom } = useSuspenseQuery(classroomQueryOptions(classroomId));
  const { data: teams } = useSuspenseQuery(teamsQueryOptions(classroomId));

  const { mutateAsync, isPending } = useJoinTeam(classroomId);

  const joinTeam = async (teamId: string) => {
    await mutateAsync(teamId);
    await navigate({
      to: "/classrooms/$classroomId",
      params: { classroomId },
    });
  };

  return (
    <div className="p-2">
      <Header title={`Join a team of ${joinedClassroom.classroom.name}`}>
        {joinedClassroom.classroom.createTeams &&
          (joinedClassroom.classroom.maxTeams === 0 || teams.length < joinedClassroom.classroom.maxTeams) && (
            <Dialog>
              <DialogTrigger asChild>
                <Button variant="default">Create</Button>
              </DialogTrigger>
              <DialogHeader>Create a new Team</DialogHeader>
              <DialogContent>
                <CreateTeamForm classroomId={classroomId} />
              </DialogContent>
            </Dialog>
          )}
      </Header>
      <TeamsTable
        teams={teams}
        isPending={isPending}
        joinTeam={joinTeam}
        maxTeamSize={joinedClassroom.classroom.maxTeamSize}
      />
    </div>
  );
}

interface TeamsTableProps {
  teams: TeamResponse[];
  isPending: boolean;
  joinTeam: (teamId: string) => Promise<void>;
  maxTeamSize: number;
}

function TeamsTable({ teams, isPending, joinTeam, maxTeamSize }: TeamsTableProps) {
  return (
    <Table>
      <TableCaption>Teams to join</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead>Name</TableHead>
          <TableHead>Member Count</TableHead>
          <TableHead>Gitlab-Link</TableHead>
          <TableHead className="text-right">Actions</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {teams.map((t) => (
          <TableRow key={t.id}>
            <TableCell>{t.name}</TableCell>
            <TableCell>{t.members.length}</TableCell>
            <TableCell>
              <a href={t.webUrl} target="_blank" rel="noreferrer">
                <Code />
              </a>
            </TableCell>
            <TableCell className="text-right">
              <Button
                disabled={isPending || t.members.length >= maxTeamSize}
                onClick={() => joinTeam(t.id!)}
                variant="outline"
              >
                Join Team
              </Button>
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
}