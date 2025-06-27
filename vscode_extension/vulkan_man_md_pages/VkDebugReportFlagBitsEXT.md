# VkDebugReportFlagBitsEXT(3) Manual Page

## Name

VkDebugReportFlagBitsEXT - Bitmask specifying events which cause a debug
report callback



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackCreateInfoEXT.html)::`flags`,
specifying events which cause a debug report, are:

``` c
// Provided by VK_EXT_debug_report
typedef enum VkDebugReportFlagBitsEXT {
    VK_DEBUG_REPORT_INFORMATION_BIT_EXT = 0x00000001,
    VK_DEBUG_REPORT_WARNING_BIT_EXT = 0x00000002,
    VK_DEBUG_REPORT_PERFORMANCE_WARNING_BIT_EXT = 0x00000004,
    VK_DEBUG_REPORT_ERROR_BIT_EXT = 0x00000008,
    VK_DEBUG_REPORT_DEBUG_BIT_EXT = 0x00000010,
} VkDebugReportFlagBitsEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DEBUG_REPORT_ERROR_BIT_EXT` specifies that the application has
  violated a valid usage condition of the specification.

- `VK_DEBUG_REPORT_WARNING_BIT_EXT` specifies use of Vulkan that **may**
  expose an application bug. Such cases may not be immediately harmful,
  such as a fragment shader outputting to a location with no attachment.
  Other cases **may** point to behavior that is almost certainly bad
  when unintended such as using an image whose memory has not been
  filled. In general if you see a warning but you know that the behavior
  is intended/desired, then simply ignore the warning.

- `VK_DEBUG_REPORT_PERFORMANCE_WARNING_BIT_EXT` specifies a potentially
  non-optimal use of Vulkan, e.g. using
  [vkCmdClearColorImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdClearColorImage.html) when setting
  [VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription.html)::`loadOp` to
  `VK_ATTACHMENT_LOAD_OP_CLEAR` would have worked.

- `VK_DEBUG_REPORT_INFORMATION_BIT_EXT` specifies an informational
  message such as resource details that may be handy when debugging an
  application.

- `VK_DEBUG_REPORT_DEBUG_BIT_EXT` specifies diagnostic information from
  the implementation and layers.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_report](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_report.html),
[VkDebugReportFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportFlagsEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDebugReportFlagBitsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
