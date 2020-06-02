package com.bitget.contract.openapi.common.domain;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 *
 * @author bitget-sdk-team
 * @date 2019-01-15
 */
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class ClientParameter {

    /**
     * 用户 api key，必填
     */
    private String apiKey;
    /**
     * 用户密钥，必填
     */
    private String secretKey;
    /**
     * 用户 passphrase，必填
     */
    private String passphrase;
    /**
     * 服务 url，非必填 默认 https://capi.bitget.com/api/v1/
     */
    private String baseUrl;
    /**
     * 链接超时时间，非必填 默认 30s
     */
    private Long timeout;
    /**
     * 语言环境
     */
    private String locale;

}
