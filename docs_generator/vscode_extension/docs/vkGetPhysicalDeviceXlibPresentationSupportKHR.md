# vkGetPhysicalDeviceXlibPresentationSupportKHR(3) Manual Page

## Name

vkGetPhysicalDeviceXlibPresentationSupportKHR - Query physical device for presentation to X11 server using Xlib



## [](#_c_specification)C Specification

To determine whether a queue family of a physical device supports presentation to an X11 server, using the Xlib client-side library, call:

```c++
// Provided by VK_KHR_xlib_surface
VkBool32 vkGetPhysicalDeviceXlibPresentationSupportKHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t                                    queueFamilyIndex,
    Display*                                    dpy,
    VisualID                                    visualID);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device.
- `queueFamilyIndex` is the queue family index.
- `dpy` is a pointer to an Xlib `Display` connection to the server.
- `visualID` is an X11 visual (`VisualID`).

## [](#_description)Description

This platform-specific function **can** be called prior to creating a surface.

Valid Usage

- [](#VUID-vkGetPhysicalDeviceXlibPresentationSupportKHR-queueFamilyIndex-01315)VUID-vkGetPhysicalDeviceXlibPresentationSupportKHR-queueFamilyIndex-01315  
  `queueFamilyIndex` **must** be less than `pQueueFamilyPropertyCount` returned by `vkGetPhysicalDeviceQueueFamilyProperties` for the given `physicalDevice`

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceXlibPresentationSupportKHR-physicalDevice-parameter)VUID-vkGetPhysicalDeviceXlibPresentationSupportKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceXlibPresentationSupportKHR-dpy-parameter)VUID-vkGetPhysicalDeviceXlibPresentationSupportKHR-dpy-parameter  
  `dpy` **must** be a valid pointer to a `Display` value

## [](#_see_also)See Also

[VK\_KHR\_xlib\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_xlib_surface.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceXlibPresentationSupportKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0