# VkPipelineCacheHeaderVersion(3) Manual Page

## Name

VkPipelineCacheHeaderVersion - Encode pipeline cache version



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of the `headerVersion` value of the pipeline cache
header are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkPipelineCacheHeaderVersion {
    VK_PIPELINE_CACHE_HEADER_VERSION_ONE = 1,
} VkPipelineCacheHeaderVersion;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_PIPELINE_CACHE_HEADER_VERSION_ONE` specifies version one of the
  pipeline cache, described by
  [VkPipelineCacheHeaderVersionOne](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheHeaderVersionOne.html).

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkPipelineCacheHeaderVersionOne](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheHeaderVersionOne.html),
[vkCreatePipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreatePipelineCache.html),
[vkGetPipelineCacheData](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineCacheData.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineCacheHeaderVersion"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
