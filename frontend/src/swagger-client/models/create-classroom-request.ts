/* tslint:disable */
/* eslint-disable */
/**
 * GitLab Classrooms – Backend API
 * This is the API for our Gitlab Classroom Webapp
 *
 * OpenAPI spec version: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */

 /**
 * 
 *
 * @export
 * @interface CreateClassroomRequest
 */
export interface CreateClassroomRequest {

    /**
     * @type {boolean}
     * @memberof CreateClassroomRequest
     */
    createTeams?: boolean;

    /**
     * @type {string}
     * @memberof CreateClassroomRequest
     */
    description?: string;

    /**
     * @type {number}
     * @memberof CreateClassroomRequest
     */
    maxTeamSize?: number;

    /**
     * @type {number}
     * @memberof CreateClassroomRequest
     */
    maxTeams?: number;

    /**
     * @type {string}
     * @memberof CreateClassroomRequest
     */
    name?: string;

    /**
     * @type {boolean}
     * @memberof CreateClassroomRequest
     */
    studentsViewAllProjects?: boolean;
}
