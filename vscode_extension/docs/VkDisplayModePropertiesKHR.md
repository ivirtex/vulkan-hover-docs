# VkDisplayModePropertiesKHR(3) Manual Page

## Name

VkDisplayModePropertiesKHR - Structure describing display mode properties



## [](#_c_specification)C Specification

The `VkDisplayModePropertiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_display
typedef struct VkDisplayModePropertiesKHR {
    VkDisplayModeKHR              displayMode;
    VkDisplayModeParametersKHR    parameters;
} VkDisplayModePropertiesKHR;
```

## [](#_members)Members

- `displayMode` is a handle to the display mode described in this structure. This handle will be valid for the lifetime of the Vulkan instance.
- `parameters` is a [VkDisplayModeParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeParametersKHR.html) structure describing the display parameters associated with `displayMode`.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_KHR\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_display.html), [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeKHR.html), [VkDisplayModeParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeParametersKHR.html), [VkDisplayModeProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeProperties2KHR.html), [vkGetDisplayModePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDisplayModePropertiesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDisplayModePropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0