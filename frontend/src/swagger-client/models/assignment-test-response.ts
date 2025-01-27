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

import { AssignmentJunitTest } from './assignment-junit-test';
import { AssignmentTestReport } from './assignment-test-report';
import { ExamplesLanguageCIExample } from './examples-language-ciexample';
 /**
 * 
 *
 * @export
 * @interface AssignmentTestResponse
 */
export interface AssignmentTestResponse {

    /**
     * @type {boolean}
     * @memberof AssignmentTestResponse
     */
    activatible: boolean;

    /**
     * @type {ExamplesLanguageCIExample}
     * @memberof AssignmentTestResponse
     */
    example: ExamplesLanguageCIExample;

    /**
     * @type {Array<AssignmentTestReport>}
     * @memberof AssignmentTestResponse
     */
    report: Array<AssignmentTestReport>;

    /**
     * @type {Array<AssignmentJunitTest>}
     * @memberof AssignmentTestResponse
     */
    selectedTests: Array<AssignmentJunitTest>;
}
