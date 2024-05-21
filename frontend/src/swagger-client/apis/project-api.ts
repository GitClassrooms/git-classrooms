// @ts-nocheck
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

import globalAxios, { AxiosResponse, AxiosInstance, AxiosRequestConfig } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
import { HTTPError } from '../models';
import { ProjectResponse } from '../models';
/**
 * ProjectApi - axios parameter creator
 * @export
 */
export const ProjectApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * GetClassroomAssignmentProject
         * @summary GetClassroomAssignmentProject
         * @param {string} classroomId Classroom ID
         * @param {string} assignmentId Assignment ID
         * @param {string} projectId Project ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getClassroomAssignmentProject: async (classroomId: string, assignmentId: string, projectId: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'classroomId' is not null or undefined
            if (classroomId === null || classroomId === undefined) {
                throw new RequiredError('classroomId','Required parameter classroomId was null or undefined when calling getClassroomAssignmentProject.');
            }
            // verify required parameter 'assignmentId' is not null or undefined
            if (assignmentId === null || assignmentId === undefined) {
                throw new RequiredError('assignmentId','Required parameter assignmentId was null or undefined when calling getClassroomAssignmentProject.');
            }
            // verify required parameter 'projectId' is not null or undefined
            if (projectId === null || projectId === undefined) {
                throw new RequiredError('projectId','Required parameter projectId was null or undefined when calling getClassroomAssignmentProject.');
            }
            const localVarPath = `/api/v2/classrooms/{classroomId}/assignments/{assignmentId}/projects/{projectId}`
                .replace(`{${"classroomId"}}`, encodeURIComponent(String(classroomId)))
                .replace(`{${"assignmentId"}}`, encodeURIComponent(String(assignmentId)))
                .replace(`{${"projectId"}}`, encodeURIComponent(String(projectId)));
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
         * GetClassroomAssignmentProjects
         * @summary GetClassroomAssignmentProjects
         * @param {string} classroomId Classroom ID
         * @param {string} assignmentId Assignment ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getClassroomAssignmentProjects: async (classroomId: string, assignmentId: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'classroomId' is not null or undefined
            if (classroomId === null || classroomId === undefined) {
                throw new RequiredError('classroomId','Required parameter classroomId was null or undefined when calling getClassroomAssignmentProjects.');
            }
            // verify required parameter 'assignmentId' is not null or undefined
            if (assignmentId === null || assignmentId === undefined) {
                throw new RequiredError('assignmentId','Required parameter assignmentId was null or undefined when calling getClassroomAssignmentProjects.');
            }
            const localVarPath = `/api/v2/classrooms/{classroomId}/assignments/{assignmentId}/projects`
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
         * GetClassroomProject
         * @summary GetClassroomProject
         * @param {string} classroomId Classroom ID
         * @param {string} projectId Project ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getClassroomProject: async (classroomId: string, projectId: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'classroomId' is not null or undefined
            if (classroomId === null || classroomId === undefined) {
                throw new RequiredError('classroomId','Required parameter classroomId was null or undefined when calling getClassroomProject.');
            }
            // verify required parameter 'projectId' is not null or undefined
            if (projectId === null || projectId === undefined) {
                throw new RequiredError('projectId','Required parameter projectId was null or undefined when calling getClassroomProject.');
            }
            const localVarPath = `/api/v2/classrooms/{classroomId}/projects/{projectId}`
                .replace(`{${"classroomId"}}`, encodeURIComponent(String(classroomId)))
                .replace(`{${"projectId"}}`, encodeURIComponent(String(projectId)));
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
         * GetClassroomProjects
         * @summary GetClassroomProjects
         * @param {string} classroomId Classroom ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getClassroomProjects: async (classroomId: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'classroomId' is not null or undefined
            if (classroomId === null || classroomId === undefined) {
                throw new RequiredError('classroomId','Required parameter classroomId was null or undefined when calling getClassroomProjects.');
            }
            const localVarPath = `/api/v2/classrooms/{classroomId}/projects`
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
         * GetClassroomTeamProject
         * @summary GetClassroomTeamProject
         * @param {string} classroomId Classroom ID
         * @param {string} teamId Team ID
         * @param {string} projectId Project ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getClassroomTeamProject: async (classroomId: string, teamId: string, projectId: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'classroomId' is not null or undefined
            if (classroomId === null || classroomId === undefined) {
                throw new RequiredError('classroomId','Required parameter classroomId was null or undefined when calling getClassroomTeamProject.');
            }
            // verify required parameter 'teamId' is not null or undefined
            if (teamId === null || teamId === undefined) {
                throw new RequiredError('teamId','Required parameter teamId was null or undefined when calling getClassroomTeamProject.');
            }
            // verify required parameter 'projectId' is not null or undefined
            if (projectId === null || projectId === undefined) {
                throw new RequiredError('projectId','Required parameter projectId was null or undefined when calling getClassroomTeamProject.');
            }
            const localVarPath = `/api/v2/classrooms/{classroomId}/teams/{teamId}/projects/{projectId}`
                .replace(`{${"classroomId"}}`, encodeURIComponent(String(classroomId)))
                .replace(`{${"teamId"}}`, encodeURIComponent(String(teamId)))
                .replace(`{${"projectId"}}`, encodeURIComponent(String(projectId)));
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
         * GetClassroomTeamProjects
         * @summary GetClassroomTeamProjects
         * @param {string} classroomId Classroom ID
         * @param {string} teamId Team ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getClassroomTeamProjects: async (classroomId: string, teamId: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'classroomId' is not null or undefined
            if (classroomId === null || classroomId === undefined) {
                throw new RequiredError('classroomId','Required parameter classroomId was null or undefined when calling getClassroomTeamProjects.');
            }
            // verify required parameter 'teamId' is not null or undefined
            if (teamId === null || teamId === undefined) {
                throw new RequiredError('teamId','Required parameter teamId was null or undefined when calling getClassroomTeamProjects.');
            }
            const localVarPath = `/api/v2/classrooms/{classroomId}/teams/{teamId}/projects`
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
 * ProjectApi - functional programming interface
 * @export
 */
export const ProjectApiFp = function(configuration?: Configuration) {
    return {
        /**
         * GetClassroomAssignmentProject
         * @summary GetClassroomAssignmentProject
         * @param {string} classroomId Classroom ID
         * @param {string} assignmentId Assignment ID
         * @param {string} projectId Project ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getClassroomAssignmentProject(classroomId: string, assignmentId: string, projectId: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => Promise<AxiosResponse<ProjectResponse>>> {
            const localVarAxiosArgs = await ProjectApiAxiosParamCreator(configuration).getClassroomAssignmentProject(classroomId, assignmentId, projectId, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs :AxiosRequestConfig = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * GetClassroomAssignmentProjects
         * @summary GetClassroomAssignmentProjects
         * @param {string} classroomId Classroom ID
         * @param {string} assignmentId Assignment ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getClassroomAssignmentProjects(classroomId: string, assignmentId: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => Promise<AxiosResponse<Array<ProjectResponse>>>> {
            const localVarAxiosArgs = await ProjectApiAxiosParamCreator(configuration).getClassroomAssignmentProjects(classroomId, assignmentId, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs :AxiosRequestConfig = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * GetClassroomProject
         * @summary GetClassroomProject
         * @param {string} classroomId Classroom ID
         * @param {string} projectId Project ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getClassroomProject(classroomId: string, projectId: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => Promise<AxiosResponse<ProjectResponse>>> {
            const localVarAxiosArgs = await ProjectApiAxiosParamCreator(configuration).getClassroomProject(classroomId, projectId, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs :AxiosRequestConfig = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * GetClassroomProjects
         * @summary GetClassroomProjects
         * @param {string} classroomId Classroom ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getClassroomProjects(classroomId: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => Promise<AxiosResponse<Array<ProjectResponse>>>> {
            const localVarAxiosArgs = await ProjectApiAxiosParamCreator(configuration).getClassroomProjects(classroomId, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs :AxiosRequestConfig = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * GetClassroomTeamProject
         * @summary GetClassroomTeamProject
         * @param {string} classroomId Classroom ID
         * @param {string} teamId Team ID
         * @param {string} projectId Project ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getClassroomTeamProject(classroomId: string, teamId: string, projectId: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => Promise<AxiosResponse<ProjectResponse>>> {
            const localVarAxiosArgs = await ProjectApiAxiosParamCreator(configuration).getClassroomTeamProject(classroomId, teamId, projectId, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs :AxiosRequestConfig = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * GetClassroomTeamProjects
         * @summary GetClassroomTeamProjects
         * @param {string} classroomId Classroom ID
         * @param {string} teamId Team ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getClassroomTeamProjects(classroomId: string, teamId: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => Promise<AxiosResponse<Array<ProjectResponse>>>> {
            const localVarAxiosArgs = await ProjectApiAxiosParamCreator(configuration).getClassroomTeamProjects(classroomId, teamId, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs :AxiosRequestConfig = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * ProjectApi - factory interface
 * @export
 */
export const ProjectApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * GetClassroomAssignmentProject
         * @summary GetClassroomAssignmentProject
         * @param {string} classroomId Classroom ID
         * @param {string} assignmentId Assignment ID
         * @param {string} projectId Project ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getClassroomAssignmentProject(classroomId: string, assignmentId: string, projectId: string, options?: AxiosRequestConfig): Promise<AxiosResponse<ProjectResponse>> {
            return ProjectApiFp(configuration).getClassroomAssignmentProject(classroomId, assignmentId, projectId, options).then((request) => request(axios, basePath));
        },
        /**
         * GetClassroomAssignmentProjects
         * @summary GetClassroomAssignmentProjects
         * @param {string} classroomId Classroom ID
         * @param {string} assignmentId Assignment ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getClassroomAssignmentProjects(classroomId: string, assignmentId: string, options?: AxiosRequestConfig): Promise<AxiosResponse<Array<ProjectResponse>>> {
            return ProjectApiFp(configuration).getClassroomAssignmentProjects(classroomId, assignmentId, options).then((request) => request(axios, basePath));
        },
        /**
         * GetClassroomProject
         * @summary GetClassroomProject
         * @param {string} classroomId Classroom ID
         * @param {string} projectId Project ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getClassroomProject(classroomId: string, projectId: string, options?: AxiosRequestConfig): Promise<AxiosResponse<ProjectResponse>> {
            return ProjectApiFp(configuration).getClassroomProject(classroomId, projectId, options).then((request) => request(axios, basePath));
        },
        /**
         * GetClassroomProjects
         * @summary GetClassroomProjects
         * @param {string} classroomId Classroom ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getClassroomProjects(classroomId: string, options?: AxiosRequestConfig): Promise<AxiosResponse<Array<ProjectResponse>>> {
            return ProjectApiFp(configuration).getClassroomProjects(classroomId, options).then((request) => request(axios, basePath));
        },
        /**
         * GetClassroomTeamProject
         * @summary GetClassroomTeamProject
         * @param {string} classroomId Classroom ID
         * @param {string} teamId Team ID
         * @param {string} projectId Project ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getClassroomTeamProject(classroomId: string, teamId: string, projectId: string, options?: AxiosRequestConfig): Promise<AxiosResponse<ProjectResponse>> {
            return ProjectApiFp(configuration).getClassroomTeamProject(classroomId, teamId, projectId, options).then((request) => request(axios, basePath));
        },
        /**
         * GetClassroomTeamProjects
         * @summary GetClassroomTeamProjects
         * @param {string} classroomId Classroom ID
         * @param {string} teamId Team ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getClassroomTeamProjects(classroomId: string, teamId: string, options?: AxiosRequestConfig): Promise<AxiosResponse<Array<ProjectResponse>>> {
            return ProjectApiFp(configuration).getClassroomTeamProjects(classroomId, teamId, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * ProjectApi - object-oriented interface
 * @export
 * @class ProjectApi
 * @extends {BaseAPI}
 */
export class ProjectApi extends BaseAPI {
    /**
     * GetClassroomAssignmentProject
     * @summary GetClassroomAssignmentProject
     * @param {string} classroomId Classroom ID
     * @param {string} assignmentId Assignment ID
     * @param {string} projectId Project ID
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ProjectApi
     */
    public async getClassroomAssignmentProject(classroomId: string, assignmentId: string, projectId: string, options?: AxiosRequestConfig) : Promise<AxiosResponse<ProjectResponse>> {
        return ProjectApiFp(this.configuration).getClassroomAssignmentProject(classroomId, assignmentId, projectId, options).then((request) => request(this.axios, this.basePath));
    }
    /**
     * GetClassroomAssignmentProjects
     * @summary GetClassroomAssignmentProjects
     * @param {string} classroomId Classroom ID
     * @param {string} assignmentId Assignment ID
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ProjectApi
     */
    public async getClassroomAssignmentProjects(classroomId: string, assignmentId: string, options?: AxiosRequestConfig) : Promise<AxiosResponse<Array<ProjectResponse>>> {
        return ProjectApiFp(this.configuration).getClassroomAssignmentProjects(classroomId, assignmentId, options).then((request) => request(this.axios, this.basePath));
    }
    /**
     * GetClassroomProject
     * @summary GetClassroomProject
     * @param {string} classroomId Classroom ID
     * @param {string} projectId Project ID
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ProjectApi
     */
    public async getClassroomProject(classroomId: string, projectId: string, options?: AxiosRequestConfig) : Promise<AxiosResponse<ProjectResponse>> {
        return ProjectApiFp(this.configuration).getClassroomProject(classroomId, projectId, options).then((request) => request(this.axios, this.basePath));
    }
    /**
     * GetClassroomProjects
     * @summary GetClassroomProjects
     * @param {string} classroomId Classroom ID
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ProjectApi
     */
    public async getClassroomProjects(classroomId: string, options?: AxiosRequestConfig) : Promise<AxiosResponse<Array<ProjectResponse>>> {
        return ProjectApiFp(this.configuration).getClassroomProjects(classroomId, options).then((request) => request(this.axios, this.basePath));
    }
    /**
     * GetClassroomTeamProject
     * @summary GetClassroomTeamProject
     * @param {string} classroomId Classroom ID
     * @param {string} teamId Team ID
     * @param {string} projectId Project ID
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ProjectApi
     */
    public async getClassroomTeamProject(classroomId: string, teamId: string, projectId: string, options?: AxiosRequestConfig) : Promise<AxiosResponse<ProjectResponse>> {
        return ProjectApiFp(this.configuration).getClassroomTeamProject(classroomId, teamId, projectId, options).then((request) => request(this.axios, this.basePath));
    }
    /**
     * GetClassroomTeamProjects
     * @summary GetClassroomTeamProjects
     * @param {string} classroomId Classroom ID
     * @param {string} teamId Team ID
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ProjectApi
     */
    public async getClassroomTeamProjects(classroomId: string, teamId: string, options?: AxiosRequestConfig) : Promise<AxiosResponse<Array<ProjectResponse>>> {
        return ProjectApiFp(this.configuration).getClassroomTeamProjects(classroomId, teamId, options).then((request) => request(this.axios, this.basePath));
    }
}
