# VkDisplayModeParametersKHR(3) Manual Page

## Name

VkDisplayModeParametersKHR - Structure describing display parameters associated with a display mode



## [](#_c_specification)C Specification

The `VkDisplayModeParametersKHR` structure is defined as:

```c++
// Provided by VK_KHR_display
typedef struct VkDisplayModeParametersKHR {
    VkExtent2D    visibleRegion;
    uint32_t      refreshRate;
} VkDisplayModeParametersKHR;
```

## [](#_members)Members

- `visibleRegion` is the 2D extents of the visible region.
- `refreshRate` is a `uint32_t` that is the number of times the display is refreshed each second multiplied by 1000.

## [](#_description)Description

Note

For example, a 60Hz display mode would report a `refreshRate` of 60,000.

Valid Usage

- [](#VUID-VkDisplayModeParametersKHR-width-01990)VUID-VkDisplayModeParametersKHR-width-01990  
  The `width` member of `visibleRegion` **must** be greater than `0`
- [](#VUID-VkDisplayModeParametersKHR-height-01991)VUID-VkDisplayModeParametersKHR-height-01991  
  The `height` member of `visibleRegion` **must** be greater than `0`
- [](#VUID-VkDisplayModeParametersKHR-refreshRate-01992)VUID-VkDisplayModeParametersKHR-refreshRate-01992  
  `refreshRate` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_KHR\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_display.html), [VkDisplayModeCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeCreateInfoKHR.html), [VkDisplayModePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModePropertiesKHR.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDisplayModeParametersKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0