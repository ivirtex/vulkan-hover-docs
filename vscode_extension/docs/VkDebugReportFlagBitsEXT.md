# VkDebugReportFlagBitsEXT(3) Manual Page

## Name

VkDebugReportFlagBitsEXT - Bitmask specifying events which cause a debug report callback



## [](#_c_specification)C Specification

Bits which **can** be set in [VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackCreateInfoEXT.html)::`flags`, specifying events which cause a debug report, are:

```c++
// Provided by VK_EXT_debug_report
typedef enum VkDebugReportFlagBitsEXT {
    VK_DEBUG_REPORT_INFORMATION_BIT_EXT = 0x00000001,
    VK_DEBUG_REPORT_WARNING_BIT_EXT = 0x00000002,
    VK_DEBUG_REPORT_PERFORMANCE_WARNING_BIT_EXT = 0x00000004,
    VK_DEBUG_REPORT_ERROR_BIT_EXT = 0x00000008,
    VK_DEBUG_REPORT_DEBUG_BIT_EXT = 0x00000010,
} VkDebugReportFlagBitsEXT;
```

## [](#_description)Description

- `VK_DEBUG_REPORT_ERROR_BIT_EXT` specifies that the application has violated a valid usage condition of the specification.
- `VK_DEBUG_REPORT_WARNING_BIT_EXT` specifies use of Vulkan that **may** expose an application bug. Such cases may not be immediately harmful, such as a fragment shader outputting to a location with no attachment. Other cases **may** point to behavior that is almost certainly bad when unintended such as using an image whose memory has not been filled. In general if you see a warning but you know that the behavior is intended/desired, then simply ignore the warning.
- `VK_DEBUG_REPORT_PERFORMANCE_WARNING_BIT_EXT` specifies a potentially non-optimal use of Vulkan, e.g. using [vkCmdClearColorImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdClearColorImage.html) when setting [VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescription.html)::`loadOp` to `VK_ATTACHMENT_LOAD_OP_CLEAR` would have worked.
- `VK_DEBUG_REPORT_INFORMATION_BIT_EXT` specifies an informational message such as resource details that may be handy when debugging an application.
- `VK_DEBUG_REPORT_DEBUG_BIT_EXT` specifies diagnostic information from the implementation and layers.

## [](#_see_also)See Also

[VK\_EXT\_debug\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_report.html), [VkDebugReportFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportFlagsEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDebugReportFlagBitsEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0