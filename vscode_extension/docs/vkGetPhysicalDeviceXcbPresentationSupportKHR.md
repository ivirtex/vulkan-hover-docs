# vkGetPhysicalDeviceXcbPresentationSupportKHR(3) Manual Page

## Name

vkGetPhysicalDeviceXcbPresentationSupportKHR - Query physical device for presentation to X11 server using XCB



## [](#_c_specification)C Specification

To determine whether a queue family of a physical device supports presentation to an X11 server, using the XCB client-side library, call:

```c++
// Provided by VK_KHR_xcb_surface
VkBool32 vkGetPhysicalDeviceXcbPresentationSupportKHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t                                    queueFamilyIndex,
    xcb_connection_t*                           connection,
    xcb_visualid_t                              visual_id);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device.
- `queueFamilyIndex` is the queue family index.
- `connection` is a pointer to an `xcb_connection_t` to the X server.
- `visual_id` is an X11 visual (`xcb_visualid_t`).

## [](#_description)Description

This platform-specific function **can** be called prior to creating a surface.

Valid Usage

- [](#VUID-vkGetPhysicalDeviceXcbPresentationSupportKHR-queueFamilyIndex-01312)VUID-vkGetPhysicalDeviceXcbPresentationSupportKHR-queueFamilyIndex-01312  
  `queueFamilyIndex` **must** be less than `pQueueFamilyPropertyCount` returned by `vkGetPhysicalDeviceQueueFamilyProperties` for the given `physicalDevice`

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceXcbPresentationSupportKHR-physicalDevice-parameter)VUID-vkGetPhysicalDeviceXcbPresentationSupportKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceXcbPresentationSupportKHR-connection-parameter)VUID-vkGetPhysicalDeviceXcbPresentationSupportKHR-connection-parameter  
  `connection` **must** be a valid pointer to an `xcb_connection_t` value

## [](#_see_also)See Also

[VK\_KHR\_xcb\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_xcb_surface.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceXcbPresentationSupportKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0