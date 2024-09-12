/* tslint:disable */
/* eslint-disable */
/**
 * GitClassrooms – Backend API
 * This is the API for our Gitlab Classroom Webapp
 *
 * OpenAPI spec version: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */

import { UserAvatar } from './user-avatar';
 /**
 * 
 *
 * @export
 * @interface GetMeResponse
 */
export interface GetMeResponse {

    /**
     * @type {UserAvatar}
     * @memberof GetMeResponse
     */
    gitlabAvatar: UserAvatar;

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
