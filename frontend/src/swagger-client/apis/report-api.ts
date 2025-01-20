// @ts-nocheck
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

import globalAxios, { AxiosResponse, AxiosInstance, AxiosRequestConfig } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
import { HTTPError } from '../models';
import { UtilsReportDataItem } from '../models';
/**
 * ReportApi - axios parameter creator
 * @export
 */
export const ReportApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * GetClassroomAssignmentReport
         * @summary GetClassroomAssignmentReport
         * @param {string} classroomId Classroom ID
         * @param {string} assignmentId Assignment ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getClassroomAssignmentReport: async (classroomId: string, assignmentId: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'classroomId' is not null or undefined
            if (classroomId === null || classroomId === undefined) {
                throw new RequiredError('classroomId','Required parameter classroomId was null or undefined when calling getClassroomAssignmentReport.');
            }
            // verify required parameter 'assignmentId' is not null or undefined
            if (assignmentId === null || assignmentId === undefined) {
                throw new RequiredError('assignmentId','Required parameter assignmentId was null or undefined when calling getClassroomAssignmentReport.');
            }
            const localVarPath = `/api/v1/classrooms/{classroomId}/assignments/{assignmentId}/grading/report`
                .replace(`{${"classroomId"}}`, encodeURIComponent(String(classroomId)))
                .replace(`{${"assignmentId"}}`, encodeURIComponent(String(assignmentId)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, 'https://example.com');
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions :AxiosRequestConfig = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            const query = new URLSearchParams(localVarUrlObj.search);
            for (const key in localVarQueryParameter) {
                query.set(key, localVarQueryParameter[key]);
            }
            for (const key in options.params) {
                query.set(key, options.params[key]);
            }
            localVarUrlObj.search = (new URLSearchParams(query)).toString();
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * GetClassroomReport
         * @summary GetClassroomReport
         * @param {string} classroomId Classroom ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getClassroomReport: async (classroomId: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'classroomId' is not null or undefined
            if (classroomId === null || classroomId === undefined) {
                throw new RequiredError('classroomId','Required parameter classroomId was null or undefined when calling getClassroomReport.');
            }
            const localVarPath = `/api/v1/classrooms/{classroomId}/grading/report`
                .replace(`{${"classroomId"}}`, encodeURIComponent(String(classroomId)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, 'https://example.com');
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions :AxiosRequestConfig = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            const query = new URLSearchParams(localVarUrlObj.search);
            for (const key in localVarQueryParameter) {
                query.set(key, localVarQueryParameter[key]);
            }
            for (const key in options.params) {
                query.set(key, options.params[key]);
            }
            localVarUrlObj.search = (new URLSearchParams(query)).toString();
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * GetClassroomTeamReport
         * @summary GetClassroomTeamReport
         * @param {string} classroomId Classroom ID
         * @param {string} teamId Team ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getClassroomTeamReport: async (classroomId: string, teamId: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'classroomId' is not null or undefined
            if (classroomId === null || classroomId === undefined) {
                throw new RequiredError('classroomId','Required parameter classroomId was null or undefined when calling getClassroomTeamReport.');
            }
            // verify required parameter 'teamId' is not null or undefined
            if (teamId === null || teamId === undefined) {
                throw new RequiredError('teamId','Required parameter teamId was null or undefined when calling getClassroomTeamReport.');
            }
            const localVarPath = `/api/v1/classrooms/{classroomId}/teams/{teamId}/grading/report`
                .replace(`{${"classroomId"}}`, encodeURIComponent(String(classroomId)))
                .replace(`{${"teamId"}}`, encodeURIComponent(String(teamId)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, 'https://example.com');
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions :AxiosRequestConfig = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            const query = new URLSearchParams(localVarUrlObj.search);
            for (const key in localVarQueryParameter) {
                query.set(key, localVarQueryParameter[key]);
            }
            for (const key in options.params) {
                query.set(key, options.params[key]);
            }
            localVarUrlObj.search = (new URLSearchParams(query)).toString();
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * ReportApi - functional programming interface
 * @export
 */
export const ReportApiFp = function(configuration?: Configuration) {
    return {
        /**
         * GetClassroomAssignmentReport
         * @summary GetClassroomAssignmentReport
         * @param {string} classroomId Classroom ID
         * @param {string} assignmentId Assignment ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getClassroomAssignmentReport(classroomId: string, assignmentId: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => Promise<AxiosResponse<Array<UtilsReportDataItem>>>> {
            const localVarAxiosArgs = await ReportApiAxiosParamCreator(configuration).getClassroomAssignmentReport(classroomId, assignmentId, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs :AxiosRequestConfig = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * GetClassroomReport
         * @summary GetClassroomReport
         * @param {string} classroomId Classroom ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getClassroomReport(classroomId: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => Promise<AxiosResponse<Array<UtilsReportDataItem>>>> {
            const localVarAxiosArgs = await ReportApiAxiosParamCreator(configuration).getClassroomReport(classroomId, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs :AxiosRequestConfig = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * GetClassroomTeamReport
         * @summary GetClassroomTeamReport
         * @param {string} classroomId Classroom ID
         * @param {string} teamId Team ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getClassroomTeamReport(classroomId: string, teamId: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => Promise<AxiosResponse<Array<UtilsReportDataItem>>>> {
            const localVarAxiosArgs = await ReportApiAxiosParamCreator(configuration).getClassroomTeamReport(classroomId, teamId, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs :AxiosRequestConfig = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * ReportApi - factory interface
 * @export
 */
export const ReportApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * GetClassroomAssignmentReport
         * @summary GetClassroomAssignmentReport
         * @param {string} classroomId Classroom ID
         * @param {string} assignmentId Assignment ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getClassroomAssignmentReport(classroomId: string, assignmentId: string, options?: AxiosRequestConfig): Promise<AxiosResponse<Array<UtilsReportDataItem>>> {
            return ReportApiFp(configuration).getClassroomAssignmentReport(classroomId, assignmentId, options).then((request) => request(axios, basePath));
        },
        /**
         * GetClassroomReport
         * @summary GetClassroomReport
         * @param {string} classroomId Classroom ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getClassroomReport(classroomId: string, options?: AxiosRequestConfig): Promise<AxiosResponse<Array<UtilsReportDataItem>>> {
            return ReportApiFp(configuration).getClassroomReport(classroomId, options).then((request) => request(axios, basePath));
        },
        /**
         * GetClassroomTeamReport
         * @summary GetClassroomTeamReport
         * @param {string} classroomId Classroom ID
         * @param {string} teamId Team ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getClassroomTeamReport(classroomId: string, teamId: string, options?: AxiosRequestConfig): Promise<AxiosResponse<Array<UtilsReportDataItem>>> {
            return ReportApiFp(configuration).getClassroomTeamReport(classroomId, teamId, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * ReportApi - object-oriented interface
 * @export
 * @class ReportApi
 * @extends {BaseAPI}
 */
export class ReportApi extends BaseAPI {
    /**
     * GetClassroomAssignmentReport
     * @summary GetClassroomAssignmentReport
     * @param {string} classroomId Classroom ID
     * @param {string} assignmentId Assignment ID
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ReportApi
     */
    public async getClassroomAssignmentReport(classroomId: string, assignmentId: string, options?: AxiosRequestConfig) : Promise<AxiosResponse<Array<UtilsReportDataItem>>> {
        return ReportApiFp(this.configuration).getClassroomAssignmentReport(classroomId, assignmentId, options).then((request) => request(this.axios, this.basePath));
    }
    /**
     * GetClassroomReport
     * @summary GetClassroomReport
     * @param {string} classroomId Classroom ID
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ReportApi
     */
    public async getClassroomReport(classroomId: string, options?: AxiosRequestConfig) : Promise<AxiosResponse<Array<UtilsReportDataItem>>> {
        return ReportApiFp(this.configuration).getClassroomReport(classroomId, options).then((request) => request(this.axios, this.basePath));
    }
    /**
     * GetClassroomTeamReport
     * @summary GetClassroomTeamReport
     * @param {string} classroomId Classroom ID
     * @param {string} teamId Team ID
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ReportApi
     */
    public async getClassroomTeamReport(classroomId: string, teamId: string, options?: AxiosRequestConfig) : Promise<AxiosResponse<Array<UtilsReportDataItem>>> {
        return ReportApiFp(this.configuration).getClassroomTeamReport(classroomId, teamId, options).then((request) => request(this.axios, this.basePath));
    }
}
