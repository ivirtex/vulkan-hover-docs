# vkReleaseDisplayEXT(3) Manual Page

## Name

vkReleaseDisplayEXT - Release access to an acquired VkDisplayKHR



## [](#_c_specification)C Specification

To release a previously acquired display, call:

```c++
// Provided by VK_EXT_direct_mode_display
VkResult vkReleaseDisplayEXT(
    VkPhysicalDevice                            physicalDevice,
    VkDisplayKHR                                display);
```

## [](#_parameters)Parameters

- `physicalDevice` The physical device the display is on.
- `display` The display to release control of.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkReleaseDisplayEXT-physicalDevice-parameter)VUID-vkReleaseDisplayEXT-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkReleaseDisplayEXT-display-parameter)VUID-vkReleaseDisplayEXT-display-parameter  
  `display` **must** be a valid [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html) handle
- [](#VUID-vkReleaseDisplayEXT-display-parent)VUID-vkReleaseDisplayEXT-display-parent  
  `display` **must** have been created, allocated, or retrieved from `physicalDevice`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_direct\_mode\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_direct_mode_display.html), [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkReleaseDisplayEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0