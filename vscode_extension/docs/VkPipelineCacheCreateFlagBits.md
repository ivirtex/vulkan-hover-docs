# VkPipelineCacheCreateFlagBits(3) Manual Page

## Name

VkPipelineCacheCreateFlagBits - Bitmask specifying the behavior of the pipeline cache



## [](#_c_specification)C Specification

Bits which **can** be set in [VkPipelineCacheCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCacheCreateInfo.html)::`flags`, specifying behavior of the pipeline cache, are:

```c++
// Provided by VK_VERSION_1_3, VK_KHR_maintenance8, VK_EXT_pipeline_creation_cache_control
typedef enum VkPipelineCacheCreateFlagBits {
  // Provided by VK_VERSION_1_3
    VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT = 0x00000001,
  // Provided by VK_KHR_maintenance8
    VK_PIPELINE_CACHE_CREATE_INTERNALLY_SYNCHRONIZED_MERGE_BIT_KHR = 0x00000008,
  // Provided by VK_EXT_pipeline_creation_cache_control
    VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT_EXT = VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT,
} VkPipelineCacheCreateFlagBits;
```

## [](#_description)Description

- `VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT` specifies that all commands that modify the created [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html) will be [externally synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fundamentals-threadingbehavior). When set, the implementation **may** skip any unnecessary processing needed to support simultaneous modification from multiple threads where allowed.
- `VK_PIPELINE_CACHE_CREATE_INTERNALLY_SYNCHRONIZED_MERGE_BIT_KHR` specifies that when the created [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html) is used as the `dstCache` parameter of [vkMergePipelineCaches](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMergePipelineCaches.html), it does not need to be [externally synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fundamentals-threadingbehavior). This flag is mutually exclusive with `VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT`.

## [](#_see_also)See Also

[VK\_EXT\_pipeline\_creation\_cache\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pipeline_creation_cache_control.html), [VK\_KHR\_maintenance8](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance8.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkPipelineCacheCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCacheCreateFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineCacheCreateFlagBits)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0