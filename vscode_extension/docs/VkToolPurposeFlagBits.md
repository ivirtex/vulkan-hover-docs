# VkToolPurposeFlagBits(3) Manual Page

## Name

VkToolPurposeFlagBits - Bitmask specifying the purposes of an active tool



## [](#_c_specification)C Specification

Bits which **can** be set in [VkPhysicalDeviceToolProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceToolProperties.html)::`purposes`, specifying the purposes of an active tool, are:

```c++
// Provided by VK_VERSION_1_3
typedef enum VkToolPurposeFlagBits {
    VK_TOOL_PURPOSE_VALIDATION_BIT = 0x00000001,
    VK_TOOL_PURPOSE_PROFILING_BIT = 0x00000002,
    VK_TOOL_PURPOSE_TRACING_BIT = 0x00000004,
    VK_TOOL_PURPOSE_ADDITIONAL_FEATURES_BIT = 0x00000008,
    VK_TOOL_PURPOSE_MODIFYING_FEATURES_BIT = 0x00000010,
  // Provided by VK_EXT_debug_report with VK_EXT_tooling_info, VK_EXT_debug_utils with VK_EXT_tooling_info
    VK_TOOL_PURPOSE_DEBUG_REPORTING_BIT_EXT = 0x00000020,
  // Provided by VK_EXT_debug_marker with VK_EXT_tooling_info, VK_EXT_debug_utils with VK_EXT_tooling_info
    VK_TOOL_PURPOSE_DEBUG_MARKERS_BIT_EXT = 0x00000040,
  // Provided by VK_EXT_tooling_info
    VK_TOOL_PURPOSE_VALIDATION_BIT_EXT = VK_TOOL_PURPOSE_VALIDATION_BIT,
  // Provided by VK_EXT_tooling_info
    VK_TOOL_PURPOSE_PROFILING_BIT_EXT = VK_TOOL_PURPOSE_PROFILING_BIT,
  // Provided by VK_EXT_tooling_info
    VK_TOOL_PURPOSE_TRACING_BIT_EXT = VK_TOOL_PURPOSE_TRACING_BIT,
  // Provided by VK_EXT_tooling_info
    VK_TOOL_PURPOSE_ADDITIONAL_FEATURES_BIT_EXT = VK_TOOL_PURPOSE_ADDITIONAL_FEATURES_BIT,
  // Provided by VK_EXT_tooling_info
    VK_TOOL_PURPOSE_MODIFYING_FEATURES_BIT_EXT = VK_TOOL_PURPOSE_MODIFYING_FEATURES_BIT,
} VkToolPurposeFlagBits;
```

or the equivalent

```c++
// Provided by VK_EXT_tooling_info
typedef VkToolPurposeFlagBits VkToolPurposeFlagBitsEXT;
```

## [](#_description)Description

- `VK_TOOL_PURPOSE_VALIDATION_BIT` specifies that the tool provides validation of API usage.
- `VK_TOOL_PURPOSE_PROFILING_BIT` specifies that the tool provides profiling of API usage.
- `VK_TOOL_PURPOSE_TRACING_BIT` specifies that the tool is capturing data about the application’s API usage, including anything from simple logging to capturing data for later replay.
- `VK_TOOL_PURPOSE_ADDITIONAL_FEATURES_BIT` specifies that the tool provides additional API features/extensions on top of the underlying implementation.
- `VK_TOOL_PURPOSE_MODIFYING_FEATURES_BIT` specifies that the tool modifies the API features/limits/extensions presented to the application.
- `VK_TOOL_PURPOSE_DEBUG_REPORTING_BIT_EXT` specifies that the tool reports additional information to the application via callbacks specified by [vkCreateDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDebugReportCallbackEXT.html) or [vkCreateDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDebugUtilsMessengerEXT.html)
- `VK_TOOL_PURPOSE_DEBUG_MARKERS_BIT_EXT` specifies that the tool consumes [debug markers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#debugging-debug-markers) or [object debug annotation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#debugging-object-debug-annotation), [queue labels](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#debugging-queue-labels), or [command buffer labels](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#debugging-command-buffer-labels)

## [](#_see_also)See Also

[VK\_EXT\_tooling\_info](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_tooling_info.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkToolPurposeFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkToolPurposeFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkToolPurposeFlagBits).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0