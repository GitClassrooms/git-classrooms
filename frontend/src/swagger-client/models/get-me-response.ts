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
 * @interface GetMeResponse
 */
export interface GetMeResponse {

    /**
     * @type {string}
     * @memberof GetMeResponse
     */
    avatarURL: string;

    /**
     * @type {string}
     * @memberof GetMeResponse
     */
    fallbackAvatarURL: string;

    /**
     * @type {string}
     * @memberof GetMeResponse
     */
    gitlabEmail: string;

    /**
     * @type {string}
     * @memberof GetMeResponse
     */
    gitlabUrl: string;

    /**
     * @type {string}
     * @memberof GetMeResponse
     */
    gitlabUsername: string;

    /**
     * @type {number}
     * @memberof GetMeResponse
     */
    id: number;

    /**
     * @type {string}
     * @memberof GetMeResponse
     */
    name: string;
}
