/**
 * Bitget Open API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 2.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class TraceSettingsRequest {
    /**
    * traderUserId
    */
    'traderUserId': string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "traderUserId",
            "baseName": "traderUserId",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return TraceSettingsRequest.attributeTypeMap;
    }
}
