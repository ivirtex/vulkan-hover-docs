# vkGetPhysicalDeviceDirectFBPresentationSupportEXT(3) Manual Page

## Name

vkGetPhysicalDeviceDirectFBPresentationSupportEXT - Query physical device for presentation with DirectFB



## [](#_c_specification)C Specification

To determine whether a queue family of a physical device supports presentation with DirectFB library, call:

```c++
// Provided by VK_EXT_directfb_surface
VkBool32 vkGetPhysicalDeviceDirectFBPresentationSupportEXT(
    VkPhysicalDevice                            physicalDevice,
    uint32_t                                    queueFamilyIndex,
    IDirectFB*                                  dfb);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device.
- `queueFamilyIndex` is the queue family index.
- `dfb` is a pointer to the `IDirectFB` main interface of DirectFB.

## [](#_description)Description

This platform-specific function **can** be called prior to creating a surface.

Valid Usage

- [](#VUID-vkGetPhysicalDeviceDirectFBPresentationSupportEXT-queueFamilyIndex-04119)VUID-vkGetPhysicalDeviceDirectFBPresentationSupportEXT-queueFamilyIndex-04119  
  `queueFamilyIndex` **must** be less than `pQueueFamilyPropertyCount` returned by `vkGetPhysicalDeviceQueueFamilyProperties` for the given `physicalDevice`

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceDirectFBPresentationSupportEXT-physicalDevice-parameter)VUID-vkGetPhysicalDeviceDirectFBPresentationSupportEXT-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceDirectFBPresentationSupportEXT-dfb-parameter)VUID-vkGetPhysicalDeviceDirectFBPresentationSupportEXT-dfb-parameter  
  `dfb` **must** be a valid pointer to an `IDirectFB` value

## [](#_see_also)See Also

[VK\_EXT\_directfb\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_directfb_surface.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceDirectFBPresentationSupportEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0