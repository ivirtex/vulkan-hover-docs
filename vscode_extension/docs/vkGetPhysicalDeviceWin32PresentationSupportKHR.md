# vkGetPhysicalDeviceWin32PresentationSupportKHR(3) Manual Page

## Name

vkGetPhysicalDeviceWin32PresentationSupportKHR - Query queue family
support for presentation on a Win32 display



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine whether a queue family of a physical device supports
presentation to the Microsoft Windows desktop, call:

``` c
// Provided by VK_KHR_win32_surface
VkBool32 vkGetPhysicalDeviceWin32PresentationSupportKHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t                                    queueFamilyIndex);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device.

- `queueFamilyIndex` is the queue family index.

## <a href="#_description" class="anchor"></a>Description

This platform-specific function **can** be called prior to creating a
surface.

Valid Usage

- <a
  href="#VUID-vkGetPhysicalDeviceWin32PresentationSupportKHR-queueFamilyIndex-01309"
  id="VUID-vkGetPhysicalDeviceWin32PresentationSupportKHR-queueFamilyIndex-01309"></a>
  VUID-vkGetPhysicalDeviceWin32PresentationSupportKHR-queueFamilyIndex-01309  
  `queueFamilyIndex` **must** be less than `pQueueFamilyPropertyCount`
  returned by `vkGetPhysicalDeviceQueueFamilyProperties` for the given
  `physicalDevice`

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceWin32PresentationSupportKHR-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceWin32PresentationSupportKHR-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceWin32PresentationSupportKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_win32_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_win32_surface.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceWin32PresentationSupportKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
