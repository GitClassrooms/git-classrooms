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

import { ModelTestReportTestSuite } from './model-test-report-test-suite';
 /**
 * 
 *
 * @export
 * @interface DatabaseJUnitTestResult
 */
export interface DatabaseJUnitTestResult {

    /**
     * @type {number}
     * @memberof DatabaseJUnitTestResult
     */
    errorCount: number;

    /**
     * @type {number}
     * @memberof DatabaseJUnitTestResult
     */
    failedCount: number;

    /**
     * @type {number}
     * @memberof DatabaseJUnitTestResult
     */
    skippedCount: number;

    /**
     * @type {number}
     * @memberof DatabaseJUnitTestResult
     */
    successCount: number;

    /**
     * @type {Array<ModelTestReportTestSuite>}
     * @memberof DatabaseJUnitTestResult
     */
    testSuites: Array<ModelTestReportTestSuite>;

    /**
     * @type {number}
     * @memberof DatabaseJUnitTestResult
     */
    totalCount: number;

    /**
     * @type {number}
     * @memberof DatabaseJUnitTestResult
     */
    totalTime: number;
}
