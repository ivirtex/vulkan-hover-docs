# VkPipelineBinaryCreateInfoKHR(3) Manual Page

## Name

VkPipelineBinaryCreateInfoKHR - Structure specifying where to retrieve data for pipeline binary creation



## [](#_c_specification)C Specification

The `VkPipelineBinaryCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_pipeline_binary
typedef struct VkPipelineBinaryCreateInfoKHR {
    VkStructureType                          sType;
    const void*                              pNext;
    const VkPipelineBinaryKeysAndDataKHR*    pKeysAndDataInfo;
    VkPipeline                               pipeline;
    const VkPipelineCreateInfoKHR*           pPipelineCreateInfo;
} VkPipelineBinaryCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pKeysAndDataInfo` is `NULL` or a pointer to a [VkPipelineBinaryKeysAndDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKeysAndDataKHR.html) structure that contains keys and data to create the pipeline binaries from.
- `pipeline` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or a `VkPipeline` that contains data to create the pipeline binaries from.
- `pPipelineCreateInfo` is `NULL` or a pointer to a [VkPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateInfoKHR.html) structure with the pipeline creation info. This is used to probe the implementation’s internal cache for pipeline binaries.

## [](#_description)Description

When `pPipelineCreateInfo` is not `NULL`, an implementation will attempt to retrieve pipeline binary data from an internal cache external to the application if [`pipelineBinaryInternalCache`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-pipelineBinaryInternalCache) is `VK_TRUE`. Applications **can** use this to determine if a pipeline **can** be created without compilation. If the implementation fails to create a pipeline binary due to missing an internal cache entry, `VK_PIPELINE_BINARY_MISSING_KHR` is returned. If creation succeeds, the resulting binary **can** be used to create a pipeline. `VK_PIPELINE_BINARY_MISSING_KHR` **may** be returned for any reason in this situation, even if creating a pipeline binary with the same parameters that succeeded earlier.

If [`pipelineBinaryPrecompiledInternalCache`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-pipelineBinaryPrecompiledInternalCache) is `VK_TRUE`, the implementation **may** be able to create pipeline binaries even when `pPipelineCreateInfo` has not been used to create binaries before by the application.

Note

On some platforms, internal pipeline caches may be pre-populated before running the application.

Valid Usage

- [](#VUID-VkPipelineBinaryCreateInfoKHR-pipeline-09607)VUID-VkPipelineBinaryCreateInfoKHR-pipeline-09607  
  If `pipeline` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pipeline` **must** have been created with `VK_PIPELINE_CREATE_2_CAPTURE_DATA_BIT_KHR`
- [](#VUID-VkPipelineBinaryCreateInfoKHR-pipeline-09608)VUID-VkPipelineBinaryCreateInfoKHR-pipeline-09608  
  If `pipeline` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), [vkReleaseCapturedPipelineDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkReleaseCapturedPipelineDataKHR.html) **must** not have been called on `pipeline` prior to this command
- [](#VUID-VkPipelineBinaryCreateInfoKHR-pipelineBinaryInternalCache-09609)VUID-VkPipelineBinaryCreateInfoKHR-pipelineBinaryInternalCache-09609  
  If [`pipelineBinaryInternalCache`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-pipelineBinaryInternalCache) is `VK_FALSE` pPipelineCreateInfo **must** be `NULL`
- [](#VUID-VkPipelineBinaryCreateInfoKHR-device-09610)VUID-VkPipelineBinaryCreateInfoKHR-device-09610  
  If `device` was created with [VkDevicePipelineBinaryInternalCacheControlKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevicePipelineBinaryInternalCacheControlKHR.html)::`disableInternalCache` set to `VK_TRUE`, `pPipelineCreateInfo` **must** be `NULL`
- [](#VUID-VkPipelineBinaryCreateInfoKHR-pKeysAndDataInfo-09619)VUID-VkPipelineBinaryCreateInfoKHR-pKeysAndDataInfo-09619  
  One and only one of `pKeysAndDataInfo`, `pipeline`, or `pPipelineCreateInfo` **must** be non-`NULL`
- [](#VUID-VkPipelineBinaryCreateInfoKHR-pPipelineCreateInfo-09606)VUID-VkPipelineBinaryCreateInfoKHR-pPipelineCreateInfo-09606  
  If `pPipelineCreateInfo` is not `NULL`, the `pNext` chain of `pPipelineCreateInfo` **must** not set [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` to a value greater than `0`

Valid Usage (Implicit)

- [](#VUID-VkPipelineBinaryCreateInfoKHR-sType-sType)VUID-VkPipelineBinaryCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_BINARY_CREATE_INFO_KHR`
- [](#VUID-VkPipelineBinaryCreateInfoKHR-pNext-pNext)VUID-VkPipelineBinaryCreateInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkPipelineBinaryCreateInfoKHR-pKeysAndDataInfo-parameter)VUID-VkPipelineBinaryCreateInfoKHR-pKeysAndDataInfo-parameter  
  If `pKeysAndDataInfo` is not `NULL`, `pKeysAndDataInfo` **must** be a valid pointer to a valid [VkPipelineBinaryKeysAndDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKeysAndDataKHR.html) structure
- [](#VUID-VkPipelineBinaryCreateInfoKHR-pipeline-parameter)VUID-VkPipelineBinaryCreateInfoKHR-pipeline-parameter  
  If `pipeline` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle
- [](#VUID-VkPipelineBinaryCreateInfoKHR-pPipelineCreateInfo-parameter)VUID-VkPipelineBinaryCreateInfoKHR-pPipelineCreateInfo-parameter  
  If `pPipelineCreateInfo` is not `NULL`, `pPipelineCreateInfo` **must** be a valid pointer to a valid [VkPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateInfoKHR.html) structure

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_binary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_binary.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkPipelineBinaryKeysAndDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKeysAndDataKHR.html), [VkPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateInfoKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreatePipelineBinariesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreatePipelineBinariesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineBinaryCreateInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0