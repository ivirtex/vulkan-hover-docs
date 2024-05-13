# vkGetPhysicalDeviceXcbPresentationSupportKHR(3) Manual Page

## Name

vkGetPhysicalDeviceXcbPresentationSupportKHR - Query physical device for
presentation to X11 server using XCB



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine whether a queue family of a physical device supports
presentation to an X11 server, using the XCB client-side library, call:

``` c
// Provided by VK_KHR_xcb_surface
VkBool32 vkGetPhysicalDeviceXcbPresentationSupportKHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t                                    queueFamilyIndex,
    xcb_connection_t*                           connection,
    xcb_visualid_t                              visual_id);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device.

- `queueFamilyIndex` is the queue family index.

- `connection` is a pointer to an `xcb_connection_t` to the X server.

- `visual_id` is an X11 visual (`xcb_visualid_t`).

## <a href="#_description" class="anchor"></a>Description

This platform-specific function **can** be called prior to creating a
surface.

Valid Usage

- <a
  href="#VUID-vkGetPhysicalDeviceXcbPresentationSupportKHR-queueFamilyIndex-01312"
  id="VUID-vkGetPhysicalDeviceXcbPresentationSupportKHR-queueFamilyIndex-01312"></a>
  VUID-vkGetPhysicalDeviceXcbPresentationSupportKHR-queueFamilyIndex-01312  
  `queueFamilyIndex` **must** be less than `pQueueFamilyPropertyCount`
  returned by `vkGetPhysicalDeviceQueueFamilyProperties` for the given
  `physicalDevice`

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceXcbPresentationSupportKHR-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceXcbPresentationSupportKHR-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceXcbPresentationSupportKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceXcbPresentationSupportKHR-connection-parameter"
  id="VUID-vkGetPhysicalDeviceXcbPresentationSupportKHR-connection-parameter"></a>
  VUID-vkGetPhysicalDeviceXcbPresentationSupportKHR-connection-parameter  
  `connection` **must** be a valid pointer to an `xcb_connection_t`
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_xcb_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_xcb_surface.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceXcbPresentationSupportKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
