# VkQueryPoolCreateInfo(3) Manual Page

## Name

VkQueryPoolCreateInfo - Structure specifying parameters of a newly
created query pool



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkQueryPoolCreateInfo` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `queryType` is a [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html) value specifying the
  type of queries managed by the pool.

- `queryCount` is the number of queries managed by the pool.

- `pipelineStatistics` is a bitmask of
  [VkQueryPipelineStatisticFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlagBits.html)
  specifying which counters will be returned in queries on the new pool,
  as described below in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-pipestats"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-pipestats</a>.

## <a href="#_description" class="anchor"></a>Description

`pipelineStatistics` is ignored if `queryType` is not
`VK_QUERY_TYPE_PIPELINE_STATISTICS`.

Valid Usage

- <a href="#VUID-VkQueryPoolCreateInfo-queryType-00791"
  id="VUID-VkQueryPoolCreateInfo-queryType-00791"></a>
  VUID-VkQueryPoolCreateInfo-queryType-00791  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineStatisticsQuery"
  target="_blank" rel="noopener"><code>pipelineStatisticsQuery</code></a>
  feature is not enabled, `queryType` **must** not be
  `VK_QUERY_TYPE_PIPELINE_STATISTICS`

- <a href="#VUID-VkQueryPoolCreateInfo-meshShaderQueries-07068"
  id="VUID-VkQueryPoolCreateInfo-meshShaderQueries-07068"></a>
  VUID-VkQueryPoolCreateInfo-meshShaderQueries-07068  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-meshShaderQueries"
  target="_blank" rel="noopener"><code>meshShaderQueries</code></a>
  feature is not enabled, `queryType` **must** not be
  `VK_QUERY_TYPE_MESH_PRIMITIVES_GENERATED_EXT`

- <a href="#VUID-VkQueryPoolCreateInfo-meshShaderQueries-07069"
  id="VUID-VkQueryPoolCreateInfo-meshShaderQueries-07069"></a>
  VUID-VkQueryPoolCreateInfo-meshShaderQueries-07069  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-meshShaderQueries"
  target="_blank" rel="noopener"><code>meshShaderQueries</code></a>
  feature is not enabled, and `queryType` is
  `VK_QUERY_TYPE_PIPELINE_STATISTICS`, `pipelineStatistics` **must** not
  contain `VK_QUERY_PIPELINE_STATISTIC_TASK_SHADER_INVOCATIONS_BIT_EXT`
  or `VK_QUERY_PIPELINE_STATISTIC_MESH_SHADER_INVOCATIONS_BIT_EXT`

- <a href="#VUID-VkQueryPoolCreateInfo-queryType-00792"
  id="VUID-VkQueryPoolCreateInfo-queryType-00792"></a>
  VUID-VkQueryPoolCreateInfo-queryType-00792  
  If `queryType` is `VK_QUERY_TYPE_PIPELINE_STATISTICS`,
  `pipelineStatistics` **must** be a valid combination of
  [VkQueryPipelineStatisticFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlagBits.html)
  values

- <a href="#VUID-VkQueryPoolCreateInfo-queryType-09534"
  id="VUID-VkQueryPoolCreateInfo-queryType-09534"></a>
  VUID-VkQueryPoolCreateInfo-queryType-09534  
  If `queryType` is `VK_QUERY_TYPE_PIPELINE_STATISTICS`,
  `pipelineStatistics` **must** not be zero

- <a href="#VUID-VkQueryPoolCreateInfo-queryType-03222"
  id="VUID-VkQueryPoolCreateInfo-queryType-03222"></a>
  VUID-VkQueryPoolCreateInfo-queryType-03222  
  If `queryType` is `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, the `pNext`
  chain **must** include a
  [VkQueryPoolPerformanceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolPerformanceCreateInfoKHR.html)
  structure

- <a href="#VUID-VkQueryPoolCreateInfo-queryCount-02763"
  id="VUID-VkQueryPoolCreateInfo-queryCount-02763"></a>
  VUID-VkQueryPoolCreateInfo-queryCount-02763  
  `queryCount` **must** be greater than 0

- <a href="#VUID-VkQueryPoolCreateInfo-queryType-07133"
  id="VUID-VkQueryPoolCreateInfo-queryType-07133"></a>
  VUID-VkQueryPoolCreateInfo-queryType-07133  
  If `queryType` is `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR`, then the
  `pNext` chain **must** include a
  [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html) structure with
  `videoCodecOperation` specifying an encode operation

- <a href="#VUID-VkQueryPoolCreateInfo-queryType-07906"
  id="VUID-VkQueryPoolCreateInfo-queryType-07906"></a>
  VUID-VkQueryPoolCreateInfo-queryType-07906  
  If `queryType` is `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR`, then the
  `pNext` chain **must** include a
  [VkQueryPoolVideoEncodeFeedbackCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolVideoEncodeFeedbackCreateInfoKHR.html)
  structure

- <a href="#VUID-VkQueryPoolCreateInfo-queryType-07907"
  id="VUID-VkQueryPoolCreateInfo-queryType-07907"></a>
  VUID-VkQueryPoolCreateInfo-queryType-07907  
  If `queryType` is `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR`, and the
  `pNext` chain includes a
  [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html) structure and a
  [VkQueryPoolVideoEncodeFeedbackCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolVideoEncodeFeedbackCreateInfoKHR.html)
  structure, then
  [VkQueryPoolVideoEncodeFeedbackCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolVideoEncodeFeedbackCreateInfoKHR.html)::`encodeFeedbackFlags`
  **must** not contain any bits that are not set in
  [VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeCapabilitiesKHR.html)::`supportedEncodeFeedbackFlags`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profiles"
  target="_blank" rel="noopener">video profile</a> described by
  [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html) and its `pNext`
  chain

Valid Usage (Implicit)

- <a href="#VUID-VkQueryPoolCreateInfo-sType-sType"
  id="VUID-VkQueryPoolCreateInfo-sType-sType"></a>
  VUID-VkQueryPoolCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO`

- <a href="#VUID-VkQueryPoolCreateInfo-pNext-pNext"
  id="VUID-VkQueryPoolCreateInfo-pNext-pNext"></a>
  VUID-VkQueryPoolCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkQueryPoolPerformanceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolPerformanceCreateInfoKHR.html),
  [VkQueryPoolPerformanceQueryCreateInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolPerformanceQueryCreateInfoINTEL.html),
  [VkQueryPoolVideoEncodeFeedbackCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolVideoEncodeFeedbackCreateInfoKHR.html),
  [VkVideoDecodeAV1ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1ProfileInfoKHR.html),
  [VkVideoDecodeH264ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264ProfileInfoKHR.html),
  [VkVideoDecodeH265ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265ProfileInfoKHR.html),
  [VkVideoDecodeUsageInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeUsageInfoKHR.html),
  [VkVideoEncodeH264ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264ProfileInfoKHR.html),
  [VkVideoEncodeH265ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265ProfileInfoKHR.html),
  [VkVideoEncodeUsageInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeUsageInfoKHR.html), or
  [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html)

- <a href="#VUID-VkQueryPoolCreateInfo-sType-unique"
  id="VUID-VkQueryPoolCreateInfo-sType-unique"></a>
  VUID-VkQueryPoolCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkQueryPoolCreateInfo-flags-zerobitmask"
  id="VUID-VkQueryPoolCreateInfo-flags-zerobitmask"></a>
  VUID-VkQueryPoolCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`

- <a href="#VUID-VkQueryPoolCreateInfo-queryType-parameter"
  id="VUID-VkQueryPoolCreateInfo-queryType-parameter"></a>
  VUID-VkQueryPoolCreateInfo-queryType-parameter  
  `queryType` **must** be a valid [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkQueryPipelineStatisticFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlags.html),
[VkQueryPoolCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateFlags.html),
[VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateQueryPool.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkQueryPoolCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
