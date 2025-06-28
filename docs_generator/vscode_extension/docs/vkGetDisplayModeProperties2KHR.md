# vkGetDisplayModeProperties2KHR(3) Manual Page

## Name

vkGetDisplayModeProperties2KHR - Query information about the available display modes.



## [](#_c_specification)C Specification

To query the properties of a device’s built-in display modes, call:

```c++
// Provided by VK_KHR_get_display_properties2
VkResult vkGetDisplayModeProperties2KHR(
    VkPhysicalDevice                            physicalDevice,
    VkDisplayKHR                                display,
    uint32_t*                                   pPropertyCount,
    VkDisplayModeProperties2KHR*                pProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device associated with `display`.
- `display` is the display to query.
- `pPropertyCount` is a pointer to an integer related to the number of display modes available or queried, as described below.
- `pProperties` is either `NULL` or a pointer to an array of `VkDisplayModeProperties2KHR` structures.

## [](#_description)Description

`vkGetDisplayModeProperties2KHR` behaves similarly to [vkGetDisplayModePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDisplayModePropertiesKHR.html), with the ability to return extended information via chained output structures.

Valid Usage (Implicit)

- [](#VUID-vkGetDisplayModeProperties2KHR-physicalDevice-parameter)VUID-vkGetDisplayModeProperties2KHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetDisplayModeProperties2KHR-display-parameter)VUID-vkGetDisplayModeProperties2KHR-display-parameter  
  `display` **must** be a valid [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html) handle
- [](#VUID-vkGetDisplayModeProperties2KHR-pPropertyCount-parameter)VUID-vkGetDisplayModeProperties2KHR-pPropertyCount-parameter  
  `pPropertyCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetDisplayModeProperties2KHR-pProperties-parameter)VUID-vkGetDisplayModeProperties2KHR-pProperties-parameter  
  If the value referenced by `pPropertyCount` is not `0`, and `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer to an array of `pPropertyCount` [VkDisplayModeProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeProperties2KHR.html) structures
- [](#VUID-vkGetDisplayModeProperties2KHR-display-parent)VUID-vkGetDisplayModeProperties2KHR-display-parent  
  `display` **must** have been created, allocated, or retrieved from `physicalDevice`

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_INCOMPLETE`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_KHR\_get\_display\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_display_properties2.html), [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html), [VkDisplayModeProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeProperties2KHR.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDisplayModeProperties2KHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0