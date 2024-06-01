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
 * @interface ManualGradingRubric
 */
export interface ManualGradingRubric {

    /**
     * @type {string}
     * @memberof ManualGradingRubric
     */
    assignmentId: string;

    /**
     * @type {string}
     * @memberof ManualGradingRubric
     */
    description: string;

    /**
     * @type {string}
     * @memberof ManualGradingRubric
     */
    id: string;

    /**
     * @type {number}
     * @memberof ManualGradingRubric
     */
    maxScore: number;

    /**
     * @type {string}
     * @memberof ManualGradingRubric
     */
    name: string;
}
