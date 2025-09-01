# vkGetPhysicalDeviceWin32PresentationSupportKHR(3) Manual Page

## Name

vkGetPhysicalDeviceWin32PresentationSupportKHR - Query queue family support for presentation on a Win32 display



## [](#_c_specification)C Specification

To determine whether a queue family of a physical device supports presentation to the Microsoft Windows desktop, call:

```c++
// Provided by VK_KHR_win32_surface
VkBool32 vkGetPhysicalDeviceWin32PresentationSupportKHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t                                    queueFamilyIndex);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device.
- `queueFamilyIndex` is the queue family index.

## [](#_description)Description

This platform-specific function **can** be called prior to creating a surface.

Valid Usage

- [](#VUID-vkGetPhysicalDeviceWin32PresentationSupportKHR-queueFamilyIndex-01309)VUID-vkGetPhysicalDeviceWin32PresentationSupportKHR-queueFamilyIndex-01309  
  `queueFamilyIndex` **must** be less than `pQueueFamilyPropertyCount` returned by `vkGetPhysicalDeviceQueueFamilyProperties` for the given `physicalDevice`

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceWin32PresentationSupportKHR-physicalDevice-parameter)VUID-vkGetPhysicalDeviceWin32PresentationSupportKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle

## [](#_see_also)See Also

[VK\_KHR\_win32\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_win32_surface.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceWin32PresentationSupportKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0