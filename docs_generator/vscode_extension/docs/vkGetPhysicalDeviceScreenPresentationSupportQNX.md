# vkGetPhysicalDeviceScreenPresentationSupportQNX(3) Manual Page

## Name

vkGetPhysicalDeviceScreenPresentationSupportQNX - Query physical device for presentation to QNX Screen



## [](#_c_specification)C Specification

To determine whether a queue family of a physical device supports presentation to a QNX Screen compositor, call:

```c++
// Provided by VK_QNX_screen_surface
VkBool32 vkGetPhysicalDeviceScreenPresentationSupportQNX(
    VkPhysicalDevice                            physicalDevice,
    uint32_t                                    queueFamilyIndex,
    struct _screen_window*                      window);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device.
- `queueFamilyIndex` is the queue family index.
- `window` is the QNX Screen `window` object.

## [](#_description)Description

This platform-specific function **can** be called prior to creating a surface.

Valid Usage

- [](#VUID-vkGetPhysicalDeviceScreenPresentationSupportQNX-queueFamilyIndex-04743)VUID-vkGetPhysicalDeviceScreenPresentationSupportQNX-queueFamilyIndex-04743  
  `queueFamilyIndex` **must** be less than `pQueueFamilyPropertyCount` returned by `vkGetPhysicalDeviceQueueFamilyProperties` for the given `physicalDevice`

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceScreenPresentationSupportQNX-physicalDevice-parameter)VUID-vkGetPhysicalDeviceScreenPresentationSupportQNX-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceScreenPresentationSupportQNX-window-parameter)VUID-vkGetPhysicalDeviceScreenPresentationSupportQNX-window-parameter  
  `window` **must** be a valid pointer to a `_screen_window` value

## [](#_see_also)See Also

[VK\_QNX\_screen\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QNX_screen_surface.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceScreenPresentationSupportQNX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0