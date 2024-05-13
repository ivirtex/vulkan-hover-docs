# VkPipelineCacheCreateFlagBits(3) Manual Page

## Name

VkPipelineCacheCreateFlagBits - Bitmask specifying the behavior of the
pipeline cache



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkPipelineCacheCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheCreateInfo.html)::`flags`,
specifying behavior of the pipeline cache, are:

``` c
// Provided by VK_EXT_pipeline_creation_cache_control
typedef enum VkPipelineCacheCreateFlagBits {
  // Provided by VK_VERSION_1_3
    VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT = 0x00000001,
  // Provided by VK_EXT_pipeline_creation_cache_control
    VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT_EXT = VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT,
} VkPipelineCacheCreateFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT` specifies that
  all commands that modify the created
  [VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html) will be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fundamentals-threadingbehavior"
  target="_blank" rel="noopener">externally synchronized</a>. When set,
  the implementation **may** skip any unnecessary processing needed to
  support simultaneous modification from multiple threads where allowed.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_pipeline_creation_cache_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_pipeline_creation_cache_control.html),
[VkPipelineCacheCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheCreateFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineCacheCreateFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
