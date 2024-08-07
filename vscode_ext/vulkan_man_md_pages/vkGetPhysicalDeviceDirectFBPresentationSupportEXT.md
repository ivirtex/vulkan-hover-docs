# vkGetPhysicalDeviceDirectFBPresentationSupportEXT(3) Manual Page

## Name

vkGetPhysicalDeviceDirectFBPresentationSupportEXT - Query physical
device for presentation with DirectFB



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine whether a queue family of a physical device supports
presentation with DirectFB library, call:

``` c
// Provided by VK_EXT_directfb_surface
VkBool32 vkGetPhysicalDeviceDirectFBPresentationSupportEXT(
    VkPhysicalDevice                            physicalDevice,
    uint32_t                                    queueFamilyIndex,
    IDirectFB*                                  dfb);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device.

- `queueFamilyIndex` is the queue family index.

- `dfb` is a pointer to the `IDirectFB` main interface of DirectFB.

## <a href="#_description" class="anchor"></a>Description

This platform-specific function **can** be called prior to creating a
surface.

Valid Usage

- <a
  href="#VUID-vkGetPhysicalDeviceDirectFBPresentationSupportEXT-queueFamilyIndex-04119"
  id="VUID-vkGetPhysicalDeviceDirectFBPresentationSupportEXT-queueFamilyIndex-04119"></a>
  VUID-vkGetPhysicalDeviceDirectFBPresentationSupportEXT-queueFamilyIndex-04119  
  `queueFamilyIndex` **must** be less than `pQueueFamilyPropertyCount`
  returned by `vkGetPhysicalDeviceQueueFamilyProperties` for the given
  `physicalDevice`

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceDirectFBPresentationSupportEXT-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceDirectFBPresentationSupportEXT-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceDirectFBPresentationSupportEXT-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceDirectFBPresentationSupportEXT-dfb-parameter"
  id="VUID-vkGetPhysicalDeviceDirectFBPresentationSupportEXT-dfb-parameter"></a>
  VUID-vkGetPhysicalDeviceDirectFBPresentationSupportEXT-dfb-parameter  
  `dfb` **must** be a valid pointer to an `IDirectFB` value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_directfb_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_directfb_surface.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceDirectFBPresentationSupportEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
