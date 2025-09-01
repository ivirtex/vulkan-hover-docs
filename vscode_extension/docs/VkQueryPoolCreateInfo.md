# VkQueryPoolCreateInfo(3) Manual Page

## Name

VkQueryPoolCreateInfo - Structure specifying parameters of a newly created query pool



## [](#_c_specification)C Specification

The `VkQueryPoolCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkQueryPoolCreateInfo {
    VkStructureType                  sType;
    const void*                      pNext;
    VkQueryPoolCreateFlags           flags;
    VkQueryType                      queryType;
    uint32_t                         queryCount;
    VkQueryPipelineStatisticFlags    pipelineStatistics;
} VkQueryPoolCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkQueryPoolCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateFlagBits.html)
- `queryType` is a [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html) value specifying the type of queries managed by the pool.
- `queryCount` is the number of queries managed by the pool.
- `pipelineStatistics` is a bitmask of [VkQueryPipelineStatisticFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPipelineStatisticFlagBits.html) specifying which counters will be returned in queries on the new pool, as described below in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-pipestats](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-pipestats).

## [](#_description)Description

`pipelineStatistics` is ignored if `queryType` is not `VK_QUERY_TYPE_PIPELINE_STATISTICS`.

Valid Usage

- [](#VUID-VkQueryPoolCreateInfo-queryType-00791)VUID-VkQueryPoolCreateInfo-queryType-00791  
  If the [`pipelineStatisticsQuery`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineStatisticsQuery) feature is not enabled, `queryType` **must** not be `VK_QUERY_TYPE_PIPELINE_STATISTICS`
- [](#VUID-VkQueryPoolCreateInfo-meshShaderQueries-07068)VUID-VkQueryPoolCreateInfo-meshShaderQueries-07068  
  If the [`meshShaderQueries`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-meshShaderQueries) feature is not enabled, `queryType` **must** not be `VK_QUERY_TYPE_MESH_PRIMITIVES_GENERATED_EXT`
- [](#VUID-VkQueryPoolCreateInfo-meshShaderQueries-07069)VUID-VkQueryPoolCreateInfo-meshShaderQueries-07069  
  If the [`meshShaderQueries`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-meshShaderQueries) feature is not enabled, and `queryType` is `VK_QUERY_TYPE_PIPELINE_STATISTICS`, `pipelineStatistics` **must** not contain `VK_QUERY_PIPELINE_STATISTIC_TASK_SHADER_INVOCATIONS_BIT_EXT` or `VK_QUERY_PIPELINE_STATISTIC_MESH_SHADER_INVOCATIONS_BIT_EXT`
- [](#VUID-VkQueryPoolCreateInfo-queryType-00792)VUID-VkQueryPoolCreateInfo-queryType-00792  
  If `queryType` is `VK_QUERY_TYPE_PIPELINE_STATISTICS`, `pipelineStatistics` **must** be a valid combination of [VkQueryPipelineStatisticFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPipelineStatisticFlagBits.html) values
- [](#VUID-VkQueryPoolCreateInfo-queryType-09534)VUID-VkQueryPoolCreateInfo-queryType-09534  
  If `queryType` is `VK_QUERY_TYPE_PIPELINE_STATISTICS`, `pipelineStatistics` **must** not be zero
- [](#VUID-VkQueryPoolCreateInfo-queryType-03222)VUID-VkQueryPoolCreateInfo-queryType-03222  
  If `queryType` is `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, the `pNext` chain **must** include a [VkQueryPoolPerformanceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolPerformanceCreateInfoKHR.html) structure
- [](#VUID-VkQueryPoolCreateInfo-queryCount-02763)VUID-VkQueryPoolCreateInfo-queryCount-02763  
  `queryCount` **must** be greater than 0
- [](#VUID-VkQueryPoolCreateInfo-pNext-10779)VUID-VkQueryPoolCreateInfo-pNext-10779  
  If the `pNext` chain includes a [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) structure and its `videoCodecOperation` member is `VK_VIDEO_CODEC_OPERATION_DECODE_VP9_BIT_KHR`, then the [`videoDecodeVP9`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-videoDecodeVP9) feature **must** be enabled
- [](#VUID-VkQueryPoolCreateInfo-queryType-07133)VUID-VkQueryPoolCreateInfo-queryType-07133  
  If `queryType` is `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR`, then the `pNext` chain **must** include a [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) structure with `videoCodecOperation` specifying an encode operation
- [](#VUID-VkQueryPoolCreateInfo-queryType-07906)VUID-VkQueryPoolCreateInfo-queryType-07906  
  If `queryType` is `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR`, then the `pNext` chain **must** include a [VkQueryPoolVideoEncodeFeedbackCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolVideoEncodeFeedbackCreateInfoKHR.html) structure
- [](#VUID-VkQueryPoolCreateInfo-queryType-07907)VUID-VkQueryPoolCreateInfo-queryType-07907  
  If `queryType` is `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR`, and the `pNext` chain includes a [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) structure and a [VkQueryPoolVideoEncodeFeedbackCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolVideoEncodeFeedbackCreateInfoKHR.html) structure, then [VkQueryPoolVideoEncodeFeedbackCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolVideoEncodeFeedbackCreateInfoKHR.html)::`encodeFeedbackFlags` **must** not contain any bits that are not set in [VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeCapabilitiesKHR.html)::`supportedEncodeFeedbackFlags`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the [video profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-profiles) described by [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) and its `pNext` chain
- [](#VUID-VkQueryPoolCreateInfo-pNext-10248)VUID-VkQueryPoolCreateInfo-pNext-10248  
  If the `pNext` chain includes a [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) structure and its `videoCodecOperation` member is `VK_VIDEO_CODEC_OPERATION_ENCODE_AV1_BIT_KHR`, then the [`videoEncodeAV1`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-videoEncodeAV1) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-VkQueryPoolCreateInfo-sType-sType)VUID-VkQueryPoolCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO`
- [](#VUID-VkQueryPoolCreateInfo-pNext-pNext)VUID-VkQueryPoolCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkQueryPoolPerformanceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolPerformanceCreateInfoKHR.html), [VkQueryPoolPerformanceQueryCreateInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolPerformanceQueryCreateInfoINTEL.html), [VkQueryPoolVideoEncodeFeedbackCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolVideoEncodeFeedbackCreateInfoKHR.html), [VkVideoDecodeAV1ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeAV1ProfileInfoKHR.html), [VkVideoDecodeH264ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264ProfileInfoKHR.html), [VkVideoDecodeH265ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265ProfileInfoKHR.html), [VkVideoDecodeUsageInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeUsageInfoKHR.html), [VkVideoDecodeVP9ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeVP9ProfileInfoKHR.html), [VkVideoEncodeAV1ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1ProfileInfoKHR.html), [VkVideoEncodeH264ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264ProfileInfoKHR.html), [VkVideoEncodeH265ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265ProfileInfoKHR.html), [VkVideoEncodeUsageInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeUsageInfoKHR.html), or [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html)
- [](#VUID-VkQueryPoolCreateInfo-sType-unique)VUID-VkQueryPoolCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkQueryPoolCreateInfo-flags-parameter)VUID-VkQueryPoolCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkQueryPoolCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateFlagBits.html) values
- [](#VUID-VkQueryPoolCreateInfo-queryType-parameter)VUID-VkQueryPoolCreateInfo-queryType-parameter  
  `queryType` **must** be a valid [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkQueryPipelineStatisticFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPipelineStatisticFlags.html), [VkQueryPoolCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateFlags.html), [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateQueryPool.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueryPoolCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0