# VkDebugUtilsMessageSeverityFlagBitsEXT(3) Manual Page

## Name

VkDebugUtilsMessageSeverityFlagBitsEXT - Bitmask specifying which severities of events cause a debug messenger callback



## [](#_c_specification)C Specification

Bits which **can** be set in [VkDebugUtilsMessengerCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerCreateInfoEXT.html)::`messageSeverity`, specifying event severities which cause a debug messenger to call the callback, are:

```c++
// Provided by VK_EXT_debug_utils
typedef enum VkDebugUtilsMessageSeverityFlagBitsEXT {
    VK_DEBUG_UTILS_MESSAGE_SEVERITY_VERBOSE_BIT_EXT = 0x00000001,
    VK_DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT = 0x00000010,
    VK_DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT = 0x00000100,
    VK_DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT = 0x00001000,
} VkDebugUtilsMessageSeverityFlagBitsEXT;
```

## [](#_description)Description

- `VK_DEBUG_UTILS_MESSAGE_SEVERITY_VERBOSE_BIT_EXT` specifies the most verbose output indicating all diagnostic messages from the Vulkan loader, layers, and drivers should be captured.
- `VK_DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT` specifies an informational message such as resource details that may be handy when debugging an application.
- `VK_DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT` specifies use of Vulkan that **may** expose an application bug. Such cases may not be immediately harmful, such as a fragment shader outputting to a location with no attachment. Other cases **may** point to behavior that is almost certainly bad when unintended such as using an image whose memory has not been filled. In general if you see a warning but you know that the behavior is intended/desired, then simply ignore the warning.
- `VK_DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT` specifies that the application has violated a valid usage condition of the specification.

Note

The values of [VkDebugUtilsMessageSeverityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html) are sorted based on severity. The higher the flag value, the more severe the message. This allows for simple boolean operation comparisons when looking at [VkDebugUtilsMessageSeverityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html) values.

For example:

```c++
    if (messageSeverity >= VK_DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT) {
        // Do something for warnings and errors
    }
```

In addition, space has been left between the enums to allow for later addition of new severities in between the existing values.

## [](#_see_also)See Also

[PFN\_vkDebugUtilsMessengerCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkDebugUtilsMessengerCallbackEXT.html), [VK\_EXT\_debug\_utils](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_utils.html), [VkDebugUtilsMessageSeverityFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageSeverityFlagsEXT.html), [vkSubmitDebugUtilsMessageEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSubmitDebugUtilsMessageEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDebugUtilsMessageSeverityFlagBitsEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0