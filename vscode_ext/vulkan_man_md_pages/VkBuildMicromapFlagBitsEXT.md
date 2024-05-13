# VkBuildMicromapFlagBitsEXT(3) Manual Page

## Name

VkBuildMicromapFlagBitsEXT - Bitmask specifying additional parameters
for micromap builds



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html)::`flags`
specifying additional parameters for micromap builds, are:

``` c
// Provided by VK_EXT_opacity_micromap
typedef enum VkBuildMicromapFlagBitsEXT {
    VK_BUILD_MICROMAP_PREFER_FAST_TRACE_BIT_EXT = 0x00000001,
    VK_BUILD_MICROMAP_PREFER_FAST_BUILD_BIT_EXT = 0x00000002,
    VK_BUILD_MICROMAP_ALLOW_COMPACTION_BIT_EXT = 0x00000004,
} VkBuildMicromapFlagBitsEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_BUILD_MICROMAP_PREFER_FAST_TRACE_BIT_EXT` indicates that the given
  micromap build **should** prioritize trace performance over build
  time.

- `VK_BUILD_MICROMAP_PREFER_FAST_BUILD_BIT_EXT` indicates that the given
  micromap build **should** prioritize build time over trace
  performance.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkBuildMicromapFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildMicromapFlagsEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBuildMicromapFlagBitsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
