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

 /**
 * 
 *
 * @export
 * @interface ModelTestReportTestCase
 */
export interface ModelTestReportTestCase {

    /**
     * @type {string}
     * @memberof ModelTestReportTestCase
     */
    classname: string;

    /**
     * @type {number}
     * @memberof ModelTestReportTestCase
     */
    executionTime: number;

    /**
     * @type {string}
     * @memberof ModelTestReportTestCase
     */
    name: string;

    /**
     * @type {string}
     * @memberof ModelTestReportTestCase
     */
    stackTrace: string;

    /**
     * @type {string}
     * @memberof ModelTestReportTestCase
     */
    status: string;

    /**
     * @type {any}
     * @memberof ModelTestReportTestCase
     */
    systemOutput: any;
}
