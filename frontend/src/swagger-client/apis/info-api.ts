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
import { GetInfoGitlabResponse } from '../models';
import { HTTPError } from '../models';
/**
 * InfoApi - axios parameter creator
 * @export
 */
export const InfoApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * GetGitlabInfo
         * @summary GetGitlabInfo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getGitlabInfo: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/info/gitlab`;
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
 * InfoApi - functional programming interface
 * @export
 */
export const InfoApiFp = function(configuration?: Configuration) {
    return {
        /**
         * GetGitlabInfo
         * @summary GetGitlabInfo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getGitlabInfo(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => Promise<AxiosResponse<GetInfoGitlabResponse>>> {
            const localVarAxiosArgs = await InfoApiAxiosParamCreator(configuration).getGitlabInfo(options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs :AxiosRequestConfig = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * InfoApi - factory interface
 * @export
 */
export const InfoApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * GetGitlabInfo
         * @summary GetGitlabInfo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getGitlabInfo(options?: AxiosRequestConfig): Promise<AxiosResponse<GetInfoGitlabResponse>> {
            return InfoApiFp(configuration).getGitlabInfo(options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * InfoApi - object-oriented interface
 * @export
 * @class InfoApi
 * @extends {BaseAPI}
 */
export class InfoApi extends BaseAPI {
    /**
     * GetGitlabInfo
     * @summary GetGitlabInfo
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof InfoApi
     */
    public async getGitlabInfo(options?: AxiosRequestConfig) : Promise<AxiosResponse<GetInfoGitlabResponse>> {
        return InfoApiFp(this.configuration).getGitlabInfo(options).then((request) => request(this.axios, this.basePath));
    }
}
