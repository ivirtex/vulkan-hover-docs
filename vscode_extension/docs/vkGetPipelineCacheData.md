# vkGetPipelineCacheData(3) Manual Page

## Name

vkGetPipelineCacheData - Get the data store from a pipeline cache



## [](#_c_specification)C Specification

Data **can** be retrieved from a pipeline cache object using the command:

```c++
// Provided by VK_VERSION_1_0
VkResult vkGetPipelineCacheData(
    VkDevice                                    device,
    VkPipelineCache                             pipelineCache,
    size_t*                                     pDataSize,
    void*                                       pData);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the pipeline cache.
- `pipelineCache` is the pipeline cache to retrieve data from.
- `pDataSize` is a pointer to a `size_t` value related to the amount of data in the pipeline cache, as described below.
- `pData` is either `NULL` or a pointer to a buffer.

## [](#_description)Description

If `pData` is `NULL`, then the maximum size of the data that **can** be retrieved from the pipeline cache, in bytes, is returned in `pDataSize`. Otherwise, `pDataSize` **must** point to a variable set by the application to the size of the buffer, in bytes, pointed to by `pData`, and on return the variable is overwritten with the amount of data actually written to `pData`. If `pDataSize` is less than the maximum size that **can** be retrieved by the pipeline cache, at most `pDataSize` bytes will be written to `pData`, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that not all of the pipeline cache was returned.

Any data written to `pData` is valid and **can** be provided as the `pInitialData` member of the [VkPipelineCacheCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCacheCreateInfo.html) structure passed to `vkCreatePipelineCache`.

Two calls to `vkGetPipelineCacheData` with the same parameters **must** retrieve the same data unless a command that modifies the contents of the cache is called between them.

The initial bytes written to `pData` **must** be a header as described in the [Pipeline Cache Header](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-cache-header) section.

If `pDataSize` is less than what is necessary to store this header, nothing will be written to `pData` and zero will be written to `pDataSize`.

Note

This query does not behave consistently with the behavior described in [Opaque Binary Data Results](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fundamentals-binaryresults), for historical reasons.

If the amount of data available is larger than the passed `pDataSize`, the query returns up to the size of the passed buffer, and signals overflow with a `VK_INCOMPLETE` success status instead of returning a `VK_ERROR_NOT_ENOUGH_SPACE_KHR` error status.

Valid Usage (Implicit)

- [](#VUID-vkGetPipelineCacheData-device-parameter)VUID-vkGetPipelineCacheData-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetPipelineCacheData-pipelineCache-parameter)VUID-vkGetPipelineCacheData-pipelineCache-parameter  
  `pipelineCache` **must** be a valid [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html) handle
- [](#VUID-vkGetPipelineCacheData-pDataSize-parameter)VUID-vkGetPipelineCacheData-pDataSize-parameter  
  `pDataSize` **must** be a valid pointer to a `size_t` value
- [](#VUID-vkGetPipelineCacheData-pData-parameter)VUID-vkGetPipelineCacheData-pData-parameter  
  If the value referenced by `pDataSize` is not `0`, and `pData` is not `NULL`, `pData` **must** be a valid pointer to an array of `pDataSize` bytes
- [](#VUID-vkGetPipelineCacheData-pipelineCache-parent)VUID-vkGetPipelineCacheData-pipelineCache-parent  
  `pipelineCache` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_INCOMPLETE`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPipelineCacheData).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0