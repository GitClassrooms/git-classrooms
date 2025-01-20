/* tslint:disable */
/* eslint-disable */
/**
 * GitClassrooms – Backend API
 * This is the API for our GitClassrooms Webapp.
 *
 * OpenAPI spec version: develop
 * Contact: info@git-classrooms.dev
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */

import { Assignment } from './assignment';
import { DatabaseJUnitTestResult } from './database-junit-test-result';
import { DatabaseStatus } from './database-status';
import { ManualGradingResult } from './manual-grading-result';
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
     * @type {string}
     * @memberof ProjectResponse
     */
    createdAt: string;

    /**
     * @type {DatabaseJUnitTestResult}
     * @memberof ProjectResponse
     */
    gradingJUnitTestResult?: DatabaseJUnitTestResult;

    /**
     * @type {Array<ManualGradingResult>}
     * @memberof ProjectResponse
     */
    gradingManualResults: Array<ManualGradingResult>;

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
     * @type {DatabaseStatus}
     * @memberof ProjectResponse
     */
    projectStatus: DatabaseStatus;

    /**
     * @type {string}
     * @memberof ProjectResponse
     */
    reportWebUrl: string;

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
