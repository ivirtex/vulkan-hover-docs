# vkGetPhysicalDeviceXlibPresentationSupportKHR(3) Manual Page

## Name

vkGetPhysicalDeviceXlibPresentationSupportKHR - Query physical device
for presentation to X11 server using Xlib



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine whether a queue family of a physical device supports
presentation to an X11 server, using the Xlib client-side library, call:

``` c
// Provided by VK_KHR_xlib_surface
VkBool32 vkGetPhysicalDeviceXlibPresentationSupportKHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t                                    queueFamilyIndex,
    Display*                                    dpy,
    VisualID                                    visualID);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device.

- `queueFamilyIndex` is the queue family index.

- `dpy` is a pointer to an Xlib `Display` connection to the server.

- `visualId` is an X11 visual (`VisualID`).

## <a href="#_description" class="anchor"></a>Description

This platform-specific function **can** be called prior to creating a
surface.

Valid Usage

- <a
  href="#VUID-vkGetPhysicalDeviceXlibPresentationSupportKHR-queueFamilyIndex-01315"
  id="VUID-vkGetPhysicalDeviceXlibPresentationSupportKHR-queueFamilyIndex-01315"></a>
  VUID-vkGetPhysicalDeviceXlibPresentationSupportKHR-queueFamilyIndex-01315  
  `queueFamilyIndex` **must** be less than `pQueueFamilyPropertyCount`
  returned by `vkGetPhysicalDeviceQueueFamilyProperties` for the given
  `physicalDevice`

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceXlibPresentationSupportKHR-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceXlibPresentationSupportKHR-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceXlibPresentationSupportKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceXlibPresentationSupportKHR-dpy-parameter"
  id="VUID-vkGetPhysicalDeviceXlibPresentationSupportKHR-dpy-parameter"></a>
  VUID-vkGetPhysicalDeviceXlibPresentationSupportKHR-dpy-parameter  
  `dpy` **must** be a valid pointer to a `Display` value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_xlib_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_xlib_surface.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceXlibPresentationSupportKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
