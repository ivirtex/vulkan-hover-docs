# vkGetPipelineCacheData(3) Manual Page

## Name

vkGetPipelineCacheData - Get the data store from a pipeline cache



## <a href="#_c_specification" class="anchor"></a>C Specification

Data **can** be retrieved from a pipeline cache object using the
command:

``` c
// Provided by VK_VERSION_1_0
VkResult vkGetPipelineCacheData(
    VkDevice                                    device,
    VkPipelineCache                             pipelineCache,
    size_t*                                     pDataSize,
    void*                                       pData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the pipeline cache.

- `pipelineCache` is the pipeline cache to retrieve data from.

- `pDataSize` is a pointer to a `size_t` value related to the amount of
  data in the pipeline cache, as described below.

- `pData` is either `NULL` or a pointer to a buffer.

## <a href="#_description" class="anchor"></a>Description

If `pData` is `NULL`, then the maximum size of the data that **can** be
retrieved from the pipeline cache, in bytes, is returned in `pDataSize`.
Otherwise, `pDataSize` **must** point to a variable set by the user to
the size of the buffer, in bytes, pointed to by `pData`, and on return
the variable is overwritten with the amount of data actually written to
`pData`. If `pDataSize` is less than the maximum size that **can** be
retrieved by the pipeline cache, at most `pDataSize` bytes will be
written to `pData`, and `VK_INCOMPLETE` will be returned instead of
`VK_SUCCESS`, to indicate that not all of the pipeline cache was
returned.

Any data written to `pData` is valid and **can** be provided as the
`pInitialData` member of the
[VkPipelineCacheCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheCreateInfo.html) structure
passed to `vkCreatePipelineCache`.

Two calls to `vkGetPipelineCacheData` with the same parameters **must**
retrieve the same data unless a command that modifies the contents of
the cache is called between them.

The initial bytes written to `pData` **must** be a header as described
in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-cache-header"
target="_blank" rel="noopener">Pipeline Cache Header</a> section.

If `pDataSize` is less than what is necessary to store this header,
nothing will be written to `pData` and zero will be written to
`pDataSize`.

Valid Usage (Implicit)

- <a href="#VUID-vkGetPipelineCacheData-device-parameter"
  id="VUID-vkGetPipelineCacheData-device-parameter"></a>
  VUID-vkGetPipelineCacheData-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetPipelineCacheData-pipelineCache-parameter"
  id="VUID-vkGetPipelineCacheData-pipelineCache-parameter"></a>
  VUID-vkGetPipelineCacheData-pipelineCache-parameter  
  `pipelineCache` **must** be a valid
  [VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html) handle

- <a href="#VUID-vkGetPipelineCacheData-pDataSize-parameter"
  id="VUID-vkGetPipelineCacheData-pDataSize-parameter"></a>
  VUID-vkGetPipelineCacheData-pDataSize-parameter  
  `pDataSize` **must** be a valid pointer to a `size_t` value

- <a href="#VUID-vkGetPipelineCacheData-pData-parameter"
  id="VUID-vkGetPipelineCacheData-pData-parameter"></a>
  VUID-vkGetPipelineCacheData-pData-parameter  
  If the value referenced by `pDataSize` is not `0`, and `pData` is not
  `NULL`, `pData` **must** be a valid pointer to an array of `pDataSize`
  bytes

- <a href="#VUID-vkGetPipelineCacheData-pipelineCache-parent"
  id="VUID-vkGetPipelineCacheData-pipelineCache-parent"></a>
  VUID-vkGetPipelineCacheData-pipelineCache-parent  
  `pipelineCache` **must** have been created, allocated, or retrieved
  from `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPipelineCacheData"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
