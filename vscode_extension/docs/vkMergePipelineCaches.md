# vkMergePipelineCaches(3) Manual Page

## Name

vkMergePipelineCaches - Combine the data stores of pipeline caches



## [](#_c_specification)C Specification

Pipeline cache objects **can** be merged using the command:

```c++
// Provided by VK_VERSION_1_0
VkResult vkMergePipelineCaches(
    VkDevice                                    device,
    VkPipelineCache                             dstCache,
    uint32_t                                    srcCacheCount,
    const VkPipelineCache*                      pSrcCaches);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the pipeline cache objects.
- `dstCache` is the handle of the pipeline cache to merge results into.
- `srcCacheCount` is the length of the `pSrcCaches` array.
- `pSrcCaches` is a pointer to an array of pipeline cache handles, which will be merged into `dstCache`. The previous contents of `dstCache` are included after the merge.

## [](#_description)Description

Note

The details of the merge operation are implementation-dependent, but implementations **should** merge the contents of the specified pipelines and prune duplicate entries.

Valid Usage

- [](#VUID-vkMergePipelineCaches-dstCache-00770)VUID-vkMergePipelineCaches-dstCache-00770  
  `dstCache` **must** not appear in the list of source caches
- [](#VUID-vkMergePipelineCaches-dstCache-10202)VUID-vkMergePipelineCaches-dstCache-10202  
  Host access to `dstCache` **must** be externally synchronized if it was not created with `VK_PIPELINE_CACHE_CREATE_INTERNALLY_SYNCHRONIZED_MERGE_BIT_KHR`

Valid Usage (Implicit)

- [](#VUID-vkMergePipelineCaches-device-parameter)VUID-vkMergePipelineCaches-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkMergePipelineCaches-dstCache-parameter)VUID-vkMergePipelineCaches-dstCache-parameter  
  `dstCache` **must** be a valid [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html) handle
- [](#VUID-vkMergePipelineCaches-pSrcCaches-parameter)VUID-vkMergePipelineCaches-pSrcCaches-parameter  
  `pSrcCaches` **must** be a valid pointer to an array of `srcCacheCount` valid [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html) handles
- [](#VUID-vkMergePipelineCaches-srcCacheCount-arraylength)VUID-vkMergePipelineCaches-srcCacheCount-arraylength  
  `srcCacheCount` **must** be greater than `0`
- [](#VUID-vkMergePipelineCaches-dstCache-parent)VUID-vkMergePipelineCaches-dstCache-parent  
  `dstCache` **must** have been created, allocated, or retrieved from `device`
- [](#VUID-vkMergePipelineCaches-pSrcCaches-parent)VUID-vkMergePipelineCaches-pSrcCaches-parent  
  Each element of `pSrcCaches` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkMergePipelineCaches).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0