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
import { AuthGetCsrfResponse } from '../models';
import { GetMeResponse } from '../models';
import { HTTPError } from '../models';
/**
 * AuthApi - axios parameter creator
 * @export
 */
export const AuthApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * Get your csrf-Token
         * @summary Show your csrf-Token
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getCsrf: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/auth/csrf`;
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
         * Get your user account
         * @summary Show your user account
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getMe: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/me`;
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
 * AuthApi - functional programming interface
 * @export
 */
export const AuthApiFp = function(configuration?: Configuration) {
    return {
        /**
         * Get your csrf-Token
         * @summary Show your csrf-Token
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getCsrf(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => Promise<AxiosResponse<AuthGetCsrfResponse>>> {
            const localVarAxiosArgs = await AuthApiAxiosParamCreator(configuration).getCsrf(options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs :AxiosRequestConfig = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * Get your user account
         * @summary Show your user account
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getMe(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => Promise<AxiosResponse<GetMeResponse>>> {
            const localVarAxiosArgs = await AuthApiAxiosParamCreator(configuration).getMe(options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs :AxiosRequestConfig = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * AuthApi - factory interface
 * @export
 */
export const AuthApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * Get your csrf-Token
         * @summary Show your csrf-Token
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getCsrf(options?: AxiosRequestConfig): Promise<AxiosResponse<AuthGetCsrfResponse>> {
            return AuthApiFp(configuration).getCsrf(options).then((request) => request(axios, basePath));
        },
        /**
         * Get your user account
         * @summary Show your user account
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getMe(options?: AxiosRequestConfig): Promise<AxiosResponse<GetMeResponse>> {
            return AuthApiFp(configuration).getMe(options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * AuthApi - object-oriented interface
 * @export
 * @class AuthApi
 * @extends {BaseAPI}
 */
export class AuthApi extends BaseAPI {
    /**
     * Get your csrf-Token
     * @summary Show your csrf-Token
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AuthApi
     */
    public async getCsrf(options?: AxiosRequestConfig) : Promise<AxiosResponse<AuthGetCsrfResponse>> {
        return AuthApiFp(this.configuration).getCsrf(options).then((request) => request(this.axios, this.basePath));
    }
    /**
     * Get your user account
     * @summary Show your user account
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AuthApi
     */
    public async getMe(options?: AxiosRequestConfig) : Promise<AxiosResponse<GetMeResponse>> {
        return AuthApiFp(this.configuration).getMe(options).then((request) => request(this.axios, this.basePath));
    }
}
