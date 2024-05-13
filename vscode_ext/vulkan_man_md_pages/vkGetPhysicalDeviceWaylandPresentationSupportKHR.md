# vkGetPhysicalDeviceWaylandPresentationSupportKHR(3) Manual Page

## Name

vkGetPhysicalDeviceWaylandPresentationSupportKHR - Query physical device
for presentation to Wayland



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine whether a queue family of a physical device supports
presentation to a Wayland compositor, call:

``` c
// Provided by VK_KHR_wayland_surface
VkBool32 vkGetPhysicalDeviceWaylandPresentationSupportKHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t                                    queueFamilyIndex,
    struct wl_display*                          display);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device.

- `queueFamilyIndex` is the queue family index.

- `display` is a pointer to the `wl_display` associated with a Wayland
  compositor.

## <a href="#_description" class="anchor"></a>Description

This platform-specific function **can** be called prior to creating a
surface.

Valid Usage

- <a
  href="#VUID-vkGetPhysicalDeviceWaylandPresentationSupportKHR-queueFamilyIndex-01306"
  id="VUID-vkGetPhysicalDeviceWaylandPresentationSupportKHR-queueFamilyIndex-01306"></a>
  VUID-vkGetPhysicalDeviceWaylandPresentationSupportKHR-queueFamilyIndex-01306  
  `queueFamilyIndex` **must** be less than `pQueueFamilyPropertyCount`
  returned by `vkGetPhysicalDeviceQueueFamilyProperties` for the given
  `physicalDevice`

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceWaylandPresentationSupportKHR-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceWaylandPresentationSupportKHR-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceWaylandPresentationSupportKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceWaylandPresentationSupportKHR-display-parameter"
  id="VUID-vkGetPhysicalDeviceWaylandPresentationSupportKHR-display-parameter"></a>
  VUID-vkGetPhysicalDeviceWaylandPresentationSupportKHR-display-parameter  
  `display` **must** be a valid pointer to a `wl_display` value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_wayland_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_wayland_surface.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceWaylandPresentationSupportKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
