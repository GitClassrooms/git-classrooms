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

import { Assignment } from './assignment';
import { Team } from './team';
 /**
 * 
 *
 * @export
 * @interface ProjectResponse
 */
export interface ProjectResponse {

    /**
     * @type {Assignment}
     * @memberof ProjectResponse
     */
    assignment: Assignment;

    /**
     * @type {boolean}
     * @memberof ProjectResponse
     */
    assignmentAccepted: boolean;

    /**
     * @type {string}
     * @memberof ProjectResponse
     */
    createdAt: string;

    /**
     * @type {string}
     * @memberof ProjectResponse
     */
    id: string;

    /**
     * @type {number}
     * @memberof ProjectResponse
     */
    projectId: number;

    /**
     * @type {Team}
     * @memberof ProjectResponse
     */
    team: Team;

    /**
     * @type {string}
     * @memberof ProjectResponse
     */
    teamId: string;

    /**
     * @type {string}
     * @memberof ProjectResponse
     */
    updatedAt: string;

    /**
     * @type {string}
     * @memberof ProjectResponse
     */
    webUrl: string;
}
