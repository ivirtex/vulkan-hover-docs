# VkBuildMicromapFlagBitsEXT(3) Manual Page

## Name

VkBuildMicromapFlagBitsEXT - Bitmask specifying additional parameters for micromap builds



## [](#_c_specification)C Specification

Bits which **can** be set in [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html)::`flags` specifying additional parameters for micromap builds, are:

```c++
// Provided by VK_EXT_opacity_micromap
typedef enum VkBuildMicromapFlagBitsEXT {
    VK_BUILD_MICROMAP_PREFER_FAST_TRACE_BIT_EXT = 0x00000001,
    VK_BUILD_MICROMAP_PREFER_FAST_BUILD_BIT_EXT = 0x00000002,
    VK_BUILD_MICROMAP_ALLOW_COMPACTION_BIT_EXT = 0x00000004,
} VkBuildMicromapFlagBitsEXT;
```

## [](#_description)Description

- `VK_BUILD_MICROMAP_PREFER_FAST_TRACE_BIT_EXT` specifies that the given micromap build **should** prioritize trace performance over build time.
- `VK_BUILD_MICROMAP_PREFER_FAST_BUILD_BIT_EXT` specifies that the given micromap build **should** prioritize build time over trace performance.

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkBuildMicromapFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildMicromapFlagsEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBuildMicromapFlagBitsEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0