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

import { Classroom } from './classroom';
import { ManualGradingRubric } from './manual-grading-rubric';
 /**
 * 
 *
 * @export
 * @interface ActiveAssignmentResponse
 */
export interface ActiveAssignmentResponse {

    /**
     * @type {Classroom}
     * @memberof ActiveAssignmentResponse
     */
    classroom: Classroom;

    /**
     * @type {string}
     * @memberof ActiveAssignmentResponse
     */
    classroomId: string;

    /**
     * @type {boolean}
     * @memberof ActiveAssignmentResponse
     */
    closed: boolean;

    /**
     * @type {string}
     * @memberof ActiveAssignmentResponse
     */
    createdAt: string;

    /**
     * @type {string}
     * @memberof ActiveAssignmentResponse
     */
    description: string;

    /**
     * @type {string}
     * @memberof ActiveAssignmentResponse
     */
    dueDate?: string;

    /**
     * @type {boolean}
     * @memberof ActiveAssignmentResponse
     */
    gradingJUnitAutoGradingActive: boolean;

    /**
     * @type {Array<ManualGradingRubric>}
     * @memberof ActiveAssignmentResponse
     */
    gradingManualRubrics: Array<ManualGradingRubric>;

    /**
     * @type {string}
     * @memberof ActiveAssignmentResponse
     */
    id: string;

    /**
     * @type {string}
     * @memberof ActiveAssignmentResponse
     */
    name: string;

    /**
     * @type {number}
     * @memberof ActiveAssignmentResponse
     */
    templateProjectId: number;

    /**
     * @type {string}
     * @memberof ActiveAssignmentResponse
     */
    updatedAt: string;
}
