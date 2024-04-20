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
 * @interface GetMe
 */
export interface GetMe {

    /**
     * @type {DatabaseUserAvatar}
     * @memberof GetMe
     */
    gitlabAvatar?: DatabaseUserAvatar;

    /**
     * @type {string}
     * @memberof GetMe
     */
    gitlabEmail?: string;

    /**
     * @type {string}
     * @memberof GetMe
     */
    gitlabWeburl?: string;

    /**
     * @type {number}
     * @memberof GetMe
     */
    id?: number;

    /**
     * @type {string}
     * @memberof GetMe
     */
    name?: string;
}