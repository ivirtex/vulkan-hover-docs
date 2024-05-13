# VkPipelineCacheCreateInfo(3) Manual Page

## Name

VkPipelineCacheCreateInfo - Structure specifying parameters of a newly
created pipeline cache



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineCacheCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkPipelineCacheCreateInfo {
    VkStructureType               sType;
    const void*                   pNext;
    VkPipelineCacheCreateFlags    flags;
    size_t                        initialDataSize;
    const void*                   pInitialData;
} VkPipelineCacheCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkPipelineCacheCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheCreateFlagBits.html)
  specifying the behavior of the pipeline cache.

- `initialDataSize` is the number of bytes in `pInitialData`. If
  `initialDataSize` is zero, the pipeline cache will initially be empty.

- `pInitialData` is a pointer to previously retrieved pipeline cache
  data. If the pipeline cache data is incompatible (as defined below)
  with the device, the pipeline cache will be initially empty. If
  `initialDataSize` is zero, `pInitialData` is ignored.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkPipelineCacheCreateInfo-initialDataSize-00768"
  id="VUID-VkPipelineCacheCreateInfo-initialDataSize-00768"></a>
  VUID-VkPipelineCacheCreateInfo-initialDataSize-00768  
  If `initialDataSize` is not `0`, it **must** be equal to the size of
  `pInitialData`, as returned by `vkGetPipelineCacheData` when
  `pInitialData` was originally retrieved

- <a href="#VUID-VkPipelineCacheCreateInfo-initialDataSize-00769"
  id="VUID-VkPipelineCacheCreateInfo-initialDataSize-00769"></a>
  VUID-VkPipelineCacheCreateInfo-initialDataSize-00769  
  If `initialDataSize` is not `0`, `pInitialData` **must** have been
  retrieved from a previous call to `vkGetPipelineCacheData`

- <a
  href="#VUID-VkPipelineCacheCreateInfo-pipelineCreationCacheControl-02892"
  id="VUID-VkPipelineCacheCreateInfo-pipelineCreationCacheControl-02892"></a>
  VUID-VkPipelineCacheCreateInfo-pipelineCreationCacheControl-02892  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineCreationCacheControl"
  target="_blank"
  rel="noopener"><code>pipelineCreationCacheControl</code></a> feature
  is not enabled, `flags` **must** not include
  `VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT`

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineCacheCreateInfo-sType-sType"
  id="VUID-VkPipelineCacheCreateInfo-sType-sType"></a>
  VUID-VkPipelineCacheCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_CACHE_CREATE_INFO`

- <a href="#VUID-VkPipelineCacheCreateInfo-pNext-pNext"
  id="VUID-VkPipelineCacheCreateInfo-pNext-pNext"></a>
  VUID-VkPipelineCacheCreateInfo-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkPipelineCacheCreateInfo-flags-parameter"
  id="VUID-VkPipelineCacheCreateInfo-flags-parameter"></a>
  VUID-VkPipelineCacheCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkPipelineCacheCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheCreateFlagBits.html)
  values

- <a href="#VUID-VkPipelineCacheCreateInfo-pInitialData-parameter"
  id="VUID-VkPipelineCacheCreateInfo-pInitialData-parameter"></a>
  VUID-VkPipelineCacheCreateInfo-pInitialData-parameter  
  If `initialDataSize` is not `0`, `pInitialData` **must** be a valid
  pointer to an array of `initialDataSize` bytes

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkPipelineCacheCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheCreateFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreatePipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreatePipelineCache.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineCacheCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
