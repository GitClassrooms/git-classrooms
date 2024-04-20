/* tslint:disable */
/* eslint-disable */
/**
 * Gitlab Classroom API
 * This is the API for our Gitlab Classroom Webapp
 *
 * OpenAPI spec version: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */

import { DatabaseUserAvatar } from './database-user-avatar';
 /**
 * 
 *
 * @export
 * @interface DefaultControllerGetMeResponse
 */
export interface DefaultControllerGetMeResponse {

    /**
     * @type {DatabaseUserAvatar}
     * @memberof DefaultControllerGetMeResponse
     */
    gitlabAvatar?: DatabaseUserAvatar;

    /**
     * @type {string}
     * @memberof DefaultControllerGetMeResponse
     */
    gitlabEmail?: string;

    /**
     * @type {string}
     * @memberof DefaultControllerGetMeResponse
     */
    gitlabWeburl?: string;

    /**
     * @type {number}
     * @memberof DefaultControllerGetMeResponse
     */
    id?: number;

    /**
     * @type {string}
     * @memberof DefaultControllerGetMeResponse
     */
    name?: string;
}