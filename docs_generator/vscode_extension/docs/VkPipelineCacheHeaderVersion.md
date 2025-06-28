# VkPipelineCacheHeaderVersion(3) Manual Page

## Name

VkPipelineCacheHeaderVersion - Encode pipeline cache version



## [](#_c_specification)C Specification

Possible values of the `headerVersion` value of the pipeline cache header are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkPipelineCacheHeaderVersion {
    VK_PIPELINE_CACHE_HEADER_VERSION_ONE = 1,
} VkPipelineCacheHeaderVersion;
```

## [](#_description)Description

- `VK_PIPELINE_CACHE_HEADER_VERSION_ONE` specifies version one of the pipeline cache, described by [VkPipelineCacheHeaderVersionOne](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCacheHeaderVersionOne.html).

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkPipelineCacheHeaderVersionOne](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCacheHeaderVersionOne.html), [vkCreatePipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreatePipelineCache.html), [vkGetPipelineCacheData](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineCacheData.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineCacheHeaderVersion)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0