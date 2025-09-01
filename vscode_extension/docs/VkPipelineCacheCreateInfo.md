# VkPipelineCacheCreateInfo(3) Manual Page

## Name

VkPipelineCacheCreateInfo - Structure specifying parameters of a newly created pipeline cache



## [](#_c_specification)C Specification

The `VkPipelineCacheCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkPipelineCacheCreateInfo {
    VkStructureType               sType;
    const void*                   pNext;
    VkPipelineCacheCreateFlags    flags;
    size_t                        initialDataSize;
    const void*                   pInitialData;
} VkPipelineCacheCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkPipelineCacheCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCacheCreateFlagBits.html) specifying the behavior of the pipeline cache.
- `initialDataSize` is the number of bytes in `pInitialData`. If `initialDataSize` is zero, the pipeline cache will initially be empty.
- `pInitialData` is a pointer to previously retrieved pipeline cache data. If the pipeline cache data is incompatible (as defined below) with the device, the pipeline cache will be initially empty. If `initialDataSize` is zero, `pInitialData` is ignored.

Valid Usage

- [](#VUID-VkPipelineCacheCreateInfo-initialDataSize-00768)VUID-VkPipelineCacheCreateInfo-initialDataSize-00768  
  If `initialDataSize` is not `0`, it **must** be equal to the size of `pInitialData`, as returned by `vkGetPipelineCacheData` when `pInitialData` was originally retrieved
- [](#VUID-VkPipelineCacheCreateInfo-initialDataSize-00769)VUID-VkPipelineCacheCreateInfo-initialDataSize-00769  
  If `initialDataSize` is not `0`, `pInitialData` **must** have been retrieved from a previous call to `vkGetPipelineCacheData`
- [](#VUID-VkPipelineCacheCreateInfo-pipelineCreationCacheControl-02892)VUID-VkPipelineCacheCreateInfo-pipelineCreationCacheControl-02892  
  If the [`pipelineCreationCacheControl`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineCreationCacheControl) feature is not enabled, `flags` **must** not include `VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT`
- [](#VUID-VkPipelineCacheCreateInfo-maintenance8-10200)VUID-VkPipelineCacheCreateInfo-maintenance8-10200  
  If the [`maintenance8`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance8) feature is not enabled, `flags` **must** not include `VK_PIPELINE_CACHE_CREATE_INTERNALLY_SYNCHRONIZED_MERGE_BIT_KHR`
- [](#VUID-VkPipelineCacheCreateInfo-flags-10201)VUID-VkPipelineCacheCreateInfo-flags-10201  
  If `flags` includes `VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT`, it **must** not include `VK_PIPELINE_CACHE_CREATE_INTERNALLY_SYNCHRONIZED_MERGE_BIT_KHR`

Valid Usage (Implicit)

- [](#VUID-VkPipelineCacheCreateInfo-sType-sType)VUID-VkPipelineCacheCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_CACHE_CREATE_INFO`
- [](#VUID-VkPipelineCacheCreateInfo-pNext-pNext)VUID-VkPipelineCacheCreateInfo-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkPipelineCacheCreateInfo-flags-parameter)VUID-VkPipelineCacheCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkPipelineCacheCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCacheCreateFlagBits.html) values
- [](#VUID-VkPipelineCacheCreateInfo-pInitialData-parameter)VUID-VkPipelineCacheCreateInfo-pInitialData-parameter  
  If `initialDataSize` is not `0`, `pInitialData` **must** be a valid pointer to an array of `initialDataSize` bytes

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkPipelineCacheCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCacheCreateFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreatePipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreatePipelineCache.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineCacheCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0