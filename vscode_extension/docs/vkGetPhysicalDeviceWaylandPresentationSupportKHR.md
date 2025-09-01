# vkGetPhysicalDeviceWaylandPresentationSupportKHR(3) Manual Page

## Name

vkGetPhysicalDeviceWaylandPresentationSupportKHR - Query physical device for presentation to Wayland



## [](#_c_specification)C Specification

To determine whether a queue family of a physical device supports presentation to a Wayland compositor, call:

```c++
// Provided by VK_KHR_wayland_surface
VkBool32 vkGetPhysicalDeviceWaylandPresentationSupportKHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t                                    queueFamilyIndex,
    struct wl_display*                          display);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device.
- `queueFamilyIndex` is the queue family index.
- `display` is a pointer to the `wl_display` associated with a Wayland compositor.

## [](#_description)Description

This platform-specific function **can** be called prior to creating a surface.

Valid Usage

- [](#VUID-vkGetPhysicalDeviceWaylandPresentationSupportKHR-queueFamilyIndex-01306)VUID-vkGetPhysicalDeviceWaylandPresentationSupportKHR-queueFamilyIndex-01306  
  `queueFamilyIndex` **must** be less than `pQueueFamilyPropertyCount` returned by `vkGetPhysicalDeviceQueueFamilyProperties` for the given `physicalDevice`

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceWaylandPresentationSupportKHR-physicalDevice-parameter)VUID-vkGetPhysicalDeviceWaylandPresentationSupportKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceWaylandPresentationSupportKHR-display-parameter)VUID-vkGetPhysicalDeviceWaylandPresentationSupportKHR-display-parameter  
  `display` **must** be a valid pointer to a `wl_display` value

## [](#_see_also)See Also

[VK\_KHR\_wayland\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_wayland_surface.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceWaylandPresentationSupportKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0